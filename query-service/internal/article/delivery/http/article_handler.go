package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	internalHttp "github.com/Reinhardjs/golang-alpha-indo-soft/internal/delivery/http"
	"github.com/Reinhardjs/golang-alpha-indo-soft/internal/dto"
	pkgErrors "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/errors"
	"github.com/Reinhardjs/golang-alpha-indo-soft/query-service/internal/article/usecases"
	"github.com/jinzhu/gorm"
)

type ArticleController struct {
	articleUsecase usecases.ArticleUsecase
}

func CreateArticleController(articleUsecase usecases.ArticleUsecase) ArticleController {
	return ArticleController{articleUsecase}
}

func (e *ArticleController) GetArticles() http.Handler {
	return internalHttp.RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		posts, err := e.articleUsecase.ReadAll()

		if err != nil {
			return err
		}

		response := dto.HttpResponse{Status: http.StatusOK, Message: "success", Data: posts}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *ArticleController) GetArticle() http.Handler {
	return internalHttp.RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			return pkgErrors.NewHTTPError(err, 400, "Invalid post id")
		}

		rw.Header().Add("Content-Type", "application/json")

		post, err := e.articleUsecase.ReadById(int(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return pkgErrors.NewHTTPError(err, 404, "record not found")
			} else {
				return err
			}
		}

		response := dto.HttpResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
