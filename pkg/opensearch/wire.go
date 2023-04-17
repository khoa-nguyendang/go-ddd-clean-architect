//go:build wireinject
// +build wireinject

package opensearch

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewConfig,
	New,
)
