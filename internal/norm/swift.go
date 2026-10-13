package norm

import (
	"regexp"
	"strings"

	"github.com/moov-io/iso3166"
)

// ISO 9362 BIC: 4-letter institution, 2-letter country, 2-alphanumeric location,
// optional 3-alphanumeric branch. XXX is the primary office (same as BIC8).
// Captures are exactly 8 or 11 characters with a word boundary so remark
// words such as "Registration" are not taken as a BIC.
const swiftBICToken = `([A-Za-z]{6}[A-Za-z0-9]{2}(?:[A-Za-z0-9]{3})?)`

var (
	swiftLabeledRe  = regexp.MustCompile(`(?i)(?:alt\.\s*)?SWIFT\s*/\s*BIC(?:\s+CODE)?[\s:,./-]+` + swiftBICToken + `\b`)
	swiftTrailingRe = regexp.MustCompile(`(?i)\b` + swiftBICToken + `[\s,]+SWIFT\s*/\s*BIC`)
)

// Slashes are left in place so the "SWIFT/BIC" label does not collapse to SWIFTBIC.
var swiftStrip = strings.NewReplacer(" ", "", "-", "", ".", "")

// SwiftBIC returns a canonical ISO 9362 BIC, or empty if input is not a BIC.
// Punctuation and spaces are stripped, letters are uppercased, and an XXX
// branch suffix is dropped so the 8-character head office form is stored.
func SwiftBIC(input string) string {
	s := swiftStrip.Replace(strings.ToUpper(strings.TrimSpace(input)))
	if !isSwiftBIC(s) {
		return ""
	}
	if len(s) == 11 && strings.HasSuffix(s, "XXX") {
		return s[:8]
	}
	return s
}

// SwiftBIC8 returns the 8-character head-office form, or empty.
func SwiftBIC8(input string) string {
	s := SwiftBIC(input)
	if len(s) >= 8 {
		return s[:8]
	}
	return s
}

// EqualSwiftBIC reports whether two values are the same institution.
// 8- and 11-character forms match on the first eight characters.
func EqualSwiftBIC(a, b string) bool {
	a8, b8 := SwiftBIC8(a), SwiftBIC8(b)
	return a8 != "" && a8 == b8
}

// IsSwiftBIC reports whether input is an 8- or 11-character ISO 9362 BIC.
func IsSwiftBIC(input string) bool {
	return SwiftBIC(input) != ""
}

func isSwiftBIC(s string) bool {
	if len(s) != 8 && len(s) != 11 {
		return false
	}
	for i := 0; i < 6; i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
	}
	for i := 6; i < len(s); i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return false
		}
	}
	// Characters 5-6 are the ISO 3166 country code. This rejects 8-letter
	// English words (PASSPORT, SWIFTBIC) pulled out of ID lists.
	return iso3166.GetName(s[4:6]) != ""
}

// ParseSwiftBICs extracts unique canonical BICs from labeled remarks
// (`SWIFT/BIC HNRBBY2X`, `CKLBCNBJ, SWIFT/BIC`) and from fields that are
// themselves a BIC.
func ParseSwiftBICs(parts ...string) []string {
	var out []string
	seen := make(map[string]struct{})
	add := func(code string) {
		n := SwiftBIC(code)
		if n == "" {
			return
		}
		if _, ok := seen[n]; ok {
			return
		}
		seen[n] = struct{}{}
		out = append(out, n)
	}

	for _, part := range parts {
		if part == "" {
			continue
		}
		for _, m := range swiftLabeledRe.FindAllStringSubmatch(part, -1) {
			if len(m) > 1 {
				add(m[1])
			}
		}
		for _, m := range swiftTrailingRe.FindAllStringSubmatch(part, -1) {
			if len(m) > 1 {
				add(m[1])
			}
		}
		add(part)
		if strings.ContainsAny(part, ";,") {
			for _, tok := range strings.FieldsFunc(part, func(r rune) bool {
				return r == ';' || r == ','
			}) {
				add(tok)
			}
		}
	}
	return out
}
