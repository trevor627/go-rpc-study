package main

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeepwe address")
	basePath = flag.String("base", "/rpc_test/Arith", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}
	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
