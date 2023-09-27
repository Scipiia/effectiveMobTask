package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/scipiia/effectivemobiletask/validate"
	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(kafkaUrl, groupID, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaUrl},
		GroupID: groupID,
		Topic:   topic,
	})
}

func NewKafkaConsume(ctx context.Context, brokerAddress, topic, groupID string) (Process, error) {
	process := &KafkaConfig{
		topic:         "FIO",
		brokerAddress: "localhost:9092",
		GroupID:       "my-group",
	}

	return process, nil
}

func (h *KafkaConfig) Consume(ctx context.Context) (Data, error) {
	// writer := NewKafkaWriter(broker2Address, topic1)
	// defer writer.Close()

	reader := NewKafkaReader(h.brokerAddress, h.GroupID, h.topic)
	defer reader.Close()

	var obj = &Data{}

	fmt.Println("start consuming ... !!")
	for i := 0; ; i++ {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("cannot read message", err)
			break
		}
		fmt.Println("received: ", string(msg.Value))

		err = json.Unmarshal(msg.Value, obj)
		if err != nil {
			fmt.Println("cannot unmarshal value %w", err)
		}

		err = validate.ValidateName(obj.Name)
		if err != nil {
			// msg := kafka.Message{
			// 	Value: []byte("FIO_FAILED empty string"),
			// }
			// writer.WriteMessages(context.Background(), msg)
		}

		fmt.Println("receved:", obj.Name)
	}

	return *obj, nil
}

// func Consume(ctx context.Context) (*Data, error) {
// 	var obj = &Data{}

// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{broker1Address, broker2Address, broker3Address},
// 		Topic:   topic,
// 		GroupID: "my-group",
// 	})

// 	r.SetOffset(42)

// 	for {
// 		msg, err := r.ReadMessage(ctx)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println("received: ", string(msg.Value))

// 		err = json.Unmarshal(msg.Value, obj)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println("recevd1", obj.Name)
// 	}

// 	err := validate.ValidateName(obj.Name)
// 	if err != nil {

// 	}

// 	return obj, nil
// }
