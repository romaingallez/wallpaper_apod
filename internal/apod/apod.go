package apod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-getter"
	"github.com/romaingallez/wallpaper_apod/internal/config"
)

func getApodIMG(config config.ConfigType) (apod ApodReturn) {

	// url := "http://api.open-notify.org/astros.json"

	u, err := url.Parse("https://api.nasa.gov/planetary/apod")
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("api_key", config.ApiKey)
	u.RawQuery = q.Encode()
	// log.Println(u)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &apod)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return apod
}

func DownloadApod(config config.ConfigType) {
	apod := getApodIMG(config)

	output := fmt.Sprintf("%s\\wallpaper.bmp", config.ImagePath)
	// log.Println(output)
	err := getter.GetFile(output, apod.URL)
	if err != nil {
		log.Panic(err)
	}

}
