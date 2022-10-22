package serviceAdapters

import (
	"context"
	"fmt"
	"github.com/quangtran88/anifni-authentication/constants"
	"github.com/quangtran88/anifni-base/libs/utils"
	"github.com/quangtran88/anifni-gateway/core/ports"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

type KafkaProducer struct {
	hosts []string
}

func NewKafkaProducer() *KafkaProducer {
	env := baseUtils.GetEnvManager()
	kafkaHostsEnv := env.GetEnv(constants.KafkaHostEnvKey)
	hosts := strings.Split(kafkaHostsEnv, ",")
	log.Printf("Init Kafka Producer with hosts %s", hosts)
	return &KafkaProducer{hosts}
}

func (p KafkaProducer) Produce(ctx context.Context, topic string, key string, value string) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	w := p.initWriter(topic)
	err := w.WriteMessages(ctx, kafkaMessage)
	if err != nil {
		log.Printf("Failed to produce kafka message: %v", err)
		return err
	}

	log.Printf("Produced kafka message to topic %s with message %s", topic, p.serializeMessages(kafkaMessage))

	err = w.Close()
	if err != nil {
		log.Printf("Failed to close writer: %v", err)
		return err
	}

	return nil
}

func (p KafkaProducer) ProduceMultiple(ctx context.Context, topic string, messages []ports.EventMessage) error {
	kafkaMessages := make([]kafka.Message, 0, len(messages))
	for _, msg := range messages {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Key:   []byte(msg.Key),
			Value: []byte(msg.Value),
		})
	}

	w := p.initWriter(topic)
	err := w.WriteMessages(ctx, kafkaMessages...)
	if err != nil {
		log.Printf("Failed to produce kafka message: %v", err)
		return err
	}

	log.Printf("Produced kafka message to topic %s with message %s", topic, p.serializeMessages(kafkaMessages...))

	err = w.Close()
	if err != nil {
		log.Printf("Failed to close writer: %v", err)
		return err
	}

	return nil
}

func (p KafkaProducer) initWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(p.hosts...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (p KafkaProducer) serializeMessages(messages ...kafka.Message) string {
	s := make([]string, 0, len(messages))
	for _, msg := range messages {
		s = append(s, fmt.Sprintf("{ Key: %s, Value: %s }", msg.Key, msg.Value))
	}
	return strings.Join(s, " ")
}
