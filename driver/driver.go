package main

import (
    "context"
//    "io"
    "log"
    "strconv"
    "flag"

    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    pb "github.com/toniz/GrpcDemo/protos"
)

const Address string = ":8000"

var guideClient pb.ControlClient
var channum int

func init() {
    flag.IntVar(&channum, "c", 1, "chan numbers selected")
}

func main() {
    flag.Parse()
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

        guideClient = pb.NewControlClient(conn)
        err = streamcall()
        if err != nil {
            log.Printf("streamcall Failure: %v", err)
            s := status.Convert(err)
            if s.Code() == codes.AlreadyExists {
                log.Fatalln("Already Exists DriverID")
            }
        }
    }
}

func streamcall() error {
    stream, err := guideClient.StreamCall(context.Background())
    if err != nil {
        log.Printf("get stream call err: %v", err)
    }
    defer stream.CloseSend()

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
        return err
    }

    for {
        res, err := stream.Recv()
        if err != nil {
            return err
        }

        log.Println(res)
  
        if res.Ping == "PING" {
            log.Printf("StreamCall Recevie PING From Server")
            continue
        }
      
        if seq != res.Seq {
            log.Printf("Seq %d != %d ", seq, res.Seq)
        }

        // test function, remove it.
        time.Sleep(time.Second)

        err = stream.Send(&pb.Request{
            DriverId: int32(channum),
            Seq: seq,
            Data: "Driver Finish: " + strconv.Itoa(int(seq)),
        })
        
        if err != nil {
            log.Printf("stream login err: %v", err)
            return err
        }

        seq++
    }
}


