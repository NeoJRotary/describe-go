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
