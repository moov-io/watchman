// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download

import (
	"cmp"
	"context"
	"fmt"
	"os"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"
)

func loadOpensanctionsRecords(ctx context.Context, logger log.Logger, config Config, responseCh chan preparedList) error {
	params := senzingDownload{
		lists:  config.OpenSanctions.Lists,
		config: config,
	}

	if len(params.lists) == 0 {
		return nil
	}

	apiKey := cmp.Or(os.Getenv("OPENSANCTIONS_API_KEY"), config.OpenSanctions.ApiKey)
	if apiKey != "" {
		params.downloadOptions = append(params.downloadOptions,
			download.WithAdditionalHeaders(opensanctionsAuthHeader(apiKey)))
	}

	return prepareSenzingRecords(ctx, logger, params, responseCh)
}

func opensanctionsAuthHeader(apiKey string) map[string]string {
	out := make(map[string]string)
	if apiKey != "" {
		out["Authorization"] = fmt.Sprintf("Token %s", apiKey)
	}
	return out
}
