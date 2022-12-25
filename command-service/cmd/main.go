package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/config"
	articleHttp "github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/delivery/http"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/repositories"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/usecases"
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
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

	// Instantiate repositories
	natsRepository := commonRepositories.NewNatsRepository(s.NatsConn)
	searchRepository := commonRepositories.NewElasticRepository(s.ElasticSearchClient)
	articleRepository := repositories.NewArticleRepository(s.PostgresConn, s.RedisConn)
	articleUsecase := usecases.NewArticleUsecase(articleRepository, natsRepository)
	articleController := articleHttp.CreateArticleController(articleUsecase)

	// --- Consume Nats Start ---
	err = natsRepository.OnArticleCreated(func(m models.ArticleCreatedMessage) {
		// Index article for searching
		article := models.Article{
			ID:        m.ID,
			Title:     m.Title,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
		if err := searchRepository.InsertArticle(context.Background(), article); err != nil {
			log.Println(err)
		}
		log.Println(m)
	})
	if err != nil {
		log.Println(err)
	}
	// --- Consume Nats End ---

	// Initiating Article's Command Service
	router := mux.NewRouter()
	router.Handle("/articles", articleController.CreateArticle()).Methods("POST")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
