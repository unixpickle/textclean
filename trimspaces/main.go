// Command trimspaces goes through every file it is passed
// and removes whitespace from the beginning and end of
// the file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/unixpickle/essentials"
)

func main() {
	if len(os.Args) == 1 {
		essentials.Die("no input files")
	}
	for _, x := range os.Args[1:] {
		fmt.Println(x)
		contents, err := ioutil.ReadFile(x)
		if err != nil {
			essentials.Die(err)
		}
		trimmed := []byte(strings.TrimSpace(string(contents)))
		if err := ioutil.WriteFile(x, trimmed, 0755); err != nil {
			essentials.Die(err)
		}
	}
}
