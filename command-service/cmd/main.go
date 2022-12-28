package main

import (
	"fmt"
	"net/http"

	"github.com/Reinhardjs/clean-cqrs-go/command-service/config"
	articleHttp "github.com/Reinhardjs/clean-cqrs-go/command-service/internal/article/delivery/http"
	"github.com/Reinhardjs/clean-cqrs-go/command-service/internal/article/repositories"
	"github.com/Reinhardjs/clean-cqrs-go/command-service/internal/article/usecases"
	"github.com/Reinhardjs/clean-cqrs-go/command-service/internal/server"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

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
	articleRepository := repositories.NewArticleRepository(s.PostgresConn, s.RedisConn)
	articleUsecase := usecases.NewArticleUsecase(articleRepository, natsRepository)
	articleController := articleHttp.CreateArticleController(articleUsecase)

	// Initiating Article's Command Service
	router := mux.NewRouter()
	router.Handle("/articles", articleController.CreateArticle()).Methods("POST")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
