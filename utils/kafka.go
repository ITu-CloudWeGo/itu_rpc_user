package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func PublishUID(uid int64) error {
	// 创建 Kafka Writer
	writer := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "uid_topic",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}

	defer writer.Close()

	// 构造 UID 消息
	message := map[string]string{
		"uid": fmt.Sprintf("%d", uid),
	}
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// 发送消息到 Kafka
	for i := 0; i < 3; i++ {
		if err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: value,
		}); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				return err
			} else {
				fmt.Printf("批量写入kafka失败：%v\n", err)
			}
		} else {
			break
		}
	}
	log.Println("UID sent to Kafka")
	return nil
}


// func ConsumeTokenResponses() {
// 	// 创建 Kafka Reader
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topic,
// 		GroupID: "user_service_group",
// 	})

// 	defer reader.Close()

// 	for {
// 		// 读取消息
// 		msg, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			log.Printf("Error reading message from Kafka: %v", err)
// 			continue
// 		}

// 		// 解析 Token 响应
// 		var tokenResponse map[string]string
// 		err = json.Unmarshal(msg.Value, &tokenResponse)
// 		if err != nil {
// 			log.Printf("Failed to unmarshal token response: %v", err)
// 			continue
// 		}

// 		uid := tokenResponse["uid"]
// 		token := tokenResponse["token"]

// 		log.Printf("Received JWT token for UID: %s", uid)
// 		// 将 Token 返回给客户端（具体逻辑根据需求实现）
// 	}
// }