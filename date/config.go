/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:00
  @note:
**/
package date

import (
	"github.com/shenyuanheli/tools/common"
	"regexp"
	"strings"
	"time"
)

var weekStartDay = common.Sunday

type DateTime struct {
	t            time.Time
	weekStartDay common.Weekday
}

const (
	year            = "2006"
	month           = "01"
	day             = "02"
	hour            = "15"
	minute          = "04"
	second          = "05"
	complete        = "2006-01-02 15:04:05"
	stringToTimeOne = "2006-01-02 15:04:05"
	stringToTimeTow = "2006-01-02"
)

func formatTimeToList(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}

//time format string
func replace(s string) (string, bool) {
	flag := false
	if strings.Contains(s, "YYYY") {
		s = strings.Replace(s, "YYYY", year, 1)
		flag = true
	}
	if strings.Contains(s, "MM") {
		s = strings.Replace(s, "MM", month, 1)
		flag = true
	}
	if strings.Contains(s, "DD") {
		s = strings.Replace(s, "DD", day, 1)
		flag = true
	}
	if strings.Contains(s, "hh") {
		s = strings.Replace(s, "hh", hour, 1)
		flag = true
	}
	if strings.Contains(s, "mm") {
		s = strings.Replace(s, "mm", minute, 1)
		flag = true
	}
	if strings.Contains(s, "ss") {
		s = strings.Replace(s, "ss", second, 1)
		flag = true
	}
	return s, flag
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// TimeFormats default time formats will be parsed as
var TimeFormats = []string{
	"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2",
	"15:4:5", "15:4", "15",
	"15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05Z0700", "2006-01-02T15:04:05Z07",
	"2006.1.2", "2006.1.2 15:04:05", "2006.01.02", "2006.01.02 15:04:05", "2006.01.02 15:04:05.999999999",
	"1/2/2006", "1/2/2006 15:4:5", "2006/01/02", "20060102", "2006/01/02 15:04:05",
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
	time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}
