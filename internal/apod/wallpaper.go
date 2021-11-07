package apod

import (
	"fmt"
	"log"

	"github.com/reujab/wallpaper"
	"github.com/romaingallez/wallpaper_apod/internal/config"
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func SetWallpaper(config config.ConfigType) {
	background, err := wallpaper.Get()
	check(err)
	fmt.Println("Current wallpaper:", background)

	// err = wallpaper.SetFromFile(fmt.Sprintf("%s\\wallpaper.bmp", config.ImagePath))
	// check(err)

	// log.Println(config.ImagePath)

	apod := GetApodIMG(config)

	err = wallpaper.SetFromURL(apod.URL)
	check(err)

	err = wallpaper.SetMode(wallpaper.Crop)
	check(err)

}
