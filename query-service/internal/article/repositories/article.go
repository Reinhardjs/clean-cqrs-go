package repositories

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
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
		return nil, err
	}

	if redisResult == nil {

		err := e.DB.Table("articles").Find(&articles).Error
		if err != nil {
			return nil, fmt.Errorf("DB error : %v", err)
		}

		articleJSON, err := json.Marshal(articles)
		if err != nil {
			return nil, err
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "article:all", articleJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, saveRedisError
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
		return nil, err
	}

	if redisResult == nil {

		errorRead := e.DB.Table("articles").Where("id = ?", id).First(article).Error

		if errorRead != nil {
			return nil, errorRead
		}

		articleJSON, err := json.Marshal(article)
		if err != nil {
			return nil, err
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "article:"+strconv.Itoa(id), articleJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, saveRedisError
		}
	} else {
		json.Unmarshal(redisResult.([]byte), &article)
	}

	return article, nil
}
