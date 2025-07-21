package test

import "testing"

func TestAssertDire(t *testing.T) {
	// setup
	dire := MockDire(t)

	// success
	AssertDire(t, dire, MockData)
}

func TestAssertFile(t *testing.T) {
	// setup
	orig := MockFile(t, "alpha.extn")

	// success
	AssertFile(t, orig, MockData["alpha.extn"])
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	AssertDire(t, dire, MockData)
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	AssertFile(t, orig, MockData["alpha.extn"])
}
