package main

import (
	"log"
	"musicLib/configs"
	"musicLib/internal/song"
	"musicLib/pkg/db"
	"musicLib/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfigs()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Автоматическая миграция таблицы Song
	db.AutoMigrate(&song.Song{})

	// Repositories || Репозитории || Уровень в API для взаимодействия с БД
	songRepo := song.NewSongRepository(db)

	// Handler || Хэндлеры || Уровень в API для обработки запросов
	song.NewSongHandler(router, song.SongHandlerDeps{
		SongRepository: songRepo,
	})

	// Middlewares || Миддлвэйры || Уровень в API для выполнения каких-то задач в промежуточных этапах обработки
	stack := middleware.Chain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    conf.Adress.OurAddr,
		Handler: stack(router),
	}

	log.Printf("INFO: Сервер запущен и слушает по адресу: %s", conf.Adress.OurAddr)
	server.ListenAndServe()
}
