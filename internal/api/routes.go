package api

import (
	"github.com/go-chi/chi/v5"
)

// API REST
func RegisterApiRestRoutes(a *ApiHandler, r chi.Router) {
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/", a.CreateRoom)
		r.Get("/", a.GetRooms)

		r.Route("/{room_id}/messages", func(r chi.Router) {
			r.Post("/", a.CreateRoomMessages)
			r.Get("/", a.GetRoomMessages)

			r.Route("/{message_id}", func(r chi.Router) {
				r.Get("/", a.GetRoomMessage)
				r.Patch("/react", a.ReactToMessage)
				r.Delete("/react", a.RemoveReactFromMessage)
				r.Patch("/answer", a.MarkMessageAsAnswered)
			})
		})
	})
}

// WEBSOCKETS
func RegisterWebsocketRoutes(a *ApiHandler, r chi.Router) {
	r.Get("/subscribe/{room_id}", a.Subscribe)
}
