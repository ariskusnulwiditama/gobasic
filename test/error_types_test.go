package test

import "testing"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func TestErrorTypes(t *testing.T) {
	err := &validationError{"validation error"}
	t.Log(err)
}