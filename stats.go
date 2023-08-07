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
