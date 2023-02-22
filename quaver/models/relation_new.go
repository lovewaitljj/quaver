package models

/*message douyin_relation_action_reqauest {
required string token = 1; // 用户鉴权token
required int64 to_user_id = 2; // 对方用户id
required int32 action_type = 3; // 1-关注，2-取消关注*/

func (u Attention) TableName() string {
	return "attention"
}

type Attention struct {
	//Token      string `json:"token"`
	//ToUserId   int64  `json:"to_user_id"`
	//ActionType int32  `json:"action_type"`
	Token      string `form:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" binding:"required"`
	ActionType int32  `form:"action_type" binding:"required"`
}

func (u CommentList) TableName() string {
	return "commentList"
}

//用户关注列表
type CommentList struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserInfo struct {
	ID              int64  `form:"id" binding:"required"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow" gorm:"-"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}
