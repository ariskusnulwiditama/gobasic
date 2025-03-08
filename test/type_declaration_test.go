package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type NoKTP string

func TestType(t *testing.T) {
	type NoKTP string
	
	var noKTP NoKTP = "1234567890"
	noSim := NoKTP("45454")
	log.Println(noKTP)
	log.Println(noSim)
}

func TestOperator(t *testing.T) {
	ff := 90
	ff++
	log.Println(ff)
	assert.Equal(t, ff, 91)
}

func TestComparation(t *testing.T) {
	word1 := "Hello"
	word2 := "Hello"

	if len(word1) > len(word2) {
		log.Println(word1)
	} else {
		log.Println(word2)
	}
}