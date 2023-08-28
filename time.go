package toolkit

import "time"

// BJTodayDate get today's date in Beijing time
func BJTodayDate() string {
	loc, _ := GetBeijingLoc()
	today := time.Now().In(loc)
	return today.Format(time.DateOnly)
}

// BJYesterdayDate get yesterday's date in Beijing time
func BJYesterdayDate() string {
	loc, _ := GetBeijingLoc()
	yesterday := time.Now().In(loc).AddDate(0, 0, -1)
	return yesterday.Format(time.DateOnly)
}

// FormatBJNowTime ...
func FormatBJNowTime() string {
	loc, _ := GetBeijingLoc()
	today := time.Now().In(loc)
	return today.Format(time.DateTime)
}

// FormatLocalNowTime ...
func FormatLocalNowTime() string {
	return time.Now().Format(time.DateTime)
}

// GetBeijingLoc ...
func GetBeijingLoc() (*time.Location, error) {
	return time.LoadLocation("Asia/Shanghai")
}
