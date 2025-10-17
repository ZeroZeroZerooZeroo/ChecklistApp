package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/models"
)

type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func (p *Producer) SendUserAction(ctx context.Context, event models.UserActionEvent) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(eventJSON),
		Key:   sarama.StringEncoder(event.Action),
	}

	_, _, err = p.producer.SendMessage(msg)
	return err
}

func (p *Producer) SendCreateAction(ctx context.Context, userIP string, taskID int, title string) error {
	event := models.UserActionEvent{
		Timestamp: time.Now(),
		Action:    "create",
		UserIP:    userIP,
		TaskID:    taskID,
		Details:   "Created task: " + title,
	}
	return p.SendUserAction(ctx, event)
}


func (p *Producer) SendUpdateAction(ctx context.Context, userIP string, taskID int) error {
	event := models.UserActionEvent{
		Timestamp: time.Now(),
		Action:    "update",
		UserIP:    userIP,
		TaskID:    taskID,
		Details:   "Updated task status",
	}
	return p.SendUserAction(ctx, event)
}

func (p *Producer) SendDeleteAction(ctx context.Context, userIP string, taskID int) error {
	event := models.UserActionEvent{
		Timestamp: time.Now(),
		Action:    "delete",
		UserIP:    userIP,
		TaskID:    taskID,
		Details:   "Deleted task",
	}
	return p.SendUserAction(ctx, event)
}

func (p *Producer) SendGetAction(ctx context.Context, userIP string) error {
	event := models.UserActionEvent{
		Timestamp: time.Now(),
		Action:    "get",
		UserIP:    userIP,
		Details:   "Retrieved tasks list",
	}
	return p.SendUserAction(ctx, event)
}