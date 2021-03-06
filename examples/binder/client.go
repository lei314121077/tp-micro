package main

import (
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/henrylee2cn/tp-micro"
)

func main() {
	tp.SetLoggerLevel("ERROR")
	cli := micro.NewClient(
		micro.CliConfig{
			Failover:        3,
			HeartbeatSecond: 4,
		},
		micro.NewStaticLinker(":9090"),
	)
	defer cli.Close()

	type Args struct {
		A int
		B int
	}

	var reply int
	rerr := cli.Pull("/static/p/divide?x=testquery_x&xy_z=testquery_xy_z", &Args{
		A: 10,
		B: 2,
	}, &reply).Rerror()
	if tp.IsConnRerror(rerr) {
		tp.Fatalf("has conn rerror: %v", rerr)
	}
	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("test 10/2=%d", reply)

	rerr = cli.Pull("/static/p/divide?x=testquery_x&xy_z=testquery_xy_z", &Args{
		A: 10,
		B: 0,
	}, &reply).Rerror()
	if tp.IsConnRerror(rerr) {
		tp.Fatalf("has conn rerror: %v", rerr)
	}
	if rerr == nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("test 10/0:%v", rerr)

	rerr = cli.Pull("/static/p/divide", &Args{
		A: 10,
		B: 5,
	}, &reply).Rerror()
	if rerr == nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("test 10/5:%v", rerr)
}
