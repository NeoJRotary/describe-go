package describe

import (
	"testing"
)

func TestStringSliceSame(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := []string{"aadfw", "fwfe"}
	if !StringSlice(a).Same(a) {
		t.Error("should be true")
	}
	if StringSlice(a).Same(b) {
		t.Error("should be false")
	}
}

func TestStringSliceHas(t *testing.T) {
	a := []string{"1", "2", "3"}
	if !StringSlice(a).Has("1") {
		t.Error("should be true")
	}
	if StringSlice(a).Has("fjf") {
		t.Error("should be false")
	}
}

func TestStringSliceTrim(t *testing.T) {
	a := []string{"1", "2", "3", "1", "1"}
	b := []string{"2", "3"}
	if ss := StringSlice(a).Trim("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"2", "3", "1"}
	if ss := StringSlice(a).Trim("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"1", "2", "3"}
	if ss := StringSlice(a).Trim("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"2", "3"}
	if ss := StringSlice(a).Trim("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{}
	if ss := StringSlice(a).Trim("1"); !ss.Same([]string{}) {
		t.Error("Get ", ss.Get(), ", should be ", []string{})
	}
}

func TestStringSliceTrimSpace(t *testing.T) {
	a := []string{"", "2", "3", "", ""}
	b := []string{"2", "3"}
	if ss := StringSlice(a).TrimSpace(); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"2", "3", ""}
	if ss := StringSlice(a).TrimSpace(); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"", "2", "3"}
	if ss := StringSlice(a).TrimSpace(); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{"2", "3"}
	if ss := StringSlice(a).TrimSpace(); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	a = []string{}
	if ss := StringSlice(a).TrimSpace(); !ss.Same([]string{}) {
		t.Error("Get ", ss.Get(), ", should be ", []string{})
	}
}

func TestStringSliceElm(t *testing.T) {
	a := []string{"121", "131"}
	b := []string{"2", "3"}
	if ss := StringSlice(a).ElmTrim("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	if ss := StringSlice(b).ElmWrapBy("1"); !ss.Same(a) {
		t.Error("Get ", ss.Get(), ", should be ", a)
	}

	a = []string{" 2 ", "    3"}
	if ss := StringSlice(a).ElmTrimSpace(); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}
}

func TestStringPush(t *testing.T) {
	a := []string{"2", "3"}
	b := []string{"2", "3", "4", "5"}

	if ss := StringSlice(a).Push("4", "5"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	if ss := StringSlice(a).Push("4").Push("5"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}
}

func TestStringShift(t *testing.T) {
	a := []string{"3", "4"}
	b := []string{"1", "2", "3", "4"}

	if ss := StringSlice(a).Shift("1", "2"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	if ss := StringSlice(a).Shift("2").Shift("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}
}

func BenchmarkNativeAppend(b *testing.B) {
	a := []string{}
	for i := 0; i < b.N; i++ {
		a = append(a, "1")
	}
}

func BenchmarkStringSlicePush(b *testing.B) {
	a := StringSlice(nil)
	for i := 0; i < b.N; i++ {
		a = a.Push("1")
	}
}
