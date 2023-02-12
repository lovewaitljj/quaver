package models

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ResponseSignUp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type ResponseLogin struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type ResponseUserInfo struct {
	Response
	User User `json:"user"`
}

type ResponseFeed struct {
	Response
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}
type ResponsePublishList struct {
	Response
	VideoList []Video `json:"video_list"`
}
