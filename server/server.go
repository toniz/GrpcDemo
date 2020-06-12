package main

import (
    //"context"
    "fmt"
    "net"
    "time"
//    "errors"

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
    fmtinStatus [5]bool
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
        fmt.Printf("net.Listen err: %v", err)
    }
    fmt.Println(Address + " net.Listing...")
    grpcServer := grpc.NewServer()
    pb.RegisterControlServer(grpcServer, &StreamService{})

    for i := range chans {
        chans[i] = SafeChannel{
            ch: make(chan string, 0),
            mux: make(chan struct{}, 1),
        }
        chans[i].mux <- struct{}{}
    }

    err = grpcServer.Serve(listener)
    if err != nil {
        fmt.Printf("grpcServer.Serve err: %v", err)
    }
}

// client call a sequence of actions
func (s *StreamService) Call(srv pb.Control_CallServer)  error {
    var driverId int32
    var cmd string
    if name, err := srv.Recv(); err != nil {
        fmt.Printf("Recv From Driver err: %v", err)
        return err
    } else {
        fmt.Printf("Driver driverId[%v] Client login", name.DriverId)
        driverId = name.DriverId
        cmd = name.Cmd
    }
    
    if fmtinStatus[driverId] == false {
        fmt.Printf("Driver[%d] Not Ready!", driverId)
        //err := srv.Send(&pb.Result{
        //    DriverId: driverId,
        //    Data: "driver not ready: " + cmd,
        //})
        return status.Errorf(codes.AlreadyExists, "Driver Not Ready!")
        //return errors.New("Driver Not Ready!")
    }
    
    select {
    case <-chans[driverId].mux:
        fmt.Println("Receive Get Lock")
    default:
        err := srv.Send(&pb.Result{
            DriverId: driverId,
            Data: "busy: " + cmd,
        })
        return err
    }
    defer func(){ 
        chans[driverId].mux <- struct{}{}
        fmt.Println("Receive Release Lock")    
    }()
        
    for {
        var driverId int32
        var cmd string
        if name, err := srv.Recv(); err != nil {
            fmt.Printf("Recv From Driver err: %v", err)
            return err
        } else {
            fmt.Printf("Driver driverId[%v] cmd", name.DriverId)
            driverId = name.DriverId
            cmd = name.Cmd
        }

        for _, v := range cmdlist[cmd] {
            select {
            case chans[driverId].ch <- v:
         //       fmt.Printf("Send %s", v)
            case <-time.After(5 * time.Second):
                fmt.Println("Send Timeout!")
                err := srv.Send(&pb.Result{
                    DriverId: driverId,
                    Data: "Call Driver Timeout: " + cmd,
                })
                return err
            }
        }
        
        srv.Send(&pb.Result{
            DriverId: driverId,
            Data: "finish: " + cmd,
        })
    }
}

// Call driver by stream
func (s *StreamService) StreamCall(srv pb.Control_StreamCallServer) error {
    var driverId int32
    var seq int32
    if name, err := srv.Recv(); err != nil {
        fmt.Printf("Recv From Driver err: %v", err)
        return err
    } else {
        fmt.Printf("Driver driverId[%v] login", name.DriverId)
        driverId = name.DriverId
        seq = name.Seq
    }

    if fmtinStatus[driverId] == true {
        fmt.Printf("Driver driverId[%v] AlReady login", driverId)
        return status.Errorf(codes.AlreadyExists, "AlReady fmtin!")
    }

    fmtinStatus[driverId] = true
    defer func(){fmtinStatus[driverId] = false}()

    for {
        var val string
        select {
            case val = <- chans[driverId].ch:
                fmt.Printf("Driver driverId[%v] Get Action [%s]!", driverId, val)
            case <-time.After(3 * time.Second):
                fmt.Printf("Driver driverId[%v] Timeout And Continue!", driverId)
                err := srv.Send(&pb.Response{
                    DriverId: driverId,
                    Seq: seq,
                    Ping: "PING",
                })
                
                if err != nil {
                    fmt.Printf("Clinet err: %v", err)
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
            fmt.Printf("Clinet err: %v", err)
            return err
        }

        res, err := srv.Recv()
        if err != nil {
            fmt.Printf("Recv From Clinet err: %v", err)
            return err
        }

        if seq != res.Seq {
            fmt.Printf("Seq %d != %d ", seq, res.Seq)
        }

        seq++
        fmt.Printf("Recv From Driver: %v", res)
    }
}


