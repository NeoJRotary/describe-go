package describe

var timeFormatCompareList = []string{
	"YYYY", "2006",
	"YY", "06",
	"DD", "02",
	"D", "2",
	"HH", "15",
	"mm", "04",
	"m", "4",
	"ss", "05",
	"s", "5",
	"MMMM", "January",
	"MMM", "Jan",
	"MM", "01",
	"M", "1",
}

// GetTimeFormat convert YYYYMMDD format to golang's time format layout
func GetTimeFormat(layout string) string {
	s := String(layout)
	for i := 0; i < len(timeFormatCompareList); i += 2 {
		s.ReplaceAll(timeFormatCompareList[i], timeFormatCompareList[i+1])
	}
	return s.Get()
}
