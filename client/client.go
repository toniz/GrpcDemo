package main

import (
    "context"
    "log"
    "flag"

    "time"
    "google.golang.org/grpc"
    pb "github.com/toniz/GrpcDemo/protos"
)

const Address string = ":8000"

var guideClient pb.ControlClient
var dreverId int
var cmdId string

func init() {
    flag.IntVar(&dreverId, "d", 1, "Driver ID selected")
    flag.StringVar(&cmdId, "c", "move", "Action ID selected")
}

func main() {
    flag.Parse()
    conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Printf("Connect Failed: %v", err)
        time.Sleep(time.Second)
    } else {
        guideClient = pb.NewControlClient(conn)
        call()
    }
    defer conn.Close()
}

func call() {
    req := pb.Command{
        DriverId: int32(dreverId),
        Cmd: cmdId,
    }

    log.Println(req)
    res, err := guideClient.Call(context.Background(), &req)
    if err != nil {
        log.Printf("Call Server err: %v", err)
    }

    log.Println(res.Data)
}

