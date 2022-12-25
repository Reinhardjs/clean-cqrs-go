package usecases

import (
	"log"

	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/repositories"
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	commonRepositories "github.com/Reinhardjs/golang-alpha-indo-soft/internal/repositories"
	"github.com/pkg/errors"
)

type ArticleUsecase interface {
	Create(article *models.Article) (*models.Article, error)
}

type articleUsecase struct {
	articleRepo repositories.ArticleRepository
	natsRepo    commonRepositories.NatsRepository
}

func NewArticleUsecase(articleRepo repositories.ArticleRepository, natsRepo commonRepositories.NatsRepository) ArticleUsecase {
	return &articleUsecase{articleRepo, natsRepo}
}

func (e *articleUsecase) Create(article *models.Article) (*models.Article, error) {
	article, err := e.articleRepo.Create(article)

	if err != nil {
		return nil, errors.Wrap(err, "command-service.internal.article.usecase.Create.articleRepo.Create")
	}

	// Publish event
	if err := e.natsRepo.PublishArticleCreated(*article); err != nil {
		log.Println(err)
	}

	return article, nil
}
