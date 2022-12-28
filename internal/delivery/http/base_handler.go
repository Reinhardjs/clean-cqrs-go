package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Reinhardjs/clean-cqrs-go/pkg/errors"
)

type RootHandler func(w http.ResponseWriter, r *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	log.Printf("An error accured: %v", err)

	clientError, ok := err.(errors.ClientError)
	if !ok {
		// the error is not ClientError
		w.WriteHeader(500) // return 500 Internal Server Error.
		fmt.Fprintf(w, "Internal server error")
		return
	}

	body, err := clientError.ResponseBody() // Get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500) // return 500 Internal Server Error.
		fmt.Fprintf(w, "Internal server error")
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
