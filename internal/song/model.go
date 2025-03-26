package song

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	SongName    string `json:"songName"`
	GroupName   string `json:"groupName"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func NewSong(songName string, groupName string, addr string) (*Song, error) {

	// Запрос на API для остальных данных

	resp, err := http.Get(addr + "/info?group=" + groupName + "&song=" + songName)
	log.Printf("DEBUG: %s\n", songName)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + ": Не удалось получить данные")
	}

	var data SongCreateApiResponce
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, errors.New("500 Internal Server Error: не удалось сохранить песню в базе")
	}

	song := &Song{
		SongName:    songName,
		GroupName:   groupName,
		ReleaseDate: data.ReleaseDate,
		Text:        data.Text,
		Link:        data.Link,
	}

	return song, nil
}
