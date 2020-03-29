package firstPackage

import "testing"

func TestWork(t *testing.T) {
	r := Work()
	if r != "My Package Works" {
		t.Error("Expected: My Package Works, got ", r)
	}
}