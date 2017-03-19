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

func (cli *Cli) Op(leaderaddr string, in *msg.Req) (*msg.Resp, error) {
	if cli.leader != leaderaddr {
		cli.getclient(leaderaddr)
	}
	return msg.NewKvClient(cli.conn).Op(context.TODO(), in)
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
