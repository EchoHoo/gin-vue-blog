package redis_ser

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/utils"
	"time"
)

const prefix = "logout_"

func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", prefix, token), "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	//判断是否在redis中存在，存在就是注销过了
	keys := global.Redis.Keys(prefix + "*").Val()
	return utils.InList(prefix+token, keys)
}
