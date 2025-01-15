package search

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareName(t *testing.T) {
	var buf bytes.Buffer

	tests := []struct {
		name          string
		query         Entity[any]
		index         Entity[any]
		expectedScore float64
		shouldMatch   bool
		exact         bool
	}{
		{
			name: "exact match",
			query: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "case insensitive match",
			query: Entity[any]{
				Name: "aerocaribbean airlines",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "punctuation differences",
			query: Entity[any]{
				Name: "ANGLO CARIBBEAN CO LTD",
			},
			index: Entity[any]{
				Name: "ANGLO-CARIBBEAN CO., LTD.",
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "slight misspelling",
			query: Entity[any]{
				Name: "AEROCARRIBEAN AIRLINES",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			expectedScore: 0.95,
			shouldMatch:   true,
			exact:         false,
		},
		{
			name: "word reordering",
			query: Entity[any]{
				Name: "AIRLINES AEROCARIBBEAN",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			expectedScore: 0.90,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "extra words in query",
			query: Entity[any]{
				Name: "THE AEROCARIBBEAN AIRLINES COMPANY",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			expectedScore: 0.85,
			shouldMatch:   false,
			exact:         false,
		},
		{
			name: "historical name match",
			query: Entity[any]{
				Name: "OLD AEROCARIBBEAN",
			},
			index: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
				HistoricalInfo: []HistoricalInfo{
					{
						Type:  "Former Name",
						Value: "OLD AEROCARIBBEAN",
					},
				},
			},
			expectedScore: 0.90,
			shouldMatch:   true,
			exact:         false,
		},
		{
			name: "alternative name match for person",
			query: Entity[any]{
				Name: "JOHN MICHAEL SMITH",
				Type: EntityPerson,
				Person: &Person{
					Name: "JOHN MICHAEL SMITH",
				},
			},
			index: Entity[any]{
				Name: "JOHN SMITH",
				Type: EntityPerson,
				Person: &Person{
					Name:     "JOHN SMITH",
					AltNames: []string{"JOHN MICHAEL SMITH", "J.M. SMITH"},
				},
			},
			expectedScore: 0.95,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "minimum term matches",
			query: Entity[any]{
				Name: "CARIBBEAN TRADING LIMITED",
			},
			index: Entity[any]{
				Name: "PACIFIC TRADING LIMITED",
			},
			expectedScore: 0.8628,
			shouldMatch:   true,
			exact:         false,
		},
		{
			name: "completely different names",
			query: Entity[any]{
				Name: "AEROCARIBBEAN AIRLINES",
			},
			index: Entity[any]{
				Name: "BANCO NACIONAL DE CUBA",
			},
			expectedScore: 0.4479,
			shouldMatch:   false,
			exact:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareName(&buf, tt.query, tt.index, 1.0)

			assert.InDelta(t, tt.expectedScore, result.score, 0.1,
				"expected score %v but got %v", tt.expectedScore, result.score)
			assert.Equal(t, tt.shouldMatch, result.matched,
				"expected matched=%v but got matched=%v", tt.shouldMatch, result.matched)
			assert.Equal(t, tt.exact, result.exact,
				"expected exact=%v but got exact=%v", tt.exact, result.exact)
		})
	}
}

func TestCompareEntityTitlesFuzzy(t *testing.T) {
	var buf bytes.Buffer

	tests := []struct {
		name          string
		query         Entity[any]
		index         Entity[any]
		expectedScore float64
		shouldMatch   bool
		exact         bool
	}{
		{
			name: "exact title match",
			query: Entity[any]{
				Person: &Person{
					Titles: []string{"Chief Executive Officer"},
				},
			},
			index: Entity[any]{
				Person: &Person{
					Titles: []string{"Chief Executive Officer"},
				},
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "abbreviated title match",
			query: Entity[any]{
				Person: &Person{
					Titles: []string{"CEO"},
				},
			},
			index: Entity[any]{
				Person: &Person{
					Titles: []string{"Chief Executive Officer"},
				},
			},
			expectedScore: 0.0, // TODO(adam): needs fixed
			shouldMatch:   false,
			exact:         false,
		},
		{
			name: "multiple titles with partial matches",
			query: Entity[any]{
				Person: &Person{
					Titles: []string{"CEO", "Director of Operations"},
				},
			},
			index: Entity[any]{
				Person: &Person{
					Titles: []string{"Chief Executive Officer", "Operations Director"},
				},
			},
			expectedScore: 0.50,
			shouldMatch:   false,
			exact:         false,
		},
		{
			name: "similar but not exact titles",
			query: Entity[any]{
				Person: &Person{
					Titles: []string{"Senior Technical Manager"},
				},
			},
			index: Entity[any]{
				Person: &Person{
					Titles: []string{"Technical Manager"},
				},
			},
			expectedScore: 0.0, // TODO(adam): needs fixed
			shouldMatch:   false,
			exact:         false,
		},
		{
			name: "no matching titles",
			query: Entity[any]{
				Person: &Person{
					Titles: []string{"Chief Financial Officer"},
				},
			},
			index: Entity[any]{
				Person: &Person{
					Titles: []string{"Sales Director", "Regional Manager"},
				},
			},
			expectedScore: 0.0,
			shouldMatch:   false,
			exact:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareEntityTitlesFuzzy(&buf, tt.query, tt.index, 1.0)

			assert.InDelta(t, tt.expectedScore, result.score, 0.1,
				"expected score %v but got %v", tt.expectedScore, result.score)
			assert.Equal(t, tt.shouldMatch, result.matched,
				"expected matched=%v but got matched=%v", tt.shouldMatch, result.matched)
			assert.Equal(t, tt.exact, result.exact,
				"expected exact=%v but got exact=%v", tt.exact, result.exact)
		})
	}
}

func TestCompareAffiliationsFuzzy(t *testing.T) {
	var buf bytes.Buffer

	tests := []struct {
		name          string
		query         Entity[any]
		index         Entity[any]
		expectedScore float64
		shouldMatch   bool
		exact         bool
	}{
		{
			name: "exact affiliation match",
			query: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			index: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "similar affiliation with related type",
			query: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "Banco Nacional Cuba",
						Type:       "owned by",
					},
				},
			},
			index: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			expectedScore: 0.90,
			shouldMatch:   true,
			exact:         true,
		},
		{
			name: "multiple affiliations with partial matches",
			query: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "CARIBBEAN TRADING CO",
						Type:       "linked to",
					},
					{
						EntityName: "BANCO CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			index: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "CARIBBEAN TRADING COMPANY",
						Type:       "associated with",
					},
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "parent company",
					},
				},
			},
			expectedScore: 1.0,
			shouldMatch:   true,
			exact:         false,
		},
		{
			name: "no matching affiliations",
			query: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "ACME CORPORATION",
						Type:       "owned by",
					},
				},
			},
			index: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			expectedScore: 0.3956,
			shouldMatch:   false,
			exact:         false,
		},
		{
			name: "empty affiliations",
			query: Entity[any]{
				Affiliations: []Affiliation{},
			},
			index: Entity[any]{
				Affiliations: []Affiliation{
					{
						EntityName: "BANCO NACIONAL DE CUBA",
						Type:       "subsidiary of",
					},
				},
			},
			expectedScore: 0.0,
			shouldMatch:   false,
			exact:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareAffiliationsFuzzy(&buf, tt.query, tt.index, 1.0)

			assert.InDelta(t, tt.expectedScore, result.score, 0.1,
				"expected score %v but got %v", tt.expectedScore, result.score)
			assert.Equal(t, tt.shouldMatch, result.matched,
				"expected matched=%v but got matched=%v", tt.shouldMatch, result.matched)
			assert.Equal(t, tt.exact, result.exact,
				"expected exact=%v but got exact=%v", tt.exact, result.exact)
		})
	}
}

