package main

import "fmt"

type Card struct {
	val   int
	descr rune // unicode character type (one or more bytes)
	suit  string
}

func (c Card) String() string {
	return fmt.Sprintf("%c", c.descr)
}
