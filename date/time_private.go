/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:00
  @note:
**/
package date

import (
	"errors"
	"time"
)

func (dateTime *DateTime) parseWithFormat(str string, location *time.Location) (t time.Time, err error) {
	for _, format := range TimeFormats {
		t, err = time.ParseInLocation(format, str, location)

		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}
