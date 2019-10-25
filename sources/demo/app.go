package main

import (
	"context"
	"flag"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/google/uuid"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
	"log"
	"time"
)

type eventData struct {
	Message string `json:"message,omitempty,string"`
}

type SampleMessage struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}


func receive(ctx context.Context, event cloudevents.Event, response *cloudevents.EventResponse) error {
	// Here is where your code to process the event will go.
	// In this example we will log the event msg
	fmt.Printf("Event received. Context: %v\n", event.Context)
	data := &SampleMessage{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("Error while extracting cloudevent Data: %s\n", err.Error())
		return err
	}

	fmt.Printf("Consumer Received %q\n", data.Message)

	return nil
}
func consumer() {

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receive))

}

func producer() {
	counter := 0
	for {

		cc , _ := kncloudevents.NewDefaultClient("http://default-broker.knative-samples.svc.cluster.local")
		data := &SampleMessage{Sequence: counter, Message: "Some event"}

		newEvent := cloudevents.NewEvent()
		newEvent.SetID(uuid.New().String())
		newEvent.SetSource("manny.test.source")
		newEvent.SetType("manny.sample.event")
		newEvent.SetSpecVersion(cloudevents.VersionV02)
		newEvent.SetData(data)

		// cloudevents.
		if _, _, err := cc.Send(context.TODO(), newEvent); err != nil {
			fmt.Println("error sending: %v", err)
		}

		fmt.Printf("Event sent")
		time.Sleep(10000 * time.Millisecond)

		counter += 1
	}
}

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "", "")
	flag.Parse()
	fmt.Println("Mode", mode)
	if mode == "consumer" {
		log.Print("Started consumer", mode)
		consumer()
	} else {
		log.Print("Started producer")
		producer()
	}
}
