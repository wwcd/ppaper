package main

type RadixTrieNode struct {
	slots [128]interface{}
	index int
}

var radixTrieNode = RadixTrieNode{index: -1}

func radixTrieInsert(index int, s string) {
	node := &radixTrieNode
	for i, v := range s {
		if node.slots[v] == nil {
			if node.index != -1 && i < len(samples[node.index]) {
				if samples[node.index][i] == byte(v) {
					node.slots[v] = &RadixTrieNode{index: node.index}
				} else {
					node.slots[v] = &RadixTrieNode{index: -1}
					node.slots[samples[node.index][i]] = &RadixTrieNode{index: node.index}
				}
			} else {
				node.slots[v] = &RadixTrieNode{index: index}
				return
			}
		} else {
			node = node.slots[v].(*RadixTrieNode)
		}
	}
	node.index = index
}
