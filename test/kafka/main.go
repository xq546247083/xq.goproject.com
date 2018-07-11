package main

import (
	"sync"

	"xq.goproject.com/test/kafka/consumer"
)

var Address = []string{"10.130.138.164:9092", "10.130.138.164:9093", "10.130.138.164:9094"}

func main() {
	topic := []string{"test"}
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	//广播式消费：消费者1
	go consumer.ClusterConsumer(wg, Address, topic, "group-1")
	//广播式消费：消费者2
	go consumer.ClusterConsumer(wg, Address, topic, "group-2")

	wg.Wait()
}
