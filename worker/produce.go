package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// const (
// 	topic          = "FIO"
// 	topic1         = "FIO_FAILED"
// 	broker1Address = "localhost:9092"
// 	broker2Address = "localhost:9094"
// 	broker3Address = "localhost:9095"
// 	groupID        = "my-group"
// )

type KafkaConfig struct {
	topic         string
	brokerAddress string
	GroupID       string
}

type Data struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func NewKafkaWriter(kafkaUrl, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaUrl),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewKafkaProduce(ctx context.Context, brokerAddress, topic string) (Process, error) {
	process := &KafkaConfig{
		topic:         "FIO",
		brokerAddress: "localhost:9092",
		GroupID:       "my-group",
	}

	return process, nil
}

func (h *KafkaConfig) Produce(ctx context.Context) error {
	writer := NewKafkaWriter(h.brokerAddress, h.topic)
	defer writer.Close()

	d := &Data{
		"Dmitriy",
		"Ushakov",
		"Vasilevich",
	}

	fmt.Println("Start producing...")

	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println("cannot marshal %w", err)
	}

	for i := 0; ; i++ {
		msg := kafka.Message{
			Value: []byte(b),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}

}
