package song

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	SongName    string `json:"songName"`
	GroupName   string `json:"groupName"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func NewSong(songName string, groupName string) *Song {
	song := &Song{
		SongName:  songName,
		GroupName: groupName,
	}

	// Запрос на API для остальных данных

	return song
}
