package rmq

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gowins/dionysus/rmq"
)

// &rmq.ConsumerConfig{
// NameSrvAddr:    []string{"http://rmq-cn-zvp2ud6lc0e.cn-hangzhou.rmq.aliyuncs.com:8080"},
// GroupName:      "hltv_g",
// ConsumerModel:  0,
// UseCredentials: true,
// AccessKey:      "5qjJ0K0SdXIvuj2t",
// SecretKey:      "3yuDlpGFwL2itNkAx2Djxe+XcK1BDLC1VFqD6b7CV+A=",
// ConsumerOrder:  false,
// }
type Cons struct {
	clientConsumer rocketmq.PushConsumer
}

var Cs Cons

func Initial(c *rmq.ConsumerConfig) (*Cons, error) {
	if c, err := rmq.NewConsumer(c); err == nil {
		Cs.clientConsumer = c
		return &Cs, nil
	} else {
		return nil, err
	}
}

func (c *Cons) SetTopicConsumer(topic string, handler func(ext *primitive.MessageExt)) {
	cc := c.clientConsumer
	cc.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, m := range ext {
			handler(m)
			fmt.Printf("receive msg topic %s tag %s broker %s queue %d body %s\n", m.Topic, m.GetTags(), m.Queue.BrokerName, m.Queue.QueueId, string(m.Body))
		}
		return consumer.ConsumeSuccess, nil
	})
}
