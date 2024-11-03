package trie_test

import (
	"testing"

	"github.com/vdinovi/advent_of_code/lib/golang/lib/trie"
)

type integer struct {
	val int
}

func TestTrie(t *testing.T) {
	tr := trie.NewTrie[integer, rune]()
	data := map[string]int{
		"abc":  1,
		"ab":   2,
		"bc":   3,
		"bcde": 4,
	}
	for k, v := range data {
		tr.Add([]rune(k), &integer{v})
	}
	entries := tr.Entries()
	if len(entries) != 4 {
		t.Fatalf("expected trie to have %d entries but had %d", 4, len(entries))
	}
	for key, val := range data {
		v := tr.Get([]rune(key), nil)
		if v == nil {
			t.Fatalf("expected trie to have value %d for key %s but did not", val, key)
		}
		if v.val != val {
			t.Fatalf("expected trie to have value %d for key %s but had %d", val, key, v.val)
		}
	}
	for _, entry := range entries {
		if entry.Val.val != data[string(entry.Keys)] {
			t.Fatalf("expected entry %v to have value %d but had %d", entry.Keys, data[string(entry.Keys)], entry.Val.val)
		}
	}
}
