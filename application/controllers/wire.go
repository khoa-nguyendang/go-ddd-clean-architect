package controllers

import "github.com/google/wire"

var Wired = wire.NewSet(
	NewConfig,
	New,
)
