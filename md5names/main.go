// Command md5names renames text files to the MD5 hashes
// of their contents.
// This makes it easy to remove duplicate files, and to
// split text files in a random but deterministic way.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/unixpickle/essentials"
)

func main() {
	if len(os.Args) == 1 {
		essentials.Die("no input files")
	}
	for _, x := range os.Args[1:] {
		contents, err := ioutil.ReadFile(x)
		if err != nil {
			essentials.Die(err)
		}
		hash := md5.Sum(contents)
		newBasename := hex.EncodeToString(hash[:]) + filepath.Ext(x)
		name := filepath.Join(filepath.Dir(x), newBasename)
		fmt.Println(x, "->", newBasename)
		if err := os.Rename(x, name); err != nil {
			essentials.Die(err)
		}
	}
}
