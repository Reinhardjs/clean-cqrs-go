package repositories

import (
	"encoding/json"
	"strconv"

	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ArticleRepository interface {
	ReadAll() (*[]models.Article, error)
	ReadById(id int) (*models.Article, error)
}

type articleRepository struct {
	DB          *gorm.DB
	RedisClient redis.Conn
}

func NewArticleRepository(DB *gorm.DB, RedisClient redis.Conn) ArticleRepository {
	return &articleRepository{DB, RedisClient}
}

func (e *articleRepository) ReadAll() (*[]models.Article, error) {
	articles := make([]models.Article, 0)

	// Get JSON blob from Redis
	redisResult, err := e.RedisClient.Do("GET", "article:all")

	if err != nil {
		// Failed getting data from redis
		return nil, errors.Wrap(err, "query-service.article.repository.redis.getAll")
	}

	if redisResult == nil {

		err := e.DB.Table("articles").Find(&articles).Error
		if err != nil {
			return nil, errors.Wrap(err, "query-service.article.repository.DB.findAll")
		}

		articleJSON, err := json.Marshal(articles)
		if err != nil {
			return nil, errors.Wrap(err, "query-service.article.repository.json.marshal")
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "article:all", articleJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, errors.Wrap(saveRedisError, "query-service.article.repository.redis.setAll")
		}

	} else {

		json.Unmarshal(redisResult.([]byte), &articles)
	}

	return &articles, nil
}

func (e *articleRepository) ReadById(id int) (*models.Article, error) {
	article := &models.Article{}

	// Get JSON blob from Redis
	redisResult, err := e.RedisClient.Do("GET", "article:"+strconv.Itoa(id))

	if err != nil {
		// Failed getting data from redis
		return nil, errors.Wrap(err, "query-service.article.repository.redis.get")
	}

	if redisResult == nil {

		errorRead := e.DB.Table("articles").Where("id = ?", id).First(article).Error

		if errorRead != nil {
			return nil, errorRead
		}

		articleJSON, err := json.Marshal(article)

		if err != nil {
			return nil, errors.Wrap(err, "query-service.article.repository.json.marshal")
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "article:"+strconv.Itoa(id), articleJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, errors.Wrap(saveRedisError, "query-service.article.repository.redis.set")
		}
	} else {

		json.Unmarshal(redisResult.([]byte), &article)
	}

	return article, nil
}
