package core

import (
	"autoFollowAnime/global"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type PostSendMsg struct {
	Action string      `json:"action"`
	Params *ParamsData `json:"params"`
	Echo   string      `json:"echo"`
}

type ParamsData struct {
	GroupId int64  `json:"group_id"`
	Message string `json:"message"`
}

type GroupMsg struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int         `json:"font"`
	GroupId     int64       `json:"group_id"`
	Message     string      `json:"message"`
	MessageId   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfId      int         `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserId   int64  `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserId  int64  `json:"user_id"`
}

func ReadMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf(err.Error())
	}
	global.Conn = conn

	for {
		var msg GroupMsg
		err := conn.ReadJSON(&msg)
		if err != nil {
			return
		}

		if msg.Message == "测试自动追番机器人" {
			WriteMessage("机器人正常部署")
		}
	}

}

func WriteMessage(message string) {
	var msg *PostSendMsg
	data := &ParamsData{
		GroupId: global.QQGroupId,
		Message: message,
	}
	msg = &PostSendMsg{
		Action: "send_msg",
		Params: data,
	}
	err := global.Conn.WriteJSON(msg)
	if err != nil {
		return
	}
}
