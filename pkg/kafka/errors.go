package kafka

import (
	"errors"
)

var (
	// ErrorZeroTopics the input kafka topics length are zero
	ErrorZeroTopics = errors.New("the input kafka topics length are zero")

	// ErrorGroupIDAlreadySubscribed group id already subscribed
	ErrorGroupIDAlreadySubscribed = errors.New("group id already subscribed")

	// ErrorReceiveFail kafka message receive fail
	ErrorReceiveFail = errors.New("kafka message receive fail")

	// ErrorCommitFail kafka message commit fail
	ErrorCommitFail = errors.New("kafka message commit fail")

	// ErrorExecuteFail callback function execute fail
	ErrorExecuteFail = errors.New("callback function execute fail")

	// ErrorZeroWriteMsg the input kafka write message length are zero
	ErrorZeroWriteMsg = errors.New("the input kafka write message length are zero")

	// ErrorNilWriteMsg the input kafka write message are nil
	ErrorNilWriteMsg = errors.New("the input kafka write message are nil")

	// ErrorNotAbleConvertToSlice the input interface is not able to convert to slice
	ErrorNotAbleConvertToSlice = errors.New("the input interface is not able to convert to slice")

	// ErrorKafkaRetryWriteTimeout ...
	ErrorKafkaRetryWriteTimeout = errors.New("error kafka write msg timeout")
)
