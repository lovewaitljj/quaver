package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
	"quaver/logic"
	"quaver/models"
	"quaver/pkg/jwt"
	"time"
)

// Feed 视频流接口
func Feed(c *gin.Context) {
	p := new(models.ParamFeed)
	// 1. 获取参数和校验参数
	err := c.ShouldBind(p)
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Feed with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2. 业务处理
	var videoList *[]models.Video
	if p.Token == "" {
		videoList, err = logic.Feed(p.LatestTime)
	} else {
		mc, err := jwt.ParseToken(p.Token) // 从token中解析出UserID
		if err != nil {
			ResponseError(c, CodeInvalidToken)
			return
		}
		videoList, err = logic.Feed(p.LatestTime, mc.UserID)
		if err != nil {
			zap.L().Error("logic.Feed failed", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	list := *videoList
	latestTime := time.Now()
	// 解析字符串格式的时间
	if len(list) > 0 {
		loc, _ := time.LoadLocation("Local")
		latestTime, _ = time.ParseInLocation("2006-01-02T15:04:05Z", list[0].CreateTime, loc)
		//latestTime, _ = time.Parse("2006-01-02T15:04:05Z", list[0].CreateTime) // 会加8小时
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, models.ResponseFeed{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		NextTime:  latestTime.Unix(),
		VideoList: list,
	})
}

// PublishList 发布列表
func PublishList(c *gin.Context) {
	p := new(models.ParamPublishList)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("PublishList with invalid param", zap.Error(err))
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
	publishList, err := logic.PublishList(currentUserID, p.UserID)
	if err != nil {
		zap.L().Error("logic.PublishList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, models.ResponsePublishList{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: *publishList,
	})
}

// Publish 发布视频
func Publish(c *gin.Context) {
	p := new(models.ParamPublish)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Publish with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		zap.L().Error("load file data failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2. 业务处理
	currentUserID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 生成文件名和保存路径
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", currentUserID, filename)
	saveFile := filepath.Join("./public/", finalName)
	// 保存文件
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		zap.L().Error("SaveUploadedFile failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	err = logic.Publish(currentUserID, p.Title, saveFile, finalName)
	if err != nil {
		zap.L().Error("logic.Publish failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, models.ResponsePublishList{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
	})
}
