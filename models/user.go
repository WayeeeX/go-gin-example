package models

import (
	"errors"
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/jinzhu/gorm"
)

type User struct {
	common.Model
	Username      string            `gorm:"unique,autoIncrement" json:"username"`
	Password      string            `json:"-"`
	Nickname      string            `json:"nickname"`
	Phone         string            `json:"phone"`
	Avatar        string            `json:"avatar"`
	Status        *int              `gorm:"default:1" json:"status"`
	Role          *int              `json:"role"`
	LastLoginTime *common.LocalTime `gorm:"autoUpdateTime" json:"last_login_time"`
	LastLoginIP   string            `json:"last_login_ip"`
}

// GetUserByName 根据用户名查询用户
func (u *User) GetByUsername(username string) (user User) {
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}

func (u *User) GetByNickname(nickname string) (user User) {
	err := DB.Where("nickname = ?", nickname).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}

func (u *User) GetByID(userID uint64) (user User) {
	err := DB.First(&user, userID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return user
}
func (u *User) GetList(req request.PageQuery) (users []User, total int) {
	keyword := "%" + req.Keyword + "%"
	DB.Table("tb_user").Where("username like ? or nickname like ? or phone like ?", keyword, keyword, keyword).Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req)).
		Find(&users)
	return users, total
}
func (u *User) Create(user User) User {
	err := DB.Create(&user).Error
	if err != nil {
		panic(err)
	}
	return user
}
func (u *User) Delete(req request.IdsJson) int {
	err := DB.Table("tb_user").Where("id IN (?)", req.Ids).Updates(map[string]interface{}{"status": -1}).Error
	if err != nil {
		panic(err)
	}
	return e.SUCCESS
}
func (u *User) UpdateStatus(req request.UpdateStatus) int {
	err := DB.Model(User{}).Where("id IN (?)", req.Ids).Updates(User{Status: req.Status}).Error
	if err != nil {
		panic(err)
	}
	return e.SUCCESS
}
func (u *User) Save(user User) bool {
	err := DB.Save(&user).Error
	if err != nil {
		panic(err)
	}
	return true
}
