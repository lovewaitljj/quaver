package mysql

import (
	"errors"
	"gorm.io/gorm"
	"quaver/models"
	"quaver/settings"
)

// Feed 视频流接口
func Feed(latestTime string, currentUserID ...int64) (videoList []*models.Video, err error) {
	videoList = make([]*models.Video, 30)
	// 从满足条件的id往前查找
	if err = db.Order("id desc").Where("create_time < ?", latestTime).Limit(30).Find(&videoList).Error; err != nil {
		return nil, err
	}

	for k, video := range videoList {
		videoList[k].PlayUrl = settings.Conf.Url + videoList[k].PlayUrl
		videoList[k].CoverUrl = settings.Conf.Url + videoList[k].CoverUrl
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
			if followed || user.ID == currentUserID[0] {
				user.IsFollow = true
			}
			// 判断currentUser是否点赞该视频
			like, err := isLike(currentUserID[0], video.ID)
			if err != nil {
				return nil, err
			}
			if like {
				videoList[k].IsFavorite = true
			}
		}
		videoList[k].Author = *user // 这里不能直接用video.Author = *user,因为video是值拷贝...
	}
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
func PublishList(userID ...int64) (publishList []*models.Video, err error) {
	publishList = make([]*models.Video, 30)
	user := new(models.User)
	if err = db.Where("user_id=?", userID[0]).Find(&publishList).Error; err != nil {
		return nil, err
	}
	if err = db.Select("id", "name", "follow_count", "follower_count").Where("id=?", userID[0]).Find(user).Error; err != nil {
		return nil, err
	}
	if len(userID) == 1 { // 查询自己的发布列表
		for k, video := range publishList { // 遍历currentUser是否点赞该视频
			if like, err := isLike(userID[0], video.ID); err != nil {
				return nil, err
			} else if like {
				publishList[k].IsFavorite = true
			}

			publishList[k].CoverUrl = settings.Conf.Url + publishList[k].CoverUrl
			publishList[k].PlayUrl = settings.Conf.Url + publishList[k].PlayUrl
			publishList[k].Author = *user
		}
	} else { // 别人的发布列表
		// 当前用户是否关注传来的user_id参数
		if followed, err := isFollow(userID[0], userID[1]); err != nil {
			return nil, err
		} else if followed {
			user.IsFollow = true
		}

		for k, video := range publishList { // 遍历currentUser是否点赞该视频
			if like, err := isLike(userID[1], video.ID); err != nil {
				return nil, err
			} else if like {
				publishList[k].IsFavorite = true
			}
			publishList[k].CoverUrl = settings.Conf.Url + publishList[k].CoverUrl
			publishList[k].PlayUrl = settings.Conf.Url + publishList[k].PlayUrl
			publishList[k].Author = *user
		}
	}
	return publishList, err
}

// Publish 发布视频
func Publish(video *models.Video) (err error) {
	return db.Create(video).Error
}

// DoFavorite 点赞操作
func DoFavorite(userID int64, p *models.Like) (err error) {
	follow := new(models.Like)
	//查询该likes表中videosId对应的点赞人有无此id
	if errors.Is(db.Where("video_id = ? and user_id = ?", p.VideoID, userID).First(&follow).Error, gorm.ErrRecordNotFound) {
		//如果没点过，则新加一个字段：
		db.Model(&models.Video{}).Where("id = ?", p.VideoID).Update("favorite_count", gorm.Expr("favorite_count + ? ", 1))
		return db.Create(&p).Error
	}
	//如果点过，则进行改变
	if follow.IsLike == 1 {
		db.Model(&models.Video{}).Where("id = ?", p.VideoID).Update("favorite_count", gorm.Expr("favorite_count - ? ", 1))
		db.Model(&models.Like{}).Where("video_id = ? and user_id = ?", p.VideoID, userID).Update("is_like", 2)
		return
	}
	db.Model(&models.Video{}).Where("id = ?", p.VideoID).Update("favorite_count", gorm.Expr("favorite_count + ? ", 1))
	db.Model(&models.Like{}).Where("video_id = ? and user_id = ?", p.VideoID, userID).Update("is_like", 1)
	return
}

// FavoriteList 喜欢列表
func FavoriteList(userID ...int64) (favoriteList []*models.Video, err error) {
	//favoriteList = make([]*models.Video, 0)
	//1.连接like和video表查出video相关信息
	if err = db.Raw("SELECT v.id,v.user_id,v.title,v.play_url,v.cover_url,v.favorite_count,v.comment_count"+
		" FROM videos v LEFT JOIN likes l ON v.id=l.video_id where l.user_id = ? and is_like = 1", userID[0]).Scan(&favoriteList).Error; err != nil {
		return nil, err
	}
	for k, video := range favoriteList {
		//2.根据查出来的视频列表的作者去user表查出作者相关信息
		user := new(models.User)
		if err = db.Select("id", "name", "follow_count", "follower_count").Where("id = ?", video.UserID).Find(&user).Error; err != nil {
			return nil, err
		}
		//3.判断currentId是否关注视频作者
		if followed, err := isFollow(user.ID, userID[0]); err != nil {
			return nil, err
		} else if followed {
			user.IsFollow = true
		}
		favoriteList[k].CoverUrl = settings.Conf.Url + favoriteList[k].CoverUrl
		favoriteList[k].PlayUrl = settings.Conf.Url + favoriteList[k].PlayUrl
		favoriteList[k].Author = *user
		favoriteList[k].IsFavorite = true
	}
	return favoriteList, err
}

// DoComment 发布评论
func DoComment(comment *models.Comment) (user *models.User, err error) {
	err = db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	//获取用户信息
	db.Select("id,name,follow_count,follower_count").Where("id=?", comment.UserID).First(&user)
	return
}

//DelComment 删除评论
func DelComment(commentId int64) error {
	return db.Delete(&models.Comment{}, commentId).Error
}
