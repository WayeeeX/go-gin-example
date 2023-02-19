package models

import "time"

type LoginRecord struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
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
