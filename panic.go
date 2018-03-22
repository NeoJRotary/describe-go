package describe

const (
	// ErrOutOfRange out of range
	ErrOutOfRange = "out of range"
	// ErrSliceIsEmpty slice is empty
	ErrSliceIsEmpty = "slice is empty"
	// ErrRangeIndexInvalid range index error : i > j or index < 0
	ErrRangeIndexInvalid = "range index error : i > j or index < 0"
	// ErrInvalidIndex invalid index < 0
	ErrInvalidIndex = "invalid index < 0"
)

// PanicErr panic error by string
func PanicErr(s string) {
	panic(NewErr(s))
}
