package user_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils/jwts"
	"gvb_server/gvb_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

// 修改登录人的ID
type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"旧密码不能为空"`
	Pwd    string `json:"pwd" binding:"required" msg:"新密码不能为空"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	//判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	//更新密码
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		res.FailWithMessage("修改密码失败", c)
		return
	}
	res.OKWithMessage("修改密码成功", c)

}
