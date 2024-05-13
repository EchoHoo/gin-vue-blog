package logTest

import (
	"fmt"
	"gvb_server/gvb_server/core"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/plugins/log_stash"
	"testing"
)

func TestLog(t *testing.T) {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	log := log_stash.New("192.168.1.100", "xxx")
	log.Error(fmt.Sprintf("%s hello ", "gch"))
}
