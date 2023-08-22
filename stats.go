package sse

type SubscriberStats struct {
	Pending  int
	Capacity int
}

type TopicStats struct {
	Subscribers int
	Pending     int
	Capacity    int

	BySubscriber map[string]SubscriberStats
}

type Stats struct {
	Topics int

	ByTopic map[string]TopicStats
}

func (s Stats) Empty() bool {
	if s.Topics == 0 {
		return true
	}

	for _, t := range s.ByTopic {
		if t.Subscribers > 0 {
			return false
		}
	}

	return true
}
