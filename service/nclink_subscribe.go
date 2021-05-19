package service

import (
	"fmt"

	"github.com/apache/dubbo-go/common/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/mk1010/idustry/task"
	"github.com/mk1010/industry_adaptor/nclink"
	"github.com/pkg/errors"
)

func (n *NcLinkServiceProvider) NCLinkSubscribe(subServer nclink.NCLinkService_NCLinkSubscribeServer) error {
	msg, err := subServer.Recv()
	if err != nil {
		return err
	}
	logger.Info("NCLinkSubscribe收到第一条msg:", msg)
	if msg.MessageKind != int32(nclink.NclinkCommandMessageKind_Subscribe) {
		err := fmt.Errorf("第一条信息非订阅信息，请检查代码逻辑")
		logger.Error(err)
		return err
	}
	topicSub := new(nclink.NCLinkTopicSub)
	err = jsoniter.Unmarshal(msg.Payload.Payload, topicSub)
	if err != nil {
		err = errors.Wrap(err, "NCLinkTopicSub消息类型解析失败")
		logger.Error(err)
		return err
	}
	switch topicSub.Topic {
	case nclink.CommandTopic:
		{
			return task.NCLinkCommandTopic(topicSub, subServer)
		}
	default:
		{
			err = errors.New("未知的主题类型")
			return err
		}
	}
	return nil
}
