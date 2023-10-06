package kafka

import (
	"time"
)

// Header message header from message queue, normally indicate message contain validation.
type Header struct {
	// header's key.
	Key string
	// header's contain.
	Value []byte
}

// Message data transform entity for receiving message from message queue.
type Message struct {
	// topic of receiving message.
	Topic string
	// source of different record on group id. each one has their own offset.
	GroupID string
	// data stream load balance.
	Partition int
	// msg receiving record in data stream.
	Offset int64
	// partition key, unset indicates random partition.
	Key []byte
	// message contain payload.
	Value []byte
	// message header payload.
	Headers []Header
	// message sending time record.
	Time time.Time
}

// WriterMessage data transform entity for sending message.
type WriterMessage struct {
	// message partition key. unset indicates random partition.
	Key []byte
	// message contain.
	Value interface{}
}

// TopicConfig message queue topic configuration
type TopicConfig struct {
	// Topic name of topic.
	Topic string
	// NumPartitions number of partitions created. -1 indicates unset.
	NumPartitions int
	// ReplicationFactor normally same as brokers number for the topic. -1 indicates unset.
	ReplicationFactor int
}

// ExecuteFunc callback func, execution after receiving message from message queue.
type ExecuteFunc func(message Message) error

// Offset enum of kafka consumer offset
type Offset int64

const (
	// FromLastOffset The most recent offset available for a partition.
	FromLastOffset Offset = -1
	// FromFirstOffset The least recent offset available for a partition.
	FromFirstOffset Offset = -2
)
