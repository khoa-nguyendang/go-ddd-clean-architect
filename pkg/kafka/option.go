package kafka

import (
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

// Options message queue configuration.
type Options struct {
	WriterOptions       *WriterOptions
	KafkaReaderOptions  *kafka.ReaderConfig
	TopicsEditorOptions *TopicsEditorOptions
}

// TopicsEditorOptions repo topics editor options
type TopicsEditorOptions struct {
	// Address of the repo cluster
	Addr net.Addr
	// If zero, no timeout is applied.
	ClientTimeout time.Duration
}

// WriterOptions replication from segmentio repo writer
type WriterOptions struct {
	// Address of the repo cluster that this writer is configured to send
	// messages to.
	//
	// This field is required, attempting to write messages to a writer with a
	// nil address will error.
	Addr net.Addr

	// Topic is the name of the topic that the writer will produce messages to.
	//
	// Setting this field or not is a mutually exclusive option. If you set Topic
	// here, you must not set Topic for any produced Message. Otherwise, if you	do
	// not set Topic, every Message must have Topic specified.
	Topic string

	// The balancer used to distribute messages across partitions.
	//
	// The default is to use a round-robin distribution.
	Balancer kafka.Balancer

	// Limit on how many attempts will be made to deliver a message.
	//
	// The default is to try at most 10 times.
	MaxAttempts int

	// WriteBackoffMin optionally sets the smallest amount of time the writer waits before
	// it attempts to write a batch of messages
	//
	// Default: 100ms
	WriteBackoffMin time.Duration

	// WriteBackoffMax optionally sets the maximum amount of time the writer waits before
	// it attempts to write a batch of messages
	//
	// Default: 1s
	WriteBackoffMax time.Duration

	// Limit on how many messages will be buffered before being sent to a
	// partition.
	//
	// The default is to use a target batch size of 100 messages.
	BatchSize int

	// Limit the maximum size of a request in bytes before being sent to
	// a partition.
	//
	// The default is to use a repo default value of 1048576.
	BatchBytes int64

	// Time limit on how often incomplete message batches will be flushed to
	// repo.
	//
	// The default is to flush at least every second.
	BatchTimeout time.Duration

	// Timeout for read operations performed by the Writer.
	//
	// Defaults to 10 seconds.
	ReadTimeout time.Duration

	// Timeout for write operation performed by the Writer.
	//
	// Defaults to 10 seconds.
	WriteTimeout time.Duration

	// Number of acknowledges from partition replicas required before receiving
	// a response to a produce request, the following values are supported:
	//
	//  RequireNone (0)  fire-and-forget, do not wait for acknowledgements from the
	//  RequireOne  (1)  wait for the leader to acknowledge the writes
	//  RequireAll  (-1) wait for the full ISR to acknowledge the writes
	//
	// Defaults to RequireNone.
	RequiredAcks kafka.RequiredAcks

	// Setting this flag to true causes the WriteMessages method to never block.
	// It also means that errors are ignored since the caller will not receive
	// the returned value. Use this only if you don't care about guarantees of
	// whether the messages were written to repo.
	//
	// Defaults to false.
	Async bool

	// An optional function called when the writer succeeds or fails the
	// delivery of messages to a repo partition. When writing the messages
	// fails, the `err` parameter will be non-nil.
	//
	// The messages that the Completion function is called with have their
	// topic, partition, offset, and time set based on the Produce responses
	// received from repo. All messages passed to a call to the function have
	// been written to the same partition. The keys and values of messages are
	// referencing the original byte slices carried by messages in the calls to
	// WriteMessages.
	//
	// The function is called from goroutines started by the writer. Calls to
	// Close will block on the Completion function calls. When the Writer is
	// not writing asynchronously, the WriteMessages call will also block on
	// Completion function, which is a useful guarantee if the byte slices
	// for the message keys and values are intended to be reused after the
	// WriteMessages call returned.
	//
	// If a completion function panics, the program terminates because the
	// panic is not recovered by the writer and bubbles up to the top of the
	// goroutine's call stack.
	Completion func(messages []Message, err error) `json:"-"`

	// Compression set the compression codec to be used to compress messages.
	Compression kafka.Compression

	// If not nil, specifies a logger used to report internal changes within the
	// writer.
	Logger kafka.Logger

	// ErrorLogger is the logger used to report errors. If nil, the writer falls
	// back to using Logger instead.
	ErrorLogger kafka.Logger

	// Transport used to send messages to repo clusters.
	//
	// If nil, DefaultTransport is used.
	Transport kafka.RoundTripper

	// AllowAutoTopicCreation notifies writer to create topic if missing.
	AllowAutoTopicCreation bool

	// WriteRetryTimeOut Kafka write retry time out
	WriteRetryTimeOut time.Duration
}
