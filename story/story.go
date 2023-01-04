package story

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func LoadData() map[string]Story {
	file, err := os.Open("./story/story.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var stories map[string]Story
	json.Unmarshal(data, &stories)
	return stories
}
