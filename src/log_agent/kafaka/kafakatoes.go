package kafaka

import (
	"encoding/json"
	"fmt"
	"log_agent/es"

	"github.com/Shopify/sarama"
)

// EsMapdata 定义需要向es输入的数据
var EsMapdata = map[int64]string{}

// KafakaToEs 从kafaka消费消息写入es
func KafakaToEs() {
	fmt.Println("从kafaka消费消息写入es")
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("app_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("app_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				EsMapdata[msg.Offset] = string(msg.Value)
				toEsData, err := json.Marshal(EsMapdata)
				if err != nil {
					fmt.Println(err)
					return
				}
				es.PutEs(string(toEsData))
			}

		}(pc)
	}
	select {}

}
