package main

import (
	"log"
	"os"

	"github.com/gen2brain/dlgs"

	"github.com/romaingallez/wallpaper_apod/internal/apod"
	"github.com/romaingallez/wallpaper_apod/internal/config"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("apod-wallpaper: ")
	log.SetOutput(os.Stderr)

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

	yes, err := dlgs.Question("Question", "Do you want to set the last apod image as wallpaper ?", false)
	if err != nil {
		log.Panic(err)
	}

	if yes {
		config := config.Config()
		apod.DownloadApod(config)
		apod.SetWallpaper(config)

	}
}
