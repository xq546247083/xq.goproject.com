package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/sdbaiguanghe/glog"
)

var (
	topics = "topics_1"
	wg     sync.WaitGroup
)

func main() {
	wg.Add(1)

	go syncProducer()
	go consumer()

	wg.Wait()
}

// consumer 消费者
func consumer() {
	groupID := "group-1"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	c, err := cluster.NewConsumer(strings.Split("localhost:9092", ","), groupID, strings.Split(topics, ","), config)
	if err != nil {
		glog.Errorf("Failed open consumer: %v", err)
		return
	}
	defer c.Close()
	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				glog.Errorln(err)
			case notiObj := <-noti:
				fmt.Println("接受到消息:", notiObj)
			}
		}
	}(c)

	for msg := range c.Messages() {
		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

// syncProducer 同步生产者
// 并发量小时，可以用这种方式
func syncProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(strings.Split("localhost:9092", ","), config)
	if err != nil {
		glog.Errorln(err)
		return
	}
	defer p.Close()

	v := "sync: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
	fmt.Fprintln(os.Stdout, v)
	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(v),
	}
	if _, _, err := p.SendMessage(msg); err != nil {
		glog.Errorln(err)
		return
	}
}

// asyncProducer 异步生产者
// 并发量大时，必须采用这种方式
func asyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(strings.Split("localhost:9092", ","), config)
	defer p.Close()
	if err != nil {
		return
	}

	// 必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					glog.Errorln(err)
				}
			case <-success:
			}
		}
	}(p)

	v := "async: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
	fmt.Fprintln(os.Stdout, v)
	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(v),
	}
	p.Input() <- msg
}
