package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/prynnekey/mit6.824/Lab1/rpc/pb"
)

type SumService struct {
}

func (s *SumService) Sum(sumRequest *pb.SumRequest, sumResponse *pb.SumResponse) error {
	sumResponse.Result = sumRequest.A + sumRequest.B

	return nil
}

func main() {
	sumService := &SumService{}
	rpc.Register(sumService)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("Listening on port 1234")
	http.Serve(listener, nil)
}
