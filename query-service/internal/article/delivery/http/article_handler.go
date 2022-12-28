package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	internalHttp "github.com/Reinhardjs/clean-cqrs-go/internal/delivery/http"
	"github.com/Reinhardjs/clean-cqrs-go/internal/dto"
	pkgErrors "github.com/Reinhardjs/clean-cqrs-go/pkg/errors"
	"github.com/Reinhardjs/clean-cqrs-go/query-service/internal/article/usecases"
	"gorm.io/gorm"
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

func (e *ArticleController) SearchArticle() http.Handler {
	return internalHttp.RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		query := r.URL.Query().Get("query")

		if len(query) == 0 {
			return pkgErrors.NewHTTPError(nil, 400, "Missing query parameter")
		}

		if err != nil {
			return pkgErrors.NewHTTPError(err, 400, "Invalid post id")
		}

		rw.Header().Add("Content-Type", "application/json")

		posts, err := e.articleUsecase.Search(query)

		if err != nil {
			return err
		}

		response := dto.HttpResponse{Status: http.StatusOK, Message: "success", Data: posts}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
