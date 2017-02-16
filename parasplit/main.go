// Command parasplit splits a text document into separate
// files, one per "paragraph".
// A paragraph is a chunk of lines separated by two
// newlines, such as:
//
//     This is paragraph 1.
//     Still paragraph 1.
//
//     This is paragraph 2.
//
// Once paragraphs have been processed, further newlines
// within the paragraphs are replaced with spaces.
//
// Very short paragraphs are ignored, since they are less
// likely to be paragraphs and more likely to be titles or
// dividers of some kind.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/unixpickle/essentials"
)

const DefaultMinParagraphLength = 80

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		essentials.Die("Usage: parasplit <file.txt> <out_dir> [min_len]")
	}
	inFile := os.Args[1]
	outDir := os.Args[2]
	minLen := DefaultMinParagraphLength
	if len(os.Args) == 4 {
		var parseErr error
		minLen, parseErr = strconv.Atoi(os.Args[3])
		if parseErr != nil {
			essentials.Die("invalid min_len argument:", os.Args[3])
		}
	}

	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		essentials.Die(err)
	}

	split := strings.Split(strings.Replace(string(contents), "\r", "", -1), "\n\n")
	parIdx := 0
	for _, p := range split {
		if len(p) < minLen {
			continue
		}
		p = strings.Replace(p, "\n", " ", -1)
		parIdx++
		outPath := filepath.Join(outDir, fmt.Sprintf("paragraph_%d.txt", parIdx))
		if err := ioutil.WriteFile(outPath, []byte(p), 0755); err != nil {
			fmt.Println()
			essentials.Die("save paragraph:", err)
		}
		fmt.Printf("\rProcessed %d paragraphs.", parIdx)
	}
	fmt.Println()
}
