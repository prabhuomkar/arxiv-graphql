package model

import "time"

type (
	// Article : Model for arXiv article of the feed
	Article struct {
		ID              string         `json:"id"`
		Updated         time.Time      `json:"updated"`
		Published       time.Time      `json:"published"`
		Title           string         `json:"title"`
		Summary         string         `json:"summary"`
		Author          []Author       `json:"author"`
		Link            []Link         `json:"link"`
		PrimaryCategory FeedCategory   `json:"primary_category"`
		Category        []FeedCategory `json:"category"`
	}

	// Response ...
	Response struct {
		Feed Feed `json:"feed"`
	}

	// SingleResponse ...
	SingleResponse struct {
		SingleFeed SingleFeed `json:"feed"`
	}

	// Feed ...
	Feed struct {
		Entry []Article `json:"entry"`
	}

	// SingleFeed ...
	SingleFeed struct {
		Entry Article `json:"entry"`
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

	// FeedCategory ...
	FeedCategory struct {
		Term   string `json:"-term"`
		Scheme string `json:"-scheme"`
	}
)
