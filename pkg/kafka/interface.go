package kafka

import "context"

// Client kafka client
type Client interface {
	// NewSubscriber create a subscriber instance, the Subscriber helping user receiving data from
	// Kafka topics, each instance has their own kafka Reader.
	NewSubscriber(ctx context.Context, topicName string) (instance Subscriber)
	// NewPublisher create a publisher instance, the Publisher helping user sending data into
	// Kafka topics, each instance has their own kafka Writer.
	NewPublisher(ctx context.Context, topicName string) (instance Publisher)
	// NewTopicsEditor create a topics editor instance, the Topic Editor helping user customize their
	// topic setting, each control will rise a new connection to Kafka cluster.
	NewTopicsEditor(ctx context.Context) (editor TopicsEditor)
}

// Subscriber client subscriber
type Subscriber interface {
	// Start trigger receiving message process
	// poolSize indicates how many go routine to concurrent receiving message
	// from message queue.
	Start(ctx context.Context, execute ExecuteFunc, groupID string, poolSize int, offsetFrom Offset) (err error)
	// Stop cancel receiving message process
	Stop(ctx context.Context) (err error)
}

// Publisher client publisher
type Publisher interface {
	// Send write the message to message queue, and user can use Message Key to determines the partition to
	// be thought of as an ordered sequence of messages
	Send(ctx context.Context, retry bool, writerMsg ...*WriterMessage) (failMsg []*WriterMessage, err error)
	//  SendByte sends a byte of data to message queue
	SendByte(ctx context.Context, retry bool, data []byte) error
	// SendWithBatch write the message items to message queue with fixed batch size.
	SendWithBatch(ctx context.Context, retry bool, size int, writerMsg *WriterMessage) (failMsg []interface{}, err error)
	// Close release the message queue publisher
	Close(ctx context.Context) (err error)
}

// TopicsEditor client topic editor
type TopicsEditor interface {
	// CreateTopics create topics with setting, the topic is a category or feed name to which message are publish/received
	// by producers/consumers
	CreateTopics(ctx context.Context, topics []*TopicConfig) (err error)
	// DeleteTopics delete topics from kafka cluster.
	DeleteTopics(ctx context.Context, topics []string) (err error)
}
