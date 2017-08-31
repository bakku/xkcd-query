package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bakku/xkcd"
)

func init() {
	// overwrite usage output
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("\txkcd [flags] <command> [optional arguments]\n")
		fmt.Println("\tFlags:")
		fmt.Println("\t\t-index\t\tpath to the index file or where the index should be saved (Default: $HOME/.xkcd_index.json)")
		fmt.Println("\tCommands:")
		fmt.Println("\t\trefresh\t\trefresh the index file (might take a while)\n")
		fmt.Println("\t\tquery\t\tqueries the index file (expects the query string as arguments after the command)\n")
	}
}

func main() {
	var indexLocation = flag.String("index", os.Getenv("HOME")+"/.xkcd_index.json", "the index file")
	flag.Parse()

	if flag.NArg() < 1 || (flag.Arg(0) == "query" && flag.NArg() == 1) {
		flag.Usage()
		return
	}

	index := xkcd.NewIndex(indexLocation)

	switch flag.Arg(0) {
	case "refresh":
		index.Refresh()
	case "query":
		comics, ok := index.Query(flag.Args()[1:])
		if !ok {
			return
		}

		printComics(comics)
	default:
		flag.Usage()
	}
}

func printComics(comics []*xkcd.Comic) {
	if len(comics) == 0 {
		fmt.Println("No results.")
		return
	}

	for i, comic := range comics {
		fmt.Printf("%d)\tNumber: %d\tLink: %s\n", i, comic.Num, comic.Img)
	}
}
