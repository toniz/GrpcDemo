package main

import (
    "context"
    "log"
    "net"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    pb "github.com/toniz/GrpcDemo/protos"

)

type StreamService struct{}

// Safe Channel is safe to use concurrently.
// Make sure that only one request controled the driver
type SafeChannel struct {
    ch  chan string
    mux chan struct{}
}

type ActionList map[string][]string

const (
    Address string = ":8000"
    Network string = "tcp"
)

var (
    chans [5]SafeChannel
    loginStatus [5]bool
    cmdlist ActionList
)

func init() {
    cmdlist = ActionList{
        "move": {"R_jointHome", "R_movel", "R_ChangeAttitude"},
        "stopmove": {"R_stopMove", "R_jointHome"},
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
        chans[i] = SafeChannel{
            ch: make(chan string, 0),
            mux: make(chan struct{}, 1),
        }
        chans[i].mux <- struct{}{}
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
    cmd := req.Cmd

    if loginStatus[driverId] == false {
        log.Printf("Driver[%d] Not Ready!", driverId)
        res := pb.Response{
            DriverId: driverId,
            Data: "driver not ready: " + req.Cmd,
        }
        return &res, nil
    }

    select {
    case <-chans[driverId].mux:
        log.Println("Receive Get Lock")
    default:
        res := pb.Response{
            DriverId: driverId,
            Data: "busy: " + req.Cmd,
        }
        return &res, nil
    }
    defer func(){ 
        chans[driverId].mux <- struct{}{}
        log.Println("Receive Release Lock")    
    }()

    for _, v := range cmdlist[action] {
        chans[driverId].ch <- v
    }
   
    res := pb.Response{
        DriverId: driverId,
        Data: "finish: " + req.Cmd,
    }       
   
    return &res, nil
}

// Call driver by stream
func (s *StreamService) StreamCall(srv pb.Guide_StreamCallServer) error {
    var driverId int32
    var seq int32
    if name, err := srv.Recv(); err != nil {
        log.Printf("Recv From Driver err: %v", err)
        return err
    } else {
        log.Printf("Driver driverId[%v] login", name.DriverId)
        driverId = name.DriverId
        seq = name.Seq
    }

    if loginStatus[driverId] == true {
        log.Printf("Driver driverId[%v] AlReady login", driverId)
        return status.Errorf(codes.AlreadyExists, "AlReady login!")
    }

    loginStatus[driverId] = true
    defer func(){loginStatus[driverId] = false}()

    for {
        var val string
        select {
            case val = <- chans[driverId].ch:
                log.Printf("Driver driverId[%v] Get Action [%s]!", driverId, val)
            case <-time.After(3 * time.Second):
                log.Printf("Driver driverId[%v] Timeout And Continue!", driverId)
                err := srv.Send(&pb.Response{
                    DriverId: driverId,
                    Seq: seq,
                    Ping: "PING",
                })
                
                if err != nil {
                    log.Printf("Clinet err: %v", err)
                    return err
                }

                continue
        }

        err := srv.Send(&pb.Response{
            DriverId: driverId,
            Seq: seq,
            Data: "Driver Do Action: " + val ,
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


