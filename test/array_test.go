package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray1(t *testing.T) {
	arr := [...]string{"a", "b", "c", "d", "e"}
	
	t.Log(arr)
	assert.Equal(t, len(arr), 5)
	assert.Equal(t, arr, [...]string{"a", "b", "c", "d", "e"})
}
