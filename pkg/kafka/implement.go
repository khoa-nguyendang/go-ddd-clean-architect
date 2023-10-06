package kafka

import (
	"app/pkg/logger"
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaPkg struct {
	options *Options
	logger  logger.Logger
}

// NewClient construct message-queue client
func NewClient(ctx context.Context, logger logger.Logger, brokers []string) Client {
	client := &KafkaPkg{
		// FIXME: most setting need move to configuration
		options: &Options{
			WriterOptions: &WriterOptions{
				Addr:              kafka.TCP(brokers...),
				BatchTimeout:      100 * time.Millisecond,
				WriteTimeout:      3 * time.Second,
				WriteRetryTimeOut: 3 * time.Second,
			},
			KafkaReaderOptions: &kafka.ReaderConfig{
				Brokers:        brokers,
				MinBytes:       1,
				MaxBytes:       10e6,
				CommitInterval: 100 * time.Millisecond,
			},
			TopicsEditorOptions: &TopicsEditorOptions{
				Addr:          kafka.TCP(brokers...),
				ClientTimeout: 5 * time.Second,
			},
		},
		logger: logger,
	}
	return client
}

// NewSubscriber create a subscriber instance
func (im *KafkaPkg) NewSubscriber(ctx context.Context, topicName string) (instance Subscriber) {
	im.options.KafkaReaderOptions.Topic = topicName
	im.logger.Info("checking subscriber options", zap.Any("subscriberOptions", im.options.KafkaReaderOptions))
	return NewSubscriber(im.options.KafkaReaderOptions)
}

// NewPublisher create a publisher instance
func (im *KafkaPkg) NewPublisher(ctx context.Context, topicName string) (instance Publisher) {
	im.options.WriterOptions.Topic = topicName
	im.logger.Info("checking publisher options", zap.Any("publisherOptions", im.options.KafkaReaderOptions))
	return NewPublisher(im.options.WriterOptions, im.logger)
}

// NewTopicsEditor create a topics editor instance
func (im *KafkaPkg) NewTopicsEditor(ctx context.Context) (editor TopicsEditor) {
	im.logger.Info("checking topics editor options", zap.Any("topicsEditorOptions", im.options.KafkaReaderOptions))
	return NewTopicsEditor(im.options.TopicsEditorOptions)
}
