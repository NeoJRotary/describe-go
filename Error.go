package describe

import "errors"

// TypeError error function collections struct
type TypeError struct {
	Obj error
}

// Error get *TypeError
func Error(e error) *TypeError {
	return &TypeError{Obj: e}
}

// ErrorMsg get *TypeError By message string
func ErrorMsg(s string) *TypeError {
	return &TypeError{Obj: errors.New(s)}
}

// Get return object error
func (e *TypeError) Get() error {
	return e.Obj
}

// Msg return object error message string
func (e *TypeError) Msg() string {
	return e.Obj.Error()
}

// ToString return object error message to TypeString
func (e *TypeError) ToString() *TypeString {
	return String(e.Obj.Error())
}

// Is whether object is an error
func (e *TypeError) Is() bool {
	return e.Obj != nil
}

// IsNot whether object is not an error
func (e *TypeError) IsNot() bool {
	return e.Obj == nil
}

// Check if object is error, panic it
func (e *TypeError) Check() {
	if e.Is() {
		panic(e)
	}
}

// Prefix prefix msg to object error
func (e *TypeError) Prefix(s string) *TypeError {
	return ErrorMsg(s + e.Msg())
}
