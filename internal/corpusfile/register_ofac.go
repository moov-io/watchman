package corpusfile

import (
	"encoding/gob"

	"github.com/moov-io/watchman/pkg/sources/ofac"
)

func init() {
	// OFAC SourceData types used when building snapshots from ofactest corpora.
	gob.Register(ofac.SDN{})
	gob.Register(ofac.Address{})
	gob.Register(ofac.AlternateIdentity{})
	gob.Register(ofac.SDNComments{})
	gob.Register([]ofac.Address{})
	gob.Register([]ofac.AlternateIdentity{})
}
