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

// Stats compiles a small set of statistics
func (t *Topics) Stats() Stats {
	t.mu.RLock()
	defer t.mu.RUnlock()

	s := Stats{
		Topics:  len(t.topics),
		ByTopic: make(map[string]TopicStats),
	}

	for name, topic := range t.topics {
		topic.mu.RLock()
		ts := TopicStats{
			Subscribers:  len(topic.subscribers),
			Pending:      len(topic.events),
			Capacity:     cap(topic.events),
			BySubscriber: map[string]SubscriberStats{},
		}

		for name, sub := range topic.subscribers {
			ts.BySubscriber[name] = SubscriberStats{
				Pending:  len(sub.events),
				Capacity: cap(sub.events),
			}
		}

		s.ByTopic[name] = ts
		topic.mu.RUnlock()
	}

	return s
}

// NewTopics returns an empty set of Topics.
func NewTopics() *Topics {
	return &Topics{
		topics: map[string]*Topic{},
		mu:     sync.RWMutex{},
	}
}
