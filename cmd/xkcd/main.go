package main

import (
	"os"
	"flag"
	"fmt"
)

func main() {
	var indexLocation = flag.String("index", os.Getenv("HOME") + "/.xkcd_index.json", "the index file")
	flag.Parse()

	fmt.Println(*indexLocation)
}
