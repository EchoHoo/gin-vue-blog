package user_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/user_ser"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	Nickname string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  //昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` //用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   //密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       //权限 1、管理员 2、普通用户 3、游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_ser.UserService{}.CreateUser(cr.Nickname, cr.UserName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	res.OKWithMessage(fmt.Sprintf("用户 %s 创建成功", cr.UserName), c)

}
