package chat_api

import (
	"encoding/json"
	"fmt"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/utils"
	"net/http"
	"strings"
	"time"

	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

type GroupRequest struct {
	Content string        `json:"content"`  // 内容
	MsgType ctype.MsgType `json:"msg_type"` // 消息类型
}

type GroupResponse struct {
	NickName    string        `json:"nick_name"`  //前端生成的
	Avatar      string        `json:"avatar"`     // 头像
	MsgType     ctype.MsgType `json:"msg_type"`   // 消息类型
	Content     string        `json:"content"`    // 内容
	OnlineCount int           `json:"user_count"` // 聊天室人数
	Data        time.Time     `json:"data"`       // 时间戳
}

func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 将http升级至websocket协议
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("D:\\gin-vue-blog\\gvb_server\\uploads\\chat_avatar\\%s.png", nickNameFirst)

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}

	ConnGroupMap[addr] = chatUser
	// 需要去生成昵称，映射头像地址，根据昵称首字母关连头像地址

	for {
		//消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天室
			SendGroupMsg(conn, GroupResponse{
				Content:     fmt.Sprintf("%s 退出聊天室", chatUser.NickName),
				OnlineCount: len(ConnGroupMap) - 1,
				Data:        time.Now(),
			})
			break
		}
		// 进行参数绑定
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     ctype.SystemMsg,
				OnlineCount: len(ConnGroupMap),
				Content:     "参数绑定失败",
			})
			continue
		}

		//判断类型
		switch request.MsgType {
		case ctype.TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupResponse{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					MsgType:     ctype.SystemMsg,
					OnlineCount: len(ConnGroupMap),
					Content:     "内容不能为空",
				})
				continue
			}
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MsgType:     ctype.TextMsg,
				OnlineCount: len(ConnGroupMap),
				Data:        time.Now(),
			})
		case ctype.InRoomMsg:
			request.Content = fmt.Sprintf("%s 进入聊天室", chatUser.NickName)
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				OnlineCount: len(ConnGroupMap),
				Data:        time.Now(),
			})
		default:
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     ctype.SystemMsg,
				OnlineCount: len(ConnGroupMap),
				Content:     "消息类型错误",
			})
		}

	}
	defer func() {
		conn.Close()
		delete(ConnGroupMap, addr)
	}()
}
func SendGroupMsg(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		MsgType:  response.MsgType,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
	})
	for _, charUser := range ConnGroupMap {
		charUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMsg 给指定用户发送消息
func SendMsg(_addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]

	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		MsgType:  response.MsgType,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	ip = addrList[0]
	addr = utils.GetAddr(ip)
	return addrList[0], addr
}
