package middleware

import (
	"fmt"
	"micro-memorandum/pkg/ctl"
	"micro-memorandum/pkg/e"
	"micro-memorandum/pkg/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		var claims *jwt.Claims
		var err error

		// 没有token
		if token == "" {
			code = 404
		} else {
			claims, err = jwt.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if claims.ExpiresAt.Before(time.Now()) {
				// 过期
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			// fmt.Println("来自middleware的json返回")
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		fmt.Println("获取claims：", claims)

		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		c.Next()
	}
}
