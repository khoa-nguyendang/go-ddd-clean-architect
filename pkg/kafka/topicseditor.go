package kafka

import (
	"app/pkg/kafka/conn"
	"app/pkg/logger"
	"context"
	"fmt"
	"net"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type topicEditor struct {
	options *TopicsEditorOptions
	logger  logger.Logger
}

// NewTopicsEditor topics editor constructor
func NewTopicsEditor(options *TopicsEditorOptions) (instance TopicsEditor) {
	instance = &topicEditor{
		options: options,
	}

	return instance
}

func (t *topicEditor) newClient(addr net.Addr, timeout time.Duration) (*kafka.Client, func()) {
	conns := &conn.WaitGroup{
		DialFunc: (&net.Dialer{}).DialContext,
	}

	// NOTE: Must not use the custom resolver in "Kafka Transport". that will make kafka-go
	// resolving the cluster domain in old process, that process contain bugs result in cluster domain
	// can't resolve in generating connection pools.
	transport := &kafka.Transport{
		Dial: conns.Dial,
	}

	client := &kafka.Client{
		Addr:      addr,
		Timeout:   timeout,
		Transport: transport,
	}

	return client, func() { transport.CloseIdleConnections(); conns.Wait() }
}

// CreateTopics create topics with setting
func (t *topicEditor) CreateTopics(ctx context.Context, topics []*TopicConfig) (err error) {
	t.logger.Info("topic editor create topics, checking CreateTopics input value", zap.Any("topics", topics))

	if len(topics) <= 0 {
		return ErrorZeroTopics
	}

	kafkaTopics := make([]kafka.TopicConfig, len(topics))
	for i, v := range topics {
		kafkaTopics[i] = kafka.TopicConfig{
			Topic:             v.Topic,
			NumPartitions:     v.NumPartitions,
			ReplicationFactor: v.ReplicationFactor,
		}
	}

	client, shutdown := t.newClient(t.options.Addr, t.options.ClientTimeout)
	defer shutdown()
	response, err := client.CreateTopics(ctx, &kafka.CreateTopicsRequest{
		Topics: kafkaTopics,
	})

	if err != nil {
		return err
	}

	kafkaErrorCodes := make([]string, 0)
	for k, v := range response.Errors {
		if v != nil {
			errCode := fmt.Sprintf("%s:%s", k, v)
			kafkaErrorCodes = append(kafkaErrorCodes, errCode)
		}
	}

	if len(kafkaErrorCodes) > 0 {
		t.logger.Error("CreateTopics response error", zap.Any("kafkaErrorCodes", topics))
	}

	return nil
}

// DeleteTopics delete topics
func (t *topicEditor) DeleteTopics(ctx context.Context, topics []string) (err error) {
	t.logger.Info("topic editor delete topics, checking DeleteTopics input value", zap.Any("topics", topics))

	if len(topics) <= 0 {
		return ErrorZeroTopics
	}

	client, shutdown := t.newClient(t.options.Addr, t.options.ClientTimeout)
	defer shutdown()
	response, err := client.DeleteTopics(ctx, &kafka.DeleteTopicsRequest{
		Topics: topics,
	})
	if err != nil {
		return err
	}

	kafkaErrorCodes := make([]string, 0)
	for k, v := range response.Errors {
		if v != nil {
			errCode := fmt.Sprintf("%s:%s", k, v)
			kafkaErrorCodes = append(kafkaErrorCodes, errCode)
		}
	}

	if len(kafkaErrorCodes) > 0 {
		t.logger.Error("DeleteTopics response error", zap.Any("kafkaErrorCode", topics))
	}

	return nil
}
