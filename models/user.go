package models

import (
	"errors"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username      string     `gorm:"unique,autoIncrement" json:"username"`
	Password      string     `json:"-"`
	Nickname      string     `json:"nickname"`
	Phone         string     `json:"phone"`
	Avatar        string     `json:"avatar"`
	Status        *int       `gorm:"default:1" json:"status"`
	Role          *int       `json:"role"`
	LastLoginTime *LocalTime `gorm:"autoUpdateTime" json:"last_login_time"`
	LastLoginIP   string     `json:"last_login_ip"`
}

// GetUserByName 根据用户名查询用户
func (u *User) GetByUsername(username string) (user User) {
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}

func (u *User) GetByNickname(nickname string) (user User) {
	err := db.Where("nickname = ?", nickname).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}

func (u *User) GetByID(userID uint) (user User) {
	err := db.First(&user, userID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}
func (u *User) GetList(req request.PageQuery) (users []User, total int) {
	db.Find(&users).Count(&total)
	db.Limit(req.PageSize).Offset(util.GetOffset(req)).
		Find(&users)
	return users, total
}
func (u *User) Create(user User) User {
	err := db.Create(&user).Error
	if err != nil {
		panic(err)
	}
	return user
}
func (u *User) Delete(req request.IdsJson) int {
	err := db.Table("tb_user").Where("id IN (?)", req.Ids).Updates(map[string]interface{}{"status": -1}).Error
	if err != nil {
		panic(err)
	}
	return e.SUCCESS
}
func (u *User) UpdateStatus(req request.UpdateStatus) int {
	err := db.Model(User{}).Where("id IN (?)", req.Ids).Updates(User{Status: req.Status}).Error
	if err != nil {
		panic(err)
	}
	return e.SUCCESS
}
func (u *User) Save(user User) bool {
	err := db.Save(&user).Error
	if err != nil {
		panic(err)
	}
	return true
}
