package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

type RabbitMQ struct {
	connection *amqp.Connection
	producerCh *amqp.Channel
	consumerCh *amqp.Channel
}
type MqInterface interface {
	PublishMessage(queueName string, message string) error
}

func NewRabbitMQ() (*RabbitMQ, error) {

	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_DEFAULT_USER"),
		os.Getenv("RABBITMQ_DEFAULT_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(connStr)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	producerCh, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a producer channel: %w", err)
	}

	consumerCh, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a consumer channel: %w", err)
	}

	return &RabbitMQ{
		connection: conn,
		producerCh: producerCh,
		consumerCh: consumerCh,
	}, nil
}

func (r *RabbitMQ) DeclareQueue(queueName string) (amqp.Queue, error) {
	q, err := r.producerCh.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return q, nil
}

func (r *RabbitMQ) PublishMessage(queueName string, message string) error {
	err := r.producerCh.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent,
		})

	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	return nil
}

func (r *RabbitMQ) StartConsuming(queueName string, handler func(amqp.Delivery)) error {
	q, err := r.DeclareQueue(queueName)
	if err != nil {
		return err
	}

	msgs, err := r.consumerCh.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	for msg := range msgs {
		handler(msg)

		err := msg.Ack(false)
		if err != nil {
			fmt.Printf("Failed to ack message: %v\n", err)
		}
	}

	return nil
}

func (r *RabbitMQ) Close() {
	r.producerCh.Close()
	r.consumerCh.Close()
	r.connection.Close()
}
