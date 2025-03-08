package test

import "testing"

type Sound interface {
	Sound()
}

type Pet interface {
	Sound
	Play()
}

type Dog struct{}

type Animal struct{}

func (d Dog) Play() {
	println("Dog is playing")
}

func (d Dog) Sound() {
	println("Woof!")
}

// interface as parameter
func MakeSound(s Sound) {
	s.Sound()
}

func MakeSoundPet(p Pet) {
	p.Sound()
	p.Play()
}

func TestInterface(t *testing.T) {
	d := Dog{}
	d.Sound()
}

func TestInterfaceAsParameter(t *testing.T) {
	d := Dog{}
	MakeSound(d)
}
