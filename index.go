package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Index struct {
	Filepath *string
	comics   []*Comic
}

func NewIndex(filepath *string) (index *Index) {
	return &Index{
		Filepath: filepath,
	}
}

func (index *Index) Refresh() {
	fmt.Print("Refreshing index. This could take a while...")

	comic, ok := GetLatestComic()
	if ok != true {
		return
	}

	var comics []*Comic
	min, max := 1, comic.Num
	for i := min; i <= max; i++ {
		if comic, ok := GetComicById(i); ok {
			comics = append(comics, comic)
		}
	}

	bytes, err := json.Marshal(comics)
	if err != nil {
		fmt.Println("Could not marshal comics to json")
		return
	}

	err = ioutil.WriteFile(*index.Filepath, bytes, 0644)
	if err != nil {
		fmt.Println("Could not write to file at:", *index.Filepath)
		return
	}

	fmt.Println("Done")
}

func (index *Index) Query(queries []string) (comics []*Comic, ok bool) {
	if ok := index.loadIndex(); !ok {
		return nil, false
	}

	for _, comic := range index.comics {
		for _, query := range queries {
			if strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(query)) {
				comics = append(comics, comic)
			}
		}
	}

	ok = true
	return
}

func (index *Index) loadIndex() (ok bool) {
	rawIndex, err := ioutil.ReadFile(*index.Filepath)
	if err != nil {
		fmt.Println("Could not load index. Have you created one? Given path:", *index.Filepath)
		return false
	}

	err = json.Unmarshal(rawIndex, &index.comics)
	if err != nil {
		fmt.Println("Could not deserialize index json file")
		return false
	}

	return true
}
