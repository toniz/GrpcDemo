package main

import (
    "context"
//    "io"
    "log"
    "strconv"
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
    flag.IntVar(&dreverId, "Driver ID", 1, "Driver ID selected")
    flag.StringVar(&actionId, "Action ID", "move", "Action ID selected")
}

func main() {
    cnt := 0
    for {
        cnt++
        log.Printf("Strat New Connection: %d", cnt)
        conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
            log.Printf("Connect Failed: %v", err)
            time.Sleep(time.Second)
            continue
        }

        defer conn.Close()

        guideClient = pb.NewGuideClient(conn)
        call()
    }
}

func call() {
    req := pb.Request{
        DriverId: int32(dreverId),
        Data: actionId,
    }

    res, err := guideClient.Call(context.Background(), &req)
    if err != nil {
        log.Printf("Call Server err: %v", err)
    }

    log.Println(res.Data)
}

