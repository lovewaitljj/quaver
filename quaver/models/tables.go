package models

// User 用户信息
type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	Token         string `json:"token,omitempty" gorm:"-"`
	IsFollow      bool   `json:"is_follow" gorm:"-"`
}

// Follow 关注信息
type Follow struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	FollowID int64 `json:"follow_id"`
	IsFollow int8  `json:"is_follow"`
}

// Video 视频信息
type Video struct {
	ID            int64  `json:"id"`
	UserID        int64  `json:"-"`
	Title         string `json:"title"`
	CreateTime    string `json:"-"`
	Author        User   `json:"author" gorm:"-"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" gorm:"-"`
}

// Like 视频点赞信息
type Like struct {
	ID      int64 `json:"id"`
	VideoID int64 `json:"video_id"`
	UserID  int64 `json:"user_id"`
	IsLike  int64 `json:"is_like"`
}
