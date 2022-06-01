package main

import (
	"autoFollowAnime/core"
	"autoFollowAnime/global"
	"log"
)

func main() {
	global.LoadConfig()
	global.InitSqlite()

	for _, rss := range global.RssAddress {
		hashs := core.GetTorrentUri(rss.Rss)
		for i, _ := range hashs {
			if core.IsDownloaded(hashs[i]) {
				log.Printf("Already Downloaded: %v", hashs[i])
			} else {
				core.PostToAria(hashs[i], rss.AppendParams)
			}
		}
	}

}
