package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"quaver/dao/mysql"
	"quaver/models"
)

func OperaAttention(p *models.Attention) (state int32, err error) {
	//业务逻辑
	attention := &models.Attention{
		Token:      p.Token,
		ToUserId:   p.ToUserId,
		ActionType: p.ActionType,
	}
	state = p.ActionType
	if state != 1 || state != 2 { //增加判断，如果不是点击关注或者是取消关注则退出
		return state, errors.New("err")
	}
	if err := mysql.Attention(attention); err != nil {
		fmt.Println(err)
		return state, err
	}
	return state, nil
}

//MatchInfo 关注列表匹配
func MatchInfo(p *models.CommentList, q *models.UserInfo) (token string, err error) {
	commen := &models.CommentList{
		UserId: p.UserId,
		Token:  p.Token,
	}
	user := &models.UserInfo{
		ID: q.ID,
	}
	if commen.UserId == user.ID {
		return commen.Token, nil
	}
	return "", errors.New("user error")
}

//CommenInfoList 关注列表查询
func CommenInfoList(s string) (ok int64) {
	//带token到mysql中进行查找，找到返回ok，否则返回失败
	Commen := make([]int64, 1)
	Commen = mysql.QuaryList(s)

	fmt.Println(Commen)
	return 1
}

//ToUser 根据自身ID以及鉴权查找关注用户ID
func ToUser(id int64, token string) (toId int64) {
	//判断id是否存在，存在再查询关注用户token，否则返回
	//用户输入不可信，需要从数据库验算
	searchtoken := mysql.UserExist(id)
	if searchtoken == "" {
		zap.L().Error("用户不存在")
		return -1
	}
	if searchtoken != token {
		zap.L().Error("用户篡改数据")
		return -1
	}
	//找到对应关注者id
	toId = mysql.SearchToUserId(token)
	return toId
}

//SearchFansList 查找用户粉丝列表
func SearchFansList() {

}
