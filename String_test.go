package describe

import (
	"testing"
)

func TestString(t *testing.T) {
	// Update
	if String("qwe").SetSliceIndex(2).Update("ggg").SliceIndex != 2 {
		t.Error("Update should return same pointer")
	}
}
