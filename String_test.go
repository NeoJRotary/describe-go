package describe

import (
	"strings"
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

func BenchmarkString_NativeReplaceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Replace("asdrdfewf333gobsddlfdfewf333nefokndedodfewf333knd", "dfewf333", "", -1)
	}
}

func BenchmarkString_ReplaceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String("asdrdfewf333gobsddlfdfewf333nefokndedodfewf333knd").ReplaceAll("dfewf333", "")
	}
}

func BenchmarkString_ReplaceAllWithCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String("asdrdfewf333gobsddlfdfewf333nefokndedodfewf333knd").Copy().ReplaceAll("dfewf333", "")
	}
}
