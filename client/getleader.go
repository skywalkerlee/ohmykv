package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/msg"
	"google.golang.org/grpc"
)

func main() {
	addr := []string{"127.0.0.1:33881", "127.0.0.1:33882", "127.0.0.1:33883"}
	conn, err := grpc.Dial(addr[rand.New(rand.NewSource(time.Now().Unix())).Intn(3)], grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	client := msg.NewKvClient(conn)
	resp, err := client.Man(context.Background(), &msg.Manreq{Op: 1})
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	logs.Info(resp.String())
}
