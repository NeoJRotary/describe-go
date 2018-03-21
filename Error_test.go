package describe

import "testing"

func TestErrorPrefix(t *testing.T) {
	e := NewErr("aaa")

	if msg := Error(e).Prefix("gg").Msg(); msg != "ggaaa" {
		t.Error("get ", msg, ", should be ggaaa")
	}
}
