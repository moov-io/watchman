package norm_test

import (
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/norm"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/publicsuffix"
)

func TestDomainKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "empty", input: "", expected: nil},
		{name: "registrable domain", input: "suex.io", expected: []string{"suex.io"}},
		{name: "www stripped then eTLD+1", input: "www.gicdf.org", expected: []string{"gicdf.org"}},
		{name: "https url", input: "https://www.gicdf.org/about", expected: []string{"gicdf.org"}},
		{name: "subdomain walk", input: "mail.ckba.net", expected: []string{"mail.ckba.net", "ckba.net"}},
		{name: "deeper walk", input: "a.b.suex.io", expected: []string{"a.b.suex.io", "b.suex.io", "suex.io"}},
		{name: "multi-label public suffix", input: "www.nitc.co.ir", expected: []string{"nitc.co.ir"}},
		{name: "subdomain of multi-label suffix", input: "mail.nitc.co.ir", expected: []string{"mail.nitc.co.ir", "nitc.co.ir"}},
		{name: "www plus multi-label suffix", input: "www.example.co.uk", expected: []string{"example.co.uk"}},
		{name: "public suffix itself", input: "co.uk", expected: nil},
		{name: "tld itself", input: "ru", expected: nil},
		{name: "ipv4", input: "127.0.0.1", expected: nil},
		{name: "trailing dot", input: "WWW.BNTD.BY.", expected: []string{"bntd.by"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := norm.DomainKeys(tc.input)
			require.Equal(t, tc.expected, got)
			if len(got) == 0 {
				return
			}
			suffix, _ := publicsuffix.PublicSuffix(got[0])
			require.NotContains(t, got, suffix)
			require.True(t, strings.HasSuffix(got[0], got[len(got)-1]))
		})
	}
}

func TestDomainKeysFromEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "empty", input: "", expected: nil},
		{name: "listed org email", input: "info@gicdf.org", expected: []string{"gicdf.org"}},
		{name: "listed subdomain mailbox", input: "press@mail.ckba.net", expected: []string{"mail.ckba.net", "ckba.net"}},
		{name: "gmail skipped", input: "khoroshev1@icloud.com", expected: nil},
		{name: "yahoo skipped", input: "user@yahoo.co.uk", expected: nil},
		{name: "mail.ru skipped", input: "name@mail.ru", expected: nil},
		{name: "no at-sign", input: "gicdf.org", expected: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, norm.DomainKeysFromEmail(tc.input))
		})
	}
}
