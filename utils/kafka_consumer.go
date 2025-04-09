package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type TokenMessage struct {
	Uid                uint64 `json:"uid"`
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
}

func StartKafkaConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "token_topic",
		GroupID:        "user_group",
		StartOffset:    kafka.FirstOffset,
		CommitInterval: 1 * time.Second,
	})

	defer reader.Close()

	for {
		// 读取消息
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		var tokenMessage TokenMessage
		err = json.Unmarshal(msg.Value, &tokenMessage)
		if err != nil {
			log.Printf("Failed to unmarshal token message: %v", err)
			continue
		}

		// 处理接收到的令牌信息
		log.Printf("Received tokens for UID: %d", tokenMessage.Uid)
		log.Printf("Access Token: %s", tokenMessage.AccessToken)
		log.Printf("Refresh Token: %s", tokenMessage.RefreshToken)
		log.Printf("Access Token Expire: %d", tokenMessage.AccessTokenExpire)
		log.Printf("Refresh Token Expire: %d", tokenMessage.RefreshTokenExpire)

		// 在这里可以添加进一步的处理逻辑，例如存储到数据库或缓存中
	}
}

// func main() {
// 	go StartKafkaConsumer()

// 	// 其他初始化代码
// 	select {}
// }
