package gostring

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	name := "John Doe"

	fmt.Println(string(name[len(name)-3:]))
}

func TestString2(t *testing.T) {
	defer func() {
		fmt.Println("Finish")
	}()
	name := "seffy diah kusumaningrum"
	word := "seffy"

	nm := strings.Fields(name)

	for _, v := range nm {
		if v == word {
			fmt.Println("Found")
			return
		} else {
			fmt.Println("Not Found")
		}
	}
}
