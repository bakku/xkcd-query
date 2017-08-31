# Query xkcd.com offline

This small tool was inspired by an exercise from the book "The Go Programming language".

Using this tool you can always find the right xkcd comic for every situation 🙂

## Building

Plain and simple

    ./build.sh

## Usage

First of all you have to create an index using the `refresh` command. The tool expects an optional flag `-index` where you can specify the location of a JSON file where the image will be stored. (by default: `$HOME/.xkcd_index.json`)

Example:

    ./xkcd refresh                                # default location
    ./xkcd -index /home/bakku/.index.json refresh # location is being specified

It might take a while for the image to be created. I might improve the download speed by adding concurrency in the future but for now each refresh will issue approx. 2000 requests and process them sequentially.

After a successful refresh you can begin to query your index.

    ./xkcd query github
    0) Number: 624    Link: https://imgs.xkcd.com/comics/branding.png
    1) Number: 1656   Link: https://imgs.xkcd.com/comics/it_begins.png

Or (depending on your index location)

    ./xkcd -index /home/bakku/.index.json query github
    0) Number: 624    Link: https://imgs.xkcd.com/comics/branding.png
    1) Number: 1656   Link: https://imgs.xkcd.com/comics/it_begins.png
