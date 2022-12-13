/**
  @author: fanyanan
  @date: 2022/6/10
  @note: //time test
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"github.com/shenyuanheli/tools/common"
	"testing"
)

//获取当前时间
func TestNow(t *testing.T) {
	fmt.Println(ztool.DateUtils.Now())
	fmt.Println(ztool.DateUtils.Now().Time())
}

//格式化时间
func TestFormat(t *testing.T) {
	fmt.Println(ztool.DateUtils.Format())
	fmt.Println(ztool.DateUtils.Format("YYYY-MM-DD hh:mm:ss"))
	fmt.Println(ztool.DateUtils.Format("YYYY#MM#DD hh:mm:ss"))
	fmt.Println(ztool.DateUtils.Format("YYYY"))
}

//设置时间
func TestSetTime(t *testing.T) {
	parse, err := ztool.DateUtils.Parse("2021-10-12")
	if err != nil {
		return
	}
	fmt.Println(ztool.DateUtils.SetTime(parse.Time()).Time())
}

//按天加减时间
func TestOperationDays(t *testing.T) {
	fmt.Printf("当前时间------->%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("计算后时间加10天------->%s\n", ztool.DateUtils.OperationDay(10).Format())
	fmt.Printf("计算后时间减10天------->%s\n", ztool.DateUtils.OperationDay(-10).Format())
}

//按月加减时间
func TestOperationMoon(t *testing.T) {
	fmt.Printf("当前时间------->%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("计算后时间加2个月------->%s\n", ztool.DateUtils.OperationMoon(2).Format())
	fmt.Printf("计算后时间减2个月------->%s\n", ztool.DateUtils.Now().OperationMoon(-2).Format())
}

//按年加减时间
func TestOperationYears(t *testing.T) {
	fmt.Printf("当前时间------->%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("计算后时间加1年------->%s\n", ztool.DateUtils.OperationYears(1).Format())
	fmt.Printf("计算后时间减1年------->%s\n", ztool.DateUtils.Now().OperationYears(-1).Format())
}

//获取当前时间小时的开始时间
func TestStartOfHour(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取当前时间小时的开始时间------>%s\n", ztool.DateUtils.StartOfHour().Format())
}

//获取一天的开始时间
func TestStartTimeOfDay(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取一天的开始时间------>%s\n", ztool.DateUtils.StartTimeOfDay().Format())
}

//获取分钟的开时间
func TestStartOfMinute(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取分钟的开始时间------>%s\n", ztool.DateUtils.StartOfMinute().Format())
}

//获取一周的开始时间
func TestStartOfWeek(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取一周的开始时间------>%s\n", ztool.DateUtils.StartOfWeek().Format())
}

//获取当月的开始时间
func TestStartOfMonth(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取当月的开始时间------>%s\n", ztool.DateUtils.StartOfMonth().Format())
}

//设置一周的开始时间从那天开始
func TestSetWeekStartDay(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("设置一周开始时间从周一开始，最后获取到的时间"+
		"------>%s\n", ztool.DateUtils.SetWeekStartDay(common.Monday).StartOfWeek().Format())
}

//获取当前季度的开始时间
func TestStartOfQuarter(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("当前季度开始时间------>%s\n", ztool.DateUtils.StartOfQuarter().Format())
}

//获取当前时间的本年度开始时间
func TestStartOfYear(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("本年度开始时间------>%s\n", ztool.DateUtils.StartOfYear().Format())
}

//获取当前分钟的结束时间
func TestEndOfMinute(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("当前分钟结束时间------>%s\n", ztool.DateUtils.EndOfMinute().Format())
}

//获取到本小时的结束时间
func TestEndOfHour(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("本小时结束时间------>%s\n", ztool.DateUtils.EndOfHour().Format())
}

//获取今天的结束时间
func TestStartOfHalf(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取今年过半的开始时间------>%s\n", ztool.DateUtils.EndOfDay().Format())
}

//获取本周结束时间
func TestEndOfWeek(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("获取今年过半的开始时间------>%s\n", ztool.DateUtils.SetWeekStartDay(common.Monday).EndOfWeek().Format())
}

//获取本月结束时间
func TestEndOfMonth(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("本月结束时间------>%s\n", ztool.DateUtils.EndOfMonth().Format())
}

//本季度结束时间
func TestEndOfQuarter(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("本季度结束时间------>%s\n", ztool.DateUtils.EndOfQuarter().Format())
}

//本年年度结束时间
func TestEndOfYear(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	fmt.Printf("本年年度结束时间------>%s\n", ztool.DateUtils.EndOfYear().Format())
}

//字符串格式化时间
func TestParse(t *testing.T) {
	fmt.Printf("当前时间------>%s\n", ztool.DateUtils.Now().Format())
	parse, err := ztool.DateUtils.Parse()
	if err != nil {
		return
	}
	fmt.Printf("字符串格式化时间无惨------>%s\n", parse.Format())
	time, err := ztool.DateUtils.Parse("2018-10-2")
	if err != nil {
		return
	}
	fmt.Printf("字符串格式化时间带参数------>%s\n", time.Format())
}
