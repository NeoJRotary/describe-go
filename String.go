package describe

import "strings"

// TypeString string function collections struct
type TypeString struct {
	Obj        string
	SliceIndex int
}

// String get *TypeString
func String(obj string) *TypeString {
	return &TypeString{Obj: obj, SliceIndex: -1}
}

// Copy get copy of type
func (s *TypeString) Copy() *TypeString {
	return &TypeString{Obj: s.Obj, SliceIndex: s.SliceIndex}
}

// Update update object
func (s *TypeString) Update(str string) *TypeString {
	s.Obj = str
	return s
}

// Get return object string
func (s *TypeString) Get() string {
	return s.Obj
}

// Len length of string
func (s *TypeString) Len() int {
	return len(s.Obj)
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
	return s.Copy().Update(strings.Replace(s.Obj, old, new, n))
}

// ReplaceAll wrapper of strings.Replace(), set n = -1
func (s *TypeString) ReplaceAll(old string, new string) *TypeString {
	return s.Copy().Update(strings.Replace(s.Obj, old, new, -1))
}

// SetInto set element into target slice by index. Panic if element doesn't have index.
func (s *TypeString) SetInto(ss *TypeStringSlice) *TypeStringSlice {
	return ss.Set(s.SliceIndex, s.Obj)
}

// SetSliceIndex set index of slice on TypeString. Panic when invalid index
func (s *TypeString) SetSliceIndex(i int) *TypeString {
	if i < 0 {
		PanicErr(ErrInvalidIndex)
	}
	s.SliceIndex = i
	return s
}

// Index find fisrt substr position in object. Retrun -1 if not found.
func (s *TypeString) Index(substr string) int {
	return strings.Index(s.Obj, substr)
}

// LastIndex find last substr position in object. Retrun -1 if not found.
func (s *TypeString) LastIndex(substr string) int {
	return strings.LastIndex(s.Obj, substr)
}

// Split wrapper of strings.Split()
func (s *TypeString) Split(sep string) *TypeStringSlice {
	return StringSlice(strings.Split(s.Obj, sep))
}

// Trim wrapper of strings.Trim()
func (s *TypeString) Trim(cutset string) *TypeString {
	return s.Copy().Update(strings.Trim(s.Obj, cutset))
}

// TrimLeft wrapper of strings.TrimLeft()
func (s *TypeString) TrimLeft(cutset string) *TypeString {
	return s.Copy().Update(strings.TrimLeft(s.Obj, cutset))
}

// TrimRight wrapper of strings.TrimRight()
func (s *TypeString) TrimRight(cutset string) *TypeString {
	return s.Copy().Update(strings.TrimRight(s.Obj, cutset))
}

// TrimPrefix wrapper of strings.TrimPrefix()
func (s *TypeString) TrimPrefix(prefix string) *TypeString {
	return s.Copy().Update(strings.TrimPrefix(s.Obj, prefix))
}

// TrimSuffix wrapper of strings.TrimSuffix()
func (s *TypeString) TrimSuffix(suffix string) *TypeString {
	return s.Copy().Update(strings.TrimSuffix(s.Obj, suffix))
}

// TrimSpace wrapper of strings.TrimSpace()
func (s *TypeString) TrimSpace() *TypeString {
	return s.Copy().Update(strings.TrimSpace(s.Obj))
}

// Empty is empty string
func (s *TypeString) Empty() bool {
	return s.Obj == ""
}

// Range get part of object from index i to j (exclude j). Panic when invalid index or out of range.
func (s *TypeString) Range(i int, j int) *TypeString {
	if i < 0 || j < 0 || i > j {
		PanicErr(ErrInvalidIndex)
	}
	if i > s.Len() {
		PanicErr(ErrOutOfRange)
	}
	if j > s.Len() {
		PanicErr(ErrOutOfRange)
	}
	return s.Copy().Update(s.Obj[i:j])
}

// RangeBetween get part of object by given first and last element. Return empty element if not found.
func (s *TypeString) RangeBetween(a, b string) *TypeString {
	i := s.Index(a)
	if i == -1 {
		return String("")
	}
	i += len(a)
	j := s.LastIndex(b)
	if j == -1 {
		return String("")
	}
	return s.Range(i, j)
}
