package describe

import (
	"strings"
	"testing"
)

func TestString_Update(t *testing.T) {
	// Update
	if String("qwe").SetSliceIndex(2).Update("ggg").SliceIndex != 2 {
		t.Error("Update should return same pointer")
	}
}

func TestString_Range(t *testing.T) {
	s := String("qwwertyuiooooo")
	if result := s.Copy().Range(1, 4).Get(); result != "wwe" {
		t.Error("get", result, "should be", "wwe")
	}

	if result := s.Copy().RangeBetween("ww", "i").Get(); result != "ertyu" {
		t.Error("get", result, "should be", "ertyu")
	}

	if result := s.Copy().RangeCut(3, 5).Get(); result != "ertyui" {
		t.Error("get", result, "should be", "ertyui")
	}

	if result := s.Copy().RangeFirst(3).Get(); result != "qww" {
		t.Error("get", result, "should be", "qww")
	}

	if result := s.Copy().RangeLast(6).Get(); result != "iooooo" {
		t.Error("get", result, "should be", "iooooo")
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
