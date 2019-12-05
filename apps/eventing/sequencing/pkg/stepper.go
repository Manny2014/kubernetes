package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"log"
)


type SampleMessage struct {
	Sequence string    `json:"id"`
	Message  string `json:"message"`
}


type stepper struct {
	step string
}

func NewStepper(step string) *stepper {
	return &stepper{step: step}
}

func (s *stepper) GotEvent(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {

	data := &SampleMessage{}


	b , _ := json.Marshal(event.Data)


	log.Println("Received Data", string(b))

	if err := event.DataAs(b); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}

	fmt.Printf("Got Data: %+v\n", data)
	fmt.Printf("Got Transport Context: %+v\n", cloudevents.HTTPTransportContextFrom(ctx))
	fmt.Printf("----------------------------\n")

	responseData := SampleMessage{
		Sequence:  s.step,
		// Just tack our step number to the Message to demo changing the event as it traverses
		// the sequence.
		Message: fmt.Sprintf("%s - Handled by %s", data.Message, s.step),
	}


	r := cloudevents.NewEvent()


	r.SetSource("manny.sample.event")
	r.SetType("manny.test.source")
	r.SetID(event.Context.GetID())
	r.SetSpecVersion(cloudevents.VersionV02)
	r.SetData(responseData)

	resp.RespondWith(200, &r)

	return nil
}
