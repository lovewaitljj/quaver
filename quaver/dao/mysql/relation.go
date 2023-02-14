package mysql

import (
	"errors"
	"gorm.io/gorm"
	"quaver/models"
)

// RelationAction 关注操作
func RelationAction(follow *models.Follow) (err error) {
	followed := new(models.Follow)
	err = db.Where("user_id = ? and follow_id = ?", follow.UserID, follow.FollowID).First(followed).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(follow).Error
	}
	followed.IsFollow = follow.IsFollow
	return db.Model(followed).Update("is_follow", followed.IsFollow).Error
}

// FollowList 关注列表
func FollowList(userID int64) (userList []*models.User, err error) {
	follows := make([]models.Follow, 10)
	userList = make([]*models.User, 10)
	followID := make([]int64, 10)
	// 从follows表查出关注用户id
	if err = db.Select("follow_id").Where("user_id = ? and is_follow = 1", userID).Find(&follows).Error; err != nil {
		return
	}
	// 根据用户id从users表查出用户信息
	for i, follow := range follows {
		followID[i] = follow.FollowID
	}
	if err = db.Select("id,name,follow_count,follower_count").Where("id IN ?", followID).Find(&userList).Error; err != nil {
		return
	}
	for i := range userList {
		userList[i].IsFollow = true
	}
	return userList, err
}
