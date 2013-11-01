package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
//	"flag"
)

type Available struct {
    Code string `xml:"available"`
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	address := os.Args[1]
	SpotifyURL := "http://ws.spotify.com/lookup/1/.json?uri=" + address
	// "spotify:track:0mLEhWB4vwLGcte4lImHNq2"
	fmt.Println(SpotifyURL)

	res, err := http.Get(SpotifyURL)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", contents)

	type TrackInfo struct {
		Available	bool
	}

	type TypeInfo struct {
		InfoType	string
	}

	type SpotifyJSON struct {
		Track	TrackInfo
		Info	TypeInfo
	}

	m := SpotifyJSON{}
	err = json.Unmarshal(contents, &m)

	fmt.Printf("Available: %b\n", m.Track.Available)

	if m.Track.Available && err == nil {
		// if we got here, there is a match and no error, so we're going to exit with 0 
		os.Exit(0)
	}

	os.Exit(1)
}
