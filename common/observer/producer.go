package observer

import "sync"

func NewEventProducer() EventProducer {
	return &BaseEventProducer{
		mutex: new(sync.Mutex),
	}
}

type EventProducer interface {
	
	
	SendMessage(Message)
	
	

	AddEventListener(EventListener)
	
	

	RemoveMessageListener(EventListener)
}

type BaseEventProducer struct {
	EventProducer
	mutex *sync.Mutex

	listeners map[string]EventListener
}

func (l *BaseEventProducer) RemoveMessageListener(e EventListener) {
	if l.listeners != nil {
		delete(l.listeners, e.Id())
	}
}

func (l *BaseEventProducer) SendMessage(message Message) {
	if l.listeners != nil {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		for _, val := range l.listeners {
			if val.ListensFor(message) {
				val.OnMessage(message)
			}
		}
	}
}

func (l *BaseEventProducer) AddEventListener(e EventListener) {
	if l.listeners == nil {
		l.listeners = make(map[string]EventListener)
	}

	l.listeners[e.Id()] = e
}
