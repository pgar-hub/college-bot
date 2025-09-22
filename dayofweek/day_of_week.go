package dayofweek

import (
	"time"
)

func DayOfWeek() string {
	_, week := time.Now().ISOWeek()
	if week%2 == 0 {
		return "числитель"
	} else {
		return "знаменатель"
	}

}

func Today() string {
	day := time.Now().Weekday()
	return RussianDayOfWeek(day)
}
func TomorrowDay() string {
	tomorrowDay := time.Now().Add(24 * time.Hour)
	return RussianDayOfWeek(tomorrowDay.Weekday())
}
func RussianDayOfWeek(t time.Weekday) string {
	days := map[time.Weekday]string{
		time.Monday:    "понедельник",
		time.Tuesday:   "вторник",
		time.Wednesday: "среда",
		time.Thursday:  "четверг",
		time.Friday:    "пятница",
		time.Saturday:  "суббота",
		time.Sunday:    "воскресенье",
	}
	return days[t]
}
