package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	baseUri     string = "https://xkcd.com"
	jsonPath    string = "info.0.json"
	archivedUri string = baseUri + "/" + jsonPath
)

func GetComicById(id int) (comic *Comic, ok bool) {
	resp, err := http.Get(baseUri + "/" + strconv.Itoa(id) + "/" + jsonPath)
	if err != nil {
		fmt.Println("Could not fetch comic with ID:", id)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// No output needed
		return nil, false
	}

	var c Comic
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		fmt.Println("Could not decode response body of comic with ID:", id)
		return nil, false
	}

	return &c, true
}

func GetLatestComic() (comic *Comic, ok bool) {
	resp, err := http.Get(archivedUri)
	if err != nil {
		fmt.Println("Could not fetch latest comic")
		return nil, false
	}
	defer resp.Body.Close()

	var c Comic
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		fmt.Println("Could not decode latest comic response")
		return nil, false
	}

	return &c, true
}
