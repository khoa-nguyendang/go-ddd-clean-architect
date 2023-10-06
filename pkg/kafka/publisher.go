package kafka

import (
	"app/pkg/logger"
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type publisher struct {
	writer            *kafka.Writer
	writeRetryTimeout time.Duration
	logger            logger.Logger
}

// NewPublisher publisher constructor
func NewPublisher(options *WriterOptions, logger logger.Logger) (instance Publisher) {
	instance = &publisher{
		writer: &kafka.Writer{
			Addr:                   options.Addr,
			Topic:                  options.Topic,
			BatchTimeout:           options.BatchTimeout,
			WriteTimeout:           options.WriteTimeout,
			AllowAutoTopicCreation: true,
		},
		writeRetryTimeout: options.WriteRetryTimeOut,
		logger:            logger,
	}

	return instance
}

// Send write the message to message queue
func (p *publisher) Send(ctx context.Context, retry bool, writerMsg ...*WriterMessage) (failMsg []*WriterMessage, err error) {
	p.logger.Info("cpublisher send. checking Send input variable", zap.Any("writerMsg", writerMsg), zap.Bool("retry", retry))
	if len(writerMsg) == 0 {
		return nil, ErrorZeroWriteMsg
	}

	msg, failMsg, err := p.convert(ctx, writerMsg...)
	if err != nil {
		return failMsg, err
	}

	if err = p.send(ctx, retry, msg...); err != nil {
		return writerMsg, err
	}

	return nil, nil
}

// SendByte sends a byte array to the publisher.
//
// It takes the following parameters:
// - ctx: A context.Context object used for context handling.
// - retry: A boolean indicating whether to retry sending the message.
// - data: A byte array containing the data to be sent.
//
// It returns an error if there was an issue sending the message.
func (p *publisher) SendByte(ctx context.Context, retry bool, data []byte) error {
	p.logger.Info("publisher SendByte - checking input variable", zap.Any("data", data), zap.Bool("retry", retry))
	return p.send(ctx, retry, kafka.Message{Value: data})
}

// SendWithBatch write the message items to message queue with single contain and retry
func (p *publisher) SendWithBatch(ctx context.Context, retry bool, size int, writerMsg *WriterMessage) (failMsg []interface{}, err error) {
	p.logger.Info("cpublisher send. checking Send input variable", zap.Any("writerMsg", writerMsg), zap.Bool("retry", retry), zap.Int("batchSize", size))
	if writerMsg == nil || writerMsg.Value == nil {
		return nil, ErrorNilWriteMsg
	}

	slice, err := p.convertToSlice(writerMsg.Value)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, ErrorZeroWriteMsg
	}
	msg, failMsg, err := p.convertPackaging(size, writerMsg.Key, slice...)
	if err != nil {
		return failMsg, err
	}

	if err = p.send(ctx, retry, msg...); err != nil {
		return slice, err
	}

	return nil, nil
}

// Close release the message queue publisher
func (p *publisher) Close(ctx context.Context) (err error) {
	p.logger.Info("publisher Close")
	return p.writer.Close()
}

func (p *publisher) convertToSlice(source interface{}) (slice []interface{}, err error) {
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(source)
		for i := 0; i < val.Len(); i++ {
			slice = append(slice, val.Index(i).Interface())
		}
	default:
		return nil, ErrorNotAbleConvertToSlice
	}
	return slice, nil
}

func (p *publisher) convertPackaging(size int, key []byte, slice ...interface{}) (msg []kafka.Message, failMsg []interface{}, err error) {
	if key == nil {
		key = []byte(uuid.New().String())
	}
	batch := make([]interface{}, 0)
	count := 0

	for _, s := range slice {
		count++
		batch = append(batch, s)
		if count >= size {
			value, err := json.Marshal(batch)
			if err != nil {
				failMsg = append(failMsg, batch...)
			} else {
				msg = append(msg, kafka.Message{
					Key:   key,
					Value: value,
				})
			}

			batch = make([]interface{}, 0)
			count = 0
		}
	}

	value, err := json.Marshal(batch)
	if err != nil {
		failMsg = append(failMsg, batch...)
	} else {
		msg = append(msg, kafka.Message{
			Key:   key,
			Value: value,
		})
	}

	return msg, failMsg, err
}

func (p *publisher) convert(ctx context.Context, writerMsg ...*WriterMessage) (msg []kafka.Message, failMsg []*WriterMessage, err error) {
	msg = make([]kafka.Message, len(writerMsg))

	for index, wMsg := range writerMsg {
		wMsgByte, err := json.Marshal(wMsg.Value)
		if err != nil {
			failMsg = append(failMsg, wMsg)
		}

		key := wMsg.Key
		//if key == nil {
		//	key = []byte(uuid.New().String())
		//}

		msg[index] = kafka.Message{
			Key:   key,
			Value: wMsgByte,
		}
	}

	return msg, failMsg, err
}

func (p *publisher) send(ctx context.Context, retry bool, writerMsg ...kafka.Message) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	if p.writeRetryTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, p.writeRetryTimeout)
	}
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return ErrorKafkaRetryWriteTimeout
		default:
			err = p.writer.WriteMessages(ctx, writerMsg...)
			if err == nil || !retry {
				return err
			}
			p.logger.Error("kafka sending fail. retry sending start ...", zap.Any("err", err))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
