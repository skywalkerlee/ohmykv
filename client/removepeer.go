package main

import (
	"context"
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/msg"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("args Error")
		return
	}
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	client := msg.NewKvClient(conn)
	resp, err := client.Man(context.Background(), &msg.Manreq{Op: 3, Value: []byte("127.0.0.1:12004")})
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	logs.Debug(resp)
}
