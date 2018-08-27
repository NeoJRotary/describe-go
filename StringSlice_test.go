package describe

import (
	"strconv"
	"strings"
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

func TestStringSlicePush(t *testing.T) {
	a := []string{"2", "3"}
	b := []string{"2", "3", "4", "5"}

	if ss := StringSlice(a).Push("4", "5"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	if ss := StringSlice(a).Push("4").Push("5"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}
}

func TestStringSliceShift(t *testing.T) {
	a := []string{"3", "4"}
	b := []string{"1", "2", "3", "4"}

	if ss := StringSlice(a).Shift("1", "2"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}

	if ss := StringSlice(a).Shift("2").Shift("1"); !ss.Same(b) {
		t.Error("Get ", ss.Get(), ", should be ", b)
	}
}

func TestStringSliceFilter(t *testing.T) {
	a := []string{"11", "12", "13", "14", "15", "21", "22"}

	ss := StringSlice(a).FilterBy(func(s *TypeString) bool {
		return s.ToInt() < 15
	})
	if ss.Len() != 4 {
		t.Error("wrong len")
	}

	ss = StringSlice(a).FilterByPrefix("1")
	if ss.Len() != 5 {
		t.Error("wrong len")
	}

	ss = StringSlice(a).FilterBySuffix("3")
	if ss.Len() != 1 {
		t.Error("wrong len")
	}
}

func TestStringSliceDelete(t *testing.T) {
	ss := StringSlice([]string{"0", "1", "2", "3", "4"})

	if i := ss.IndexOf("1"); i != 1 {
		t.Error("wrong index", i)
	}

	if i := ss.IndexOf("9"); i != -1 {
		t.Error("wrong index", i)
	}

	ss.Delete(1)
	if s := ss.ElmAt(1).Get(); s != "2" {
		t.Error("delete error, get", s)
	}

	ss.DeleteSame("0", "4")
	if !ss.Same([]string{"2", "3"}) {
		t.Error("DeleteSame error, get", ss.Get())
	}
}

func BenchmarkStringSlice_NativeAppend(b *testing.B) {
	a := []string{}
	for i := 0; i < b.N; i++ {
		a = append(a, "1", "2", "3")
	}
}

func BenchmarkStringSlice_Push(b *testing.B) {
	a := StringSlice(nil)
	for i := 0; i < b.N; i++ {
		a = a.Push("1", "2", "3")
	}
}

func BenchmarkStringSlice_PushWithCopy(b *testing.B) {
	a := StringSlice(nil)
	for i := 0; i < b.N; i++ {
		a = a.Copy().Push("1", "2", "3")
	}
}

func BenchmarkStringSlice_FilterWithNativeString(b *testing.B) {
	longSlice := []string{}
	for i := 0; i < 10000; i++ {
		longSlice = append(longSlice, strconv.Itoa(i))
	}

	filter := func(s string) bool {
		return strings.HasSuffix(s, "34")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newss := []string{}
		for _, s := range longSlice {
			if filter(s) {
				newss = append(newss, s)
			}
		}
	}
}

func BenchmarkStringSlice_FilterWithTypeString(b *testing.B) {
	longSlice := []string{}
	for i := 0; i < 10000; i++ {
		longSlice = append(longSlice, strconv.Itoa(i))
	}

	filter := func(s *TypeString) bool {
		return s.HasSuffix("34")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newss := StringSlice(nil)
		for _, s := range longSlice {
			if filter(String(s)) {
				newss.Push(s)
			}
		}
	}
}
