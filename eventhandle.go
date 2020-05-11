package eventhandle

import (
	"sync"
)

type EventHandle struct {
	id          string
	subscribers []chan interface{}
	lock        *sync.RWMutex
}

func NewHandle(ID string, initCap int) *EventHandle {
	eh := &EventHandle{
		id:          ID,
		subscribers: make([]chan interface{}, 0, initCap),
		lock:        new(sync.RWMutex),
	}
	return eh
}

func (eh *EventHandle) ID() string {
	return eh.id
}

func (eh *EventHandle) Subscribe(bufferSize int) (newSubscriber chan interface{}) {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	newSubscriber = make(chan interface{}, bufferSize)
	eh.subscribers = append(eh.subscribers, newSubscriber)
	return newSubscriber
}

func (eh *EventHandle) Unsubscribe(oldSubscriber chan interface{}) {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	for i, s := range eh.subscribers {
		if s != oldSubscriber {
			continue
		}

		if i == len(eh.subscribers)-1 {
			eh.subscribers = eh.subscribers[:len(eh.subscribers)-1]
			break
		}
		if i == 0 {
			eh.subscribers = eh.subscribers[1:len(eh.subscribers)]
			break
		}
		eh.subscribers = append(eh.subscribers[:i], eh.subscribers[i+1:]...)
		break
	}
}

func (eh *EventHandle) Publish(event interface{}) {
	eh.lock.RLock()
	defer eh.lock.RUnlock()
	for _, s := range eh.subscribers {
		s <- event
	}
}
