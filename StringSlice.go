package describe

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
	return &TypeStringSlice{Obj: ss.Obj}
}

// ElmAt return element of slice at position in describe.Type. Panic if slice is empty or out of range.
func (ss *TypeStringSlice) ElmAt(i int) *TypeString {
	if ss.Empty() {
		panic(ErrSliceIsEmpty)
	}
	if i >= len(ss.Obj) {
		panic(ErrSliceOutOfRange)
	}
	return String(ss.Obj[i]).IndexAt(i)
}

// Get return object slice
func (ss *TypeStringSlice) Get() []string {
	return ss.Obj
}

// Len return len(slice)
func (ss *TypeStringSlice) Len() int {
	return len(ss.Obj)
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
	if i > j {
		panic(ErrSliceRangeIndexInvalid)
	}
	if i > ss.Len() {
		panic(ErrSliceOutOfRange)
	}
	if j > ss.Len() {
		panic(ErrSliceOutOfRange)
	}
	return StringSlice(ss.Obj[i:j])
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

// First return first element of slice in describe.Type. Panic if slice is empty.
func (ss *TypeStringSlice) First() *TypeString {
	if ss.Empty() {
		panic(ErrSliceIsEmpty)
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
		panic(ErrSliceIsEmpty)
	}
	return String(ss.Obj[len(ss.Obj)-1])
}

// Set set element at position of slice. Panic if index is invalid.
func (ss *TypeStringSlice) Set(i int, s string) *TypeStringSlice {
	ss.ElmAt(i)
	cp := ss.Copy()
	cp.Obj[i] = s
	return cp
}
