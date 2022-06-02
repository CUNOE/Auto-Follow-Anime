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

	go wsServer()
	c := cron.New()
	_, err := c.AddFunc("@hourly", checkRss)
	if err != nil {
		log.Fatal("Add CheeckRss Func To Cron error")
		return
	}
	c.Start()

}

func checkRss() {
	for _, rss := range global.RssAddress {
		hashs, files := core.GetTorrentUri(rss.Rss)
		for i, _ := range hashs {
			if core.IsDownloaded(hashs[i]) {
				log.Printf("Already Downloaded: %v", hashs[i])
			} else {
				core.PostToAria(hashs[i], rss.AppendParams, files[i])
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
