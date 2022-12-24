package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/internal/article/usecases"
	internalHttp "github.com/Reinhardjs/golang-alpha-indo-soft/internal/delivery/http"
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/dto"
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/models"
	pkgErrors "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/errors"
)

type ArticleController struct {
	articleUsecase usecases.ArticleUsecase
}

func CreateArticleController(articleUsecase usecases.ArticleUsecase) ArticleController {
	return ArticleController{articleUsecase}
}

func (e *ArticleController) CreateArticle() http.Handler {
	return internalHttp.RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		post := &models.Article{}
		decodeError := json.NewDecoder(r.Body).Decode(post)

		if decodeError != nil {
			return pkgErrors.NewHTTPError(nil, 400, "Invalid request body format")
		}

		if message, ok := post.Validate(); !ok {
			return pkgErrors.NewHTTPError(nil, 400, message)
		}

		result, err := e.articleUsecase.Create(post)

		if err != nil {
			return err
		}

		response := dto.HttpResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
