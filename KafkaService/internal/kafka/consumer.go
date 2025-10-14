package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type UserActionEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	UserIP    string    `json:"user_ip,omitempty"`
	TaskID    int       `json:"task_id,omitempty"`
	Details   string    `json:"details,omitempty"`
}

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	topic         string
}

func NewConsumer(brokers []string, topic, groupID string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumerGroup: consumerGroup,
		topic:         topic,
	}, nil
}

func (c *Consumer) Close() error {
	return c.consumerGroup.Close()
}

func (c *Consumer) ProcessMessages(ctx context.Context) {
	handler := &consumerGroupHandler{}

	for {
		err := c.consumerGroup.Consume(ctx, []string{c.topic}, handler)
		if err != nil {
			if err == context.Canceled {
				log.Println("Consumer stopped due to context cancellation")
				return
			}
			log.Printf("Error from consumer: %v", err)
			time.Sleep(5 * time.Second)
		}

		if ctx.Err() != nil {
			return
		}
	}
}

type consumerGroupHandler struct{}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var event UserActionEvent
		if err := json.Unmarshal(message.Value, &event); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		log.Printf("USER_ACTION: Time=%s, Action=%s, TaskID=%d, IP=%s, Details=%s",
			event.Timestamp.Format(time.RFC3339),
			event.Action,
			event.TaskID,
			event.UserIP,
			event.Details,
		)

		session.MarkMessage(message, "")

	}
	return nil
}
