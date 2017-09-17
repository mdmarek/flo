package graph

import (
	"time"

	"github.com/lytics/flo/merger"
	"github.com/lytics/flo/sink"
	"github.com/lytics/flo/source"
	"github.com/lytics/flo/trigger"
	"github.com/lytics/flo/window"
)

type Event struct {
	Time time.Time   // Event time.
	ID   string      // Event ID.
	Msg  interface{} // Event message.
}

type KeyedEvent struct {
	Time time.Time   // Event time.
	Key  string      // Event key.
	Msg  interface{} // Event message.
}

func New() *Graph {
	return &Graph{
		window: window.All(),
	}
}

type Graph struct {
	from      source.Sources
	transform func(interface{}) ([]Event, error)
	group     func(interface{}) (string, error)
	merger    merger.Merger
	window    window.Window
	trigger   trigger.Trigger
	into      sink.Sinks
}

// From defines the sources of data.
func (g *Graph) From(ss source.Sources) {
	g.from = ss
}

// Transform defines how to create an event from each datum.
func (g *Graph) Transform(f func(interface{}) ([]Event, error)) {
	g.transform = f
}

// Group defines how to group datum by a string key.
func (g *Graph) Group(f func(interface{}) (string, error)) {
	g.group = f
}

// Window defines how to calculate which windows of time
// an event belongs to.
func (g *Graph) Window(w window.Window) {
	g.window = w
}

// Merger defines the function to use for merging events
// mapped to the same key.
func (g *Graph) Merger(f merger.Merger) {
	g.merger = f
}

// Trigger defines how to translate event-time events into
// process-time events.
func (g *Graph) Trigger(t trigger.Trigger) {
	g.trigger = t
}

// Into defines where to sink data.
func (g *Graph) Into(ss sink.Sinks) {
	g.into = ss
}

// Definition of the graph, which can be called
// after From, Transform, Group, Window, Merger
// Trigger, and Into have been set.
func (g *Graph) Definition() *Definition {
	return &Definition{g}
}

// Definition of the graph.
type Definition struct {
	g *Graph
}

// From definition, in other words, were to source data.
func (def *Definition) From() source.Sources {
	return def.g.from
}

// Transform v into a slice of events. A given value v can be
// transformed into multiple events or zero events.
func (def *Definition) Transform(v interface{}) ([]Event, error) {
	return def.g.transform(v)
}

// GroupAndWindowBy takes a value v and extracts a key by using
// the "group by" of the graph, and calculates the windows of
// value v using the "window by" of the graph. The result is
// a slice of keyed events, one for each window calcualted.
// Each window actuall gets the same event.
func (def *Definition) GroupAndWindowBy(id string, ts time.Time, v interface{}) ([]KeyedEvent, error) {
	if def.g.group != nil {
		key, err := def.g.group(v)
		if err != nil {
			return nil, err
		}
		return []KeyedEvent{KeyedEvent{
			Time: ts,
			Key:  key,
			Msg:  v,
		}}, nil
	} else if def.g.window != nil {
		var events []KeyedEvent
		windows := def.g.window.Apply(ts)
		for _, w := range windows {
			key := w.String()
			events = append(events, KeyedEvent{
				Time: ts,
				Key:  key,
				Msg:  v,
			})
		}
		return events, nil
	} else {
		key := id
		return []KeyedEvent{KeyedEvent{
			Time: ts,
			Key:  key,
			Msg:  v,
		}}, nil
	}
}

// Merge the new keyed event ke, into existing windows representing
// the same key.
func (def *Definition) Merge(ke *KeyedEvent, windows map[window.Span][]interface{}) error {
	f := merger.Cons()
	if def.g.merger != nil {
		f = merger.Fold(def.g.merger)
	}
	return def.g.window.Merge(ke.Time, []interface{}{ke.Msg}, windows, f)
}

// Trigger definition.
func (def *Definition) Trigger() trigger.Trigger {
	return def.g.trigger
}

// Into definition, in other words, were to sink data.
func (def *Definition) Into() sink.Sinks {
	return def.g.into
}
