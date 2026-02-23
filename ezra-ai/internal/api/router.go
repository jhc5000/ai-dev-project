package api

import (
	"log"
	"net/http"
)

type Router struct {
	chatHandler   *ChatHandler
	healthHandler *HealthHandler
}

func NewRouter(chatHandler *ChatHandler) *Router {
	return &Router{
		chatHandler:   chatHandler,
		healthHandler: &HealthHandler{},
	}
}

func (r *Router) Start(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/chat", r.chatHandler.Handle)
	mux.HandleFunc("/health", r.healthHandler.Handle)

	log.Printf("EzraAI server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
