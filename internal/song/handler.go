package song

import (
	"log"
	"musicLib/configs"
	"musicLib/pkg/responce"
	"net/http"
	"strconv"
)

type SongHandler struct {
	SongRepository *SongRepository
	Conf           *configs.Config
}

type SongHandlerDeps struct {
	SongRepository *SongRepository
	Conf           *configs.Config
}

func NewSongHandler(router *http.ServeMux, deps SongHandlerDeps) {
	handler := &SongHandler{
		SongRepository: deps.SongRepository,
		Conf:           deps.Conf,
	}

	router.HandleFunc("POST /songs", handler.Create())
	router.HandleFunc("GET /songs/all", handler.GetAll())
}

func (handler *SongHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := r.URL.Query().Get("group")
		name := r.URL.Query().Get("song")
		song, err := NewSong(name, group, handler.Conf.Adress.ApiAddr)
		if err != nil {
			log.Println("DEBUG: Punkt 1")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createdSong, err := handler.SongRepository.Create(song)
		if err != nil {
			log.Println("DEBUG: Punkt 2")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responce.Json(w, createdSong, http.StatusCreated)
	}
}

func parseQuery(r *http.Request) (string, string, string, int, int, error) {
	group := r.URL.Query().Get("group")
	name := r.URL.Query().Get("song")
	releaseDate := r.URL.Query().Get("release_date")
	pageStr := r.URL.Query().Get("page")
	var page, size int
	var err error
	if pageStr == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return "", "", "", 0, 0, err
		}
	}

	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		size = 0
	} else {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			return "", "", "", 0, 0, err
		}
	}

	return group, name, releaseDate, page, size, nil
}

func (handler *SongHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, name, releaseDate, page, size, err := parseQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if page <= 0 {
			page = 1
		}

		if size <= 0 {
			size = 8
		}

		songs, totalCount, err := handler.SongRepository.GetAll(group, name, releaseDate, page, size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := SongAllResponce{
			Page:       page,
			Size:       size,
			TotalCount: totalCount,
			Songs:      songs,
		}

		responce.Json(w, resp, http.StatusOK)
	}
}
