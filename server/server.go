package main

import (
    "context"
//    "io"
    "log"
    "net"
//    "strconv"
    "google.golang.org/grpc"

    "time"
    pb "github.com/toniz/GrpcDemo/protos"

)

type StreamService struct{}

const (
    Address string = ":8000"
    Network string = "tcp"
)

var chans [5]chan string


func main() {
    listener, err := net.Listen(Network, Address)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }
    log.Println(Address + " net.Listing...")
    grpcServer := grpc.NewServer()
    pb.RegisterGuideServer(grpcServer, &StreamService{})

    for i := range chans {
        chans[i] = make(chan string)
    }

    go func(){
       chans[1] <- "test 1"
       time.Sleep(1)
    }()

    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatalf("grpcServer.Serve err: %v", err)
    }
}

// 单步调用
func (s *StreamService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    res := pb.Response{
        Data: "hello " + req.Data,
    }
    return &res, nil
}

// 流式调用
func (s *StreamService) StreamCall(srv pb.Guide_StreamCallServer) error {
    var code int32
    var seq int32
    if name, err := srv.Recv(); err != nil {
        log.Printf("Recv From Client err: %v", err)
        return err
    } else {
        log.Printf("Client code[%v] login", name.Client_id)
        code = name.Client_id
    }

    for {

        val := <- chans[code]

        err := srv.Send(&pb.Response{
            Client_id: code,
            Seq: seq,
            Data: "The Server CChan Val: " + strconv.Itoa(val) ,
        })

        if err != nil {
            return err
        }

        req, err := srv.Recv()
        if err != nil {
            log.Printf("Recv From Clinet err: %v", err)
            return err
        }

        seq++
        log.Printf("Recv From Client: %v", req)
    }
}


