package main

import (
	"fmt"
	"log"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	// 连接 nsqd 的 tcp 连接
	producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 发布消息
	var count int
	for {
		count++
		body := fmt.Sprintf("chenjie.info %d", count)
		if err := producer.Publish("chenjie.info", []byte(body)); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
