package middlewares

import (
	"github.com/gin-gonic/gin"
	"quaver/controller"
	"quaver/pkg/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		//authHeader := c.Request.Header.Get("Authorization")
		//我们设计到查询参数放在URI
		authHeader := c.Query("token")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		//parts := strings.SplitN(authHeader, " ", 2)
		//if !(len(parts) == 2 && parts[0] == "Bearer") {
		//	controller.ResponseError(c, controller.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		//mc, err := jwt.ParseToken(parts[1])
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
