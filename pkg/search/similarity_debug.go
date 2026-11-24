package search

import (
	"fmt"
	"strings"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"
)

// ScoreStep is one adjustment on a score trail. Delta is the change from the previous step.
// There is no API stability guarantee.
type ScoreStep struct {
	Stage  string  `json:"stage"`
	Score  float64 `json:"score"`
	Delta  float64 `json:"delta,omitzero"`
	Reason string  `json:"reason,omitempty"`
}

// NameTransform is one Normalize() rewrite of a name when the value changed.
// There is no API stability guarantee.
type NameTransform struct {
	Stage string `json:"stage"`
	Value string `json:"value"`
}

// NameSide is one side of a name comparison.
// Original is the entity display name. Compared is the winning name when it
// differs from Original (an AKA or a former name). Prepared and Tokens are
// what was actually scored.
// There is no API stability guarantee.
type NameSide struct {
	Original string          `json:"original,omitempty"`
	Compared string          `json:"compared,omitempty"`
	Prepared string          `json:"prepared,omitempty"`
	Tokens   []string        `json:"tokens,omitempty"`
	Steps    []NameTransform `json:"steps,omitempty"`
}

// TokenPair is a greedy query-token to index-token alignment.
// An empty Query or Index means that token was unmatched.
// There is no API stability guarantee.
type TokenPair struct {
	Query string  `json:"query,omitempty"`
	Index string  `json:"index,omitempty"`
	Score float64 `json:"score"`
}

const (
	nameAgainstPrimary    = "primary"
	nameAgainstAlt        = "alt"
	nameAgainstHistorical = "historical"
)

// NameDebug is the name-comparison breakdown: which names were compared, token pairs, and score steps.
// There is no API stability guarantee.
type NameDebug struct {
	Score float64 `json:"score"`
	// Exact is true when Score is at least 0.99 after historical / quality / close-enough penalties.
	Exact      bool        `json:"exact,omitempty"`
	Historical bool        `json:"historical,omitempty"`
	QueryName  string      `json:"queryName,omitempty"` // primary or historical
	IndexName  string      `json:"indexName,omitempty"` // primary, alt, or historical
	Query      NameSide    `json:"query"`
	Index      NameSide    `json:"index"`
	Pairs      []TokenPair `json:"pairs,omitempty"`
	Steps      []ScoreStep `json:"steps,omitempty"`
}

func (ss SimilarityScore) String() string {
	var b strings.Builder
	switch ss.RankedBy {
	case RankedByEmbeddings:
		fmt.Fprintf(&b, "ranking  %.4f  embeddings\n", ss.RankingScore)
		fmt.Fprintf(&b, "similarity  %.4f  similarity diagnostic\n", ss.FinalScore)
	default:
		fmt.Fprintf(&b, "match  %.4f\n", ss.FinalScore)
	}

	if ss.Name != nil {
		n := ss.Name
		fmt.Fprintf(&b, "\nname   %.4f", n.Score)
		if n.Exact {
			b.WriteString("  exact")
		}
		if n.Historical {
			b.WriteString("  historical")
		}
		if n.QueryName != "" || n.IndexName != "" {
			fmt.Fprintf(&b, "  query=%s index=%s", emptyDash(n.QueryName), emptyDash(n.IndexName))
		}
		b.WriteByte('\n')
		writeNameSide(&b, "query", n.Query)
		writeNameSide(&b, "index", n.Index)
		writePairs(&b, n.Pairs)
		writeSteps(&b, "  name steps", n.Steps)
	}

	writeSteps(&b, "\nscore steps", ss.Steps)

	if len(ss.Pieces) > 0 {
		b.WriteString("\npieces\n")
		for _, piece := range ss.Pieces {
			if piece.FieldsCompared == 0 && piece.Score == 0 && !piece.Matched {
				continue
			}
			fmt.Fprintf(&b, "  %-16s  %.4f  w=%.0f", piece.PieceType, piece.Score, piece.Weight)
			if piece.Required {
				b.WriteString("  required")
			}
			if piece.Matched {
				b.WriteString("  matched")
			}
			if piece.Exact && piece.Score >= exactMatchThreshold {
				b.WriteString("  exact")
			}
			if piece.UniqueIdentity {
				b.WriteString("  unique-id")
			}
			b.WriteByte('\n')
		}
	}

	return strings.TrimSpace(b.String())
}

