package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"quaver/logic"
	"quaver/models"
)

// RelationAction 关注操作
func RelationAction(c *gin.Context) {
	p := new(models.ParamRelationAction)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("RelationAction with invalid param", zap.Error(err))
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
	if err := logic.RelationAction(currentUserID, p); err != nil {
		zap.L().Error("logic.RelationAction failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		})
	}
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	p := new(models.ParamFollowList)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("FollowList with invalid param", zap.Error(err))
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
	if followList, err := logic.FollowList(p.UserId); err != nil {
		zap.L().Error("logic.FollowList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		c.JSON(http.StatusOK, models.ResponseFollowList{
			Response: models.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			UserList: followList,
		})
	}
}
