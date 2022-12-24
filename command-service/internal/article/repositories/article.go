package repositories

import (
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ArticleRepository interface {
	Create(article *models.Article) (*models.Article, error)
}

type articleRepository struct {
	DB          *gorm.DB
	RedisClient redis.Conn
}

func NewArticleRepository(DB *gorm.DB, RedisClient redis.Conn) ArticleRepository {
	return &articleRepository{DB, RedisClient}
}

func (e *articleRepository) Create(article *models.Article) (*models.Article, error) {
	result := e.DB.Model(&models.Article{}).Create(article)

	if result.Error != nil {
		return &models.Article{}, errors.Wrap(result.Error, "command-service.article.repository.DB.create")
	}

	_, redisDeleteAllErr := e.RedisClient.Do("DEL", "article:all")

	if redisDeleteAllErr != nil {
		// Failed deleting data (article:all) from redis
		return nil, errors.Wrap(redisDeleteAllErr, "command-service.article.repository.redis.deleteAll")
	}

	return article, nil
}
