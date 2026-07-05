package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aryan735/SOCIAL-GO/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)

	})
	return r
}
func (app *application) run(r http.Handler) error {
	srv := &http.Server{Addr: app.config.addr, Handler: r, WriteTimeout: 30 * time.Second, ReadTimeout: 10 * time.Second, IdleTimeout: time.Minute}

	log.Printf("Server has started on port  :%s", app.config.addr)
	return srv.ListenAndServe()
}
