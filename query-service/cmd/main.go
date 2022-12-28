package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Reinhardjs/clean-cqrs-go/query-service/config"
	articleHttp "github.com/Reinhardjs/clean-cqrs-go/query-service/internal/article/delivery/http"
	"github.com/Reinhardjs/clean-cqrs-go/query-service/internal/article/repositories"
	"github.com/Reinhardjs/clean-cqrs-go/query-service/internal/article/usecases"
	"github.com/Reinhardjs/clean-cqrs-go/query-service/internal/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/Reinhardjs/clean-cqrs-go/internal/models"
	commonRepositories "github.com/Reinhardjs/clean-cqrs-go/internal/repositories"
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
	articleUsecase := usecases.NewArticleUsecase(articleRepository, searchRepository)
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

	// Initiating Article's Query Service
	router := mux.NewRouter()
	router.Handle("/articles", articleController.GetArticles()).Methods("GET")
	router.Handle("/articles/search", articleController.SearchArticle()).Methods("GET")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
