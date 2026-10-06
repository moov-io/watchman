package norm

import "strings"

// Website returns a comparable host from a website, URL, or hostname.
// Scheme, userinfo, path, query, fragment, default ports, a trailing dot,
// and a leading "www." are dropped. The original Contact.Websites value is
// left unchanged; this is stored on PreparedFields.
func Website(input string) string {
	s := strings.TrimSpace(strings.ToLower(input))
	if s == "" {
		return ""
	}

	if i := strings.Index(s, "://"); i >= 0 {
		s = s[i+3:]
	}
	if i := strings.LastIndex(s, "@"); i >= 0 {
		s = s[i+1:]
	}
	if i := strings.IndexAny(s, "/?#"); i >= 0 {
		s = s[:i]
	}

	s = strings.TrimSuffix(s, ".")
	s = strings.TrimSuffix(s, ":443")
	s = strings.TrimSuffix(s, ":80")
	s = strings.TrimPrefix(s, "www.")
	return strings.TrimSpace(s)
}
