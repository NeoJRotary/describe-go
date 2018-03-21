package describe

import "testing"

func TestFileExist(t *testing.T) {
	if !FileExist("./file.go") {
		t.Error("should exist")
	}

	if FileExist("./fileeee.go") {
		t.Error("should not exist")
	}
}
