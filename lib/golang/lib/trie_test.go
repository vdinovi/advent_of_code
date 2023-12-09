package lib_test

import (
	"testing"

	"github.com/vdinovi/advent_of_code/lib/golang/lib"
)

type integer struct {
	val int
}

func TestTrie(t *testing.T) {
	trie := lib.NewTrie[integer, rune]()
	data := map[string]int{
		"abc":  1,
		"ab":   2,
		"bc":   2,
		"bcde": 2,
	}
	for k, v := range data {
		trie.Add([]rune(k), &integer{v})
	}
	entries := trie.Entries()
	if len(entries) != 4 {
		t.Fatalf("expected trie to have %d entries but had %d", 4, len(entries))
	}
	for _, entry := range entries {
		if entry.Val.val != data[string(entry.Keys)] {
			t.Errorf("expected entry %v to have value %d but had %d", entry.Keys, data[string(entry.Keys)], entry.Val.val)
		}
	}
}
