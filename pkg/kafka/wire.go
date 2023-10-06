//go:build wireinject
// +build wireinject

package kafka

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewClient,
	NewPublisher,
	NewSubscriber,
	NewTopicsEditor,
)
