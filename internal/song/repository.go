package song

import "musicLib/pkg/db"

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

	return song, nil
}
