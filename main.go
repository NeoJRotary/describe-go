package describe

import (
	"errors"
	"fmt"
	"os"
)

// IsErr is error or not
func IsErr(e error) bool {
	return e != nil
}

// CheckErr panic if it is error
func CheckErr(e error) {
	if IsErr(e) {
		panic(e)
	}
}

// NewErr return new error
func NewErr(msg ...interface{}) error {
	return errors.New(fmt.Sprint(msg))
}

// RecoverErr run function when recover an error, it will panic again if it is not error
func RecoverErr(f func(error)) {
	if r := recover(); r != nil {
		e, is := r.(error)
		if is {
			panic(r)
		}
		if f != nil {
			f(e)
		}
	}
}

// GetENV get env or use default
func GetENV(name string, defaultVal string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultVal
	}
	return v
}

// Found check Slice.Find() return valid element or not
func Found(elm interface{}) bool {
	return elm != nil
}

// NotFound check Slice.Find() return valid element or not
func NotFound(elm interface{}) bool {
	return elm == nil
}
