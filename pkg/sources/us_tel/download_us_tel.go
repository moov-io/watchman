// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package us_tel

import (
	"fmt"
	"context"
	"io"
	"os"
	"encoding/json"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/download"
)

var (
	publicUSTelDownloadUrl = strx.Or(os.Getenv("US_TEL_URL"), "https://data.opensanctions.org/datasets/latest/us_state_terrorist_exclusion/targets.nested.json")
)
func DownloadUsTel(ctx context.Context, logger log.Logger, initialDir string) (map[string]io.ReadCloser, error) {
	dl := download.New(logger, nil)

	telNameAndSourceMap := make(map[string]string)
	telNameAndSourceMap["us_tel.json"] = publicUSTelDownloadUrl
	jsonBytes, _ := json.Marshal(telNameAndSourceMap)
	
	fmt.Println("Downloading US State Terrorist Exclusion List", string(jsonBytes))
	return dl.GetFiles(ctx, initialDir, telNameAndSourceMap)

}