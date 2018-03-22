package describe

import (
	"testing"
)

func TestString(t *testing.T) {
	// Update
	if String("qwe").IndexAt(2).Update("ggg").Index != 2 {
		t.Error("Update should return same pointer")
	}
}
