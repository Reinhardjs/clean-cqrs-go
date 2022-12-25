package main

import (
	"fmt"
	"net/http"

	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/config"
	articleHttp "github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/article/delivery/http"
	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/article/repositories"
	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/article/usecases"
	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	commonRepositories "github.com/Reinhardjs/golang-alpha-indo-soft/internal/repositories"
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
	searchRepository := commonRepositories.NewElasticRepository(s.ElasticSearchClient)
	articleRepository := repositories.NewArticleRepository(s.PostgresConn, s.RedisConn)
	articleUsecase := usecases.NewArticleUsecase(articleRepository, searchRepository)
	articleController := articleHttp.CreateArticleController(articleUsecase)

	router := mux.NewRouter()

	router.Handle("/articles", articleController.GetArticles()).Methods("GET")
	router.Handle("/articles/search", articleController.SearchArticle()).Methods("GET")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
