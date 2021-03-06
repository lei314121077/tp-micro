// Code generated by 'ant gen' command.
// DO NOT EDIT!

package api

import (
	tp "github.com/henrylee2cn/teleport"
)

// Route registers handlers to router.
func Route(root string, router *tp.Router) {
	// root router group
	group := router.SubRoute(root)

	// custom router
	customRoute(group.ToRouter())

	// automatically generated router

	// PULL APIs...
	group.RoutePullFunc(Home)
	group.RoutePull(new(Math))

	// PUSH APIs...
	group.RoutePushFunc(Stat)
}
