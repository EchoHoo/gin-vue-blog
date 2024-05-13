package redis_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

// Set设置某一个id的点赞数
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err()
	return err
}
func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return newData, err
}
