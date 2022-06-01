package global

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm/logger"
	"log"
	"time"

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func InitSqlite() {
	db, err := gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln(err)
	}
	err = db.AutoMigrate(&Anime{})
	if err != nil {
		log.Printf("AutoMigrate Error: %v", err.Error())
		return
	}
	DB = db
}

type Anime struct {
	ID           int       `gorm:"primaryKey"`
	DownloadedAt time.Time `gorm:"autoCreateTime"`
	Code         int
	VerifyId     string
}
