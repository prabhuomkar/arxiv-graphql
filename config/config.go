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
		Environment string  `json:"environment"`
		Service     Service `json:"service"`
		ArxivAPIURL string  `json:"arXivURL"`
		Fields      []Field `json:"fields"`
	}

	// Service ...
	Service struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

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

// Load ...
func Load(cfgFilePath string) error {
	data, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Config)
	return err
}
