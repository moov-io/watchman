package prepare

import (
	"runtime/debug"
	"testing"

	"github.com/abadojack/whatlanggo"
)

// TestStopwordUnionIsCurrent guards against the generated union silently drifting
// from the bbalet/stopwords lists. If the dependency is upgraded the union must be
// regenerated (go generate ./internal/prepare/...), otherwise the fast path could
// skip a token that the new version treats as a stopword.
func TestStopwordUnionIsCurrent(t *testing.T) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		t.Skip("build info unavailable")
	}
	for _, dep := range info.Deps {
		if dep.Path != "github.com/bbalet/stopwords" {
			continue
		}
		if dep.Replace != nil {
			t.Skip("bbalet/stopwords is replaced")
		}
		if dep.Version != bbaletStopwordsVersion {
			t.Fatalf("union_stopwords.go was generated from bbalet/stopwords %s but the module is %s; run go generate ./internal/prepare/...",
				bbaletStopwordsVersion, dep.Version)
		}
		return
	}
	t.Skip("github.com/bbalet/stopwords not found in build info")
}

// languageSpread covers Latin, Cyrillic and other scripts, plus Cmn (Chinese) which
// bbalet does not have a stopword list for, exercising the no-removal branch.
var languageSpread = []whatlanggo.Lang{
	whatlanggo.Eng, whatlanggo.Spa, whatlanggo.Fra, whatlanggo.Deu, whatlanggo.Ita,
	whatlanggo.Por, whatlanggo.Nld, whatlanggo.Swe, whatlanggo.Dan, whatlanggo.Pol,
	whatlanggo.Ron, whatlanggo.Ces, whatlanggo.Hun, whatlanggo.Fin, whatlanggo.Tur,
	whatlanggo.Ell, whatlanggo.Bul, whatlanggo.Lav, whatlanggo.Ind, whatlanggo.Rus,
	whatlanggo.Arb, whatlanggo.Pes, whatlanggo.Jpn, whatlanggo.Tha, whatlanggo.Khm,
	whatlanggo.Cmn,
}

// TestFastPathIsLanguageIndependent is the core correctness check. Whenever the fast
// path triggers, the result must equal the normal detect-then-clean output for every
// language. whatlanggo.Detect is non-deterministic on ties, so we assert this
// invariant directly rather than comparing detection results.
func TestFastPathIsLanguageIndependent(t *testing.T) {
	inputs := []string{
		"acme corporation",
		"john smith",
		"ryong gang",
		"o'brien holdings",
		"argo 1",
		"yv2040",
		"unit 5 building 3",
		"global trading limited",
		"shenzhen electronics",
		"abc 123 xyz",
	}

	triggered := 0
	for _, in := range inputs {
		out, ok := trivialNormalization(in)
		if !ok {
			continue
		}
		triggered++
		for _, lang := range languageSpread {
			if got := removeStopwords(in, lang); got != out {
				t.Errorf("fast path %q = %q, but cleaning with %s = %q", in, out, lang.Iso6391(), got)
			}
		}
	}
	if triggered == 0 {
		t.Fatal("no input exercised the fast path")
	}
}

// TestTrivialNormalizationGuards documents which inputs skip detection and which fall
// through to it, so the gating cannot regress unnoticed.
func TestTrivialNormalizationGuards(t *testing.T) {
	cases := []struct {
		in       string
		wantOK   bool
		wantText string // only checked when wantOK
	}{
		{"acme corporation", true, "acme corporation"},
		{"JOHN SMITH", true, "john smith"},
		{"yv2040", true, "yv2040"},     // single token matched by numberRegex, left alone
		{"bank of america", false, ""}, // "of" is a stopword
		{"argo 1", false, ""},          // bare "1" is language dependent, so detection is kept
		{"banco de mexico", false, ""}, // "de" is a stopword
		{"the smith group", false, ""}, // "the" is a stopword
		{"8th street", false, ""},      // "8th" is not a single clean segment
	}
	for _, tc := range cases {
		got, ok := trivialNormalization(tc.in)
		if ok != tc.wantOK {
			t.Errorf("trivialNormalization(%q) ok = %v, want %v", tc.in, ok, tc.wantOK)
			continue
		}
		if ok && got != tc.wantText {
			t.Errorf("trivialNormalization(%q) = %q, want %q", tc.in, got, tc.wantText)
		}
	}
}
