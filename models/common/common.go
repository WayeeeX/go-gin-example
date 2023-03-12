package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Model struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	CreateTime *LocalTime `gorm:"autoCreateTime,<-:create" json:"create_time"`
	UpdateTime *LocalTime `gorm:"autoUpdateTime" json:"update_time"`
}
type LocalTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	millis, err := strconv.ParseInt(string(data), 10, 64)

	*t = LocalTime(time.Unix(0, millis*int64(time.Millisecond)))
	return err
}

// MarshalJSON implements json marshal interface.
func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}
func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

// Value ...
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Date time.Time

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date(nullTime.Time)
	return
}

func (date Date) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}

func (date Date) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *Date) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(date).Format("2006-01-02") + `"`), nil
}

func (t *Date) UnmarshalJSON(data []byte) error {
	// 如果日期字符串为空，那么直接将 LocalDate 类型设置为 nil
	if string(data) == "null" {
		*t = Date{}
		return nil
	}

	// 解析日期字符串
	parsedTime, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	if err != nil {
		return err
	}

	// 只保留日期信息
	parsedDate := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, parsedTime.Location())

	// 转换成 LocalDate 类型
	*t = Date(parsedDate)
	return nil
}
