package main

import (
	"context"

	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

type Arith1 int

func (t *Arith1) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

type Arith2 int

func (t *Arith2) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B * 10
	return nil
}

func main() {
	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith1), "")
		s.Serve("tcp", "localhost:8972")
	}()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith2), "")
		s.Serve("tcp", "localhost:8973")
	}()

	select {}
}
