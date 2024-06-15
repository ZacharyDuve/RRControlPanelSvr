package main

import (
	"log"
	"time"

	"github.com/ZacharyDuve/RRControlPanelSvr/src/event"
	"github.com/ZacharyDuve/RRControlPanelSvr/src/hardware"
)

func main() {
	log.Println("Starting Railroad Control Panel, wish us luck")

	heyButton := hardware.NewTestButton("HeyButton")

	heyButton.AddEventListener(func(e *event.Event[hardware.ButtonState, hardware.Button]) {
		log.Println("Got an event from", e.Payload().Name(), e.Payload().State())
	})

	heyButton.Press()

	time.Sleep(time.Second)

	heyButton.Press()

	time.Sleep(time.Second)

	heyButton.Release()

	time.Sleep(time.Second)

	heyButton.Press()
}
