package main

import (
	"context"
	"time"

	"github.com/cloudwego/netpoll"
)

var eventLoop netpoll.EventLoop

func main() {
	//listener, err := net.Listen(network, address)
	//listener, err := netpoll.CreateListener(network, address)
	//if err != nil {
	//	panic("create net listener failed")
	//}
	//_ = listener.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = eventLoop.Shutdown(ctx)
}
