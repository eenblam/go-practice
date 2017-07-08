package main

import "errors"
import "fmt"
import "log"
import "net"
import "net/rpc"
import "net/rpc/jsonrpc"

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Add(args *Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

func (t *Arith) Subtract(args *Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func startServer(ch chan<- bool) {
    arith := new(Arith)

    server := rpc.NewServer()
    server.Register(arith)

    listener, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }

    // Conn
    defer listener.Close()
    // Accept loop
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }

        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
        ch <- true
    }

    var input string
    fmt.Scanln(&input)
}

func main() {
    ch := make(chan bool)

    go startServer(ch)

    for {
        <-ch
        log.Println("Closed")
    }
}