func emptyDash(s string) string {
	if s == "" {
		return "—"
	}
	return s
}

func writeNameSide(b *strings.Builder, label string, side NameSide) {
	fmt.Fprintf(b, "  %-5s  %s\n", label, emptyDash(side.Original))
	current := side.Original
	for _, step := range side.Steps {
		if step.Stage == "original" {
			continue
		}
		fmt.Fprintf(b, "         %-14s %s\n", step.Stage, step.Value)
		if step.Stage != "extracted" {
			current = step.Value
		}
	}
	if side.Compared != "" {
		fmt.Fprintf(b, "         %-14s %s\n", "compared", side.Compared)
		current = side.Compared
	}
	if side.Prepared != "" && side.Prepared != current && side.Prepared != strings.ToLower(strings.TrimSpace(side.Original)) {
		fmt.Fprintf(b, "         %-14s %s\n", "prepared", side.Prepared)
	}
	if len(side.Tokens) > 0 {
		fmt.Fprintf(b, "         %-14s [%s]\n", "tokens", strings.Join(side.Tokens, " "))
	}
}

func writePairs(b *strings.Builder, pairs []TokenPair) {
	if len(pairs) == 0 {
		return
	}
	b.WriteString("  pairs\n")
	for _, pair := range pairs {
		switch {
		case pair.Query == "":
			fmt.Fprintf(b, "    unmatched index  %s\n", pair.Index)
		case pair.Index == "":
			fmt.Fprintf(b, "    unmatched query  %s\n", pair.Query)
		default:
			fmt.Fprintf(b, "    %-16s ↔ %-16s  %.3f\n", pair.Query, pair.Index, pair.Score)
		}
	}
}

func writeSteps(b *strings.Builder, heading string, steps []ScoreStep) {
	if len(steps) == 0 {
		return
	}
	b.WriteString(heading)
	b.WriteByte('\n')
	for _, step := range steps {
		fmt.Fprintf(b, "    %-22s  %.4f", step.Stage, step.Score)
		if step.Delta != 0 {
			fmt.Fprintf(b, "  %+.4f", step.Delta)
		}
		if step.Reason != "" {
			fmt.Fprintf(b, "  %s", step.Reason)
		}
		b.WriteByte('\n')
	}
}

func nameSideFrom(entityOriginal, comparedOriginal, prepared string, tokens []string, typ EntityType) NameSide {
	original := entityOriginal
	if original == "" {
		original = comparedOriginal
	}
	side := NameSide{
		Original: original,
		Prepared: prepared,
		Tokens:   tokens,
		Steps:    namePrepareTrace(original, comparedOriginal, typ),
	}
	if comparedOriginal != "" && !strings.EqualFold(strings.TrimSpace(comparedOriginal), strings.TrimSpace(original)) {
		side.Compared = comparedOriginal
	}
	return side
}

