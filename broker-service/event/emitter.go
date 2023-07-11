package event

import amqp "github.com/rabbitmq/amqp091-go"

type Emitter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewEmitter(conn *amqp.Connection) (*Emitter, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	emitter := &Emitter{
		connection: conn,
		channel:    channel,
	}

	return emitter, nil
}

func (e *Emitter) Close() error {
	return e.channel.Close()
}

func (e *Emitter) Push(event string, severity string) error {
	err := e.channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	return err
}
