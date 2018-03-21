package describe

import "strings"

// TypeString string function collections struct
type TypeString struct {
	Obj string
}

// String get *TypeString
func String(obj string) *TypeString {
	return &TypeString{Obj: obj}
}

// Get return object string
func (s *TypeString) Get() string {
	return s.Obj
}

// HasPrefix wrapper of strings.HasPrefix()
func (s *TypeString) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.Obj, prefix)
}

// HasSuffix wrapper of strings.HasSuffix()
func (s *TypeString) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.Obj, suffix)
}

// Replace wrapper of strings.Replace()
func (s *TypeString) Replace(old string, new string, n int) *TypeString {
	return String(strings.Replace(s.Obj, old, new, n))
}

// ReplaceAll wrapper of strings.Replace(), set n = -1
func (s *TypeString) ReplaceAll(old string, new string) *TypeString {
	return String(strings.Replace(s.Obj, old, new, -1))
}

// Split wrapper of strings.Split()
func (s *TypeString) Split(sep string) *TypeStringSlice {
	return StringSlice(strings.Split(s.Obj, sep))
}

// Trim wrapper of strings.Trim()
func (s *TypeString) Trim(cutset string) *TypeString {
	return String(strings.Trim(s.Obj, cutset))
}

// TrimSpace wrapper of strings.TrimSpace()
func (s *TypeString) TrimSpace() *TypeString {
	return String(strings.TrimSpace(s.Obj))
}
