package request

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type PageQuery struct {
	PageSize int    `form:"page_size" binding:"required,number,min=1"'`
	PageNum  int    `form:"page_num" binding:"required,number,min=1"`
	Keyword  string `form:"keyword"`
}
type UpdateStatus struct {
	IdsJson
	Status *int `json:"status" binding:"required,number,oneof=-1 0 1"`
}
type LocalTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

// UnmarshalJSON implements json unmarshal interface.
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
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

type IdsJson struct {
	Ids []uint64 `json:"ids" binding:"required,min=1"`
}

type IdQuery struct {
	Id uint64 `form:"id" binding:"required,number,min=1"`
}

type IdPrimaryKey struct {
	ID uint64 `json:"id" binding:"required,number,gt=0"`
}
