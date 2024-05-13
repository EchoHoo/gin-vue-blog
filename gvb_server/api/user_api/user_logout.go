package user_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OKWithMessage("注销成功", c)
}
