package trie

type Trie[T any, K comparable] struct {
	root trieNode[T, K]
}

func NewTrie[T any, K comparable]() *Trie[T, K] {
	return &Trie[T, K]{
		root: trieNode[T, K]{
			val:      nil,
			children: make(map[K]*trieNode[T, K]),
		},
	}
}

func (t *Trie[T, K]) Get(key []K, val *T) *T {
	var ok bool
	iter := &t.root
	for _, k := range key {
		if iter, ok = iter.children[k]; !ok {
			return nil
		}
	}
	return iter.val
}

func (t *Trie[T, K]) Add(key []K, val *T) {
	iter := &t.root
	for _, k := range key[:len(key)-1] {
		child, ok := iter.children[k]
		if ok {
			iter = child
			continue
		}
		iter = iter.addChild(k, nil)
	}
	last := key[len(key)-1]
	if child, ok := iter.children[last]; ok {
		child.val = val
	} else {
		iter.addChild(last, val)
	}
}

type TrieEntry[T any, K comparable] struct {
	Keys []K
	Val  T
}

func (t *Trie[T, K]) Entries() []TrieEntry[T, K] {
	entries := make([]TrieEntry[T, K], 0)
	t.root.entries([]K{}, &entries)
	return entries
}

type trieNode[T any, K comparable] struct {
	val      *T
	children map[K]*trieNode[T, K]
}

func (t *trieNode[T, K]) addChild(key K, val *T) *trieNode[T, K] {
	t.children[key] = &trieNode[T, K]{
		val:      val,
		children: make(map[K]*trieNode[T, K]),
	}
	return t.children[key]
}

func (t *trieNode[T, K]) entries(keys []K, entries *[]TrieEntry[T, K]) {
	if t.val != nil {
		*entries = append(*entries, TrieEntry[T, K]{Keys: keys, Val: *t.val})
	}
	for key, child := range t.children {
		child.entries(append(keys, key), entries)
	}
}
