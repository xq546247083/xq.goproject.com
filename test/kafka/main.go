package main

import (
	"sync"
	"time"

	"xq.goproject.com/test/kafka/consumer"
	"xq.goproject.com/test/kafka/producer"
)

var address = []string{"127.0.0.1:9092"}

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(2)

	// 启动消费者
	topic := []string{"test"}
	go producer.AsyncProducer(wg, address)
	time.Sleep(5 * time.Second)
	go consumer.ClusterConsumer(wg, address, topic, "Group1")

	wg.Wait()
}
