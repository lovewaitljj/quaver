package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"quaver/dao/mysql"
	"quaver/logic"
	"quaver/models"
)

// Register 用户注册
func Register(c *gin.Context) {
	p := new(models.ParamRegister)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	if user, err := logic.Register(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		//c.JSON(http.StatusOK, models.ResponseSignUp{
		//	Response: models.Response{
		//		StatusCode: 0,
		//		StatusMsg:  "success",
		//	},
		//	UserId: user.ID,
		//	Token:  user.Token,
		//})
		ResponseSuccess(c, models.ResponseSignUp{
			Response: models.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			UserId: user.ID,
			Token:  user.Token,
		})
	}
}

// Login 用户登录
func Login(c *gin.Context) {
	p := new(models.ParamLogin)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, models.ResponseLogin{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: user.ID,
		Token:  user.Token,
	})
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	p := new(models.ParamUserInfo)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("UserInfo with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	currentUserID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	user, err := logic.UserInfo(p.UserID, currentUserID)
	if err != nil {
		zap.L().Error("logic.UserInfo failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, models.ResponseUserInfo{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: models.User{
			ID:            user.ID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
	})
}
