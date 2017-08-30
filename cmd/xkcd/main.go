package main

import (
	"flag"
	"os"
	"fmt"

	"github.com/bakku/xkcd"
)

func init() {
	// overwrite usage output
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("\txkcd [flags] <command>\n")
		fmt.Println("\tCommands:")
		fmt.Println("\t\trefresh\trefresh the index file (might take a while)\n")
		fmt.Println("\tFlags:")
		fmt.Println("\t\t-index\tpath to the index file or where the index should be saved (Default: $HOME/.xkcd_index.json)")
	}
}

func main() {
	var indexLocation = flag.String("index", os.Getenv("HOME")+"/.xkcd_index.json", "the index file")
	flag.Parse()

	index := xkcd.NewIndex(indexLocation)
	index.Refresh()
}
