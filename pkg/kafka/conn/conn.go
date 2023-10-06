package conn

import (
	"context"
	"net"
	"sync"
)

// WaitGroup connect wait group for net dial
type WaitGroup struct {
	DialFunc func(context.Context, string, string) (net.Conn, error)
	sync.WaitGroup
}

// Dial callback dial function in wait group
func (g *WaitGroup) Dial(ctx context.Context, network, address string) (net.Conn, error) {
	c, err := g.DialFunc(ctx, network, address)
	if err != nil {
		return nil, err
	}
	g.Add(1)
	return &groupConn{Conn: c, group: g}, nil
}

type groupConn struct {
	net.Conn
	group *WaitGroup
	once  sync.Once
}

func (c *groupConn) Close() error {
	defer c.once.Do(c.group.Done)
	return c.Conn.Close()
}
