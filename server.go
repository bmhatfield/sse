package sse

import (
	"fmt"
	"net/http"
)

// EventServer is the event server.
type EventServer struct {
	topics *Topics
}

// ServeHTTP implements http.Handler to serve events.
func (e *EventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topic := e.topics.Create(r.URL.Query().Get("stream"))

	s := NewSubscriber()
	topic.Subscribe(s)
	go topic.Unsubscribe(r.Context(), s)

	if err := s.WriteEvents(w); err != nil {
		http.Error(w, fmt.Sprintf("unable to write events: %s", err), http.StatusInternalServerError)
		return
	}
}

// Broadcast sends an event to all subscribers of the named topic.
func (e *EventServer) Broadcast(name string, event Event) error {
	topic, err := e.topics.get(name)
	if err != nil {
		return err
	}
	return topic.Broadcast(event)
}

// Create creates a new topic.
func (e *EventServer) Create(name string) {
	e.topics.Create(name)
}

func (e *EventServer) Stats() Stats {
	return e.topics.Stats()
}

// NewEventServer returns an empty EventServer.
func NewEventServer() *EventServer {
	return &EventServer{
		topics: NewTopics(),
	}
}
