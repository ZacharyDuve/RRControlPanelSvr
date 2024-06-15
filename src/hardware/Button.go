package hardware

import (
	"github.com/ZacharyDuve/RRControlPanelSvr/src/event"
	"github.com/google/uuid"
)

type ButtonState uint8

const (
	Released ButtonState = iota
	Pressed
)

type Button interface {
	//uuid of button
	UUID() uuid.UUID
	Name() string
	State() ButtonState
	AddEventListener(func(*event.Event[ButtonState, Button]))
}

type baseButton struct {
	id    uuid.UUID
	name  string
	state ButtonState
	event.EventListenerManager[ButtonState, Button]
}

func (this *baseButton) UUID() uuid.UUID {
	return this.id
}

func (this *baseButton) Name() string {
	return this.name
}

func (this *baseButton) State() ButtonState {
	return this.state
}

func (this *baseButton) updateState(newState ButtonState) {
	if newState != this.state {
		this.state = newState
		this.SendEvent(event.NewEvent(newState, Button(this)))
	}
}

type TestButton interface {
	Button
	Press()
	Release()
}

func NewTestButton(name string) TestButton {
	tB := &baseButton{}

	tB.id = uuid.New()
	tB.name = name
	tB.state = Released

	return tB
}

func (this *baseButton) Press() {
	this.updateState(Pressed)
}

func (this *baseButton) Release() {
	this.updateState(Released)
}
