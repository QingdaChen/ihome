package main

import (
	"context"
	"ihome/web/user/kitex_gen"
)

// SMSServiceImpl implements the last service interface defined in the IDL.
type SMSServiceImpl struct{}

// SendSMS implements the SMSServiceImpl interface.
func (s *SMSServiceImpl) SendSMS(ctx context.Context, req *kitex_gen.Request) (resp *kitex_gen.Response, err error) {
	// TODO: Your code here...
	return
}

// SaveSMS implements the SMSServiceImpl interface.
func (s *SMSServiceImpl) SaveSMS(ctx context.Context, req *kitex_gen.Request) (resp *kitex_gen.Response, err error) {
	// TODO: Your code here...
	return
}
