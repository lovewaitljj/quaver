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
