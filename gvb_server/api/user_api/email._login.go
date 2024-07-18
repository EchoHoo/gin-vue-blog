package user_api

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/plugins/log_stash"
	"gvb_server/gvb_server/utils"
	"gvb_server/gvb_server/utils/jwts"
	"gvb_server/gvb_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	// 做登录的逻辑
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	log := log_stash.NewLogByGin(c)
	// 验证用户名和密码
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.Username, cr.Username).Error
	if err != nil {
		global.Log.Warn("用户不存在")
		// 记录日志
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.Username))
		//用户不存在
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		//密码错误
		global.Log.Warn("用户密码错误")
		log.Warn(fmt.Sprintf("%s 密码错误", cr.Username))
		res.FailWithMessage("用户名或密码错误", c)

		return
	}
	// 登录成功
	// 生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error("生成token失败", err)
		log.Error(fmt.Sprintf("生成token失败%s", err.Error()))
		res.FailWithMessage("登录失败", c)
		return
	}
	// 返回token

	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(c.ClientIP(), token)
	log.Info(fmt.Sprintf("%s 登录成功", cr.Username))
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)
}
