package producer

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

//异步消费者(Select)：同一线程内，通过select同时发送消息 和 处理errors计数。
//该方式效率较低，如果有大量消息发送， 很容易导致success和errors的case无法执行，从而阻塞一定时间。
//当然可以通过设置config.Producer.Return.Successes=false;config.Producer.Return.Errors=false来解决
func AsyncProducerSelect(wg *sync.WaitGroup, address []string) {
	defer wg.Done()

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}

	var enqueued, successes, errors int
	asrcValue := "async-select: this is a message. index=%d"
	var i int
	for {
		i++

		value := fmt.Sprintf(asrcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.ByteEncoder(value),
		}

		select {
		case p.Input() <- msg:
			fmt.Fprintln(os.Stdout, value)
			enqueued++
		case <-p.Successes():
			successes++
		case err := <-p.Errors():
			log.Printf("%v 发送失败，err：%s\n", err.Msg, err.Err)
			errors++
		}

		fmt.Fprintf(os.Stdout, "发送数=%d，发送失败数=%d \n", enqueued, errors)
		time.Sleep(2 * time.Second)
	}
}
