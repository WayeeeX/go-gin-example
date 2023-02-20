package models

import "time"

type LoginRecord struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64
	LastLoginIP string
	LoginTime   *time.Time `gorm:"autoCreateTime"`
}

func (l *LoginRecord) Create(loginRecord LoginRecord) LoginRecord {
	err := db.Create(&loginRecord).Error
	if err != nil {
		panic(err)
	}
	return loginRecord
}
