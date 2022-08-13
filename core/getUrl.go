package core

import (
	"github.com/beevik/etree"
	"github.com/go-resty/resty/v2"
	"log"
	"regexp"
)

func GetTorrentUri(uri string) (hashes []string, files []string) {
	c := resty.New()
	//c.SetProxy("http://127.0.0.1:61111")
	xml, err := c.R().Get(uri)

	if err != nil {
		log.Printf("err: %v", err.Error())
		return
	}

	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(xml.Body()); err != nil {
		log.Printf("err, %v,", err.Error())
		return
	}

	root := doc.SelectElement("rss")
	if root == nil {
		log.Printf("RootFindErr")
		return
	}

	urlTorrent := root.FindElements("./channel/item/enclosure")
	if urlTorrent == nil {
		log.Printf("torrentFindErr")
		return
	}

	for _, u := range urlTorrent {
		hash := regexp.MustCompile("hash=").ReplaceAllString(regexp.MustCompile("hash=.*$").FindString(u.SelectAttr("url").Value), "")
		hashes = append(hashes, hash)

	}

	filesGet := root.FindElements("./channel/item")
	for i, f := range filesGet {
		t := f.FindElement("title").Text()
		files = append(files, t)
		log.Printf("GetFile: %v & Hash: %v", t, hashes[i])
	}

	return
}
