package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	ceClient "github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
	"gopkg.in/go-playground/webhooks.v3"
)

const (
	bbRequestUUID = "Request-UUID"
	bbEventKey    = "Event-Key"
)

// Adapter converts incoming BitBucket webhook events to CloudEvents and
// then sends them to the specified Sink.
type Adapter struct {
	Sink   string
	client ceClient.Client

	initClientOnce sync.Once
}

// HandleEvent is invoked whenever an event comes in from BitBucket.
func (a *Adapter) HandleEvent(payload interface{}, header webhooks.Header) {
	hdr := http.Header(header)
	err := a.handleEvent(payload, hdr)
	if err != nil {
		log.Printf("unexpected error handling BitBucket event: %s", err)
	}
}

func (a *Adapter) handleEvent(payload interface{}, hdr http.Header) error {
	var err error

	a.initClientOnce.Do(func() {
		a.client, err = kncloudevents.NewDefaultClient(a.Sink)
	})

	if a.client == nil {
		return fmt.Errorf("failed to create cloudevent client: %s", err)
	}

	cloudEventType := fmt.Sprintf("%s.%s", "manny", "test")
	eventID := hdr.Get("X-" + bbRequestUUID)
	source := types.ParseURLRef("http://superfake.com")
	event := cloudevents.Event{
		Context: cloudevents.EventContextV02{
			ID:     eventID,
			Type:   cloudEventType,
			Source: *source,
		}.AsV02(),
		Data: payload,
	}

	_, err = a.client.Send(context.TODO(), event)

	return err
}
