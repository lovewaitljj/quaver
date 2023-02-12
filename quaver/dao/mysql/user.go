package mysql

import (
	"errors"
	"gorm.io/gorm"
	"quaver/models"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	if errors.Is(db.Where("name = ?", username).First(&models.User{}).Error, gorm.ErrRecordNotFound) {
		return
	}
	//sqlStr := `select count(id) from users where username = ? `
	//var count int
	//if err = db.Get(&count, sqlStr, username); err != nil {
	//	return err
	//}
	//if count > 0 {
	//	return ErrorUserExist
	//}
	return ErrorUserExist
}

// Register 保存进数据库
func Register(user *models.User) (userId int64, err error) {
	err = db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	db.Where("name = ?", user.Name).First(&user)
	return user.ID, err
}

// Login 用户登录数据库校验
func Login(user *models.User) (err error) {
	oPassword := user.Password
	// 判断用户是否存在
	if errors.Is(db.Where("name = ?", user.Name).First(&user).Error, gorm.ErrRecordNotFound) {
		return ErrorUserNoExist
	}
	// 判断密码是否正确
	if oPassword != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// UserInfo 用户信息
func UserInfo(userID int64, currentUserID int64) (user *models.User, err error) {
	// 判断用户是否存在
	if errors.Is(db.Where("id = ?", userID).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, ErrorUserNoExist
	}
	// 判断currentUser是否关注user
	followed, err := isFollow(userID, currentUserID)
	if err != nil {
		return nil, err
	}
	if followed {
		user.IsFollow = true
		return user, nil
	}
	return user, nil
}

// isFollow currentUserID是否关注userID
func isFollow(userID int64, currentUserID int64) (followed bool, err error) {
	follow := new(models.Follow)
	if errors.Is(db.Where("user_id = ? and follow_id = ?", currentUserID, userID).First(&follow).Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if follow.IsFollow == 1 {
		return true, nil
	}
	return false, nil
}
