package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/pkg/errors"
)

type Server struct {
	srv *http.Server
}

func New(as Adder, ls Lister, es Editer, ds Deleter, port, auth, web string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin"},
		AllowedMethods:   []string{"DELETE", "GET", "OPTIONS", "PUT", "POST"},
		AllowedOrigins:   []string{web},
		MaxAge:           86400,
	}))
	r.Use(headersMiddleware, authorizationMiddleware(auth))

	h := NewHandler(as, ls, es, ds)

	r.Route("/plants", func(r chi.Router) {
		r.Post("/", h.Add)
		r.Get("/", h.ListAll)
		r.Get("/{id}", h.ListOne)
		r.Put("/{id}", h.Edit)
		r.Delete("/{id}", h.Delete)
	})

	s := &http.Server{}

	s.Addr = ":" + port
	s.Handler = r
	s.IdleTimeout = 120 * time.Second
	s.MaxHeaderBytes = 1 << 20 // 1 MB
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second

	return &Server{
		srv: s,
	}
}

func (s *Server) Start() error {
	fmt.Println("Plants server ready ✅")

	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")

	return s.srv.Shutdown(ctx)
}
