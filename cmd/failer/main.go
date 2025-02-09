

package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/kelseyhightower/envconfig"
)



type envConfig struct {
	DefaultResponseCode int `envconfig:"DEFAULT_RESPONSE_CODE" required:"true"`
}

var (
	env envConfig
)

// We just want a payload to result with.
type payload struct {
	ResponseCode int `json:"responsecode,omitempty"`
}

type failer struct {
	defaultResponseCode int
}

func NewFailer(defaultResponseCode int) *failer {
	return &failer{defaultResponseCode: defaultResponseCode}
}

func (f *failer) gotEvent(inputEvent event.Event) (*event.Event, error) {
	data := &payload{}
	rc := f.defaultResponseCode
	err := inputEvent.DataAs(data)
	if err != nil {
		log.Println("Got error while unmarshalling data, using default response code: ", err.Error())
	} else if data.ResponseCode != 0 {
		rc = data.ResponseCode
	}
	log.Printf("using response code: %d\n", rc)
	return nil, cloudevents.NewHTTPResult(rc, "Responding with: %d", rc)
}

func main() {
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal("[ERROR] Failed to process env var: ", err)
	}

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatal("failed to create client: ", err)
	}

	f := NewFailer(env.DefaultResponseCode)
	log.Println("listening on 8080, default response code ", env.DefaultResponseCode)
	log.Fatalf("failed to start receiver: %s\n", c.StartReceiver(context.Background(), f.gotEvent))
}