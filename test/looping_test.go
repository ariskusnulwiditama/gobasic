package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLooping1(t *testing.T) {
	

	for i := 0 ;i < 10; i++ {
		if i % 2 == 0 {
			continue
		}

		fmt.Println(i)
	}
	
}

func TestLooping2(t *testing.T) {
	i := 0

	for {
		
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}
}

func TestLooping3(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	assert.Equal(t, 5, len(arr))
}

func TestVarArgs(t *testing.T) {
	res := varArgs(1, 2, 3, 4, 5)
	assert.Equal(t, 15, res)
}

func varArgs(num ...int) int{
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}
