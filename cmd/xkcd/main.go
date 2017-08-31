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
		fmt.Println("\tFlags:")
		fmt.Println("\t\t-index\t\tpath to the index file or where the index should be saved (Default: $HOME/.xkcd_index.json)")
		fmt.Println("\tCommands:")
		fmt.Println("\t\trefresh\t\trefresh the index file (might take a while)\n")
	}
}

func main() {
	var indexLocation = flag.String("index", os.Getenv("HOME")+"/.xkcd_index.json", "the index file")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	index := xkcd.NewIndex(indexLocation)

	switch flag.Arg(0) {
		case "refresh":
			index.Refresh()
		default:
			flag.Usage()
	}
}
