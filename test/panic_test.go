package test

import (
	"fmt"
	"testing"
)

func endError() {
	errMsg := recover()
	fmt.Print("error occured", errMsg)
}

func resultError(error bool) {

	defer endError()
	if error {
		panic("error occured")
	}
}

func TestPanic(t *testing.T) {
	resultError(true)
	t.Log("TestPanic completed")
}