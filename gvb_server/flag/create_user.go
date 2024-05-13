package flag

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/service/user_ser"
)

func CreateUser(permission string) {
	//创建用户的逻辑
	//用户名 昵称 密码 确认密码 邮箱
	var (
		userName        string
		nickName        string
		password        string
		confirmPassword string
		email           string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱：")
	fmt.Scanln() // 清空缓冲区
	fmt.Scanln(&email)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请输入确认密码：")
	fmt.Scan(&confirmPassword)

	var userModel models.UserModel
	//判断用户名是否存在
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		//查到了
		global.Log.Error("用户名已存在,请重新输入...")
		return
	}
	//判断密码是否一致
	if password != confirmPassword {
		global.Log.Error("两次密码输入不一致,请重新输入...")
		return
	}

	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	err = user_ser.UserService{}.CreateUser(
		userName,
		nickName,
		password,
		role,
		email,
		"127.0.0.1",
	)

	if err != nil {
		global.Log.Error("用户创建失败...")
		return
	}
	global.Log.Info("用户创建成功...")
}
