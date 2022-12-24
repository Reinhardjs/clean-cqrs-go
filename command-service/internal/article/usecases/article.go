package usecases

import (
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/repositories"
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
)

type ArticleUsecase interface {
	Create(article *models.Article) (*models.Article, error)
}

type articleUsecase struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleUsecase(articleRepo repositories.ArticleRepository) ArticleUsecase {
	return &articleUsecase{articleRepo}
}

func (e *articleUsecase) Create(article *models.Article) (*models.Article, error) {
	return e.articleRepo.Create(article)
}
