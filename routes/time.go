package routes

import (
	"time"
)

// =================================
// ========== [[ CONV ]] ===========
// =================================

func amOrPMFormatter(hour int, minute int, second int) string {
	var status string
	var mins string
	var secs string

	if hour >= 12 {
		status = "PM"
	} else if hour == 0 || hour < 12 {
		status = "AM"
	}

	if minute < 10 {
		mins = stringAppender("0", []string{
			IntToString(minute),
		})
	} else if minute > 9 {
		mins = IntToString(minute)
	}

	if second < 10 {
		secs = stringAppender("0", []string{
			IntToString(second),
		})
	} else if second > 9 {
		secs = IntToString(second)
	}

	return FormatString("{0}:{1}:{2} {3}", []string{
		IntToString(hour),
		mins,
		secs,
		status,
	})
}

func dayFormatter(day int) string {
	var result string

	if day == 1 || day == 21 || day == 31 {
		result = IntToString(day) + "st"
	} else if day == 2 || day == 22 {
		result = IntToString(day) + "nd"
	} else if day == 3 || day == 23 {
		result = IntToString(day) + "rd"
	} else {
		result = IntToString(day) + "th"
	}

	return result
}

func monthFormatter(month time.Month) string {
	var result string

	switch month {
	case time.January:
		result = "January"

	case time.February:
		result = "February"

	case time.March:
		result = "March"

	case time.April:
		result = "April"

	case time.May:
		result = "May"

	case time.June:
		result = "June"

	case time.July:
		result = "July"

	case time.August:
		result = "August"

	case time.September:
		result = "September"

	case time.October:
		result = "October"

	case time.November:
		result = "November"

	case time.December:
		result = "December"
	}

	return result
}

// =================================
// ========== [[ TIME ]] ===========
// =================================

func Date() string {
	var (
		str string = ""
		now        = time.Now()
	)
	var (
		amOrPm = amOrPMFormatter(now.Hour(), now.Minute(), now.Second())
		day    = dayFormatter(now.Day())
		month  = monthFormatter(now.Month())
		year   = IntToString(now.Year())
	)

	str = FormatString("{0} of {1} {2} | {3}", []string{
		day,
		month,
		year,
		amOrPm,
	})

	return str
}