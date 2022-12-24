package usecases

import (
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/article/repositories"
)

type ArticleUsecase interface {
	ReadAll() (*[]models.Article, error)
	ReadById(id int) (*models.Article, error)
}

type articleUsecase struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleUsecase(articleRepo repositories.ArticleRepository) ArticleUsecase {
	return &articleUsecase{articleRepo}
}

func (e *articleUsecase) ReadAll() (*[]models.Article, error) {
	return e.articleRepo.ReadAll()
}

func (e *articleUsecase) ReadById(id int) (*models.Article, error) {
	return e.articleRepo.ReadById(id)
}
