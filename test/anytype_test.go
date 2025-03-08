package test

import "testing"

func TestAny(t *testing.T) {
	fruits := []any{
		map[string]any{"name": "apple", "color": "red"},
		[]string{"apple", "banana", "cherry"},
		"kiwi",
	}

	for _, each := range fruits {
		t.Log(each)
	}
}