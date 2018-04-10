package describe

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimeFormat(t *testing.T) {
	datetime := time.Date(2018, 04, 23, 17, 59, 8, 0, time.Now().Location())
	s := datetime.Format(GetTimeFormat("YYYY-MM-DD hh:mm:ss"))
	if s != "2018-04-23 17:59:08" {
		fmt.Println("get ", s)
	}
}
