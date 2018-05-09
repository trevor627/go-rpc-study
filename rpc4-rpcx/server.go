package main

import (
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	"golang.org/x/net/context"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.Server{}
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)
}
