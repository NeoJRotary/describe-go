package describe

import (
	"testing"
)

func TestStringUpdate(t *testing.T) {
	// Update
	if String("qwe").SetSliceIndex(2).Update("ggg").SliceIndex != 2 {
		t.Error("Update should return same pointer")
	}
}

func TestStringRange(t *testing.T) {
	s := String("qwwertyuiooooo")
	if result := s.Range(1, 4).Get(); result != "wwe" {
		t.Error("get", result, "should be", "wwe")
	}

	if result := s.RangeBetween("ww", "i").Get(); result != "ertyu" {
		t.Error("get", result, "should be", "ertyu")
	}
}
