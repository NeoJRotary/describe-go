package describe

import "testing"

func TestFileExist(t *testing.T) {
	if !FileExist("./File_test.go") {
		t.Error("should exist")
	}

	if FileExist("./fileeee.go") {
		t.Error("should not exist")
	}
}
