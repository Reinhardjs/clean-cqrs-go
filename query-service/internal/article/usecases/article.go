package usecases

import (
	"context"

	"github.com/Reinhardjs/clean-cqrs-go/internal/models"
	commonRepositories "github.com/Reinhardjs/clean-cqrs-go/internal/repositories"
	"github.com/Reinhardjs/clean-cqrs-go/query-service/internal/article/repositories"
)

type ArticleUsecase interface {
	ReadAll() (*[]models.Article, error)
	ReadById(id int) (*models.Article, error)
	Search(query string) (*[]models.Article, error)
}

type articleUsecase struct {
	articleRepo       repositories.ArticleRepository
	elasticSearchRepo commonRepositories.ElasticSearchRepository
}

func NewArticleUsecase(articleRepo repositories.ArticleRepository, elasticSearchRepo commonRepositories.ElasticSearchRepository) ArticleUsecase {
	return &articleUsecase{articleRepo, elasticSearchRepo}
}

func (e *articleUsecase) ReadAll() (*[]models.Article, error) {
	return e.articleRepo.ReadAll()
}

func (e *articleUsecase) ReadById(id int) (*models.Article, error) {
	return e.articleRepo.ReadById(id)
}

func (e *articleUsecase) Search(query string) (*[]models.Article, error) {
	return e.elasticSearchRepo.SearchArticles(context.Background(), query, uint64(0), uint64(100))
}
