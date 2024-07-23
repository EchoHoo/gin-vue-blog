package user_api

import (
	"fmt"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/common"
	"gvb_server/gvb_server/utils/desense"
	"gvb_server/gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}
type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []UserResponse
	fmt.Println(page.Page, page.Limit)
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//这就是管理员
			user.UserName = ""
		}
		//脱敏手机号
		user.Tel = desense.DesensitizationTel(user.Tel)
		//脱敏邮箱
		user.Email = desense.DesensitizationEmail(user.Email)
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)

}
