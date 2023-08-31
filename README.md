# LTI Test

## Test Brief:


Write a client-server program using gRPC in golang. The client should send an arithmetic expression to the server, which will evaluate them using goroutines and channels and return the output.

## Test instructions

Each time the client is invoked it should increment the counter.

```
go run ./server
go run ./client
```