// Helper function tests

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard name",
			input:    "AEROCARIBBEAN AIRLINES",
			expected: "aerocaribbean airlines",
		},
		{
			name:     "name with punctuation",
			input:    "ANGLO-CARIBBEAN CO., LTD.",
			expected: "anglo caribbean co ltd",
		},
		{
			name:     "extra whitespace",
			input:    "  BANCO   NACIONAL  DE   CUBA  ",
			expected: "banco nacional de cuba",
		},
		{
			name:     "mixed case with special chars",
			input:    "Banco.Nacional_de@Cuba",
			expected: "banco nacional de cuba",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special chars",
			input:    ".,!@#$%^&*()",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNormalizeTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard title",
			input:    "Chief Executive Officer",
			expected: "chief executive officer",
		},
		{
			name:     "title with punctuation",
			input:    "Sr. Vice-President, Operations",
			expected: "sr vice-president operations",
		},
		{
			name:     "abbreviated title",
			input:    "CEO & CFO",
			expected: "ceo cfo",
		},
		{
			name:     "extra whitespace",
			input:    "  Senior   Technical   Manager  ",
			expected: "senior technical manager",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeTitle(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateNameScore(t *testing.T) {
	tests := []struct {
		name          string
		queryName     string
		indexName     string
		expectedScore float64
	}{
		{
			name:          "exact match",
			queryName:     "banco nacional de cuba",
			indexName:     "banco nacional de cuba",
			expectedScore: 1.0,
		},
		{
			name:          "close match",
			queryName:     "banco nacional cuba",
			indexName:     "banco nacional de cuba",
			expectedScore: 0.95,
		},
		{
			name:          "partial match",
			queryName:     "banco cuba",
			indexName:     "banco nacional de cuba",
			expectedScore: 0.9210,
		},
		{
			name:          "word reordering",
			queryName:     "nacional banco cuba",
			indexName:     "banco nacional de cuba",
			expectedScore: 0.9842,
		},
		{
			name:          "completely different",
			queryName:     "aerocaribbean airlines",
			indexName:     "banco nacional de cuba",
			expectedScore: 0.0,
		},
		{
			name:          "empty query",
			queryName:     "",
			indexName:     "banco nacional de cuba",
			expectedScore: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := calculateNameScore(tt.queryName, tt.indexName)
			assert.InDelta(t, tt.expectedScore, score, 0.1)
		})
	}
}

func TestCalculateTypeScore(t *testing.T) {
	tests := []struct {
		name          string
		queryType     string
		indexType     string
		expectedScore float64
	}{
		{
			name:          "exact match",
			queryType:     "owned by",
			indexType:     "owned by",
			expectedScore: 1.0,
		},
		{
			name:          "same group - ownership",
			queryType:     "owned by",
			indexType:     "subsidiary of",
			expectedScore: 0.8,
		},
		{
			name:          "same group - control",
			queryType:     "controlled by",
			indexType:     "operates",
			expectedScore: 0.8,
		},
		{
			name:          "same group - association",
			queryType:     "linked to",
			indexType:     "associated with",
			expectedScore: 0.8,
		},
		{
			name:          "same group - leadership",
			queryType:     "led by",
			indexType:     "headed by",
			expectedScore: 0.8,
		},
		{
			name:          "different groups",
			queryType:     "owned by",
			indexType:     "linked to",
			expectedScore: 0.0,
		},
		{
			name:          "case insensitive",
			queryType:     "OWNED BY",
			indexType:     "owned by",
			expectedScore: 1.0,
		},
		{
			name:          "with extra spaces",
			queryType:     "  owned  by  ",
			indexType:     "owned by",
			expectedScore: 1.0,
		},
		{
			name:          "unknown types",
			queryType:     "unknown relation",
			indexType:     "other relation",
			expectedScore: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := calculateTypeScore(tt.queryType, tt.indexType)
			assert.InDelta(t, tt.expectedScore, score, 0.1)
		})
	}
}

