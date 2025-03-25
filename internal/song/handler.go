package song

import (
	"musicLib/pkg/request"
	"musicLib/pkg/responce"
	"net/http"
)

type SongHandler struct {
	SongRepository *SongRepository
}

type SongHandlerDeps struct {
	SongRepository *SongRepository
}

func NewSongHandler(router *http.ServeMux, deps SongHandlerDeps) {
	handler := &SongHandler{
		SongRepository: deps.SongRepository,
	}

	router.HandleFunc("POST /songs", handler.Create())
}

func (handler *SongHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SongCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		song := NewSong(body.SongName, body.GroupName)
		createdSong, err := handler.SongRepository.Create(song)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responce.Json(w, createdSong, http.StatusCreated)
	}
}
