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

// Get return object []string
func (ss *TypeStringSlice) Get() []string {
	return ss.Obj
}

// Same whether it is totally same as target []string
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

// Has whether it has element (string)
func (ss *TypeStringSlice) Has(elm string) bool {
	for _, v := range ss.Obj {
		if v == elm {
			return true
		}
	}
	return false
}

// Trim remove target element (string) from both sides of []slice
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

// TrimSpace remove empty string from both sides of []slice
func (ss *TypeStringSlice) TrimSpace() *TypeStringSlice {
	return ss.Trim("")
}
