package describe

import (
	"strconv"
	"strings"
)

// TypeString string function collections struct
type TypeString struct {
	Obj        string
	SliceIndex int
}

// String get *TypeString
func String(obj ...string) *TypeString {
	if len(obj) > 0 {
		return &TypeString{Obj: obj[0], SliceIndex: -1}
	}
	return &TypeString{Obj: "", SliceIndex: -1}
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

// Has obejct string contain all elements
func (s *TypeString) Has(elm string, more ...string) bool {
	if s.Index(elm) == -1 {
		return false
	}
	for _, elmS := range more {
		if s.Index(elmS) == -1 {
			return false
		}
	}
	return true
}

// HasOne obejct string contain one of elements
func (s *TypeString) HasOne(elm string, more ...string) bool {
	if s.Index(elm) != -1 {
		return true
	}
	for _, elmS := range more {
		if s.Index(elmS) != -1 {
			return true
		}
	}
	return false
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
	return s.Update(strings.Replace(s.Obj, old, new, n))
}

// ReplaceAll wrapper of strings.Replace(), set n = -1
func (s *TypeString) ReplaceAll(old string, new string) *TypeString {
	return s.Update(strings.Replace(s.Obj, old, new, -1))
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
	return s.Update(strings.Trim(s.Obj, cutset))
}

// TrimLeft wrapper of strings.TrimLeft()
func (s *TypeString) TrimLeft(cutset string) *TypeString {
	return s.Update(strings.TrimLeft(s.Obj, cutset))
}

// TrimRight wrapper of strings.TrimRight()
func (s *TypeString) TrimRight(cutset string) *TypeString {
	return s.Update(strings.TrimRight(s.Obj, cutset))
}

// TrimPrefix wrapper of strings.TrimPrefix()
func (s *TypeString) TrimPrefix(prefix string) *TypeString {
	return s.Update(strings.TrimPrefix(s.Obj, prefix))
}

// TrimSuffix wrapper of strings.TrimSuffix()
func (s *TypeString) TrimSuffix(suffix string) *TypeString {
	return s.Update(strings.TrimSuffix(s.Obj, suffix))
}

// TrimSpace wrapper of strings.TrimSpace()
func (s *TypeString) TrimSpace() *TypeString {
	return s.Update(strings.TrimSpace(s.Obj))
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
	return s.Update(s.Obj[i:j])
}

// RangeBetween get part of object by given first and last element. Panic if a or b not found.
func (s *TypeString) RangeBetween(a, b string) *TypeString {
	i := s.Index(a)
	if i == -1 {
		PanicErr(ErrInvalidIndex)
	}
	i += len(a)
	j := s.LastIndex(b)
	if j == -1 {
		PanicErr(ErrInvalidIndex)
	}
	return s.Range(i, j)
}

// RangeCut get part of obejct by cut (n) chars from start and (m) chars from end
func (s *TypeString) RangeCut(n, m int) *TypeString {
	return s.Range(n, s.Len()-m)
}

// RangeFirst get first n chars
func (s *TypeString) RangeFirst(n int) *TypeString {
	return s.Range(0, n)
}

// RangeLast get last n chars
func (s *TypeString) RangeLast(n int) *TypeString {
	return s.Range(s.Len()-n, s.Len())
}

// RangeFrom get part of string from index of sub (exclude sub)
func (s *TypeString) RangeFrom(sub string) *TypeString {
	i := s.Index(sub)
	if i == -1 {
		return s.Update("")
	}
	return s.Range(i+len(sub), s.Len())
}

// RangeUntil get part of string until index of sub (exclude sub)
func (s *TypeString) RangeUntil(sub string) *TypeString {
	i := s.Index(sub)
	if i == -1 {
		return s.Update("")
	}
	return s.Range(0, i)
}

// WrapBy wrap string by wrapper
func (s *TypeString) WrapBy(wrapper string) *TypeString {
	return s.Update(wrapper + s.Obj + wrapper)
}

// ToInt convert string to int. Panic if strconv throw error.
func (s *TypeString) ToInt() int {
	return int(s.ToInt32())
}

// ToInt8 convert string to int8. Panic if strconv throw error.
func (s *TypeString) ToInt8() int8 {
	i, e := strconv.ParseInt(s.Obj, 10, 8)
	CheckErr(e)
	return int8(i)
}

// ToInt16 convert string to int16. Panic if strconv throw error.
func (s *TypeString) ToInt16() int16 {
	i, e := strconv.ParseInt(s.Obj, 10, 16)
	CheckErr(e)
	return int16(i)
}

// ToInt32 convert string to int32. Panic if strconv throw error.
func (s *TypeString) ToInt32() int32 {
	i, e := strconv.ParseInt(s.Obj, 10, 32)
	CheckErr(e)
	return int32(i)
}

// ToInt64 convert string to int64. Panic if strconv throw error.
func (s *TypeString) ToInt64() int64 {
	i, e := strconv.ParseInt(s.Obj, 10, 64)
	CheckErr(e)
	return i
}

// ToUint convert string to uint. Panic if strconv throw error.
func (s *TypeString) ToUint() uint {
	return uint(s.ToUint32())
}

// ToUint8 convert string to uint8. Panic if strconv throw error.
func (s *TypeString) ToUint8() uint8 {
	i, e := strconv.ParseUint(s.Obj, 10, 8)
	CheckErr(e)
	return uint8(i)
}

// ToUint16 convert string to uint16. Panic if strconv throw error.
func (s *TypeString) ToUint16() uint16 {
	i, e := strconv.ParseUint(s.Obj, 10, 16)
	CheckErr(e)
	return uint16(i)
}

// ToUint32 convert string to uint32. Panic if strconv throw error.
func (s *TypeString) ToUint32() uint32 {
	i, e := strconv.ParseUint(s.Obj, 10, 32)
	CheckErr(e)
	return uint32(i)
}

// ToUint64 convert string to uint64. Panic if strconv throw error.
func (s *TypeString) ToUint64() uint64 {
	i, e := strconv.ParseUint(s.Obj, 10, 64)
	CheckErr(e)
	return i
}

// ToFloat32 convert string to float32. Panic if strconv throw error.
func (s *TypeString) ToFloat32() float32 {
	i, e := strconv.ParseFloat(s.Obj, 32)
	CheckErr(e)
	return float32(i)
}

// ToFloat64 convert string to float64. Panic if strconv throw error.
func (s *TypeString) ToFloat64() float64 {
	i, e := strconv.ParseFloat(s.Obj, 64)
	CheckErr(e)
	return i
}
