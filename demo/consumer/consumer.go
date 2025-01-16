package main

import (
	"log"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("chenjie.info", "chenjie.info", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 处理信息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))

	// 1 连接 nsqd 的 tcp 连接
	//if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	//	log.Fatal(err)
	//}

	// 2 连接 nsqd 的 tcp 连接【多个】,在多个 nsqd上消费
	//if err := consumer.ConnectToNSQDs([]string{"127.0.0.1:4150"}); err != nil {
	//	log.Fatal(err)
	//}

	// 3 连接nsqlookupd 的 http连接
	//if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
	//	log.Fatal(err)
	//}

	// 4 连接nsqlookupd 的 http连接【多个】，多个 nsqlookupd上检索 nsqd信息，各个 nsqlookupd互相独立，对于消费者而言是备用关系
	if err := consumer.ConnectToNSQLookupds([]string{"127.0.0.1:41610", "127.0.0.1:4161", "127.0.0.1:41611"}); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
