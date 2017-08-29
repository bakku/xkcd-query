package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Index struct {
	Filepath *string
}

func NewIndex(filepath *string) (index *Index) {
	return &Index{
		Filepath: filepath,
	}
}

func (index *Index) Populate() {
	fmt.Print("Populating index...")

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
