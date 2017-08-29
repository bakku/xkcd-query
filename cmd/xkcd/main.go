package main

import (
	"flag"
	"os"

	"github.com/bakku/xkcd"
)

func main() {
	var indexLocation = flag.String("index", os.Getenv("HOME")+"/.xkcd_index.json", "the index file")
	flag.Parse()

	index := xkcd.NewIndex(indexLocation)
	index.Populate()
}
