package routing

import (
	"context"

	"v2ray.com/core/common/net"
	"v2ray.com/core/common/vio"
	"v2ray.com/core/features"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a V2Ray instance to make V2Ray function properly.
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*vio.Link, error)
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
