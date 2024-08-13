package user_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils/jwts"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
}

func (UserApi) UserUpdateNicknameView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr UserUpdateNicknameRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var newMaps = map[string]any{}
	maps := structs.Map(&cr)
	for k, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[k] = val
		}
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&userModel).Updates(newMaps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("更新失败", c)
		return
	}
	res.OKWithMessage("修改个人信息成功", c)
}
