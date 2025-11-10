package test

import "testing"

func double(x *int) {
	*x += *x
	// x = nil
}

func TestPointer(t *testing.T) {
	a := 3
	double(&a)
	t.Log(a)
	t.Log(&a)
	p := &a
	t.Log(p)
	t.Log(*p)
}