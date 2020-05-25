package main

import (
    "context"
//    "io"
    "log"
//    "strconv"
    "flag"

    "time"
    "google.golang.org/grpc"
    pb "github.com/toniz/GrpcDemo/protos"
)

const Address string = ":8000"

var guideClient pb.GuideClient
var dreverId int
var actionId string

func init() {
    flag.IntVar(&dreverId, "d", 1, "Driver ID selected")
    flag.StringVar(&actionId, "a", "move", "Action ID selected")
}

func main() {
    flag.Parse()
    conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Printf("Connect Failed: %v", err)
        time.Sleep(time.Second)
    } else {
        guideClient = pb.NewGuideClient(conn)
        call()
    }
    defer conn.Close()
}

func call() {
    req := pb.Request{
        DriverId: int32(dreverId),
        Data: actionId,
    }

    log.Println(req)
    res, err := guideClient.Call(context.Background(), &req)
    if err != nil {
        log.Printf("Call Server err: %v", err)
    }

    log.Println(res.Data)
}

