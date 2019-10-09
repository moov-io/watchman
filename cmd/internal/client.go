package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/ofac"
	moov "github.com/moov-io/ofac/client"
)

const (
	DefaultApiAddress = "https://api.moov.io"
)

// addr reads flagApiAddress and flagLocal to compute the HTTP address used for connecting with OFAC.
func addr(address string, local bool) string {
	if local {
		// If '-local and -address <foo>' use <foo>
		if address != DefaultApiAddress {
			return strings.TrimSuffix(address, "/")
		} else {
			return "http://localhost" + bind.HTTP("ofac")
		}
	} else {
		address = strings.TrimSuffix(address, "/")
		// -address isn't changed, so assume Moov's API (needs extra path added)
		if address == DefaultApiAddress {
			return address + "/v1/ofac"
		}
		return address
	}
}

func Config(address string, local bool) *moov.Configuration {
	conf := moov.NewConfiguration()
	conf.BasePath = addr(address, local)

	conf.UserAgent = fmt.Sprintf("moov/ofactest:%s", ofac.Version)
	conf.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 1 * time.Minute,
		},
	}

	return conf
}
