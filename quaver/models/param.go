package models

// ParamRegister 注册请求参数
type ParamRegister struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// ParamUserInfo 获取用户消息请求参数
type ParamUserInfo struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// ParamFeed 视频流请求参数
type ParamFeed struct {
	LatestTime string `form:"latest_time"`
	Token      string `form:"token"`
}

// ParamPublishList 发布列表请求参数
type ParamPublishList struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// ParamPublish 发布视频请求参数
type ParamPublish struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}

// ParamFavorite 赞操作请求参数
type ParamFavorite struct {
	Token      string `form:"token" binding:"required"`
	VideoID    int64  `form:"video_id" binding:"required"`
	ActionType int32  `form:"action_type" binding:"required"` //1-点赞，2-取消点赞
}

// ParamRelationAction 关注操作请求参数
type ParamRelationAction struct {
	Token      string `form:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" binding:"required"`
	ActionType int32  `form:"action_type" binding:"required"`
}

// ParamFollowList 关注操作请求参数
type ParamFollowList struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// ParamFavoriteList 发布列表请求参数
type ParamFavoriteList struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// ParamComment 评论请求参数
type ParamComment struct {
	Token       string `form:"token" binding:"required"`
	VideoID     int64  `form:"video_id" binding:"required"`
	ActionType  int32  `form:"action_type" binding:"required"` //1-发布评论，2-删除评论
	CommentText string `form:"comment_text"`                   //发布评论时使用
	Comment_id  int64  `form:"comment_id"`                     //删除评论时使用
}

// ParamCommentList 评论列表参数
type ParamCommentList struct {
	Token   string `form:"token" binding:"required"`
	VideoID int64  `form:"video_id" binding:"required"`
}
