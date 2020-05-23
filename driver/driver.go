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
var channum int

func init() {
    flag.IntVar(&channum, "channum", 1, "chan numbers selected")
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
        streamcall()
    }
}


func streamcall() {
    stream, err := guideClient.StreamCall(context.Background())
    if err != nil {
        log.Printf("get stream call err: %v", err)
    }

    log.Printf("Current Driver: %d", channum)

    var seq int32
    seq = 0
    // login to server
    err = stream.Send(&pb.Request{
        DriverId: int32(channum),
        Seq: seq,
        Data: "stream client rpc "+strconv.Itoa(int(seq)),
    })

    if err != nil {
        log.Printf("stream login err: %v", err)
    }

    for {
        res, err := stream.Recv()
        if err != nil {
            log.Printf("StreamCall Recv err: %v", err)
            break
        }

        log.Println(res)
        
        if seq != res.Seq {
            log.Printf("Seq %d != %d ", seq, res.Seq)
        }
        time.Sleep(time.Second)

        err = stream.Send(&pb.Request{
            DriverId: int32(channum),
            Seq: seq,
            Data: "stream client rpc " + strconv.Itoa(int(seq)),
        })
        
        if err != nil {
            log.Printf("stream login err: %v", err)
        }

        seq++
    }

    err = stream.CloseSend()
    if err != nil {
        log.Printf("Conversations close stream err: %v", err)
    }
}


