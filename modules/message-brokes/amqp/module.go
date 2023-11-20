package amqp

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"
)

var amqpURI = "amqp://micro:micro@localhost:5672"
var TOPIC = "example.topic"

func AmqpConfigProvider() (amqpCfg amqp.Config) {
	return amqp.NewDurableQueueConfig(amqpURI)
}

func AmqpSubscriber(amqpCfg amqp.Config) (subscriber *amqp.Subscriber, err error) {

	// TODO: put this to env variables

	subscriber, err = amqp.NewSubscriber(
		amqpCfg,
		watermill.NewStdLogger(true, true),
	)

	if err != nil {
		fmt.Printf("Cannot create subscriber: %+v\n", err)
	}

	return subscriber, err
}

func AmqpPublisher(amqpCfg amqp.Config) (publisher *amqp.Publisher, err error) {
	return amqp.NewPublisher(amqpCfg, watermill.NewStdLogger(false, false))
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func Module() fx.Option {
	return fx.Module("Amqp-Message-Broker",
		fx.Provide(AmqpConfigProvider),
		fx.Provide(AmqpSubscriber),
		fx.Invoke(func(lg fx.Lifecycle, subscriber *amqp.Subscriber) {
			lg.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					messages, err := subscriber.Subscribe(ctx, TOPIC)
					if err != nil {
						_ = fmt.Errorf("cannot subscribe TOPIC %s", TOPIC)
					}

					go process(messages)
					fmt.Printf(" Subscriber %+v\n", subscriber)
					return nil
				},
			})
		}),
	)
}
