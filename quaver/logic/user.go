package logic

import (
	"quaver/dao/mysql"
	"quaver/models"
	"quaver/pkg/jwt"
)

// Register 用户注册逻辑
func Register(p *models.ParamRegister) (user *models.User, err error) {
	// 1.判断当前用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return nil, err
	}
	// 2.生成userID
	//userID := snowflake.GenID()
	// 构造一个User实例
	user = &models.User{
		Name:     p.Username,
		Password: p.Password,
	}

	// 3.保存进数据库
	userId, err := mysql.Register(user)
	if err != nil {
		return nil, err
	}
	user.ID = userId
	// 生成JWT
	token, err := jwt.GenToken(user.ID, user.Name)
	if err != nil {
		return
	}
	user.Token = token
	return
}

// Login 用户登录逻辑
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Name:     p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.ID, user.Name)
	if err != nil {
		return
	}
	user.Token = token
	return
}

// UserInfo 用户信息
func UserInfo(userID int64, currentUserID int64) (user *models.User, err error) {
	return mysql.UserInfo(userID, currentUserID)
}
