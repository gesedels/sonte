package test

import "testing"

func TestAssertDire(t *testing.T) {
	// setup
	dire := MockDire(t)

	// success
	AssertDire(t, dire, MockNotes)
}

func TestAssertFile(t *testing.T) {
	// setup
	orig := MockFile(t, "alpha.extn")

	// success
	AssertFile(t, orig, MockNotes["alpha.extn"])
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	AssertDire(t, dire, MockNotes)
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	AssertFile(t, orig, MockNotes["alpha.extn"])
}
