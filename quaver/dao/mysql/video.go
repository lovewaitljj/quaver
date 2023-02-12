package mysql

import (
	"errors"
	"gorm.io/gorm"
	"quaver/models"
)

// Feed 视频流接口
func Feed(latestTime string, currentUserID ...int64) (videoList *[]models.Video, err error) {
	videos := make([]models.Video, 30)
	// 从满足条件的id往前查找
	if err = db.Order("id desc").Where("create_time < ?", latestTime).Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}

	for k, video := range videos {
		user := new(models.User)
		if err = db.Select("id", "name", "follow_count", "follower_count").Where("id=?", video.UserID).Find(user).Error; err != nil {
			return nil, err
		}
		if len(currentUserID) != 0 {
			// 判断currentUser是否关注user
			followed, err := isFollow(user.ID, currentUserID[0])
			if err != nil {
				return nil, err
			}
			if followed {
				user.IsFollow = true
			}
			// 判断currentUser是否点赞该视频
			like, err := isLike(currentUserID[0], video.ID)
			if err != nil {
				return nil, err
			}
			if like {
				videos[k].IsFavorite = true
			}
		}
		videos[k].Author = *user // 这里不能直接用video.Author = *user,因为video是值拷贝...
	}
	videoList = &videos
	return videoList, err
}

// isLike currentUser是否点赞视频
func isLike(currentUserID int64, videoID int64) (followed bool, err error) {
	follow := new(models.Like)
	if errors.Is(db.Where("video_id = ? and user_id = ?", videoID, currentUserID).First(&follow).Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if follow.IsLike == 1 {
		return true, nil
	}
	return false, nil
}

// PublishList 发布列表
func PublishList(userID ...int64) (publishList *[]models.Video, err error) {
	videos := make([]models.Video, 30)
	user := new(models.User)
	if err = db.Where("user_id=?", userID[0]).Find(&videos).Error; err != nil {
		return nil, err
	}
	if err = db.Select("id", "name", "follow_count", "follower_count").Where("id=?", userID[0]).Find(user).Error; err != nil {
		return nil, err
	}
	if len(userID) == 1 { // 查询自己的发布列表
		for k, video := range videos { // 遍历currentUser是否点赞该视频
			if like, err := isLike(userID[0], video.ID); err != nil {
				return nil, err
			} else if like {
				videos[k].IsFavorite = true
			}
			videos[k].Author = *user
		}
	} else { // 别人的发布列表
		// 当前用户是否关注传来的user_id参数
		if followed, err := isFollow(userID[0], userID[1]); err != nil {
			return nil, err
		} else if followed {
			user.IsFollow = true
		}

		for k, video := range videos { // 遍历currentUser是否点赞该视频
			if like, err := isLike(userID[1], video.ID); err != nil {
				return nil, err
			} else if like {
				videos[k].IsFavorite = true
			}
			videos[k].Author = *user
		}
	}
	publishList = &videos
	return publishList, err
}

// Publish 发布视频
func Publish(video *models.Video) (err error) {
	return db.Create(video).Error
}
