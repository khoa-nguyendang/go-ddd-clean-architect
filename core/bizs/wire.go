//go:build wireinject
// +build wireinject

package bizs

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewJob,
)
