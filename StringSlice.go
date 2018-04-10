package describe

import "strings"

// TypeStringSlice []string function collections struct
type TypeStringSlice struct {
	Obj []string
}

// StringSlice get *TypeStringSlice
func StringSlice(obj []string) *TypeStringSlice {
	if obj == nil {
		obj = []string{}
	}
	return &TypeStringSlice{Obj: obj}
}

// Copy get copy of type
func (ss *TypeStringSlice) Copy() *TypeStringSlice {
	newSlice := make([]string, len(ss.Obj))
	copy(newSlice, ss.Obj)
	return &TypeStringSlice{Obj: newSlice}
}

// ElmAt return element of slice at position in describe.Type. Panic if slice is empty or out of range.
func (ss *TypeStringSlice) ElmAt(i int) *TypeString {
	if ss.Empty() {
		PanicErr(ErrSliceIsEmpty)
	}
	if i >= len(ss.Obj) {
		PanicErr(ErrOutOfRange)
	}
	return String(ss.Obj[i]).SetSliceIndex(i)
}

// ElmTrim trim all elements in slice
func (ss *TypeStringSlice) ElmTrim(cutset string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).Trim(cutset).Get()
	}
	return ss
}

// ElmTrimLeft trim all elements in slice
func (ss *TypeStringSlice) ElmTrimLeft(cutset string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).TrimLeft(cutset).Get()
	}
	return ss
}

// ElmTrimRight trim all elements in slice
func (ss *TypeStringSlice) ElmTrimRight(cutset string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).TrimRight(cutset).Get()
	}
	return ss
}

// ElmTrimPrefix trim all elements in slice
func (ss *TypeStringSlice) ElmTrimPrefix(prefix string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).TrimPrefix(prefix).Get()
	}
	return ss
}

// ElmTrimSuffix trim all elements in slice
func (ss *TypeStringSlice) ElmTrimSuffix(suffix string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).TrimSuffix(suffix).Get()
	}
	return ss
}

// ElmTrimSpace trim space all elements in slice
func (ss *TypeStringSlice) ElmTrimSpace() *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).TrimSpace().Get()
	}
	return ss
}

// ElmWrapBy wrap all elements in slice
func (ss *TypeStringSlice) ElmWrapBy(wrapper string) *TypeStringSlice {
	for i, elm := range ss.Obj {
		ss.Obj[i] = String(elm).WrapBy(wrapper).Get()
	}
	return ss
}

// Get return object slice
func (ss *TypeStringSlice) Get() []string {
	return ss.Obj
}

// Len return len(slice)
func (ss *TypeStringSlice) Len() int {
	return len(ss.Obj)
}

// Push push new elements into slice. Return new slice
func (ss *TypeStringSlice) Push(elm string, more ...string) *TypeStringSlice {
	// cp := append(append(append([]string{}, ss.Obj...), elm), more...)
	// return StringSlice(cp)
	ss.Obj = append(append(ss.Obj, elm), more...)
	return ss
}

// Shift shift elements into slice. Return new slice
func (ss *TypeStringSlice) Shift(elm string, more ...string) *TypeStringSlice {
	// cp := append(append([]string{elm}, more...), ss.Obj...)
	ss.Obj = append(append([]string{elm}, more...), ss.Obj...)
	return ss
}

// Empty is empty slice or not
func (ss *TypeStringSlice) Empty() bool {
	return len(ss.Obj) == 0
}

// From return slice[i:], if out of range return empty slice
func (ss *TypeStringSlice) From(i int) *TypeStringSlice {
	if i > ss.Len() {
		i = ss.Len()
	}
	return StringSlice(ss.Obj[i:])
}

// To return slice[:i], if out of range return slice[:last]
func (ss *TypeStringSlice) To(i int) *TypeStringSlice {
	if i > ss.Len() {
		i = ss.Len()
	}
	return StringSlice(ss.Obj[:i])
}

