package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type SpotifySearchResult struct {
	Tracks SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	Items []SpotifyItem `json:"items"`
}

type SpotifyItem struct {
	Data SpotifyTrackData `json:"data"`
}

type SpotifyTrackData struct {
	Id string `json:"id"`
}

type SpotifyRecommendResult struct {
	Tracks []SpotifyRecommendTracks `json:"tracks"`
}

type SpotifyRecommendTracks struct {
	Name string `json:"name"`
}

func spotifySearch(songName string) (string, error) {
	songName = strings.ReplaceAll(songName, " ", "%20")
	url := "https://spotify23.p.rapidapi.com/search/?q=" + songName + "&type=tracks&offset=0&limit=1&numberOfTopResults=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "469fb232bbmshb8d76096a7a7be3p12663cjsn447945abc208")
	req.Header.Add("X-RapidAPI-Host", "spotify23.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var searchResult SpotifySearchResult
	err := json.Unmarshal(body, &searchResult)
	if err != nil {
		log.Fatal(err)
	}
	id := searchResult.Tracks.Items[0].Data.Id
	return id, err
}

func spotifyRecommend(seedName string) ([]string, error) {
	url := "https://spotify23.p.rapidapi.com/recommendations/?limit=3&seed_tracks=" + seedName

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "469fb232bbmshb8d76096a7a7be3p12663cjsn447945abc208")
	req.Header.Add("X-RapidAPI-Host", "spotify23.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var recommendresult SpotifyRecommendResult
	err := json.Unmarshal(body, &recommendresult)
	if err != nil {
		log.Fatal(err)
	}
	var songs [3]string
	for i := 0; i < 3; i++ {
		songs[i] = recommendresult.Tracks[i].Name
	}

	return songs[:], err

}