func TestFilterSignificantTerms(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "all significant terms",
			input:    []string{"banco", "nacional", "cuba"},
			expected: []string{"banco", "nacional", "cuba"},
		},
		{
			name:     "with noise terms",
			input:    []string{"the", "banco", "of", "nacional", "and", "cuba"},
			expected: []string{"banco", "nacional", "cuba"},
		},
		{
			name:     "with short terms",
			input:    []string{"al", "banco", "de", "nacional"},
			expected: []string{"banco", "nacional"},
		},
		{
			name:     "only noise terms",
			input:    []string{"the", "of", "and", "in", "at"},
			expected: []string{},
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "mixed case terms",
			input:    []string{"THE", "Banco", "OF", "Nacional"},
			expected: []string{"banco", "nacional"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterSignificantTerms(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateCombinedScore(t *testing.T) {
	tests := []struct {
		name          string
		nameScore     float64
		typeScore     float64
		expectedScore float64
	}{
		{
			name:          "perfect match",
			nameScore:     1.0,
			typeScore:     1.0,
			expectedScore: 1.0,
		},
		{
			name:          "high name score with exact type",
			nameScore:     0.9,
			typeScore:     1.0,
			expectedScore: 1.0, // With exactTypeBonus but capped at 1.0
		},
		{
			name:          "high name score with related type",
			nameScore:     0.9,
			typeScore:     0.8,
			expectedScore: 0.98, // With relatedTypeBonues
		},
		{
			name:          "high name score with mismatched type",
			nameScore:     0.9,
			typeScore:     0.0,
			expectedScore: 0.75, // With typeMatchPenalty
		},
		{
			name:          "low scores",
			nameScore:     0.3,
			typeScore:     0.0,
			expectedScore: 0.15, // With typeMatchPenalty
		},
		{
			name:          "zero name score",
			nameScore:     0.0,
			typeScore:     1.0,
			expectedScore: 0.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := calculateCombinedScore(tt.nameScore, tt.typeScore)
			assert.InDelta(t, tt.expectedScore, score, 0.1)
		})
	}
}

func TestCalculateFinalAffiliateScore(t *testing.T) {
	tests := []struct {
		name          string
		matches       []affiliationMatch
		expectedScore float64
	}{
		{
			name: "single perfect match",
			matches: []affiliationMatch{
				{nameScore: 1.0, typeScore: 1.0, finalScore: 1.0, exactMatch: true},
			},
			expectedScore: 1.0,
		},
		{
			name: "multiple high quality matches",
			matches: []affiliationMatch{
				{nameScore: 0.95, typeScore: 1.0, finalScore: 0.95, exactMatch: false},
				{nameScore: 0.90, typeScore: 0.8, finalScore: 0.90, exactMatch: false},
			},
			expectedScore: 0.93,
		},
		{
			name: "mixed quality matches",
			matches: []affiliationMatch{
				{nameScore: 0.95, typeScore: 1.0, finalScore: 0.95, exactMatch: false},
				{nameScore: 0.50, typeScore: 0.0, finalScore: 0.35, exactMatch: false},
			},
			expectedScore: 0.85,
		},
		{
			name:          "no matches",
			matches:       []affiliationMatch{},
			expectedScore: 0.0,
		},
		{
			name: "all low quality matches",
			matches: []affiliationMatch{
				{nameScore: 0.3, typeScore: 0.0, finalScore: 0.15, exactMatch: false},
				{nameScore: 0.4, typeScore: 0.0, finalScore: 0.25, exactMatch: false},
			},
			expectedScore: 0.21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := calculateFinalAffiliateScore(tt.matches)
			assert.InDelta(t, tt.expectedScore, score, 0.1)
		})
	}
}
