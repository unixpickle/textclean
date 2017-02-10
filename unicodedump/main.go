// Command unicodedump prints a sorted list of the
// non-ASCII characters in a text file.
// The list is sorted by frequency.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/unixpickle/essentials"
)

func main() {
	if len(os.Args) != 2 {
		essentials.Die("Usage: unicodedump <file.txt>")
	}
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		essentials.Die(err)
	}
	counts := map[rune]int{}
	for _, r := range []rune(string(contents)) {
		if r > 0x7f {
			counts[r]++
		}
	}
	s := sorter{}
	for ch, count := range counts {
		s.chars = append(s.chars, ch)
		s.freqs = append(s.freqs, count)
	}
	sort.Sort(&s)

	for i, ch := range s.chars {
		fmt.Println(string(ch), "-", s.freqs[i])
	}
}

type sorter struct {
	chars []rune
	freqs []int
}

func (s *sorter) Len() int {
	return len(s.chars)
}

func (s *sorter) Swap(i, j int) {
	s.chars[i], s.chars[j] = s.chars[j], s.chars[i]
	s.freqs[i], s.freqs[j] = s.freqs[j], s.freqs[i]
}

func (s *sorter) Less(i, j int) bool {
	return s.freqs[i] < s.freqs[j]
}
