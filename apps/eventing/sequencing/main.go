package main

import (
	"context"
	baseCode "demo/pkg"
	"flag"
	cloudevents "github.com/cloudevents/sdk-go"
	"log"
)


func main(){
	var step string
	var port int
	flag.StringVar(&step, "step", "0", "")
	flag.IntVar(&port, "port", 9090, "")
	flag.Parse()

	ctx := context.Background()

	// Create HTTP Transport
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithPath("/"),
	)

	if err != nil {
		log.Fatalf("failed to create transport: %s", err.Error())
	}

	// Create cloud event client
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)

	if err != nil {
		log.Fatalf("failed to create client: %s", err.Error())
	}

	stepper := baseCode.NewStepper(step)

	log.Println("starting event handler on port", port)
	if err := c.StartReceiver(ctx, stepper.GotEvent); err != nil {
		log.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}