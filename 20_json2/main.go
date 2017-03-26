package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
)

func main() {
	var baseUrl, query, searchType string
	query = "human%20clay"
	searchType = "album"
	baseUrl = `https://api.spotify.com/v1/search?`

	res, err := http.Get(baseUrl + "q=" + query + "&type=" + searchType)
	if err != nil {
		fmt.Println(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	fmt.Printf("%s \n", bs)

	file, err := os.Create("../pracJson/albums.json")
	enc := json.NewEncoder(file)
	enc.Encode(bs)

	stringSpotify := `{"albums":{"href" : string,
    "items" : [ {
      "album_type" : string",
      "artists" : [ {
        "external_urls" : {
          "spotify" : "https://open.spotify.com/artist/43sZBwHjahUvgbx1WNIkIz"
        },
        "href" : "https://api.spotify.com/v1/artists/43sZBwHjahUvgbx1WNIkIz",
        "id" : "43sZBwHjahUvgbx1WNIkIz",
        "name" : "Creed",
        "type" : "artist",
        "uri" : "spotify:artist:43sZBwHjahUvgbx1WNIkIz"
      } ],
      "available_markets" : [ "CA", "US" ],
      "external_urls" : {
        "spotify" : "https://open.spotify.com/album/3Nyjm9NBEdiaiWr2BEaV46"
      },
      "href" : "https://api.spotify.com/v1/albums/3Nyjm9NBEdiaiWr2BEaV46",
      "id" : "3Nyjm9NBEdiaiWr2BEaV46",
      "images" : [ {
        "height" : 640,
        "url" : "https://i.scdn.co/image/6f30a62ab93d8cb4caae35d9cb3d961c071c0265",
        "width" : 640
      }, {
        "height" : 300,
        "url" : "https://i.scdn.co/image/a9f28022b3c6112c4507076bb144ba745bb6e6de",
        "width" : 300
      }, {
        "height" : 64,
        "url" : "https://i.scdn.co/image/212f283b93a7747ce0117946de64cf549fe1cde1",
        "width" : 64
      } ],
      "name" : "Human Clay",
      "type" : "album",
      "uri" : "spotify:album:3Nyjm9NBEdiaiWr2BEaV46"
    }, {
      "album_type" : "album",
      "artists" : [ {
        "external_urls" : {
          "spotify" : "https://open.spotify.com/artist/1b9C3cy1b1jNUVpkhvJ2uC"
        },
        "href" : "https://api.spotify.com/v1/artists/1b9C3cy1b1jNUVpkhvJ2uC",
        "id" : "1b9C3cy1b1jNUVpkhvJ2uC",
        "name" : "Human Clay",
        "type" : "artist",
        "uri" : "spotify:artist:1b9C3cy1b1jNUVpkhvJ2uC"
      } ],
      "available_markets" : [ "AR", "AU", "BO", "BR", "CA", "CL", "CO", "CR", "CY", "DO", "EC", "GT", "HK", "HN", "ID", "JP", "MX", "MY", "NI", "NZ", "PA", "PE", "PH", "PY", "SE", "SG", "SV", "TR", "TW", "US", "UY" ],
      "external_urls" : {
        "spotify" : "https://open.spotify.com/album/4gb5ubh8c2PSyWddYTb7t0"
      },
      "href" : "https://api.spotify.com/v1/albums/4gb5ubh8c2PSyWddYTb7t0",
      "id" : "4gb5ubh8c2PSyWddYTb7t0",
      "images" : [ {
        "height" : 640,
        "url" : "https://i.scdn.co/image/8764c751cd05a90d16d429cce61baf6abd29f54e",
        "width" : 640
      }, {
        "height" : 300,
        "url" : "https://i.scdn.co/image/7e470984b95467f672bd06cc4aa46afbed7c46a4",
        "width" : 300
      }, {
        "height" : 64,
        "url" : "https://i.scdn.co/image/83bb752571b7f1010f75bd2ede284228c77b35fc",
        "width" : 64
      } ],
      "name" : "The Complete Recordings",
      "type" : "album",
      "uri" : "spotify:album:4gb5ubh8c2PSyWddYTb7t0"
    }, {
      "album_type" : "album",
      "artists" : [ {
        "external_urls" : {
          "spotify" : "https://open.spotify.com/artist/5plAhvNTuJakch26Y0HG97"
        },
        "href" : "https://api.spotify.com/v1/artists/5plAhvNTuJakch26Y0HG97",
        "id" : "5plAhvNTuJakch26Y0HG97",
        "name" : "Clay Venezia",
        "type" : "artist",
        "uri" : "spotify:artist:5plAhvNTuJakch26Y0HG97"
      } ],
      "available_markets" : [ "AD", "AR", "AT", "AU", "BE", "BG", "BO", "BR", "CA", "CH", "CL", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "EC", "EE", "ES", "FI", "FR", "GB", "GR", "GT", "HK", "HN", "HU", "ID", "IE", "IS", "IT", "JP", "LI", "LT", "LU", "LV", "MC", "MT", "MX", "MY", "NI", "NL", "NO", "NZ", "PA", "PE", "PH", "PL", "PT", "PY", "SE", "SG", "SK", "SV", "TR", "TW", "US", "UY" ],
      "external_urls" : {
        "spotify" : "https://open.spotify.com/album/3Yh6dI35tzYPVCBzUlzGSv"
      },
      "href" : "https://api.spotify.com/v1/albums/3Yh6dI35tzYPVCBzUlzGSv",
      "id" : "3Yh6dI35tzYPVCBzUlzGSv",
      "images" : [ {
        "height" : 640,
        "url" : "https://i.scdn.co/image/1518b4bd0ca09353052d971b194fc184a5b5062e",
        "width" : 640
      }, {
        "height" : 300,
        "url" : "https://i.scdn.co/image/1412dcc34f7f52629f3c270409d0ee03d164286b",
        "width" : 300
      }, {
        "height" : 64,
        "url" : "https://i.scdn.co/image/e86f21b82ec67a5aacdf8ec743f48cfe6ed4a6b7",
        "width" : 64
      } ],
      "name" : "The Human Condition",
      "type" : "album",
      "uri" : "spotify:album:3Yh6dI35tzYPVCBzUlzGSv"
    }, {
      "album_type" : "album",
      "artists" : [ {
        "external_urls" : {
          "spotify" : "https://open.spotify.com/artist/7xfQbd5vIsdeZvdUrjHNjn"
        },
        "href" : "https://api.spotify.com/v1/artists/7xfQbd5vIsdeZvdUrjHNjn",
        "id" : "7xfQbd5vIsdeZvdUrjHNjn",
        "name" : "Clay-Doh the World Destroyer",
        "type" : "artist",
        "uri" : "spotify:artist:7xfQbd5vIsdeZvdUrjHNjn"
      } ],
      "available_markets" : [ "AD", "AR", "AT", "AU", "BE", "BG", "BO", "BR", "CA", "CH", "CL", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "EC", "EE", "ES", "FI", "FR", "GB", "GR", "GT", "HK", "HN", "HU", "ID", "IE", "IS", "IT", "JP", "LI", "LT", "LU", "LV", "MC", "MT", "MX", "MY", "NI", "NL", "NO", "NZ", "PA", "PE", "PH", "PL", "PT", "PY", "SE", "SG", "SK", "SV", "TR", "TW", "US", "UY" ],
      "external_urls" : {
        "spotify" : "https://open.spotify.com/album/1wZspjsn4ZGXhvIkdKQ54A"
      },
      "href" : "https://api.spotify.com/v1/albums/1wZspjsn4ZGXhvIkdKQ54A",
      "id" : "1wZspjsn4ZGXhvIkdKQ54A",
      "images" : [ {
        "height" : 640,
        "url" : "https://i.scdn.co/image/0a3b9f44d89f8b15d1dcf2786ab6efb0e13bd3d2",
        "width" : 640
      }, {
        "height" : 300,
        "url" : "https://i.scdn.co/image/1fb0705c464ef27367725600756b89122965030f",
        "width" : 300
      }, {
        "height" : 64,
        "url" : "https://i.scdn.co/image/7f4cbf4f70de2f45ce761d6ba6830459fe5a460a",
        "width" : 64
      } ],
      "name" : "Human Sacrifice",
      "type" : "album",
      "uri" : "spotify:album:1wZspjsn4ZGXhvIkdKQ54A"
    } ],
    "limit" : 20,
    "next" : null,
    "offset" : 0,
    "previous" : null,
    "total" : 4
  }
}`
	// put into this struct
	var album Album
	fmt.Println(album)
	jsonSpotify, err := json.Marshal(stringSpotify)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonSpotify, &album)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(album)
}
