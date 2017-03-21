package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Album struct {
	AlbumType string
	Artists []Artist
	Available_Markets []string
	External_Urls map[string]string
	Href string
	Id string
	Images []Image
	Name string
	Type string
	Uri string
}

type Image struct {
	Height int
	Url string
	Width int
}

type Artist struct {
	External_Urls map[string]string
	Href string
	Id string
	Name string
	Type string
	Uri string
}

type Track struct {
	Album Album
	Artists []Artist
	Available_Markets []string
	Disc_Number int
	Duration_Ms int
	Explicit bool
	External_Ids map[string]string
	External_Urls map[string]string
	Href string
	Id string
	Name string
	Popularity int
	Preview_Url string
	Track_Number int
	Type string
	Uri string
}

type Response struct {
	Tracks []Track
}

func main() {
	var id string = `43ZHCT0cAZBISjO8DG9PnE`

	// ****** End point related stuff *******
	const baseUrl = `https://api.spotify.com/`

	// Get Spotify catalog information about an artist’s top tracks by country.
	route := baseUrl + `v1/artists/` + id + `/top-tracks?country=SE`
	fmt.Println(route)

	var resp Response
	res, err := http.Get(route)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Header)

	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()
	err = dec.Decode(&resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Tracks[0].Artists[0].Name)
}
