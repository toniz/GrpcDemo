package main

import (
    "context"
    "io"
    "log"
    "net"
    "strconv"
    "google.golang.org/grpc"

    pb "https://github.com/toniz/GrpcDemo/protos"

)

type StreamService struct{}

const (
    Address string = ":8000"
    Network string = "tcp"
)

func main() {
    listener, err := net.Listen(Network, Address)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }
    log.Println(Address + " net.Listing...")
    grpcServer := grpc.NewServer()
    pb.RegisterStreamServer(grpcServer, &StreamService{})

    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatalf("grpcServer.Serve err: %v", err)
    }
}

// 单步调用
func (s *StreamService) Control(ctx context.Context, req *pb.DriverRequest) (*pb.DriverResponse, error) {
    res := pb.SimpleResponse{
        Value: "hello " + req.Data,
    }
    return &res, nil
}

// 流式调用
func (s *StreamService) StreamControl(srv pb.Stream_ConversationsServer) error {
    n := 1
    for {
        req, err := srv.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
        err = srv.Send(&pb.StreamResponse{
            Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
        })
        if err != nil {
            return err
        }
        n++
        log.Printf("from stream client question: %s", req.Question)
    }
}


