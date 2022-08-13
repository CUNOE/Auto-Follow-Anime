package core

import (
	"autoFollowAnime/global"
	"autoFollowAnime/model"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

func PostToAria(torrent string, appendParams json.RawMessage, file string) {
	c := resty.New()
	r := model.JsonrpcBack{}
	magnetURL := "magnet:?xt=urn:btih:" + torrent

	var params []interface{}

	if appendParams != nil {
		params = []interface{}{"token:" + global.JsonRpcToken, []string{magnetURL}, appendParams}
	} else {
		params = []interface{}{"token:" + global.JsonRpcToken, []string{magnetURL}}
	}

	body := global.JsonRpcModel
	body.Params = params

	j, _ := json.Marshal(body)

	log.Printf("SubmitBody: %v", string(j))

	_, err := c.R().SetResult(r).SetBody(body).Post(global.JsonRpcServer)
	if err != nil {
		log.Printf("PostToAria Error：%v", err.Error())
		return
	}
	if r.Error != "" {
		log.Printf("PostToAria Error")
		AddToDatabase(torrent, 0)
		return
	}

	AddToDatabase(torrent, 1)

	WriteMessage("您追的动漫更新辣！\n" + file)
	time.Sleep(1 * time.Second)

	return

}
