package toolkit

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println("today date:", BJTodayDate())
	fmt.Println("yesterday date:", BJYesterdayDate())
	fmt.Println("now time:", FormatBJNowTime())
	fmt.Println("loc now time:", FormatLocalNowTime())
}
