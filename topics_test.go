package sse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopics_Create(t *testing.T) {
	topics := NewTopics()

	topic := topics.Create("foo")
	assert.NotNil(t, topic)

	topic2 := topics.Create("foo")
	assert.Equal(t, topic, topic2)
}

func TestTopics_Get(t *testing.T) {
	topics := NewTopics()

	topic := topics.Create("foo")
	assert.NotNil(t, topic)

	topic2, err := topics.get("foo")
	assert.NoError(t, err)
	assert.Equal(t, topic, topic2)

	_, err = topics.get("bar")
	assert.Error(t, err)
}

func TestTopics_Stats(t *testing.T) {
	topics := NewTopics()
	assert.NotNil(t, topics)

	topic := topics.Create("foo")
	assert.NotNil(t, topic)

	sub := NewSubscriber()
	topic.Subscribe(sub)

	s := topics.Stats()
	assert.Equal(t, 1, s.Topics)
	assert.Len(t, s.ByTopic, 1)
	assert.Equal(t, 1, s.ByTopic["foo"].Subscribers)
	assert.Len(t, s.ByTopic["foo"].BySubscriber, 1)
}
