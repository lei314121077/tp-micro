// Code generated by 'ant gen' command.
// DO NOT EDIT!

package sdk

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"
	micro "github.com/henrylee2cn/tp-micro"
	"github.com/henrylee2cn/tp-micro/discovery"
	"github.com/henrylee2cn/tp-micro/discovery/etcd"

	"github.com/henrylee2cn/tp-micro/examples/sample/types"
)

var client *micro.Client

// Init initializes client with configs.
func Init(cliConfig micro.CliConfig, etcdConfing etcd.EasyConfig) {
	client = micro.NewClient(
		cliConfig,
		discovery.NewLinker(etcdConfing),
	)
}

// InitWithClient initializes client with specified object.
func InitWithClient(cli *micro.Client) {
	client = cli
}

// Home comment...
func Home(args *struct{}, setting ...socket.PacketSetting) (*types.HomeReply, *tp.Rerror) {
	reply := new(types.HomeReply)
	rerr := client.Pull("/sample/home", args, reply, setting...).Rerror()
	return reply, rerr
}

// Math_Divide handler
func Math_Divide(args *types.DivideArgs, setting ...socket.PacketSetting) (*types.DivideReply, *tp.Rerror) {
	reply := new(types.DivideReply)
	rerr := client.Pull("/sample/math/divide", args, reply, setting...).Rerror()
	return reply, rerr
}

// Stat comment...
func Stat(args *types.StatArgs, setting ...socket.PacketSetting) *tp.Rerror {
	return client.Push("/sample/stat", args, setting...)
}
