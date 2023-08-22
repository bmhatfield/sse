package sse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStats_Empty(t *testing.T) {
	t.Run("no topics", func(t *testing.T) {
		s := Stats{}
		assert.True(t, s.Empty())
	})

	t.Run("no subscribers", func(t *testing.T) {
		s := Stats{
			Topics: 1,
			ByTopic: map[string]TopicStats{
				"foo": {
					Subscribers: 0,
				},
			},
		}
		assert.True(t, s.Empty())
	})

	t.Run("nonzero subscribers", func(t *testing.T) {
		s := Stats{
			Topics: 1,
			ByTopic: map[string]TopicStats{
				"foo": {
					Subscribers: 1,
				},
			},
		}
		assert.False(t, s.Empty())
	})
}