// Range return slice[i:j], panic when out of range or i > j
func (ss *TypeStringSlice) Range(i int, j int) *TypeStringSlice {
	if i < 0 || j < 0 || i > j {
		PanicErr(ErrRangeIndexInvalid)
	}
	if i > ss.Len() {
		PanicErr(ErrOutOfRange)
	}
	if j > ss.Len() {
		PanicErr(ErrOutOfRange)
	}
	return StringSlice(ss.Obj[i:j])
}

// Trim remove target element from both sides of slice
func (ss *TypeStringSlice) Trim(elm string) *TypeStringSlice {
	var start, end int
	for start = 0; start < len(ss.Obj); start++ {
		if ss.Obj[start] != elm {
			break
		}
	}
	for end = len(ss.Obj) - 1; end >= 0; end-- {
		if ss.Obj[end] != elm {
			break
		}
	}
	return StringSlice(ss.Obj[start : end+1])
}

// TrimSpace remove empty string from both sides of slice
func (ss *TypeStringSlice) TrimSpace() *TypeStringSlice {
	return ss.Trim("")
}

// Same whether it is totally same as target slice
func (ss *TypeStringSlice) Same(target []string) bool {
	if len(ss.Obj) != len(target) {
		return false
	}
	for i, v := range ss.Obj {
		if target[i] != v {
			return false
		}
	}
	return true
}

// Has whether it has element
func (ss *TypeStringSlice) Has(elm string) bool {
	for _, v := range ss.Obj {
		if v == elm {
			return true
		}
	}
	return false
}

// NotHas whether it has element
func (ss *TypeStringSlice) NotHas(elm string) bool {
	return !ss.Has(elm)
}

// First return first element of slice in describe.Type. Panic if slice is empty.
func (ss *TypeStringSlice) First() *TypeString {
	if ss.Empty() {
		PanicErr(ErrSliceIsEmpty)
	}
	return String(ss.Obj[0])
}

// Find find first element in describe.Type when func passed. Return nil if not found.
func (ss *TypeStringSlice) Find(f func(*TypeString) bool) *TypeString {
	for _, v := range ss.Obj {
		sv := String(v)
		if f(sv) {
			return sv
		}
	}
	return nil
}

// FindHasPrefix find first string has prefix. Return nil if not found.
func (ss *TypeStringSlice) FindHasPrefix(prefix string) *TypeString {
	for _, v := range ss.Obj {
		sv := String(v)
		if sv.HasPrefix(prefix) {
			return sv
		}
	}
	return nil
}

// FindHasSuffix find first string has prefix. Return nil if not found.
func (ss *TypeStringSlice) FindHasSuffix(suffix string) *TypeString {
	for _, v := range ss.Obj {
		sv := String(v)
		if sv.HasSuffix(suffix) {
			return sv
		}
	}
	return nil
}

// Last return last element of slice in describe.Type. Panic if slice is empty.
func (ss *TypeStringSlice) Last() *TypeString {
	if ss.Empty() {
		PanicErr(ErrSliceIsEmpty)
	}
	return String(ss.Obj[len(ss.Obj)-1])
}

// Set set element at position of slice. Panic if index is invalid.
func (ss *TypeStringSlice) Set(i int, s string) *TypeStringSlice {
	// cp := ss.Copy()
	// cp.Obj[i] = s
	// return cp
	ss.Obj[i] = s
	return ss
}

// Join join slice. retrun TypeString
func (ss *TypeStringSlice) Join(sep string) *TypeString {
	return String(strings.Join(ss.Obj, sep))
}

// Map update elements by func, return a new slice
func (ss *TypeStringSlice) Map(fc func(*TypeString) *TypeString) *TypeStringSlice {
	cp := ss.Copy()
	for i, v := range cp.Obj {
		cp.Obj[i] = fc(String(v)).Get()
	}
	return cp
}
