package main

import "testing"

//Go’s naming convention for test files is that they end in _test.go. This suffix tells Go that this is a file to be run when tests execute, and excluded when the application is built, as shown in the next listing.
func TestName(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Respone from getName is unexpected value")
	}
}
