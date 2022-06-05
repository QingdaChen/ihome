// Code generated by Kitex v0.3.2. DO NOT EDIT.

package smsservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"ihome/service/user/kitex_gen"
)

func serviceInfo() *kitex.ServiceInfo {
	return sMSServiceServiceInfo
}

var sMSServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SMSService"
	handlerType := (*kitex_gen.SMSService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendSMS": kitex.NewMethodInfo(sendSMSHandler, newSendSMSArgs, newSendSMSResult, false),
		"SaveSMS": kitex.NewMethodInfo(saveSMSHandler, newSaveSMSArgs, newSaveSMSResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func sendSMSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.SMSService).SendSMS(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendSMSArgs:
		success, err := handler.(kitex_gen.SMSService).SendSMS(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendSMSResult)
		realResult.Success = success
	}
	return nil
}
func newSendSMSArgs() interface{} {
	return &SendSMSArgs{}
}

func newSendSMSResult() interface{} {
	return &SendSMSResult{}
}

type SendSMSArgs struct {
	Req *kitex_gen.Request
}

func (p *SendSMSArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendSMSArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SendSMSArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendSMSArgs_Req_DEFAULT *kitex_gen.Request

func (p *SendSMSArgs) GetReq() *kitex_gen.Request {
	if !p.IsSetReq() {
		return SendSMSArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendSMSArgs) IsSetReq() bool {
	return p.Req != nil
}

type SendSMSResult struct {
	Success *kitex_gen.Response
}

var SendSMSResult_Success_DEFAULT *kitex_gen.Response

func (p *SendSMSResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendSMSResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SendSMSResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendSMSResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return SendSMSResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendSMSResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *SendSMSResult) IsSetSuccess() bool {
	return p.Success != nil
}

func saveSMSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.SMSService).SaveSMS(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SaveSMSArgs:
		success, err := handler.(kitex_gen.SMSService).SaveSMS(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SaveSMSResult)
		realResult.Success = success
	}
	return nil
}
func newSaveSMSArgs() interface{} {
	return &SaveSMSArgs{}
}

func newSaveSMSResult() interface{} {
	return &SaveSMSResult{}
}

type SaveSMSArgs struct {
	Req *kitex_gen.Request
}

func (p *SaveSMSArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SaveSMSArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SaveSMSArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SaveSMSArgs_Req_DEFAULT *kitex_gen.Request

func (p *SaveSMSArgs) GetReq() *kitex_gen.Request {
	if !p.IsSetReq() {
		return SaveSMSArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SaveSMSArgs) IsSetReq() bool {
	return p.Req != nil
}

type SaveSMSResult struct {
	Success *kitex_gen.Response
}

var SaveSMSResult_Success_DEFAULT *kitex_gen.Response

func (p *SaveSMSResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SaveSMSResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SaveSMSResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveSMSResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return SaveSMSResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SaveSMSResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *SaveSMSResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendSMS(ctx context.Context, Req *kitex_gen.Request) (r *kitex_gen.Response, err error) {
	var _args SendSMSArgs
	_args.Req = Req
	var _result SendSMSResult
	if err = p.c.Call(ctx, "SendSMS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SaveSMS(ctx context.Context, Req *kitex_gen.Request) (r *kitex_gen.Response, err error) {
	var _args SaveSMSArgs
	_args.Req = Req
	var _result SaveSMSResult
	if err = p.c.Call(ctx, "SaveSMS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
