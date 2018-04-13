// Code generated by 'ant gen' command.
// DO NOT EDIT!

package types

import (
	"github.com/henrylee2cn/tp-micro/examples/sample/logic/model"
)

// HomeReply home reply
type HomeReply struct {
	Content string `json:"content"` // text
}

type (
	// DivideArgs divide api args
	DivideArgs = model.DivideArgs
	// DivideReply divide api result
	DivideReply struct {
		// quotient
		C float64 `json:"c"`
	}
)

// StatArgs stat handler args
type StatArgs struct {
	Ts int64 `json:"ts"` // timestamps
}

// User user info
type User = model.User
