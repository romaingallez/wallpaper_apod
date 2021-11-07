package main

import (
	"log"
	"os"

	"github.com/romaingallez/wallpaper_apod/internal/apod"
	"github.com/romaingallez/wallpaper_apod/internal/config"
	"github.com/romaingallez/wallpaper_apod/internal/wallpaper"
)

func init() {
	//Create wallpaper storage dir

	config := config.Config()

	if _, err := os.Stat(config.ImagePath); os.IsNotExist(err) {
		err = os.Mkdir(config.ImagePath, 0777)
		if err != nil {
			log.Println(err)
		}
	}

}

func main() {
	config := config.Config()
	apod.DownloadApod(config)
	wallpaper.SetWallpaper(config)

}
