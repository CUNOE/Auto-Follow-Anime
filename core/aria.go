package core

import (
	"autoFollowAnime/global"
	"autoFollowAnime/model"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

func PostToAria(torrent string, appendParams json.RawMessage) {
	c := resty.New()
	r := model.JsonrpcBack{}
	metalink := "magnet:?xt=urn:btih:" + torrent

	var params []interface{}

	if appendParams != nil {
		params = []interface{}{"token:" + global.JsonRpcToken, []string{metalink}, appendParams}
	} else {
		params = []interface{}{"token:" + global.JsonRpcToken, []string{metalink}}
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

	j, _ = json.Marshal(r)

	log.Printf("ResponseResult: %v", string(j))
	AddToDatabase(torrent, 1)
	return

}
