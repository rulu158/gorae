package gorae

type WordDefinition struct {
	Word    string                 `json:"word"`
	Entries []*WordDefinitionEntry `json:"entries"`
}

type WordDefinitionEntry struct {
	Num        int    `json:"num"`
	Types      string `json:"type"`
	Definition string `json:"definition"`
	Synonyms   string `json:"synonyms"`
	Antonyms   string `json:"antonyms"`
	Examples   string `json:"examples"`
}
