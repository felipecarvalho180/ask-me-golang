package api

import (
	"net/http"

	"github.com/felipecarvalho180/ask-me-golang/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	queries *pgstore.Queries
	router  *chi.Mux
}

func (handler apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func NewHandler(queries *pgstore.Queries) http.Handler {
	api := apiHandler{
		queries: queries,
	}

	router := chi.NewRouter()

	api.router = router
	return api
}
