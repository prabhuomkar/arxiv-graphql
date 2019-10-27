package model

import "time"

type (
	// FeedEntry : Model for arXiv entry of the feed
	FeedEntry struct {
		ID              string     `json:"id"`
		Updated         time.Time  `json:"updated"`
		Published       time.Time  `json:"published"`
		Title           string     `json:"title"`
		Summary         string     `json:"summary"`
		Author          []Author   `json:"author"`
		Link            []Link     `json:"link"`
		PrimaryCategory Category   `json:"primary_category"`
		Category        []Category `json:"category"`
	}

	// Response ...
	Response struct {
		Feed Feed `json:"feed"`
	}

	// Feed ...
	Feed struct {
		Entry []FeedEntry `json:"entry"`
	}

	// Author ...
	Author struct {
		Name string `json:"name"`
	}

	// Link ...
	Link struct {
		Href  string `json:"-href"`
		Rel   string `json:"-rel"`
		Title string `json:"-title"`
	}

	// Category ...
	Category struct {
		Term   string `json:"-term"`
		Scheme string `json:"-scheme"`
	}
)
