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

// Safe Channel is safe to use concurrently.
// Make sure that only one request controled the driver
type SafeChannel struct {
    ch  chan string
    inuse bool
    mux sync.Mutex
}

type ActionList map[string][]string

const (
    Address string = ":8000"
    Network string = "tcp"
)

var (
    chans [5]SafeChannel
    cmdlist ActionList
)

func init() {
    cmdlist = ActionList{
        "move": ["R_jointHome", "R_movel", "R_ChangeAttitude"],
        "stopmove": ["R_stopMove", "R_jointHome"],
    }
}

func main() {
    listener, err := net.Listen(Network, Address)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }
    log.Println(Address + " net.Listing...")
    grpcServer := grpc.NewServer()
    pb.RegisterGuideServer(grpcServer, &StreamService{})

    for i := range chans {
        chans[i] = make(chan string, 0)
    }

    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatalf("grpcServer.Serve err: %v", err)
    }
}

// client call a sequence of actions
func (s *StreamService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    log.Println("Receive Request: ", req)
    driverId := req.DriverId
    if chans[driverId].inuse {
        res := pb.Response{
            DriverId: driverId,
            Data: "busy: " + req.Data,
        }
        return &res, nil
    }

    chans[driverId].mux.Lock()
    chans[driverId].inuse = true

    for v := range cmdlist[driverId] {
        chans[req.DriverId] <- v
    }

    res := pb.Response{
        DriverId: driverId,
        Data: "finish: " + req.Data,
    }       

    chans[driverId].inuse = false
    chans[driverId].mux.Unlock()
    return &res, nil
}

// 流式调用
func (s *StreamService) StreamCall(srv pb.Guide_StreamCallServer) error {
    var code int32
    var seq int32
    if name, err := srv.Recv(); err != nil {
        log.Printf("Recv From Driver err: %v", err)
        return err
    } else {
        log.Printf("Driver code[%v] login", name.DriverId)
        code = name.DriverId
    }

    for {
        val := <- chans[code]
        err := srv.Send(&pb.Response{
            DriverId: code,
            Seq: seq,
            Data: "The Server CChan Val: " + val ,
        })

        if err != nil {
            log.Printf("Clinet err: %v", err)
            return err
        }

        res, err := srv.Recv()
        if err != nil {
            log.Printf("Recv From Clinet err: %v", err)
            return err
        }

        if seq != res.Seq {
            log.Printf("Seq %d != %d ", seq, res.Seq)
        }

        seq++
        log.Printf("Recv From Driver: %v", res)
    }
}


