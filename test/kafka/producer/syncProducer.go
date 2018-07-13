package producer

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

//同步消息模式
func SyncProducer(wg *sync.WaitGroup, address []string) {
	defer wg.Done()

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()

	topic := "test"
	srcValue := "sync: this is a message. index=%d"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}

		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("生产者："+"send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, "生产者："+value+"发送成功，partition=%d, offset=%d \n", part, offset)
		}

		time.Sleep(time.Second)
	}
}
