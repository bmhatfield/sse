package sse

import "encoding/json"

// Event is an interface for our core Event type.
// Implement on your own types to serialize custom events.
type Event interface {
	// Type is optional.
	// Empty values are interpreted by the client as "message" events.
	// Non-empty values require a matching event listener on the client.
	Type() string

	// Data is required. It must conform to UTF-8.
	Data() []byte
}

type event struct {
	kind string
	data []byte
}

func (e event) Type() string {
	return e.kind
}

func (e event) Data() []byte {
	return e.data
}

// JSONEvent encodes the provided object to JSON and returns as an Event.
// Kind may be omitted by providing an empty string.
func JSONEvent(kind string, data any) (Event, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return event{
		kind: kind,
		data: b,
	}, nil
}

// NewEvent returns an Event with the provided kind and data.
// Kind may be omitted by providing an empty string.
func NewEvent(kind string, data []byte) Event {
	return event{
		kind: kind,
		data: data,
	}
}
