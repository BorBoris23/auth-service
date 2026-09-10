package events

import (
	"context"
	"sync"
)

type Event interface {
	Name() string
}

type Listener interface {
	Handle(ctx context.Context, event Event) error
}

type Dispatcher struct {
	mu        sync.RWMutex
	listeners map[string][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]Listener),
	}
}

func (d *Dispatcher) AddListener(
	eventName string,
	listener Listener,
) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.listeners[eventName] = append(
		d.listeners[eventName],
		listener,
	)
}

func (d *Dispatcher) Dispatch(
	ctx context.Context,
	event Event,
) error {
	d.mu.RLock()
	listeners := append(
		[]Listener(nil),
		d.listeners[event.Name()]...,
	)
	d.mu.RUnlock()

	for _, listener := range listeners {
		if err := listener.Handle(ctx, event); err != nil {
			return err
		}
	}

	return nil
}
