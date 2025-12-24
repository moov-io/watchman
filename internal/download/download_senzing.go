package download

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/senzing"
)

func loadSenzingRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	locations := make(map[string]string)

	for _, loc := range conf.Senzing {
		locations[string(loc.SourceList)] = loc.Location
	}

	dl := download.New(logger, nil)
	initialDir := initialDataDirectory(conf)

	files, err := dl.GetFiles(ctx, initialDir, locations)
	if err != nil {
		return fmt.Errorf("loading senzing files: %v", err)
	}
	defer files.Close()

	for src, contents := range files {
		source := search.SourceList(src)

		rc, hashbuf := hashWriter(contents)

		entities, err := senzing.ReadEntities(rc, source)
		if err != nil {
			return fmt.Errorf("parsing %s failed: %w", source, err)
		}

		if len(entities) == 0 && conf.ErrorOnEmptyList {
			return fmt.Errorf("no entities parsed from senzing lists: %#v", source)
		}

		responseCh <- preparedList{
			ListName: source,
			Entities: entities,
			Hash:     calculateHash(hashbuf.Bytes()),
		}
	}

	return nil
}

func calculateHash(input []byte) string {
	h := sha256.Sum256(input)
	return hex.EncodeToString(h[:])
}

func hashWriter(rc io.ReadCloser) (io.Reader, *bytes.Buffer) {
	var buf bytes.Buffer
	r := io.TeeReader(rc, &buf)
	return r, &buf
}
