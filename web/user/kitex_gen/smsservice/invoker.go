// Code generated by Kitex v0.3.2. DO NOT EDIT.

package smsservice

import (
	"github.com/cloudwego/kitex/server"
	"ihome/service/user/kitex_gen"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler kitex_gen.SMSService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
