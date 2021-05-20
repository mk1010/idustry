package rmq

import (
	"context"
	"sync"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	jsoniter "github.com/json-iterator/go"
	"github.com/mk1010/idustry/config"
	"github.com/mk1010/industry_adaptor/nclink"
	"github.com/mk1010/industry_adaptor/nclink/util"
)

var RmqClientMap = make(map[string]rocketmq.PushConsumer)

var RmqMu sync.Mutex

var RmqMsgMap = make(map[string]*sync.Map) // map[topic]map[adaptor ID]chan *nclink.NCLinkTopicMessage

func RmqInitTopic(topic string) {
	RmqMu.Lock()
	defer RmqMu.Unlock()
	if _, ok := RmqClientMap[topic]; !ok {
		c, err := rocketmq.NewPushConsumer(
			consumer.WithNameServer(config.ConfInstance.RMQNamingService),
			consumer.WithConsumerModel(consumer.BroadCasting),
		)
		if err != nil {
			return
		}
		RmqClientMap[topic] = c
		if _, ok := RmqMsgMap[topic]; !ok {
			RmqMsgMap[topic] = &sync.Map{}
		}
		util.GoSafely(func() {
			NCLinkMsgDistribution(topic, c)
		}, nil)
	}
}

func topicDelete(topic string) {
	RmqMu.Lock()
	defer RmqMu.Unlock()
	delete(RmqClientMap, topic)
}

// todo 修改结构
type RmqMsg struct {
	AdaptorId   []string               `protobuf:"bytes,1,opt,name=adaptor_id,json=adaptorId,proto3" json:"adaptor_id,omitempty"`
	MessageId   string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageKind int32                  `protobuf:"varint,2,opt,name=message_kind,json=messageKind,proto3" json:"message_kind,omitempty"`
	Payload     *nclink.NCLinkPayloads `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func NCLinkMsgDistribution(topic string, c rocketmq.PushConsumer) {
	c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			rmqMsg := new(RmqMsg)
			err := jsoniter.Unmarshal(msg.Body, rmqMsg)
			if err != nil {
				logger.Errorf("rmq消息无法解析 err:%v", err)
				continue
			}
			if rmqMsg.AdaptorId == nil {
				topicMsg := &nclink.NCLinkTopicMessage{
					MessageId:   rmqMsg.MessageId,
					MessageKind: rmqMsg.MessageKind,
					Payload:     rmqMsg.Payload,
				}
				util.GoSafely(func() {
					RmqMsgMap[topic].Range(func(key, value interface{}) bool {
						ch, ok := value.(chan *nclink.NCLinkTopicMessage)
						if ok {
							ch <- topicMsg
						}
						return true
					})
				}, nil)
			} else {
				util.GoSafely(func() {
					topicMsg := &nclink.NCLinkTopicMessage{
						MessageId:   rmqMsg.MessageId,
						MessageKind: rmqMsg.MessageKind,
						Payload:     rmqMsg.Payload,
					}
					syncMap, ok := RmqMsgMap[topic]
					if !ok {
						return
					}
					for _, adaID := range rmqMsg.AdaptorId {
						{
							if val, ok := syncMap.Load(adaID); ok {
								ch, ok := val.(chan *nclink.NCLinkTopicMessage)
								if ok {
									ch <- topicMsg
								}
							}
						}
					}
				}, nil)
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err := c.Start(); err != nil {
		logger.Errorf("rmq客户端启动失败 topic:%s err:%v", topic, err)
		topicDelete(topic)
	}
}
