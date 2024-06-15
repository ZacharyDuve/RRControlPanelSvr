package event

type Event[N comparable, T any] struct {
	eventName N
	payload   T
}

func NewEvent[N comparable, T any](eName N, payload T) *Event[N, T] {
	return &Event[N, T]{eventName: eName, payload: payload}
}

func (this *Event[N, T]) EventName() N {
	return this.eventName
}

func (this *Event[N, T]) Payload() T {
	return this.payload
}

// type EventListener[N comparable, T any] interface {
// 	HandleEvent(*Event[N, T])
// }

type EventListenerManager[N comparable, T any] struct {
	eventListenerFuncs []func(e *Event[N, T])
}

func (this *EventListenerManager[N, T]) AddEventListener(eLF func(e *Event[N, T])) {
	if this.eventListenerFuncs == nil {
		this.eventListenerFuncs = make([]func(e *Event[N, T]), 1)
		this.eventListenerFuncs[0] = eLF
	} else {
		this.eventListenerFuncs = append(this.eventListenerFuncs, eLF)
	}
}

func (this *EventListenerManager[N, T]) SendEvent(e *Event[N, T]) {
	for _, curELF := range this.eventListenerFuncs {
		curELF(e)
	}
}
