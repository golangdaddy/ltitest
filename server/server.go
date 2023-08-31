package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "example.com/project/protos/arithmetic"
)

const (
	CONST_PORT = "8588"
)

type App struct {
	additionSendChannel    chan int
	additionRespondChannel chan int
	pb.UnimplementedArithmeticServiceServer
}

func (app *App) SendArithmetic(ctx context.Context, req *pb.ArithmeticRequest) (*pb.ArithmeticResponse, error) {
	arg := req.GetArg()
	app.additionSendChannel <- int(arg)
	return &pb.ArithmeticResponse{Result: int32(<-app.additionRespondChannel)}, nil
}

func main() {

	app := &App{
		additionSendChannel:    make(chan int),
		additionRespondChannel: make(chan int),
	}

	go app.handleArithmetic()

	lis, err := net.Listen("tcp", ":"+CONST_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(s, app)

	log.Println("Server started on port " + CONST_PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (app *App) handleArithmetic() {

	var runningTotal int

	for msg := range app.additionSendChannel {

		runningTotal += msg

		app.additionRespondChannel <- runningTotal

	}

}
