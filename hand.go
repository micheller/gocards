package main

import (
	"fmt"
)

const (
	SPADES   = "\u2660"
	CLUBS    = "\u2663"
	DIAMONDS = "\u2666"
	HEARTS   = "\u2665"
)

type Hand []Card

func (h Hand) String() string {
	ret := ""
	oldsuit := h[0].suit
	for _, c := range h {
		if c.suit != oldsuit {
		        ret += "\n"
			oldsuit = c.suit
		}
		ret += fmt.Sprintf("%v ", c)
	}
	return ret
}

func suitless(a, b string) bool {
	mapping := make(map[string]int)
	mapping[SPADES] = 0
	mapping[CLUBS] = 1
	mapping[DIAMONDS] = 2
	mapping[HEARTS] = 3
	return mapping[a] < mapping[b]
}

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	if h[i].suit != h[j].suit {
		return suitless(h[i].suit, h[j].suit)
	} else {
		return h[i].val < h[j].val
	}
}
