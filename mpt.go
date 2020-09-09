package main

type MptNode struct {
	slots [66]interface{}
	index int
}

var mptNode = MptNode{index: -1}

func mptInsert(index int, s string) {
	node := &mptNode
	for i, v := range s {
		if node.slots[v-65] == nil {
			if node.index != -1 && i < len(samples[node.index]) {
				if samples[node.index][i]-65 == byte(v)-65 {
					node.slots[v-65] = &MptNode{index: node.index}
				} else {
					node.slots[v-65] = &MptNode{index: -1}
					node.slots[samples[node.index][i]-65] = &MptNode{index: node.index}
				}
			} else {
				node.slots[v-65] = &MptNode{index: index}
				return
			}
		} else {
			node = node.slots[v-65].(*MptNode)
		}
	}
	node.index = index
}
