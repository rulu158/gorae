package gorae

type WordDefinition struct {
	Word    string                 `json:"word"`
	Entries []*WordDefinitionEntry `json:"entries"`
}

type WordDefinitionEntry struct {
	Num        int    `json:"num"`
	Types      string `json:"type"`
	Definition string `json:"definition"`
	Examples   string `json:"examples"`
}
