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

// ParamFeed 获取用户消息请求参数
type ParamFeed struct {
	LatestTime string `form:"latest_time"`
	Token      string `form:"token"`
}

// ParamPublishList 获取用户消息请求参数
type ParamPublishList struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// ParamPublish 发布视频请求参数
type ParamPublish struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}
