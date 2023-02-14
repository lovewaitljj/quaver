package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quaver/controller"
	"quaver/logger"
	"quaver/middlewares"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//r.Use(middlewares.Cors()) // 后端解决跨域问题
	apiRouter := r.Group("/douyin")
	{
		apiRouter.POST("/user/register/", controller.Register) // 用户注册
		apiRouter.POST("/user/login/", controller.Login)       // 用户登录
		apiRouter.GET("/feed/", controller.Feed)               // 视频流接口
	}
	apiRouter.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		apiRouter.GET("/user/", controller.UserInfo)            // 用户信息
		apiRouter.GET("/publish/list/", controller.PublishList) // 发布列表
		apiRouter.POST("/publish/action/", controller.Publish)  // 发布视频

		apiRouter.POST("/relation/action/", controller.RelationAction) // 关注操作
		apiRouter.GET("/relation/follow/list/", controller.FollowList) // 关注列表

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
