package service

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (u *UserService) AdminLogin(username string, password string, ip string) (resLogin response.AdminLogin, code int) {
	//根据名称查询用户
	user := userModel.GetByUsername(username)
	if user.ID == 0 {
		return resLogin, e.ERROR_USER_NOT_EXIST
	}
	//验证密码
	if !ComparePwd(user.Password, password) || *user.Role == 0 {
		return resLogin, e.ERROR_AUTH_PARAMS
	}
	//登录成功 生成登录记录
	loginRecordModel.Create(models.LoginRecord{UserID: user.ID, LastLoginIP: ip})
	user.LastLoginIP = ip
	userModel.Save(user)
	//生成token
	claims := util.Claims{UserID: user.ID, Role: user.Role}
	token, err := util.GenerateToken(claims)
	if err != nil {
		return resLogin, e.ERROR_AUTH_TOKEN
	}
	resLogin = util.CopyProperties[response.AdminLogin](user)
	resLogin.Token = token
	return resLogin, e.SUCCESS
}
func (u *UserService) Login(username string, password string, ip string) (resLogin response.Login, code int) {
	//根据名称查询用户
	user := userModel.GetByUsername(username)
	if user.ID == 0 {
		return resLogin, e.ERROR_USER_NOT_EXIST
	}
	//验证密码
	if !ComparePwd(user.Password, password) {
		return resLogin, e.ERROR_AUTH_PARAMS
	}
	//登录成功 生成登录记录
	loginRecordModel.Create(models.LoginRecord{UserID: user.ID, LastLoginIP: ip})
	user.LastLoginIP = ip
	userModel.Save(user)
	//生成token
	claims := util.Claims{UserID: user.ID, Role: user.Role}
	token, err := util.GenerateToken(claims)
	if err != nil {
		return resLogin, e.ERROR_AUTH_TOKEN
	}
	resLogin.User.Nickname = user.Nickname
	resLogin.User.Avatar = user.Avatar
	resLogin.User.Phone = user.Phone
	resLogin.User.ID = user.ID
	resLogin.User.Role = user.Role
	resLogin.User.Status = user.Status
	resLogin.User.Username = user.Username
	resLogin.Token = token
	return resLogin, e.SUCCESS
}
func (u *UserService) Register(username string, password string, nickname string) (user models.User, code int) {
	//检查用户名是否存在
	if exist := u.CheckUserExistByUsername(username); exist {
		return user, e.ERROR_USERNAME_EXIST
	}
	//检查昵称是否存在
	if exist := u.CheckUserExistByNickname(nickname); exist {
		return user, e.ERROR_NICKNAME_EXIST
	}
	encrypted, _ := GetPwd(password)
	user = userModel.Create(models.User{Username: username, Password: string(encrypted), Nickname: nickname})
	return user, e.SUCCESS
}
func (u *UserService) CheckUserExistByUsername(username string) bool {
	user := userModel.GetByUsername(username)
	return user.ID != 0
}
func (u *UserService) CheckUserExistByNickname(nickname string) bool {
	user := userModel.GetByNickname(nickname)
	return user.ID != 0
}
func (u *UserService) GetUserDetailByID(userID uint) (user models.User) {
	return userModel.GetByID(userID)
}
func (u *UserService) GetUserList(req request.PageQuery) (users []models.User, total int) {
	return userModel.GetList(req)
}
func (u *UserService) DeleteUsers(req request.IdsJson) (code int) {
	return userModel.Delete(req)
}
func (u *UserService) UpdateUserStatus(req request.UpdateStatus) (code int) {
	return userModel.UpdateStatus(req)
}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
