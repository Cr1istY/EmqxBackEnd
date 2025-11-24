package jobs

import (
	"context"
	"fmt"
	"log"

	"EmqxBackEnd/mqtt"
)

// MqttPublishTask MQTT消息发布任务
func MqttPublishTask(ctx context.Context, params map[string]interface{}) error {
	// 参数校验
	topic, ok := params["topic"].(string)
	if !ok || topic == "" {
		return fmt.Errorf("缺少必填参数: topic")
	}

	message, ok := params["message"].(string)
	if !ok {
		return fmt.Errorf("缺少必填参数: message")
	}

	// 检查MQTT连接
	if !mqtt.IsConnected() {
		return fmt.Errorf("MQTT客户端未连接")
	}

	// 发布参数
	qos := byte(0)
	if q, ok := params["qos"].(float64); ok {
		qos = byte(q)
	}

	retained := false
	if r, ok := params["retained"].(bool); ok {
		retained = r
	}

	// 发布消息
	client := mqtt.GetClient()
	token := client.Publish(topic, qos, retained, message)
	token.Wait()

	if token.Error() != nil {
		return fmt.Errorf("MQTT发布失败: %w", token.Error())
	}

	log.Printf("📡 MQTT消息已发送 - Topic: %s, 长度: %d bytes", topic, len(message))
	return nil
}

// MqttBatchPublishTask 批量发布MQTT消息
func MqttBatchPublishTask(ctx context.Context, params map[string]interface{}) error {
	topics, ok := params["topics"].([]interface{})
	if !ok || len(topics) == 0 {
		return fmt.Errorf("缺少参数: topics")
	}

	message := params["message"].(string)

	for _, t := range topics {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		topic := t.(string)
		singleParams := map[string]interface{}{
			"topic":    topic,
			"message":  message,
			"qos":      params["qos"],
			"retained": params["retained"],
		}

		if err := MqttPublishTask(ctx, singleParams); err != nil {
			log.Printf("批量发布失败[%s]: %v", topic, err)
			continue
		}
	}

	return nil
}
