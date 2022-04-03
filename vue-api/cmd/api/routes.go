package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/techarm/go-vue/vue-api/internal/data"
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

	// test
	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		users, err := app.models.User.GetAll()
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		app.writeJSON(w, http.StatusOK, users)
	})

	mux.Get("/users/id", func(w http.ResponseWriter, r *http.Request) {
		user, err := app.models.User.GetByID(1)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		app.writeJSON(w, http.StatusOK, user)
	})

	mux.Post("/users/email", func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := app.readJSON(w, r, &user)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		app.models.User.GetByEmail(user.Email)
	})

	mux.Post("/users/insert", func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := app.readJSON(w, r, &user)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		id, err := user.Insert()
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		app.writeJSON(w, http.StatusOK, id)
	})

	mux.Put("/users/update", func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := app.readJSON(w, r, &user)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		err = user.Update()
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	})

	mux.Delete("/users/delete", func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := app.readJSON(w, r, &user)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		user.Delete()
	})

	return mux
}
