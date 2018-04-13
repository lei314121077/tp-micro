// Code generated by 'ant gen' command.
// The temporary code used to ensure successful compilation!
// When the project is completed, it should be removed!

package logic

import (
	tp "github.com/henrylee2cn/teleport"

	"github.com/henrylee2cn/tp-micro/examples/sample/types"
	// "github.com/henrylee2cn/tp-micro/examples/sample/rerrs"
)

// Home comment...
func Home(ctx tp.PullCtx, args *struct{}) (*types.HomeReply, *tp.Rerror) {
	return new(types.HomeReply), nil
}

// Math_Divide handler
func Math_Divide(ctx tp.PullCtx, args *types.DivideArgs) (*types.DivideReply, *tp.Rerror) {
	return new(types.DivideReply), nil
}

// Stat comment...
func Stat(ctx tp.PushCtx, args *types.StatArgs) *tp.Rerror {
	return nil
}