package test

import (
	"fmt"
	"testing"
)

// Reader is an interface that wraps the basic Read method.


type deck []string
type Reader struct {
	s string
}

func showDeck() deck{
	return deck{"Hearts", "Diamonds", "Clubs", "Spades"}
}

func TestDeck(t *testing.T) {
	res := showDeck()
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func NewReader(s string) *Reader {
	return &Reader{s}
}

func (r *Reader) Len() int {
	return len(r.s)
}