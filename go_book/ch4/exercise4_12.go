package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Run once to create a json file and
const XKCD_URL = "https://xkcd.com/"
const FILE_NAME = "/Users/srujanreddy/Documents/Coding Projects/Golang/go_book/ch4/comic_info.json"

type comic struct {
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	ImageUrl   string `json:"img"`
}

func update_comic_data() {
	for i := 1; i < 100; i++ {
		url := XKCD_URL + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Errorf("search query failed: %s", resp.Status)
		}
		var newComic comic
		if err := json.NewDecoder(resp.Body).Decode(&newComic); err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}
		json_format, err := json.MarshalIndent(newComic, "", "  ")
		if err != nil {
			fmt.Printf(err.Error())
		}
		comic_file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := comic_file.Write(json_format); err != nil {
			log.Fatal(err)
		}
		if err := comic_file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func search_comic(search_term string) (*comic, error) {

}
func main() {
	update_comic_data()
}
