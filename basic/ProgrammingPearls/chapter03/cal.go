package main

var months = []uint32{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

//日期结构
type MyDate struct {
	year uint32
	mon  uint32
	day  uint32
}

//初始化
func NewMyDate(year, mon, day uint32) *MyDate {
	return &MyDate{
		year: year,
		mon:  mon,
		day:  day}
}

//返回当天是这一年的第几天
func (self *MyDate) yearDay() uint32 {
	sum := self.day

	for i := 1; i < int(self.mon); i++ {
		sum += months[i]
	}

	if self.isLeap() && self.mon > 2 {
		sum++
	}
	return sum
}

//是否是闰年
func (self *MyDate) isLeap() bool {
	return (self.year%4 == 0 && self.year%100 == 0) || (self.year%400 == 0)
}

//两个日期之间相差的天数
func distanceDays(d1, d2 *MyDate) uint32 {
	sum := -(d1.yearDay())
	for ; d1.year < d2.year; d1.year++ {
		if d1.isLeap() {
			sum += 366
		} else {
			sum += 365
		}
	}
	return sum + d2.yearDay()
}

func weekday(d *MyDate) uint32 {
	demo := NewMyDate(1900, 1, 1)
	return distanceDays(demo, d) % 7
}

func main() {
	today := NewMyDate(2017, 10, 29)
	testDay := NewMyDate(2017, 11, 11)
	println("相隔天数:", distanceDays(today, testDay))

	wd := weekday(testDay)
	println("星期", wd)
}
