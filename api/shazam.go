package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ShazamSearch struct {
	Track Track `json:"track"`
}

type Track struct {
	Title string `json:"title"`
}

func shazamApi(filename string) ShazamSearch {

	url := "https://shazam-api-free.p.rapidapi.com/shazam/recognize/"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("koshser")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("upload_file", filepath.Base("Emotionally Scarred.mp3"))
	if err != nil {
		log.Fatal("ksoskfa")
	}
	_, err = io.Copy(part, file)
	_ = writer.Close()
	req, _ := http.NewRequest("POST", url, body)

	req.Header.Add("content-type", writer.FormDataContentType()) //"multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("X-RapidAPI-Key", "469fb232bbmshb8d76096a7a7be3p12663cjsn447945abc208")
	req.Header.Add("X-RapidAPI-Host", "shazam-api-free.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body2, _ := io.ReadAll(res.Body)
	var result ShazamSearch
	err = json.Unmarshal(body2, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
