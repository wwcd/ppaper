package main

type TrieNode struct {
	slots [128]interface{}
	index int
}

var trieRoot TrieNode

func trieInsert(index int, s string) {
	node := &trieRoot
	for _, v := range s {
		if node.slots[v] == nil {
			node.slots[v] = new(TrieNode)
		}
		node = node.slots[v].(*TrieNode)
	}

	node.index = index
}
