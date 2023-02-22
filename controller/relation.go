package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"quaver/logic"
	"quaver/models"
)

const (
	Attention    = 1
	NotAttention = 2
	Success      = 3
)

//Relation 登录用户对其他用户进行关注或取消关注。
func Relation(c *gin.Context) {
	//1.获取参数
	p := new(models.Attention)
	//校验参数
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Relation bind err")
		return
	}
	//2.业务处理
	if state, err := logic.OperaAttention(p); err != nil {
		//逻辑点击失败，返回响应错误
		//3.返回响应
		zap.L().Error("关注失败", zap.Error(err))
		RelationPonseFaild(c, "关注失败，传输值出问题")
		return
	} else {
		//关注成功或者取消关注成功
		if state == Attention {
			//关注成功
			zap.L().Error("关注成功")
			RelationPonseSuccess(c, "关注成功")
		}
		if state == NotAttention {
			//取消关注成功
			zap.L().Error("取消关注成功")
			RelationPonseSuccess(c, "取消关注成功")
		}
		//关注类型错误，返回错误响应
	}

}

func ListManger(c *gin.Context) {
	//1.获取参数，校验参数
	//绑定关注列表
	p := new(models.CommentList)
	q := new(models.UserInfo)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("bind CommentList faild")
		return
	}
	if err := c.ShouldBind(q); err != nil {
		zap.L().Error("bind UserInfo faild")
		return
	}
	//2.业务处理
	//判断id信息是否一致，是否为同一用户
	token, err := logic.MatchInfo(p, q)
	if err != nil {
		//3.返回响应
		//id信息不一致，不可查询
		CommenListPonseFaild(c, "查询失败")
		return
	}
	ok := logic.CommenInfoList(token) //带着token查找关注列表
	fmt.Println("ok", ok)
	if ok != Success {
		CommenListPonseSuccess(c, "查询成功", q)
		//找到对应列表
		return
	}

}

//UserFans 用户粉丝列表
/*分析：每一个粉丝也为一个单独的用户，根据当前用户找到数据鉴权，
再根据数据鉴权找到关注的人，通过关注的人可以查看对应的信息，
所关注的人也是一个单独的User
*/
func UserFans(c *gin.Context) {
	//1.获取参数
	//var user models.ParamUserInfo
	user := new(models.ParamUserInfo)
	if err := c.ShouldBind(user); err != nil {
		zap.L().Error("绑定失败")
		return
	}
	fmt.Println(user.UserID, user.Token)
	//2.业务处理
	//利用获取到的用户id以及token鉴权查找关注的id
	toUserId := logic.ToUser(user.UserID, user.Token)
	//根据获取到的用户ID查询对应粉丝用户
	fmt.Println(toUserId)
	fansList := logic.SearchFansList(toUserId)
	/*
		userId, err := getCurrentUserID(c) //获取当前用户id
		if err != nil {
			zap.L().Error("获取本机自身id失败")
			fmt.Println("本机自身id失败为", userId)
			return
		}
		fmt.Println("本机自身id成功为", userId)
	*/

	//需要获取当前用户的关注用户，再从关注用户获取其对应信息，找出其粉丝列表

	//3.返回响应

}
