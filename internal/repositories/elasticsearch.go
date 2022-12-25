package repositories

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearchRepository interface {
	Close()
	InsertArticle(ctx context.Context, article models.Article) error
	SearchArticles(ctx context.Context, query string, skip uint64, take uint64) ([]models.Article, error)
}

type elasticSearchRepository struct {
	client *elastic.Client
}

func NewElasticRepository(client *elastic.Client) ElasticSearchRepository {

	return &elasticSearchRepository{client}
}

func (r *elasticSearchRepository) Close() {
}

func (r *elasticSearchRepository) InsertArticle(ctx context.Context, article models.Article) error {
	body, _ := json.Marshal(article)
	_, err := r.client.Index(
		"articles",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(article.ID),
		r.client.Index.WithRefresh("wait_for"),
	)
	return err
}

func (r *elasticSearchRepository) SearchArticles(ctx context.Context, query string, skip uint64, take uint64) (result []models.Article, err error) {
	var buf bytes.Buffer

	reqBody := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"body"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}
	if err = json.NewEncoder(&buf).Encode(reqBody); err != nil {
		return nil, errors.Wrap(err, "query-service.article.repository.elasticsearch.json.encode")
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("articles"),
		r.client.Search.WithFrom(int(skip)),
		r.client.Search.WithSize(int(take)),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, errors.Wrap(err, "query-service.article.repository.elasticsearch.client.Search")
	}

	defer func() {
		if err = res.Body.Close(); err != nil {
			result = nil
		}
	}()

	if res.IsError() {
		return nil, errors.Wrap(errors.New("elastic search failed"), "query-service.article.repository.elasticsearch.client.search.res.isError")
	}

	type Response struct {
		Took int64
		Hits struct {
			Total struct {
				Value int64
			}
			Hits []*struct {
				Source models.Article `json:"_source"`
			}
		}
	}

	resBody := Response{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, errors.Wrap(err, "query-service.article.repository.elasticsearch.json.decode")
	}

	var articles []models.Article
	for _, hit := range resBody.Hits.Hits {
		articles = append(articles, hit.Source)
	}

	return articles, nil
}
