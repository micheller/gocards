package main

import (
	"fmt"
	"sort"
)

func main() {
	np := NewPack(52, nil)
	fmt.Println(np)
	fmt.Println("pack of ", np.num, " cards")
	fmt.Print("\nShuffling...")
	np.Shuffle()
	fmt.Println(np)
	hands := np.Handout()
	for i, h := range hands {
		fmt.Print("Hand ", i, ": \n")
		sort.Sort(h)
		fmt.Println(h)
		fmt.Println()
	}
}
