/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package idcard

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type Idcard struct {
}

//15位身份证转为18位
var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

var area = make(map[string]string)

// initConfig locale configuration
func initConfig() {
	c := "./area.json"
	raw, err := ioutil.ReadFile(c)
	err = json.Unmarshal(raw, &area)
	if err != nil {
		err = fmt.Errorf("Failed to parse basic configuration file：%s\n", err.Error())
		return
	}
}

//18-digit ID verification code
func checkValidNo18(id string) (bool, error) {
	//string -> []byte
	id18 := []byte(id)
	nSum := 0
	for i := 0; i < len(id18)-1; i++ {
		n, _ := strconv.Atoi(string(id18[i]))
		nSum += n * weight[i]
	}
	//mod gets 18-bit ID verification code
	mod := nSum % 11
	if validValue[mod] == id18[17] {
		return true, nil
	}

	return false, errors.New("身份证不正确请进行核实")
}

// Verify birthday
func checkBirthdayCode(birthday string) (bool, error) {
	year, _ := strconv.Atoi(birthday[:4])
	month, _ := strconv.Atoi(birthday[4:6])
	day, _ := strconv.Atoi(birthday[6:])
	curYear, curMonth, curDay := time.Now().Date()
	//Birth date greater than current date
	if year < 1900 || year > curYear || month <= 0 || month > 12 || day <= 0 || day > 31 {
		return false, errors.New("请检查生日部分的日期是否正确")
	}

	if year == curYear {
		if month > int(curMonth) {
			return false, errors.New("当前日期月份小于您身份证上的月份")
		} else if month == int(curMonth) && day > curDay {
			return false, errors.New("当前日期天数小于您身份证上的天数")
		}
	}
	//出生日期在2月份
	if 2 == month {
		if isLeapYear(year) && day > 29 {
			return false, errors.New("闰年2月只有29号")
		} else if day > 28 {
			return false, errors.New("非闰年2月只有28号")
		}
	} else if 4 == month || 6 == month || 9 == month || 11 == month {
		if day > 30 {
			return false, errors.New("小月只有30号")
		}
	}
	return true, nil
}

// 判断是否为闰年
func isLeapYear(year int) bool {
	if year <= 0 {
		return false
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
