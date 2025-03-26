package song

import (
	"log"
	"musicLib/configs"
	"musicLib/pkg/request"
	"musicLib/pkg/responce"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
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
	router.HandleFunc("GET /songs", handler.GetText())
	router.HandleFunc("PATCH /songs", handler.Update())
}

func (handler *SongHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := r.URL.Query().Get("group")
		name := r.URL.Query().Get("song")
		song, err := NewSong(name, group, handler.Conf.Adress.ApiAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createdSong, err := handler.SongRepository.Create(song)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responce.Json(w, createdSong, http.StatusCreated)
	}
}

func (handler *SongHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		releaseDate := r.URL.Query().Get("release_date")
		group, name, page, size, err := parseQuery(r)
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
			http.Error(w, err.Error(), http.StatusNotFound)
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

func (handler *SongHandler) GetText() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, name, page, size, err := parseQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if (group == "") || (name == "") {
			http.Error(w, "No group or song name", http.StatusBadRequest)
			return
		}

		song, err := handler.SongRepository.GetSong(group, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		text := strings.Split(song.Text, "\n\n")

		if size == 0 {
			size = len(text)
		}
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * size
		if offset > len(text) {
			offset = len(text)
		}

		end := offset + size
		if end > len(text) {
			end = len(text)
		}

		log.Println(text[offset:end])

		resp := SongTextResponce{
			TotalCount: int64(len(text)),
			Size:       size,
			Page:       page,
			Text:       text[offset:end],
		}

		responce.Json(w, resp, http.StatusOK)
	}
}

func parseQuery(r *http.Request) (string, string, int, int, error) {
	group := r.URL.Query().Get("group")
	name := r.URL.Query().Get("song")
	pageStr := r.URL.Query().Get("page")
	var page, size int
	var err error
	if pageStr == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return "", "", 0, 0, err
		}
	}

	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		size = 0
	} else {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			return "", "", 0, 0, err
		}
	}

	return group, name, page, size, nil
}

func (handler *SongHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SongUpdateRequest](&w, r)
		if err != nil {
			return
		}

		group := r.URL.Query().Get("group")
		name := r.URL.Query().Get("song")

		if (group == "") || (name == "") {
			http.Error(w, "No group or song name", http.StatusBadRequest)
			return
		}

		song, err := handler.SongRepository.GetSong(group, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		song, err = handler.SongRepository.Update(&Song{
			Model:       gorm.Model{ID: song.ID},
			Text:        body.Text,
			SongName:    body.SongName,
			GroupName:   body.GroupName,
			Link:        body.Link,
			ReleaseDate: body.ReleaseDate,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responce.Json(w, song, http.StatusOK)
	}
}
