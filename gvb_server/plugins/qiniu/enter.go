package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gvb_server/gvb_server/config"
	"gvb_server/gvb_server/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = false
	//上传是否使用cdn加速
	cfg.UseCdnDomains = false
	return cfg
}

// UploadImage 上传图片， 文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启动qiniu云")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey与secretKey")
	}
	//控制上床七牛的图片大小
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过设定大小")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s.png", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
