/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:00
  @note:
**/
package date

import (
	"github.com/shenyuanheli/tools/common"
	"time"
)

// Time return time structure
func (dateTime *DateTime) Time() time.Time {
	return dateTime.t
}

// Now get standard time
// If time Now is empty, the single-sign time is obtained, otherwise the set time
func (dateTime *DateTime) Now() *DateTime {
	dateTime.t = time.Now()
	return dateTime
}

// Format format time
func (dateTime *DateTime) Format(s ...string) string {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	if len(s) > 0 {
		replace, b := replace(s[0])
		if !b {
			return dateTime.t.Format(complete)
		} else {
			return dateTime.t.Format(replace)
		}
	}
	return dateTime.t.Format(complete)
}

// SetTime set time
func (dateTime *DateTime) SetTime(t ...time.Time) *DateTime {
	dateTime.t = time.Now()
	if len(t) > 0 {
		dateTime.t = t[0]
	}
	return dateTime
}

// OperationDay Add and subtract time by day
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationDay(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(0, 0, num)
	return dateTime
}

// OperationMoon Add and subtract time by Moon
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationMoon(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(0, num, 0)
	return dateTime
}

// OperationYears Add and subtract time by Years
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationYears(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(num, 0, 0)
	return dateTime
}

//StartOfHour start of hour
func (dateTime *DateTime) StartOfHour() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, d := dateTime.t.Date()
	dateTime.t = time.Date(y, m, d, dateTime.t.Hour(), 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// StartTimeOfDay start time of day
// This method can get the start time of the day, which is year:month:day 00:00:00
func (dateTime *DateTime) StartTimeOfDay() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, d := dateTime.t.Date()
	dateTime.t = time.Date(y, m, d, 0, 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// StartOfMinute beginning of minute
func (dateTime *DateTime) StartOfMinute() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.Truncate(time.Minute)
	return dateTime
}

// StartOfWeek Weekly start time
func (dateTime *DateTime) StartOfWeek() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
		dateTime.weekStartDay = weekStartDay
	}
	t := dateTime.StartTimeOfDay()
	weekday := int(t.Time().Weekday())
	if dateTime.weekStartDay != common.Sunday {
		weekStartDayInt := int(dateTime.weekStartDay)
		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	dateTime.t = t.t.AddDate(0, 0, -weekday)
	return dateTime
}

// StartOfMonth beginning of month
// This method user gets the start time of each month
func (dateTime *DateTime) StartOfMonth() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, _ := dateTime.t.Date()
	dateTime.t = time.Date(y, m, 1, 0, 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// SetWeekStartDay Set the start time of the week, the default is Sunday. China generally uses a week to start on Monday
//Use this method to quickly set up The value range is 0-6, please use the enumeration preset by the system
func (dateTime *DateTime) SetWeekStartDay(weekday common.Weekday) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	if weekday < 0 || weekday > 6 {
		weekday = common.Sunday
	}
	dateTime.weekStartDay = weekday
	return dateTime
}

// StartOfQuarter Start time of current quarter
func (dateTime *DateTime) StartOfQuarter() *DateTime {
	month := dateTime.StartOfMonth()
	offset := (int(month.t.Month()) - 1) % 3
	dateTime.t = month.t.AddDate(0, -offset, 0)
	return dateTime
}

// StartOfYear Start time of current year
func (dateTime *DateTime) StartOfYear() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, _, _ := dateTime.t.Date()
	dateTime.t = time.Date(y, time.January, 1, 0, 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// EndOfMinute end of minute
func (dateTime *DateTime) EndOfMinute() *DateTime {
	dateTime.t = dateTime.StartOfMinute().t.Add(time.Minute - time.Nanosecond)
	return dateTime
}

//EndOfHour end of hour
func (dateTime *DateTime) EndOfHour() *DateTime {
	dateTime.t = dateTime.StartOfHour().t.Add(time.Hour - time.Nanosecond)
	return dateTime
}

//EndOfDay end of day
func (dateTime *DateTime) EndOfDay() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, d := dateTime.t.Date()
	dateTime.t = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), dateTime.t.Location())
	return dateTime
}

//EndOfWeek end of week
func (dateTime *DateTime) EndOfWeek() *DateTime {
	dateTime.t = dateTime.StartOfWeek().t.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return dateTime
}

// EndOfMonth end of month
func (dateTime *DateTime) EndOfMonth() *DateTime {
	dateTime.t = dateTime.StartOfMonth().t.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return dateTime
}

// EndOfQuarter end of quarter
func (dateTime *DateTime) EndOfQuarter() *DateTime {
	dateTime.t = dateTime.StartOfQuarter().t.AddDate(0, 3, 0).Add(-time.Nanosecond)
	return dateTime
}

// EndOfYear end of year
func (dateTime *DateTime) EndOfYear() *DateTime {
	dateTime.t = dateTime.StartOfYear().t.AddDate(1, 0, 0).Add(-time.Nanosecond)
	return dateTime
}

// Parse parse string to time
func (dateTime *DateTime) Parse(strFormat ...string) (*DateTime, error) {
	var err error
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	var (
		setCurrentTime  bool
		parseTime       []int
		currentLocation = dateTime.t.Location()
		onlyTimeInStr   = true
		currentTime     = formatTimeToList(dateTime.t)
	)

	for _, str := range strFormat {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err := dateTime.parseWithFormat(str, currentLocation); err == nil {
			location := t.Location()
			parseTime = formatTimeToList(t)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3],
				parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = formatTimeToList(t)
			dateTime.t = t
		}
	}
	return dateTime, err
}

// MustParse must parse string to time or it will panic
func (dateTime *DateTime) MustParse(strs ...string) (d *DateTime) {
	d, err := dateTime.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return d
}
