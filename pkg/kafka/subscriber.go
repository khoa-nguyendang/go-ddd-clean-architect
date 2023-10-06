package kafka

import (
	"app/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type subscriber struct {
	options  *kafka.ReaderConfig
	reader   map[string]*kafka.Reader
	canceler map[string]func()
	logger   logger.Logger
}

// NewSubscriber subscriber constructor
func NewSubscriber(options *kafka.ReaderConfig) (instance Subscriber) {
	instance = &subscriber{
		options:  options,
		reader:   make(map[string]*kafka.Reader),
		canceler: make(map[string]func()),
	}

	return instance
}

// Start trigger receiving message process
func (s *subscriber) Start(ctx context.Context, execute ExecuteFunc, groupID string, poolSize int, offsetFrom Offset) (err error) {
	s.logger.Info("subscriber start. checking Start input variable", zap.Int64("offset", int64(offsetFrom)), zap.Int("poolSize", poolSize), zap.String("groupID", groupID))
	if _, ok := s.reader[groupID]; ok {
		return ErrorGroupIDAlreadySubscribed
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        s.options.Brokers,
		Topic:          s.options.Topic,
		GroupID:        groupID,
		MinBytes:       s.options.MinBytes,
		MaxBytes:       s.options.MaxBytes,
		CommitInterval: s.options.CommitInterval,
		StartOffset:    int64(offsetFrom),
	})
	s.reader[groupID] = reader

	for idx := 0; idx < poolSize; idx++ {
		go s.autoReconnect(ctx, reader, groupID, idx, execute)
	}

	return nil
}

// Stop cancel receiving message process
func (s *subscriber) Stop(ctx context.Context) (err error) {
	s.logger.Info("subscriber Stopped")
	// stop all auto reconnect
	for _, cancel := range s.canceler {
		cancel()
	}

	// stop all reader
	for _, r := range s.reader {
		if err = r.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (s *subscriber) autoReconnect(ctx context.Context, reader *kafka.Reader, groupID string, poolIndex int, execute ExecuteFunc) {
	autoCTX, cancel := context.WithCancel(context.Background())
	s.canceler[fmt.Sprintf("%s_%d", groupID, poolIndex)] = cancel

	for {
		select {
		case <-autoCTX.Done():
			return
		default:
			if err := s.consumer(ctx, reader, groupID, execute); err == ErrorExecuteFail {
				return
			}
		}
	}
}

func (s *subscriber) consumer(ctx context.Context, reader *kafka.Reader, groupID string, execute ExecuteFunc) error {
	for {
		kafkaMsg, err := reader.FetchMessage(ctx)
		if err == context.Canceled {
			return nil
		}
		if err != nil {
			s.logger.Error("consumer error fetching msg from cluster", zap.Any("err", err))
			time.Sleep(time.Second)
			return ErrorReceiveFail
		}

		msg := Message{
			Topic:     kafkaMsg.Topic,
			GroupID:   groupID,
			Partition: kafkaMsg.Partition,
			Offset:    kafkaMsg.Offset,
			Key:       kafkaMsg.Key,
			Value:     kafkaMsg.Value,
			Headers:   make([]Header, len(kafkaMsg.Headers)),
			Time:      kafkaMsg.Time,
		}
		for i, h := range kafkaMsg.Headers {
			msg.Headers[i] = Header{
				Key:   h.Key,
				Value: h.Value,
			}
		}

		if err = execute(msg); err != nil {
			s.logger.Error("consumer error execute callback fail", zap.Any("err", err))
			time.Sleep(time.Second)
			return ErrorExecuteFail
		}
		if err = reader.CommitMessages(ctx, kafkaMsg); err != nil {
			s.logger.Error("consumer error commit message", zap.Any("err", err))
			return ErrorCommitFail
		}
	}
}
