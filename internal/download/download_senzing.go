package download

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"slices"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/senzing"
)

func loadSenzingRecords(ctx context.Context, logger log.Logger, config Config, ignoredLists []search.SourceList, responseCh chan preparedList) error {
	params := senzingDownload{
		lists:        config.Senzing,
		config:       config,
		ignoredLists: ignoredLists,
	}
	return prepareSenzingRecords(ctx, logger, params, responseCh)
}

type senzingDownload struct {
	lists        []SenzingList
	config       Config
	ignoredLists []search.SourceList

	downloadOptions []download.Option
}

func prepareSenzingRecords(ctx context.Context, logger log.Logger, params senzingDownload, responseCh chan preparedList) error {
	dl := download.New(logger, nil, params.downloadOptions...)
	initialDir := initialDataDirectory(params.config)

	// Download and process each list individually so that failures are
	// per-source.  This avoids a mixed ignored/non-ignored batch where an
	// ignored list's download failure would prevent non-ignored lists from
	// loading.
	for _, loc := range params.lists {
		if err := processSenzingList(ctx, logger, dl, initialDir, params, loc, responseCh); err != nil {
			return err
		}
	}

	return nil
}

// processSenzingList downloads and parses one Senzing-format list (used for
// both config.Senzing and config.OpenSanctions lists). It sends a preparedList
// to responseCh on success. For lists present in params.ignoredLists, download,
// read, or empty-list errors are logged at warn and suppressed.
func processSenzingList(ctx context.Context, logger log.Logger, dl *download.Downloader, initialDir string, params senzingDownload, loc SenzingList, responseCh chan preparedList) error {
	source := normalizeListName(loc.SourceList)
	ignored := slices.Contains(params.ignoredLists, source)

	locations := map[string]string{
		string(source): loc.Location,
	}

	files, err := dl.GetFiles(ctx, initialDir, locations)
	if err != nil {
		if ignored {
			logger.Warn().Logf("ignoring download error for %s: %v", source, err)
			return nil
		}
		return fmt.Errorf("loading senzing file %s: %w", source, err)
	}
	defer files.Close()

	contents, ok := files[string(source)]
	if !ok {
		if ignored {
			logger.Warn().Logf("ignoring download failure for %s: file not available", source)
			return nil
		}
		return fmt.Errorf("download failed for senzing list %s: file not available", source)
	}

	r, hashbuf := hashWriter(contents)

	entities, err := senzing.ReadEntities(r, source)
	if err != nil {
		if ignored {
			logger.Warn().Logf("ignoring error parsing %s: %v", source, err)
			return nil
		}
		return fmt.Errorf("parsing %s failed: %w", source, err)
	}

	if len(entities) == 0 && params.config.ErrorOnEmptyList {
		if ignored {
			logger.Warn().Logf("ignoring empty list for %s", source)
			return nil
		}
		return fmt.Errorf("no entities parsed from senzing list: %#v", source)
	}

	responseCh <- preparedList{
		ListName: source,
		Entities: entities,
		Hash:     calculateHash(hashbuf.Bytes()),
	}
	return nil
}

func calculateHash(input []byte) string {
	h := sha256.Sum256(input)
	return hex.EncodeToString(h[:])
}

func hashWriter(r io.Reader) (io.Reader, *bytes.Buffer) {
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)
	return tee, &buf
}
