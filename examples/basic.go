// +build ignore

package main

import (
	"compress/gzip"
	"context"
	"os"

	"github.com/wuhan005/obelisk"
)

func main() {
	// Create archive
	req := obelisk.Request{
		URL: "https://www.globalwitness.org/en/blog/how-the-rsf-got-their-4x4-technicals-the-open-source-intelligence-techniques-behind-our-sudan-expos%C3%A9",
	}

	arc := obelisk.Archiver{EnableLog: true}
	arc.Validate()

	result, _, err := arc.Archive(context.Background(), req)
	checkError(err)

	// Create destination file
	f, err := os.Create("globalwitness.html.gz")
	checkError(err)
	defer f.Close()

	// Create gzipper
	gz := gzip.NewWriter(f)
	gz.Write(result)
	gz.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
