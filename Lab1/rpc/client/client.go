package main

import (
	"fmt"
	"net/rpc"

	"github.com/prynnekey/mit6.824/Lab1/rpc/pb"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	sumRequest := &pb.SumRequest{
		A: 10,
		B: 88,
	}
	sumResponse := &pb.SumResponse{}

	err = conn.Call("SumService.Sum", sumRequest, sumResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println("data", sumResponse.Result)
}
