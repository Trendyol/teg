package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type TriggerFunc func(context.Context, chan<- struct{}) error

func PeriodicTrigger(duration time.Duration) TriggerFunc {
	return func(ctx context.Context, sub chan<- struct{}) error {
		t := time.NewTicker(duration)

		for {
			select {
			case <-t.C:
				sub <- struct{}{}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func AmqpTrigger(amqpURL, queueName string) (TriggerFunc, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't dial amqp url: %s. err: %w", amqpURL, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("couldn't get channel err: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't declare queue (%s) err: %w", queueName, err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't consume queue (%s) err: %w", queueName, err)
	}

	return func(ctx context.Context, sub chan<- struct{}) error {
		for {
			select {
			case _, ok := <-msgs:
				if !ok {
					return fmt.Errorf("(%s) amqp channel is close", queueName)
				}
				sub <- struct{}{}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}, nil
}
