package core

import "autoFollowAnime/global"

func IsDownloaded(hash string) bool {
	var anime global.Anime

	r := global.DB.Where("verify_id = ?", hash).Find(&anime)

	if r.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

func AddToDatabase(hash string, code int) {
	animeAdd := &global.Anime{
		Code:     code,
		VerifyId: hash,
	}
	global.DB.Create(animeAdd)
}
