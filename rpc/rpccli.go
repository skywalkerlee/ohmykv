package rpc

import (
	"context"

	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/msg"
	"google.golang.org/grpc"
)

type Cli struct {
	lock   *sync.Mutex
	leader string
	conn   *grpc.ClientConn
}

var cli *Cli

func NewCli() *Cli {
	if cli == nil {
		cli = &Cli{
			lock: new(sync.Mutex),
		}
	}
	return cli
}

func (cli *Cli) getclient(leaderaddr string) error {
	cli.lock.Lock()
	defer cli.lock.Unlock()
	if cli.leader == leaderaddr {
		return nil
	}
	if cli.conn != nil {
		cli.conn.Close()
	}
	conn, err := grpc.Dial(leaderaddr, grpc.WithInsecure())
	if err != nil {
		logs.Error(err)
		return err
	}
	cli.leader = leaderaddr
	cli.conn = conn
	return nil
}

func (cli *Cli) Write(leaderaddr string, in *msg.Writereq) (*msg.Writeresp, error) {
	if cli.leader != leaderaddr {
		cli.getclient(leaderaddr)
	}
	return msg.NewKvClient(cli.conn).Write(context.Background(), in)
}

func (cli *Cli) Man(leaderaddr string, in *msg.Manreq) (*msg.Manresp, error) {
	if cli.leader != leaderaddr {
		cli.getclient(leaderaddr)
	}
	return msg.NewKvClient(cli.conn).Man(context.TODO(), in)
}
