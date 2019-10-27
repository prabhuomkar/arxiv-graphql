package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
var Config Configuration

type (
	// Configuration ...
	Configuration struct {
		Environment string     `json:"environment"`
		Service     Service    `json:"service"`
		ArxivAPIURL string     `json:"arXivURL"`
		Categories  []Category `json:"categories"`
	}

	// Service ...
	Service struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	// Category ...
	Category struct {
		Name          string        `json:"name"`
		Tag           string        `json:"tag"`
		SubCategories []SubCategory `json:"subcategories"`
	}

	// SubCategory ...
	SubCategory struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	}
)

// Load ...
func Load(cfgFilePath string) error {
	data, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Config)
	return err
}
