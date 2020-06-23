package kafaka

import (
	"fmt"
	"log_agent/logtailf"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
	err    error
	msg    *sarama.ProducerMessage
)

//初始化Kafaka

func Init(kafakaIpaddr string, Topic string) {
	fmt.Println("初始化Kafaka")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	//	defer client.Close()
	fmt.Println("kafaka连接成功")
	msg = &sarama.ProducerMessage{}
	msg.Topic = Topic

}

//往kakafa发送消息

func SendMsg() {
	for {
		taifmsg := <-logtailf.Msgchan
		msg.Value = sarama.StringEncoder(taifmsg)
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		fmt.Printf("接受到的消息是:%s\n", taifmsg)
		//	fmt.Println("消息发送成功")
	}

}
