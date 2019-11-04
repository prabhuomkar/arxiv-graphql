package model

type (
	// Field ...
	Field struct {
		Name       string     `json:"name"`
		Tag        string     `json:"tag"`
		Categories []Category `json:"categories"`
	}

	// Category ...
	Category struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	}
)
