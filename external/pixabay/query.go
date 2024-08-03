package pixabay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ImageSearch(q string) []string {
	buffer := []string{}
	m := imageSearch(q)
	images := m["hits"].([]any)
	for _, item := range images {
		thing := item.(map[string]any)
		src := thing["largeImageURL"].(string)
		bytes := int64(thing["imageSize"].(float64))
		if bytes >= 3337749 {
			buffer = append(buffer, src)
		}
	}
	return buffer
}

/*
category
Accepted values: backgrounds, fashion, nature, science, education, feelings, health, people, religion, places, animals, industry, computer, food, sports, transportation, travel, buildings, business, music

colors 	str 	Filter images by color properties. A comma separated list of values may be used to select multiple properties.
Accepted values: "grayscale", "transparent", "red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"
*/
func imageSearch(q string) map[string]any {
	urlS := "https://pixabay.com/api?q=" + url.QueryEscape(q) + "&key=" + apiKey +
		"&category=feelings&image_type=photo&orientation=horizontal&min_width=600&editors_choice=true"
	urlS = "https://pixabay.com/api?q=" + url.QueryEscape(q) + "&key=" + apiKey +
		"&orientation=horizontal&image_type=photo&min_width=600&editors_choice=false"
	req, _ := http.NewRequest("GET", urlS, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	//jsonString := string(body)
	//fmt.Println(jsonString)
	var m map[string]any
	json.Unmarshal(body, &m)
	return m
}
