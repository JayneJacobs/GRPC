package main

import (
	"context"
"github.com/JayneJacobs/grpc/echo")

type EchoServer struct {}

func (e *EchoServer) Echo(context.Context, *echo.EchoRequest) (*echo.EchoResponse, error) {

}

func main() {

}
