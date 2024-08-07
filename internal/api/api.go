package api

import (
	"context"
	"net/http"
	"sync"

	"github.com/felipecarvalho180/ask-me-golang/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type ApiHandler struct {
	queries     *pgstore.Queries
	router      *chi.Mux
	upgrader    websocket.Upgrader
	subscribers map[string]map[*websocket.Conn]context.CancelFunc
	mu          *sync.Mutex
}

func (handler ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func NewHandler(queries *pgstore.Queries) http.Handler {
	a := ApiHandler{
		queries:     queries,
		upgrader:    websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		subscribers: make(map[string]map[*websocket.Conn]context.CancelFunc),
		mu:          &sync.Mutex{},
	}

	r := chi.NewRouter()
	ConfigureMiddlewares(r)

	r.Route("/api", func(r chi.Router) {
		RegisterWebsocketRoutes(&a, r)
		RegisterApiRestRoutes(&a, r)
	})

	a.router = r
	return a
}
