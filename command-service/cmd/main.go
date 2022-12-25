package main

import (
	"fmt"
	"net/http"

	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/config"
	articleHttp "github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/delivery/http"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/repositories"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/usecases"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {

	cfg, err := config.InitConfig()

	if err != nil {
		fmt.Println(errors.Wrap(err, "config.InitConfig()"))
	}

	s := server.NewServer(cfg)
	err = s.Run()

	if err != nil {
		fmt.Println(errors.Wrap(err, "server.NewServer()"))
	}

	// Initiating Article's Command Service
	articleRepository := repositories.NewArticleRepository(s.PostgresConn, s.RedisConn)
	articleUsecase := usecases.NewArticleUsecase(articleRepository)
	articleController := articleHttp.CreateArticleController(articleUsecase)

	router := mux.NewRouter()

	router.Handle("/posts", articleController.CreateArticle()).Methods("POST")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
