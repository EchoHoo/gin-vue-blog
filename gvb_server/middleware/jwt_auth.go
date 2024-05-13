package middleware

import (
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/redis_ser"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			//没有登录过
			res.FailWithMessage("未携带tokne", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token无效", c)
			c.Abort()
			return
		}
		//判断是否在redis中存在，存在就是注销过了
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已经失效", c)
			c.Abort()
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//如何判断是管理员
		token := c.Request.Header.Get("token")
		if token == "" {
			//没有登录过
			res.FailWithMessage("未携带tokne", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token无效", c)
			c.Abort()
			return
		}
		//判断是否在redis中存在，存在就是注销过了
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已经失效", c)
			c.Abort()
			return
		}

		//登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
