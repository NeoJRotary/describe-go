package describe

import (
	"testing"
)

func TestGetENV(t *testing.T) {
	if env := GetENV("PATH", "default"); env == "default" {
		t.Error("GetENV PATH shouldn't get default")
	}
	if env := GetENV("SOMETHINGISNOTAENVINYOUROS", "default"); env != "default" {
		t.Error("should get default")
	}
}

func TestRecoverErr(t *testing.T) {
	count := 0
	func() {
		defer RecoverErr(func(error) {
			count++
		})
		panic(NewErr(123))
	}()
	if count != 1 {
		t.Fatal("should run recover func")
	}
}
