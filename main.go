package main

import (
	"autoFollowAnime/core"
	"autoFollowAnime/global"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	global.LoadConfig()
	global.InitSqlite()

	c := cron.New()
	_, err := c.AddFunc(global.Crontab, checkRss)
	if err != nil {
		log.Fatal("Add CheckRss Func To Cron error")
		return
	}
	c.Start()

	wsServer()

}

func checkRss() {
	for _, rss := range global.RssAddress {
		hashes, files := core.GetTorrentUri(rss.Rss)
		for i, _ := range hashes {
			if core.IsDownloaded(hashes[i]) {
				log.Printf("Already Downloaded: %v", hashes[i])
			} else {
				core.PostToAria(hashes[i], rss.AppendParams, files[i])
			}
		}
	}
}

func wsServer() {

	r := gin.Default()
	r.GET("/ws", core.ReadMessage)
	err := r.Run(":" + global.WsPort)
	if err != nil {
		log.Printf("Error To Run Ws Server! Cannot send message via ws")
	}
	log.Printf("listen in %v", global.WsPort)
}
