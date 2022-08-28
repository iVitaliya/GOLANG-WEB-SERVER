package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
)

var Rnd *renderer.Render

func TodoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", deleteTodo)
	})

	return rg
}