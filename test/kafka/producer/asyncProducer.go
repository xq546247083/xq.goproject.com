package producer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"xq.goproject.com/commonTools/logTool"

	"github.com/Shopify/sarama"
)

//异步消费者(Goroutines)：用不同的goroutine异步读取Successes和Errors channel
func AsyncProducer(address []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	//config.Producer.Partitioner = 默认为message的hash
	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}

	//Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	var enqueued, successes, errors int
	wg.Add(2) //2 goroutine

	// 发送成功message计数
	go func() {
		defer wg.Done()
		for range p.Successes() {
			successes++
		}
	}()

	// 发送失败计数
	go func() {
		defer wg.Done()
		for err := range p.Errors() {
			logTool.LogError(fmt.Sprintf("%v 发送失败，err：%s\n", err.Msg, err.Err))
			errors++
		}
	}()

	// 循环发送信息
	asrcValue := "async-goroutine: this is a message. index=%d"
	var i int
Loop:
	for {
		i++
		value := fmt.Sprintf(asrcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.ByteEncoder(value),
		}
		select {
		case p.Input() <- msg: // 发送消息
			enqueued++
			fmt.Fprintln(os.Stdout, value)
		case <-signals: // 中断信号
			p.AsyncClose()
			break Loop
		}
		time.Sleep(2 * time.Second)
	}
	wg.Wait()

	fmt.Fprintf(os.Stdout, "发送数=%d，发送成功数=%d，发送失败数=%d \n", enqueued, successes, errors)

}
