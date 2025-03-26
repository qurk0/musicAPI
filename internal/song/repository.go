package song

import (
	"log"
	"musicLib/pkg/db"
)

type SongRepository struct {
	Db db.Db
}

func NewSongRepository(database *db.Db) *SongRepository {
	return &SongRepository{
		Db: *database,
	}
}

func (repo *SongRepository) Create(song *Song) (*Song, error) {
	result := repo.Db.DB.Create(song)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Printf("Name - %s\nGroup - %s\nText - %s\nReleaseDate - %s\nLink - %s\n", song.SongName, song.GroupName, song.Text, song.ReleaseDate, song.Link)

	return song, nil
}

func (repo *SongRepository) GetAll(group, name, releaseDate string, page, size int) ([]SongResponce, int64, error) {
	var songs []SongResponce
	var totalCount int64
	query := repo.Db.DB.Model(&Song{})

	if group != "" {
		query = query.Where("group_name = ?", group)
	}

	if name != "" {
		query = query.Where("song_name = ?", name)
	}

	if releaseDate != "" {
		query = query.Where("release_date = ?", releaseDate)
	}

	query.Count(&totalCount)
	offset := (page - 1) * size
	query = query.Offset(offset).Limit(size)

	result := query.Find(&songs)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return songs, totalCount, nil
}

func (repo *SongRepository) GetSong(group, name string) (*Song, error) {
	log.Printf("DEBUG: %s %s", group, name)
	var song Song
	result := repo.Db.DB.Where("group_name = ? AND song_name = ?", group, name).Find(&song)

	if result.Error != nil {
		return nil, result.Error
	}
	return &song, nil
}
