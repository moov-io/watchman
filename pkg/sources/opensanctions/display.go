// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"fmt"
)

const (
	baseDetailsURL = "https://www.opensanctions.org/entities/"
)

// DetailsURL returns the OpenSanctions entity page URL for a given entity ID
func DetailsURL(entityID string) string {
	return fmt.Sprintf("%s%s/", baseDetailsURL, entityID)
}
