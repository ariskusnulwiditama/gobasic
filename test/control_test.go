package test

import (
	"math/rand"
	"testing"
	"time"
)

func TestIfElse(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	if n := rand.Int(); n%2 == 0 {
		t.Log(n, "is even")
	} else {
		t.Log(n, "is odd")
	}

	n := rand.Int() % 2
	if n == 0 {
		t.Log("An even number")
	}

	if ; n % 2 != 0 {
		t.Log("An odd number")
	}
	
}