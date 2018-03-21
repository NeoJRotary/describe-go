package describe

import (
	"bytes"
	"testing"
)

func TestJSONBytes(t *testing.T) {
	v := map[string]interface{}{"key": "value"}
	shouldBe := []byte(`{"key":"value"}`)
	if !bytes.Equal(JSONBytes(v), shouldBe) {
		t.Errorf("%+v should be %v, get : %v", v, shouldBe, JSONBytes(v))
	}

	var a chan string
	if len(JSONBytes(a)) != 0 {
		t.Error("should be empty []byte")
	}
}

func TestJSONString(t *testing.T) {
	v := map[string]interface{}{"key": "value"}
	shouldBe := `{"key":"value"}`
	if JSONString(v) != shouldBe {
		t.Errorf("%+v should be %v", v, shouldBe)
	}

	var a chan string
	if JSONString(a) != "" {
		t.Error("should be empty string")
	}
}
