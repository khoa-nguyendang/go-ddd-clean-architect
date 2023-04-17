//go:build wireinject
// +build wireinject

//
package mysql

import "github.com/google/wire"

var Wired = wire.NewSet(
	New,
)
