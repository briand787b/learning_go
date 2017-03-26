package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Image struct {
	Height int
	Url string
	width int
}

type Follower struct {
	Href string
	Total int
}

type Artist struct {
	External_Urls map[string]string
	Followers Follower
	Genres []string
	Href string
	Id string
	Images []Image
	Name string
	Popularity int
	Type string
	Uri string
}

func main() {

	res, err := http.Get(`https://api.spotify.com/v1/artists/0OdUWJ0sBjDrqHygGUXeCF`)
	if err != nil {
		fmt.Println(err)
	}

	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()

	var artist Artist
	dec.Decode(&artist)
	fmt.Printf("%v\n", artist.Followers.Total)
}