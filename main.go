package main

import "time"

func main() {
	for i, v := range samples {
		// trieInsert(i, v)
		// radixTrieInsert(i, v)
		mptInsert(i, v)
	}

	for {
		time.Sleep(time.Second)
	}
}
