package sse

import (
	"fmt"
	"sync"
)

const backlog = 10

// Topics is a collection of topics.
type Topics struct {
	topics map[string]*Topic

	mu sync.RWMutex
}

// get returns a topic by name.
func (t *Topics) get(name string) (*Topic, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	tp, ok := t.topics[name]
	if !ok {
		return nil, fmt.Errorf("topic does not exist: %s", name)
	}

	return tp, nil
}

// Create creates a new topic.
func (t *Topics) Create(name string) *Topic {
	t.mu.Lock()
	defer t.mu.Unlock()

	topic, ok := t.topics[name]
	if ok {
		return topic
	}

	topic = NewTopic(backlog)
	t.topics[name] = topic
	return topic
}

// NewTopics returns an empty set of Topics.
func NewTopics() *Topics {
	return &Topics{
		topics: map[string]*Topic{},
		mu:     sync.RWMutex{},
	}
}
