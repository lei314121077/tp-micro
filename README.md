# TP-Micro [![GitHub release](https://img.shields.io/github/release/henrylee2cn/tp-micro.svg?style=flat-square)](https://github.com/henrylee2cn/tp-micro/releases) [![report card](https://goreportcard.com/badge/github.com/henrylee2cn/tp-micro?style=flat-square)](http://goreportcard.com/report/henrylee2cn/tp-micro) [![github issues](https://img.shields.io/github/issues/henrylee2cn/tp-micro.svg?style=flat-square)](https://github.com/henrylee2cn/tp-micro/issues?q=is%3Aopen+is%3Aissue) [![github closed issues](https://img.shields.io/github/issues-closed-raw/henrylee2cn/tp-micro.svg?style=flat-square)](https://github.com/henrylee2cn/tp-micro/issues?q=is%3Aissue+is%3Aclosed) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/henrylee2cn/tp-micro) [![view examples](https://img.shields.io/badge/learn%20by-examples-00BCD4.svg?style=flat-square)](https://github.com/henrylee2cn/tp-micro/tree/master/samples)
<!-- [![view Go网络编程群](https://img.shields.io/badge/官方QQ群-Go网络编程(42730308)-27a5ea.svg?style=flat-square)](http://jq.qq.com/?_wv=1027&k=fzi4p1) -->


TP-Micro is a simple, powerful micro service framework based on [Teleport](https://github.com/henrylee2cn/teleport).

[简体中文](https://github.com/henrylee2cn/tp-micro/blob/master/README_ZH.md)


## Install


```
go version ≥ 1.9
```

```sh
go get -u -f github.com/henrylee2cn/tp-micro
```

## Feature

- Support auto service-discovery
- Supports custom service linker
- Support load balancing
- Support NIO and connection pool
- Support custom protocol
- Support custom body codec
- Support plug-in expansion
- Support heartbeat mechanism
- Detailed log information, support print input and output details
- Support for setting slow operation alarm thresholds
- Support for custom log
- Support smooth shutdown and update
- Support push handler
- Support network list: `tcp`, `tcp4`, `tcp6`, `unix`, `unixpacket` and so on
- Client support automatically redials after disconnection

## Platform Case

[Ants](https://github.com/xiaoenai/ants): A highly available micro service platform based on [TP-Micro](https://github.com/henrylee2cn/tp-micro) and [Teleport](https://github.com/henrylee2cn/teleport).


## Demo

- server

```go
package main

import (
  micro "github.com/henrylee2cn/tp-micro"
  tp "github.com/henrylee2cn/teleport"
)

// Args args
type Args struct {
  A int
  B int `param:"<range:1:>"`
}

// P handler
type P struct {
  tp.PullCtx
}

// Divide divide API
func (p *P) Divide(args *Args) (int, *tp.Rerror) {
  return args.A / args.B, nil
}

func main() {
  srv := micro.NewServer(micro.SrvConfig{
    ListenAddress: ":9090",
  })
  srv.RoutePull(new(P))
  srv.Listen()
}
```

- client

```go
package main

import (
  micro "github.com/henrylee2cn/tp-micro"
  tp "github.com/henrylee2cn/teleport"
)

func main() {
  cli := micro.NewClient(
    micro.CliConfig{},
    micro.NewStaticLinker(":9090"),
  )
  defer cli.Close()

  type Args struct {
    A int
    B int
  }

  var reply int
  rerr := cli.Pull("/p/divide", &Args{
    A: 10,
    B: 2,
  }, &reply).Rerror()
  if rerr != nil {
    tp.Fatalf("%v", rerr)
  }
  tp.Infof("10/2=%d", reply)
  rerr = cli.Pull("/p/divide", &Args{
    A: 10,
    B: 0,
  }, &reply).Rerror()
  if rerr == nil {
    tp.Fatalf("%v", rerr)
  }
  tp.Infof("test binding error: ok: %v", rerr)
}
```

[More Samples](https://github.com/henrylee2cn/tp-micro/tree/master/samples)

## Project Management

Command ant is deployment tools of ant microservice frameware.

- Quickly create a ant project
- Run ant project with hot compilation

### Install Ant Command

```sh
go get -u -f -d github.com/xiaoenai/ants/...
cd $GOPATH/src/github.com/xiaoenai/ants/ant
go install
```

### Generate project

`ant gen` command help:

```
NAME:
   ant gen - Generate an ant project

USAGE:
   ant gen [command options] [arguments...]

OPTIONS:
   --template value, -t value    The template for code generation(relative/absolute)
   --app_path value, -p value  The path(relative/absolute) of the project
```

example: `ant gen -t ./__ant__tpl__.go -p ./myant` or default `ant gen myant`

- template file `__ant__tpl__.go` demo:

```go
// package __ANT__TPL__ is the project template
package __ANT__TPL__

// __API__PULL__ register PULL router:
//  /home
//  /math/divide
type __API__PULL__ interface {
  Home(*struct{}) *HomeReply
  Math
}

// __API__PUSH__ register PUSH router:
//  /stat
type __API__PUSH__ interface {
  Stat(*StatArgs)
}

// Math controller
type Math interface {
  // Divide handler
  Divide(*DivideArgs) *DivideReply
}

// HomeReply home reply
type HomeReply struct {
  Content string // text
}

type (
  // DivideArgs divide api args
  DivideArgs struct {
    // dividend
    A float64
    // divisor
    B float64 `param:"<range: 0.01:100000>"`
  }
  // DivideReply divide api result
  DivideReply struct {
    // quotient
    C float64
  }
)

// StatArgs stat handler args
type StatArgs struct {
  Ts int64 // timestamps
}
```

- The template generated by `ant gen` command.

```
├── README.md
├── main.go
├── api
│   ├── handlers.gen.go
│   ├── handlers.go
│   ├── router.gen.go
│   └── router.go
├── logic
│   └── tmp_code.gen.go
├── sdk
│   ├── rpc.gen.go
│   ├── rpc.gen_test.go
│   ├── rpc.go
│   └── rpc_test.go
└── types
    ├── types.gen.go
    └── types.go
```

Desc:

- add `.gen` suffix to the file name of the automatically generated file
- `tmp_code.gen.go` is temporary code used to ensure successful compilation!<br>When the project is completed, it should be removed!

### Run project

`ant run` command help:

```
NAME:
   ant run - Compile and run gracefully (monitor changes) an any existing go project

USAGE:
   ant run [options] [arguments...]
 or
   ant run [options except -app_path] [arguments...] {app_path}

OPTIONS:
   --watch_exts value, -x value  Specified to increase the listening file suffix (default: ".go", ".ini", ".yaml", ".toml", ".xml")
   --app_path value, -p value    The path(relative/absolute) of the project
```

example: `ant run -x .yaml -p myant` or `ant run`

[More Ant Command](https://github.com/xiaoenai/ants/tree/master/ant)

## Usage

### Peer(server or client) Demo

```go
// Start a server
var peer1 = tp.NewPeer(tp.PeerConfig{
    ListenAddress: "0.0.0.0:9090", // for server role
})
peer1.Listen()

...

// Start a client
var peer2 = tp.NewPeer(tp.PeerConfig{})
var sess, err = peer2.Dial("127.0.0.1:8080")
```

### Pull-Controller-Struct API template

```go
type Aaa struct {
    tp.PullCtx
}
func (x *Aaa) XxZz(args *<T>) (<T>, *tp.Rerror) {
    ...
    return r, nil
}
```

- register it to root router:

```go
// register the pull route: /aaa/xx_zz
peer.RoutePull(new(Aaa))

// or register the pull route: /xx_zz
peer.RoutePullFunc((*Aaa).XxZz)
```

### Pull-Handler-Function API template

```go
func XxZz(ctx tp.PullCtx, args *<T>) (<T>, *tp.Rerror) {
    ...
    return r, nil
}
```

- register it to root router:

```go
// register the pull route: /xx_zz
peer.RoutePullFunc(XxZz)
```

### Push-Controller-Struct API template

```go
type Bbb struct {
    tp.PushCtx
}
func (b *Bbb) YyZz(args *<T>) *tp.Rerror {
    ...
    return nil
}
```

- register it to root router:

```go
// register the push route: /bbb/yy_zz
peer.RoutePush(new(Bbb))

// or register the push route: /yy_zz
peer.RoutePushFunc((*Bbb).YyZz)
```

### Push-Handler-Function API template

```go
// YyZz register the route: /yy_zz
func YyZz(ctx tp.PushCtx, args *<T>) *tp.Rerror {
    ...
    return nil
}
```

- register it to root router:

```go
// register the push route: /yy_zz
peer.RoutePushFunc(YyZz)
```

### Unknown-Pull-Handler-Function API template

```go
func XxxUnknownPull (ctx tp.UnknownPullCtx) (interface{}, *tp.Rerror) {
    ...
    return r, nil
}
```

- register it to root router:

```go
// register the unknown pull route: /*
peer.SetUnknownPull(XxxUnknownPull)
```

### Unknown-Push-Handler-Function API template

```go
func XxxUnknownPush(ctx tp.UnknownPushCtx) *tp.Rerror {
    ...
    return nil
}
```

- register it to root router:

```go
// register the unknown push route: /*
peer.SetUnknownPush(XxxUnknownPush)
```

### Plugin Demo

```go
// NewIgnoreCase Returns a ignoreCase plugin.
func NewIgnoreCase() *ignoreCase {
    return &ignoreCase{}
}

type ignoreCase struct{}

var (
    _ tp.PostReadPullHeaderPlugin = new(ignoreCase)
    _ tp.PostReadPushHeaderPlugin = new(ignoreCase)
)

func (i *ignoreCase) Name() string {
    return "ignoreCase"
}

func (i *ignoreCase) PostReadPullHeader(ctx tp.ReadCtx) *tp.Rerror {
    // Dynamic transformation path is lowercase
    ctx.Url().Path = strings.ToLower(ctx.Url().Path)
    return nil
}

func (i *ignoreCase) PostReadPushHeader(ctx tp.ReadCtx) *tp.Rerror {
    // Dynamic transformation path is lowercase
    ctx.Url().Path = strings.ToLower(ctx.Url().Path)
    return nil
}
```

### Register above handler and plugin

```go
// add router group
group := peer.SubRoute("test")
// register to test group
group.RoutePull(new(Aaa), NewIgnoreCase())
peer.RoutePullFunc(XxZz, NewIgnoreCase())
group.RoutePush(new(Bbb))
peer.RoutePushFunc(YyZz)
peer.SetUnknownPull(XxxUnknownPull)
peer.SetUnknownPush(XxxUnknownPush)
```

### Config

```go
// SrvConfig server config
type SrvConfig struct {
	TlsCertFile       string        `yaml:"tls_cert_file"        ini:"tls_cert_file"        comment:"TLS certificate file path"`
	TlsKeyFile        string        `yaml:"tls_key_file"         ini:"tls_key_file"         comment:"TLS key file path"`
	DefaultSessionAge time.Duration `yaml:"default_session_age"  ini:"default_session_age"  comment:"Default session max age, if less than or equal to 0, no time limit; ns,µs,ms,s,m,h"`
	DefaultContextAge time.Duration `yaml:"default_context_age"  ini:"default_context_age"  comment:"Default PULL or PUSH context max age, if less than or equal to 0, no time limit; ns,µs,ms,s,m,h"`
	SlowCometDuration time.Duration `yaml:"slow_comet_duration"  ini:"slow_comet_duration"  comment:"Slow operation alarm threshold; ns,µs,ms,s ..."`
	DefaultBodyCodec  string        `yaml:"default_body_codec"   ini:"default_body_codec"   comment:"Default body codec type id"`
	PrintBody         bool          `yaml:"print_body"           ini:"print_body"           comment:"Is print body or not"`
	CountTime         bool          `yaml:"count_time"           ini:"count_time"           comment:"Is count cost time or not"`
	Network           string        `yaml:"network"              ini:"network"              comment:"Network; tcp, tcp4, tcp6, unix or unixpacket"`
	ListenAddress     string        `yaml:"listen_address"       ini:"listen_address"       comment:"Listen address; for server role"`
	EnableHeartbeat   bool          `yaml:"enable_heartbeat"     ini:"enable_heartbeat"     comment:"enable heartbeat"`
}

// CliConfig client config
type CliConfig struct {
	TlsCertFile         string        `yaml:"tls_cert_file"          ini:"tls_cert_file"          comment:"TLS certificate file path"`
	TlsKeyFile          string        `yaml:"tls_key_file"           ini:"tls_key_file"           comment:"TLS key file path"`
	DefaultSessionAge   time.Duration `yaml:"default_session_age"    ini:"default_session_age"    comment:"Default session max age, if less than or equal to 0, no time limit; ns,µs,ms,s,m,h"`
	DefaultContextAge   time.Duration `yaml:"default_context_age"    ini:"default_context_age"    comment:"Default PULL or PUSH context max age, if less than or equal to 0, no time limit; ns,µs,ms,s,m,h"`
	DefaultDialTimeout  time.Duration `yaml:"default_dial_timeout"   ini:"default_dial_timeout"   comment:"Default maximum duration for dialing; for client role; ns,µs,ms,s,m,h"`
	RedialTimes         int           `yaml:"redial_times"           ini:"redial_times"           comment:"The maximum times of attempts to redial, after the connection has been unexpectedly broken; for client role"`
	Failover            int           `yaml:"failover"               ini:"failover"               comment:"The maximum times of failover"`
	SlowCometDuration   time.Duration `yaml:"slow_comet_duration"    ini:"slow_comet_duration"    comment:"Slow operation alarm threshold; ns,µs,ms,s ..."`
	DefaultBodyCodec    string        `yaml:"default_body_codec"     ini:"default_body_codec"     comment:"Default body codec type id"`
	PrintBody           bool          `yaml:"print_body"             ini:"print_body"             comment:"Is print body or not"`
	CountTime           bool          `yaml:"count_time"             ini:"count_time"             comment:"Is count cost time or not"`
	Network             string        `yaml:"network"                ini:"network"                comment:"Network; tcp, tcp4, tcp6, unix or unixpacket"`
	HeartbeatSecond     int           `yaml:"heartbeat_second"       ini:"heartbeat_second"       comment:"When the heartbeat interval(second) is greater than 0, heartbeat is enabled; if it's smaller than 3, change to 3 default"`
	SessMaxQuota        int           `yaml:"sess_max_quota"         ini:"sess_max_quota"         comment:"The maximum number of sessions in the connection pool"`
	SessMaxIdleDuration time.Duration `yaml:"sess_max_idle_duration" ini:"sess_max_idle_duration" comment:"The maximum time period for the idle session in the connection pool; ns,µs,ms,s,m,h"`
}
```

### Optimize

- SetPacketSizeLimit sets max packet size.
  If maxSize<=0, set it to max uint32.

    ```go
    func SetPacketSizeLimit(maxPacketSize uint32)
    ```

- SetSocketKeepAlive sets whether the operating system should send
  keepalive messages on the connection.

    ```go
    func SetSocketKeepAlive(keepalive bool)
    ```

- SetSocketKeepAlivePeriod sets period between keep alives.

    ```go
    func SetSocketKeepAlivePeriod(d time.Duration)
    ```

- SetSocketNoDelay controls whether the operating system should delay
  packet transmission in hopes of sending fewer packets (Nagle's
  algorithm).  The default is true (no delay), meaning that data is
  sent as soon as possible after a Write.

    ```go
    func SetSocketNoDelay(_noDelay bool)
    ```

- SetSocketReadBuffer sets the size of the operating system's
  receive buffer associated with the connection.

    ```go
    func SetSocketReadBuffer(bytes int)
    ```

- SetSocketWriteBuffer sets the size of the operating system's
  transmit buffer associated with the connection.

    ```go
    func SetSocketWriteBuffer(bytes int)
    ```

[More Usage](https://github.com/henrylee2cn/teleport)

## License

Ant is under Apache v2 License. See the [LICENSE](https://github.com/henrylee2cn/tp-micro/raw/master/LICENSE) file for the full license text
