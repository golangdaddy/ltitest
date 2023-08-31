package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "example.com/project/protos/arithmetic"
)

const (
	CONST_PORT = "8588"
)

func main() {
	conn, err := grpc.Dial("localhost:"+CONST_PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewArithmeticServiceClient(conn)

	req := &pb.ArithmeticRequest{
		Arg: 1,
	}

	res, err := c.SendArithmetic(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Println("Response:", res.GetResult())
}