func namePrepareTrace(display, scored string, typ EntityType) []NameTransform {
	if strings.TrimSpace(display) == "" {
		return nil
	}
	steps := []NameTransform{{Stage: "original", Value: display}}
	current := display
	add := func(stage, v string) {
		if v != current {
			steps = append(steps, NameTransform{Stage: stage, Value: v})
			current = v
		}
	}
	prefix := func(raw string) string {
		v := prepare.StripLeadingPartyLabels(raw)
		add("party-labels", v)
		primary, former := prepare.SplitExNames(v)
		add("former-names", primary)
		var extracted []string
		for _, name := range former {
			if prepareEntityName(name, typ) == "" {
				continue
			}
			extracted = append(extracted, name)
		}
		if len(extracted) > 0 {
			steps = append(steps, NameTransform{Stage: "extracted", Value: strings.Join(extracted, "; ")})
		}
		return primary
	}

	rest := prefix(display)
	if scored != "" && !strings.EqualFold(strings.TrimSpace(scored), strings.TrimSpace(display)) {
		current = scored
		rest = prefix(scored)
	}

	if typ == EntityVessel {
		next := prepare.StripTrailingVesselPlace(rest)
		add("vessel-place", next)
		rest = next
	}

	out := prepare.LowerAndRemovePunctuation(rest)
	add("normalize", out)

	if typ == EntityVessel {
		next := prepare.StripVesselNamePrefixes(out)
		add("vessel-prefixes", next)
		out = next
	}
	if typ == EntityBusiness || typ == EntityOrganization {
		next := prepare.CanonicalizeEnglishLegalForms(out)
		add("legal-forms", next)
		out = next
	}

	tokens := strings.Join(removeStopwords(out), " ")
	add("stopwords", tokens)

	if len(steps) == 1 {
		return nil
	}
	return steps
}

func originalAltNames[T any](e Entity[T]) []string {
	var alts []string
	if e.Person != nil {
		alts = e.Person.AltNames
	}
	if e.Business != nil {
		alts = e.Business.AltNames
	}
	if e.Organization != nil {
		alts = e.Organization.AltNames
	}
	if e.Aircraft != nil {
		alts = e.Aircraft.AltNames
	}
	if e.Vessel != nil {
		alts = e.Vessel.AltNames
	}
	return alts
}

func originalHistoricalNames[T any](e Entity[T]) []string {
	var out []string
	for _, hist := range e.HistoricalInfo {
		if !strings.EqualFold(hist.Type, "Former Name") || hist.Value == "" {
			continue
		}
		out = append(out, hist.Value)
	}
	for _, former := range prepare.ExNames(e.Name) {
		if prepareEntityName(former, e.Type) == "" {
			continue
		}
		out = append(out, former)
	}
	return out
}

func tokenPairsFromExplain(pairs []stringscore.TokenPair) []TokenPair {
	if len(pairs) == 0 {
		return nil
	}
	out := make([]TokenPair, len(pairs))
	for i, pair := range pairs {
		out[i] = TokenPair{Query: pair.Search, Index: pair.Index, Score: pair.Score}
	}
	return out
}

func identityTokenPairs(tokens []string) []TokenPair {
	if len(tokens) == 0 {
		return nil
	}
	out := make([]TokenPair, len(tokens))
	for i, tok := range tokens {
		out[i] = TokenPair{Query: tok, Index: tok, Score: 1}
	}
	return out
}

type scoreTrail struct {
	score   float64
	steps   []ScoreStep
	collect bool
}

func (t *scoreTrail) base(score float64) {
	t.score = score
	if t.collect {
		t.steps = append(t.steps, ScoreStep{Stage: "base", Score: score})
	}
}

func (t *scoreTrail) apply(stage, reason string, score float64) {
	t.applyStep(stage, reason, score, false)
}

func (t *scoreTrail) applyForced(stage, reason string, score float64) {
	t.applyStep(stage, reason, score, true)
}

func (t *scoreTrail) applyStep(stage, reason string, score float64, force bool) {
	if !force && score == t.score {
		return
	}
	delta := score - t.score
	t.score = score
	if t.collect {
		t.steps = append(t.steps, ScoreStep{Stage: stage, Score: score, Delta: delta, Reason: reason})
	}
}

func appendNameStep(steps []ScoreStep, stage, reason string, prev, score float64) []ScoreStep {
	step := ScoreStep{Stage: stage, Score: score, Reason: reason}
	if len(steps) > 0 {
		step.Delta = score - prev
	}
	return append(steps, step)
}
