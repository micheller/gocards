package main

import (
	"fmt"
	"math/rand"
	"time"
	// "strconv"
)


type Pack struct {
	num   int
	cards []Card
	order Suite_order
}

func (p Pack) String() string {
	ret := ""
	for i, c := range p.cards {
		if i%(p.num/4) == 0 {
			ret += "\n"
		}
		ret += fmt.Sprintf("%v ", c)
	}
	return ret
}

func NewPack(num int, order Suite_order) *Pack {
	if order == nil {
		order = Suite_order{
			{0x1F0A1, SPADES},
			{0x1F0A1+48, CLUBS},
			{0x1F0A1+32, DIAMONDS},
			{0x1F0A1+16, HEARTS}}
	}
	p := new(Pack)
	p.num = num
	p.order = order
	for _, s := range p.order {
		for i := 14 - (num / 4) + 1; i <= 14; i++ {
			desc :=  s.unicode_base + i - 1
			switch i {
			case 14:
				desc -= i-1
			case 12, 13:
				desc++
			}
			p.cards = append(p.cards, Card{i, rune(desc), s.repr})
		}
	}
	return p
}

func (p Pack) Shuffle() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < p.num; i++ {
		a := rand.Intn(p.num)
		b := rand.Intn(p.num)
		p.cards[a], p.cards[b] = p.cards[b], p.cards[a]
	}
}

func (p Pack) Handout() []Hand {
	hands := make([]Hand, 4)
	for i, c := range p.cards {
		hnum := i % 4
		hands[hnum] = append(hands[hnum], c)
	}
	return hands
}
