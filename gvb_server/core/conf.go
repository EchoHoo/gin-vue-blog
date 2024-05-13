package core

import (
	"fmt"
	"gvb_server/gvb_server/config"
	"gvb_server/gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "D:\\gin-vue-blog\\gvb_server\\settings.yaml"

// 读取配置 yaml文件的配置
func InitConf() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, &c)
	if err != nil {
		log.Fatal("config Init Unmaershal", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	//如果您的 byteData 恰好与 YAML 文件中某一配置项的大小相等
	//并且您使用 ioutil.WriteFile 写入到文件中，那么它会精确地覆盖该配置项的所有内容。
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
