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
