package sse

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// Topic is a collection of subscribers.
type Topic struct {
	events chan Event

	subscribers map[string]Subscriber

	mu sync.RWMutex
}

// fanout broadcasts events from the event channel to this topic's subscribers.
func (t *Topic) fanout() {
	for event := range t.events {
		t.mu.RLock()
		for _, sub := range t.subscribers {
			if err := sub.Push(event); err != nil {
				log.Printf("could not push event: %s", err)
			}
		}
		t.mu.RUnlock()
	}
}

// Broadcast enqueues an event for all subscribers of this topic.
func (t *Topic) Broadcast(event Event) error {
	select {
	case t.events <- event:
		return nil
	default:
		return fmt.Errorf("could not broadcast event, topic full")
	}
}

// Subscribe adds a subscriber to this topic.
func (t *Topic) Subscribe(sub Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.subscribers[sub.id] = sub
}

// Unsubscribe removes a subscriber from this topic when the context is cancelled.
func (t *Topic) Unsubscribe(ctx context.Context, sub Subscriber) {
	<-ctx.Done()

	t.mu.Lock()
	defer t.mu.Unlock()

	sub, ok := t.subscribers[sub.id]
	if ok {
		sub.Close()
		delete(t.subscribers, sub.id)
	}
}

// NewTopic returns a new Topic.
// The backlog parameter determines how many events are buffered before blocking.
// The fanout goroutine is started immediately.
func NewTopic(backlog int) *Topic {
	t := &Topic{
		events:      make(chan Event, backlog),
		subscribers: map[string]Subscriber{},
		mu:          sync.RWMutex{},
	}

	go t.fanout()
	return t
}
