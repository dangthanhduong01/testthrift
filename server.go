package main

import (
	"context"
	"fmt"

	"thriftwithgo/gen-go/trythrift"

	"github.com/apache/thrift/lib/go/thrift"
)

type Calculator struct{}

func (h *Calculator) Add(ctx context.Context, num1, num2 int32) (int32, error) {
	return num1 + num2, nil
}

func (h *Calculator) Multiply(ctx context.Context, num1, num2 int32) (int32, error) {
	return num1 * num2, nil
}

func main() {
	handler := &Calculator{}
	processor := trythrift.NewCalculatorProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		fmt.Println("Error!", err)
		return
	}
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	fmt.Println("Starting the server...")
	server.Serve()
}
