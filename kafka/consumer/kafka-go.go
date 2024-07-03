package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 定义 Kafka brokers 和 topic
	brokers := []string{"172.26.80.1:9092"}
	topic := "my-topic"
	groupID := "my-consumer-group"

	// 创建 Kafka 阅读器配置
	config := kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	// 创建 Kafka 阅读器
	reader := kafka.NewReader(config)
	defer reader.Close()

	log.Println("Starting Kafka consumer")

	// 处理接收到的信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		// 读取消息
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		// 打印消息
		log.Printf("Message at offset %d: key=%s value=%s\n", msg.Offset, string(msg.Key), string(msg.Value))

		// 在这里处理消息，可以添加你的业务逻辑

		// 提交偏移量
		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf("Error committing offset: %v", err)
		}

		select {
		case <-signals:
			log.Println("Received termination signal, exiting...")
			return
		default:
		}
	}
}
