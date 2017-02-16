// Command htmlparas reads the paragraphs from a set of
// HTML documents and outputs them as lines to standard
// output.
//
// This can be used to, for instance, extract paragraphs
// from ebooks.
package main

import (
	"fmt"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	if len(os.Args) == 1 {
		essentials.Die("no input files")
	}

	for _, in := range os.Args[1:] {
		f, err := os.Open(in)
		if err != nil {
			essentials.Die(err)
		}
		parsed, err := html.Parse(f)
		f.Close()
		if err != nil {
			essentials.Die("parse "+in+":", err)
		}
		paras := scrape.FindAll(parsed, scrape.ByTag(atom.P))
		for _, p := range paras {
			fmt.Println(scrape.Text(p))
			fmt.Println()
		}
	}
}
