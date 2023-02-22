package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quaver/models"
)

type RelationPonse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommenListPonse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	User       models.UserInfo `json:"user"`
}

func RelationPonseSuccess(c *gin.Context, StatusMsg string) {
	c.JSON(http.StatusOK, &RelationPonse{
		StatusCode: int32(CodeSuccess),
		StatusMsg:  CodeSuccess.Msg(),
	})
}

func RelationPonseFaild(c *gin.Context, StatusMsg string) {
	c.JSON(http.StatusNotFound, &RelationPonse{
		StatusCode: int32(CodeServerBusy),
		StatusMsg:  CodeServerBusy.Msg(),
	})
}

func CommenListPonseSuccess(c *gin.Context, StatusMsg string, q *models.UserInfo) {
	c.JSON(http.StatusOK, &CommenListPonse{
		StatusCode: int32(CodeSuccess),
		StatusMsg:  CodeSuccess.Msg(),
		User: models.UserInfo{
			ID:            q.ID,
			Name:          q.Name,
			FollowCount:   q.FollowCount,
			FollowerCount: q.FollowerCount,
			IsFollow:      q.IsFollow,
		},
	})
}

func CommenListPonseFaild(c *gin.Context, StatusMsg string) {
	c.JSON(http.StatusNotFound, &CommenListPonse{
		StatusCode: int32(CodeServerBusy),
		StatusMsg:  CodeServerBusy.Msg(),
	})
}
