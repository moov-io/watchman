package us_tel

type TELRecord struct {
	ID         string        `json:"id"`
	Caption    string        `json:"caption"`
	Schema     string        `json:"schema"`
	Referents  []string      `json:"referents"`
	Datasets   []string      `json:"datasets"`
	Origin     []string      `json:"origin"`
	FirstSeen  string        `json:"first_seen"`
	LastSeen   string        `json:"last_seen"`
	LastChange string        `json:"last_change"`
	Properties TELProperties `json:"properties"`
	Target     bool          `json:"target"`
}