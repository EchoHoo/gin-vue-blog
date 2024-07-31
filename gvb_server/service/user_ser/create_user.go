package user_ser

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/utils"
	"gvb_server/gvb_server/utils/pwd"
)

const Avatar = "/uploads/avatar/default.jpg"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	//创建用户的逻辑
	var userModel models.UserModel
	//判断用户名是否存在
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		//查到了
		global.Log.Error("用户名已存在,请重新输入...")
		return err
	}
	//对密码进行加密
	hashPwd := pwd.HashPwd(password)

	//头像问题
	avatar := Avatar

	addr := utils.GetAddr(ip)
	fmt.Println("addr:-->>", addr)
	//入库
	err = global.DB.Create(&models.UserModel{
		UserName:   userName,
		NickName:   nickName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error("用户创建失败...")
		return err
	}
	global.Log.Info("用户创建成功...")
	return nil
}
