package task

import (
	"errors"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/mk1010/idustry/modules/rmq"
	"github.com/mk1010/industry_adaptor/nclink"
)

func NCLinkCommandTopic(subMsg *nclink.NCLinkTopicSub, subServer nclink.NCLinkService_NCLinkSubscribeServer) error {
	_, ok := rmq.RmqClientMap[nclink.CommandTopic]
	if !ok {
		rmq.RmqInitTopic(nclink.CommandTopic)
	}
	msgMap := rmq.RmqMsgMap[subMsg.Topic]
	msgChan := make(chan *nclink.NCLinkTopicMessage, 5)
	msgMap.Store(subMsg.AdaptorId, msgChan)
	defer func() {
		c, ok := msgMap.LoadAndDelete(subMsg.AdaptorId)
		if ok {
			msgc, ok1 := c.(chan *nclink.NCLinkTopicMessage)
			if ok1 && msgc != msgChan {
				msgMap.LoadOrStore(subMsg.AdaptorId, msgc)
			}
		}
	}()
	asyncChan := make(chan *nclink.NCLinkTopicMessage)
	for {
		select {
		case <-subServer.Context().Done():
			{
				logger.Infof("适配器取消了nclink command topic的订阅 订阅msg：%v", subMsg)
				return nil

			}
		case msg, ok := <-msgChan:
			{
				if !ok {
					logger.Infof("nclink command topic的消息队列被关闭 订阅msg：%v", subMsg)
					return nil
				}
				// do sth msg
				err := subServer.Send(msg)
				if err != nil {
					logger.Infof("适配器取消了nclink command topic的订阅 订阅msg：%v", subMsg)
					return err
				}
			}
		case msg, ok := <-asyncChan:
			{

				if !ok {
					return errors.New("代理接收消息异常 具体错误请查询代理日志")
				}
				switch msg.MessageKind {
				// ignore now
				}
				logger.Infof("收到来自适配器%s的msg:%v", subMsg.AdaptorId, msg)
			}
		}
	}
	return nil
}

func asyncSubscribeRecv(done chan struct{}, subServer nclink.NCLinkService_NCLinkSubscribeServer, subMsg *nclink.NCLinkTopicSub, asyncChan chan *nclink.NCLinkTopicMessage) {
	for {
		select {
		case <-done:
			{
				close(asyncChan)
				return
			}
		default:
			{
				msg, err := subServer.Recv()
				if err != nil {
					logger.Errorf("适配器取消或异常退出了nclink command topic的订阅 订阅msg：%v", subMsg)
					close(asyncChan)
				}
				asyncChan <- msg
				return
			}
		}
	}
}
