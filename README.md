[![Go Reference](https://pkg.go.dev/badge/github.com/bmhatfield/sse.svg)](https://pkg.go.dev/github.com/bmhatfield/sse)

# github.com/bmhatfield/sse
[Server sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events) publisher in Go

## Install

`go get github.com/bmhatfield/sse`

## Usage

Import sse

```go
import "github.com/bmhatfield/sse"
```

Create the event server. The event server is an async topic publisher, with events pushed to active subscribers.

```go
// Create event server.
events := sse.NewEventServer()

// Create initial topic
events.Create("topic-name")
```

Create an event, and then broadcast it to subscribers. Messages will be sent to active subscribers. Broadcasts to non-existent topics will fail. Broadcasts to topics without subscribers will be dropped.

```go
// Create an event
// See also: JSONEvent, or implement your own Event!
event := sse.NewEvent("optional-kind", []byte("hi!"))

// Broadcast!
if err := events.Broadcast("topic-name", event); err != nil {
    log.Printf("failed to broadcast: %s", err)
}
```

The event server is a simple HTTP handler. Use with your favorite router, or serve with the standard library.

```go
// Serve event stream
http.Handle("/events", events)
if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
}
```

Use EventSource in `js` to listen for events. Clients can select their stream with the `stream` query param.

```js
const stream = new EventSource("/events?stream=points");
```
