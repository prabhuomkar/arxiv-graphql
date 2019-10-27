package resolver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prabhuomkar/arXiv/config"
)

// BuildAPIURL : Builds API URL for ArXiV based on parameters
func BuildAPIURL(category, sortBy, sortOrder string, start, maxResults int) string {
	apiURL := config.Config.ArxivAPIURL + APIPath
	if category != "" {
		apiURL += "cat:" + category
	}
	if sortBy != "" {
		apiURL += "&sortBy=" + sortBy
	}
	if sortOrder != "" {
		apiURL += "&sortOrder=" + sortOrder
	} else {
		apiURL += "&sortOrder=" + SortOrderDescending
	}
	apiURL += "&start=" + fmt.Sprintf("%d", start)
	if maxResults != 0 {
		apiURL += "&max_results=" + fmt.Sprintf("%d", maxResults)
	} else {
		apiURL += "&max_results=" + fmt.Sprintf("%d", resultsLimit)
	}
	return apiURL
}

// FetchResponse : Gets response for HTTP GET request made on API URL
func FetchResponse(apiURL string) ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
