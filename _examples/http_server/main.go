package main

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/exp/slog"

	"github.com/nikoksr/onelog"
	slogadapter "github.com/nikoksr/onelog/adapter/slog"
)

// HelloHandler handles "/hello" requests and logs the name provided in the request.
type HelloHandler struct {
	name   string
	logger onelog.Logger
}

// ServeHTTP reveals how the logger behaves when logging informational messages.
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info().
		Time("time", time.Now()).
		Str("service", h.name).
		Msgf("Saying hello to %s", r.URL.Query().Get("name"))

	w.Write([]byte("Hello " + r.URL.Query().Get("name")))
	w.WriteHeader(http.StatusOK)
}

// ErrorProvokerHandler intentionally produces an error and logs it.
type ErrorProvokerHandler struct {
	logger onelog.Logger
}

// ServeHTTP demonstrates how the logger handles error messages.
func (e *ErrorProvokerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := errors.New("simulated error")
	e.logger.Error().
		Time("time", time.Now()).
		Err(err).
		Msg("An error has occurred")

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// SimpleServer routes the incoming requests to the appropriate handler.
type SimpleServer struct {
	mux *http.ServeMux
}

// NewSimpleServer returns a new SimpleServer.
func NewSimpleServer() *SimpleServer {
	return &SimpleServer{mux: http.NewServeMux()}
}

// RegisterHandler ties a given path to a specific handler.
func (s *SimpleServer) RegisterHandler(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

// ServeHTTP delegates the request to the appropriate handler.
func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// Logger instantiation and server setup are done in the main function.
func main() {
	logger := slogadapter.NewAdapter(slog.Default())
	helloHandler := &HelloHandler{name: "simple_server", logger: logger}
	errorProvokerHandler := &ErrorProvokerHandler{logger: logger}

	server := NewSimpleServer()
	server.RegisterHandler("/hello", helloHandler)
	server.RegisterHandler("/error", errorProvokerHandler)

	if err := http.ListenAndServe(":8080", server); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
