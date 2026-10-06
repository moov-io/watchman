package norm_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/norm"

	"github.com/stretchr/testify/require"
)

func TestWebsite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "hostname", input: "suex.io", expected: "suex.io"},
		{name: "www hostname", input: "www.gicdf.org", expected: "gicdf.org"},
		{name: "https url", input: "https://www.gicdf.org/about", expected: "gicdf.org"},
		{name: "http url with port", input: "http://www.dialog.info:80/", expected: "dialog.info"},
		{name: "trailing dot", input: "WWW.BNTD.BY.", expected: "bntd.by"},
		{name: "uppercase", input: "SUEX.IO", expected: "suex.io"},
		{name: "query string", input: "https://example.com/path?q=1#frag", expected: "example.com"},
		{name: "userinfo", input: "https://user:pass@mail.ckba.net/inbox", expected: "mail.ckba.net"},
		{name: "whitespace", input: "  www.tidewaterco.com  ", expected: "tidewaterco.com"},
		{name: "https default port", input: "https://suex.io:443", expected: "suex.io"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, norm.Website(tc.input))
		})
	}
}
