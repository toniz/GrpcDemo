package main

import (
    "context"
    "io"
    "log"
    "strconv"

    "time"
    "google.golang.org/grpc"
    pb "github.com/toniz/GrpcDemo/protos"
)

const Address string = ":8000"

var guideClient pb.GuideClient

var chans [5]chan string
for i := range chans {
   chans[i] = make(chan string)
}

func main() {

    seq := 0
    for {
        seq++
        log.Printf("Strat New Connection: %d", seq)
        conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
            log.Printf("Connect Failed: %v", err)
            time.Sleep(time.Second)
            continue
        }

        defer conn.Close()
        log.Printf("Connect Success: %v", conn)

        guideClient = pb.NewGuideClient(conn)
        call()
        stream()
    }
}


func call() {
    req := pb.Request{
        Data: "grpc",
    }

    res, err := guideClient.Call(context.Background(), &req)
    if err != nil {
        log.Printf("Call Server err: %v", err)
    }

    log.Println(res.Data)
}

func stream() {
    stream, err := guideClient.StreamCall(context.Background())
    if err != nil {
        log.Printf("get stream call err: %v", err)
    }

    err := stream.Send(&pb.Request{Data: "stream client rpc " + strconv.Itoa(n)})
    if err != nil {
        log.Printf("stream request err: %v", err)
    }

    for n := 0; ; n++ {
        cmd := 
        err := stream.Send(&pb.Request{Data: "stream client rpc " + strconv.Itoa(n)})
        if err != nil {
            log.Printf("stream request err: %v", err)
        }
        res, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("StreamCall Recv err: %v", err)
            break
        }
        log.Println(res.Data)
    }

    err = stream.CloseSend()
    if err != nil {
        log.Printf("Conversations close stream err: %v", err)
    }
}



