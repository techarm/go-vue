package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)
	mux.Post("/validate-token", app.ValidateToken)

	mux.Get("/books", app.GetAllBooks)
	mux.Get("/books/{slug}", app.GetBook)

	mux.Route("/admin", func(r chi.Router) {
		r.Use(app.AuthoTokenMiddleware)

		r.Get("/users", app.GetUsers)
		r.Get("/users/{id}", app.GetUser)
		r.Post("/users", app.CreateUser)
		r.Put("/users/{id}", app.UpdateUser)
		r.Delete("/users/{id}", app.DeleteUser)
		r.Post("/log-user-out/{id}", app.LogUserOutAndSetInactive)

		r.Get("/books/{id}", app.GetBookByID)
		r.Get("/authors", app.GetAllAuthors)
		r.Get("/genres", app.GetAllGenres)
	})

	fileServerHandler := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServerHandler))
	return mux
}
