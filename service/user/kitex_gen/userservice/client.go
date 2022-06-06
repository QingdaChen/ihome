// Code generated by Kitex v0.3.2. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"ihome/service/user/kitex_gen"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendSMS(ctx context.Context, Req *kitex_gen.SMSRequest, callOptions ...callopt.Option) (r *kitex_gen.Response, err error)
	Register(ctx context.Context, Req *kitex_gen.RegRequest, callOptions ...callopt.Option) (r *kitex_gen.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) SendSMS(ctx context.Context, Req *kitex_gen.SMSRequest, callOptions ...callopt.Option) (r *kitex_gen.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendSMS(ctx, Req)
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *kitex_gen.RegRequest, callOptions ...callopt.Option) (r *kitex_gen.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}
