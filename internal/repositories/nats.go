package repositories

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	"github.com/nats-io/nats.go"
)

type NatsEventStore interface {
	Close()
	PublishArticleCreated(article models.Article) error
	OnArticleCreated(f func(models.ArticleCreatedMessage)) error
}

type natsEventStore struct {
	nc                         *nats.Conn
	articleCreatedSubscription *nats.Subscription
	articleCreatedChan         chan models.ArticleCreatedMessage
}

func NewNats(url string) (NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &natsEventStore{nc: nc}, nil
}

func (es *natsEventStore) writeMessage(m models.Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (es *natsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (es *natsEventStore) OnArticleCreated(f func(models.ArticleCreatedMessage)) (err error) {
	m := models.ArticleCreatedMessage{}
	es.articleCreatedSubscription, err = es.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := es.readMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})
	return
}

func (es *natsEventStore) PublishArticleCreated(article models.Article) error {
	m := models.ArticleCreatedMessage{ID: article.ID, Title: article.Title, Content: article.Content, CreatedAt: article.CreatedAt, UpdatedAt: article.UpdatedAt}
	data, err := es.writeMessage(&m)
	if err != nil {
		return err
	}
	return es.nc.Publish(m.Key(), data)
}

func (es *natsEventStore) Close() {
	if es.nc != nil {
		es.nc.Close()
	}
	if es.articleCreatedSubscription != nil {
		if err := es.articleCreatedSubscription.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(es.articleCreatedChan)
}
