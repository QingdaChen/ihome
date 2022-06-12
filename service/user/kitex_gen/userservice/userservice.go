// Code generated by Kitex v0.3.2. DO NOT EDIT.

package userservice

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
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*kitex_gen.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendSMS":        kitex.NewMethodInfo(sendSMSHandler, newSendSMSArgs, newSendSMSResult, false),
		"Register":       kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"Login":          kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"SessionAuth":    kitex.NewMethodInfo(sessionAuthHandler, newSessionAuthArgs, newSessionAuthResult, false),
		"GetSessionInfo": kitex.NewMethodInfo(getSessionInfoHandler, newGetSessionInfoArgs, newGetSessionInfoResult, false),
		"DeleteSession":  kitex.NewMethodInfo(deleteSessionHandler, newDeleteSessionArgs, newDeleteSessionResult, false),
		"GetUserInfo":    kitex.NewMethodInfo(getUserInfoHandler, newGetUserInfoArgs, newGetUserInfoResult, false),
		"UpdateUserInfo": kitex.NewMethodInfo(updateUserInfoHandler, newUpdateUserInfoArgs, newUpdateUserInfoResult, false),
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
		req := new(kitex_gen.SMSRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).SendSMS(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendSMSArgs:
		success, err := handler.(kitex_gen.UserService).SendSMS(ctx, s.Req)
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
	Req *kitex_gen.SMSRequest
}

func (p *SendSMSArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendSMSArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SendSMSArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SMSRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendSMSArgs_Req_DEFAULT *kitex_gen.SMSRequest

func (p *SendSMSArgs) GetReq() *kitex_gen.SMSRequest {
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.RegRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(kitex_gen.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *kitex_gen.RegRequest
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.RegRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *kitex_gen.RegRequest

func (p *RegisterArgs) GetReq() *kitex_gen.RegRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type RegisterResult struct {
	Success *kitex_gen.Response
}

var RegisterResult_Success_DEFAULT *kitex_gen.Response

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.LoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(kitex_gen.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *kitex_gen.LoginRequest
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.LoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *kitex_gen.LoginRequest

func (p *LoginArgs) GetReq() *kitex_gen.LoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type LoginResult struct {
	Success *kitex_gen.Response
}

var LoginResult_Success_DEFAULT *kitex_gen.Response

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func sessionAuthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.SessionAuthRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).SessionAuth(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SessionAuthArgs:
		success, err := handler.(kitex_gen.UserService).SessionAuth(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SessionAuthResult)
		realResult.Success = success
	}
	return nil
}
func newSessionAuthArgs() interface{} {
	return &SessionAuthArgs{}
}

func newSessionAuthResult() interface{} {
	return &SessionAuthResult{}
}

type SessionAuthArgs struct {
	Req *kitex_gen.SessionAuthRequest
}

func (p *SessionAuthArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SessionAuthArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SessionAuthArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SessionAuthRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SessionAuthArgs_Req_DEFAULT *kitex_gen.SessionAuthRequest

func (p *SessionAuthArgs) GetReq() *kitex_gen.SessionAuthRequest {
	if !p.IsSetReq() {
		return SessionAuthArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SessionAuthArgs) IsSetReq() bool {
	return p.Req != nil
}

type SessionAuthResult struct {
	Success *kitex_gen.Response
}

var SessionAuthResult_Success_DEFAULT *kitex_gen.Response

func (p *SessionAuthResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SessionAuthResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SessionAuthResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SessionAuthResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return SessionAuthResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SessionAuthResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *SessionAuthResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getSessionInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.SessionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).GetSessionInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetSessionInfoArgs:
		success, err := handler.(kitex_gen.UserService).GetSessionInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetSessionInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetSessionInfoArgs() interface{} {
	return &GetSessionInfoArgs{}
}

func newGetSessionInfoResult() interface{} {
	return &GetSessionInfoResult{}
}

type GetSessionInfoArgs struct {
	Req *kitex_gen.SessionRequest
}

func (p *GetSessionInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetSessionInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetSessionInfoArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SessionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetSessionInfoArgs_Req_DEFAULT *kitex_gen.SessionRequest

func (p *GetSessionInfoArgs) GetReq() *kitex_gen.SessionRequest {
	if !p.IsSetReq() {
		return GetSessionInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetSessionInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetSessionInfoResult struct {
	Success *kitex_gen.Response
}

var GetSessionInfoResult_Success_DEFAULT *kitex_gen.Response

func (p *GetSessionInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetSessionInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetSessionInfoResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSessionInfoResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return GetSessionInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetSessionInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *GetSessionInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteSessionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.SessionDeleteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).DeleteSession(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteSessionArgs:
		success, err := handler.(kitex_gen.UserService).DeleteSession(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteSessionResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteSessionArgs() interface{} {
	return &DeleteSessionArgs{}
}

func newDeleteSessionResult() interface{} {
	return &DeleteSessionResult{}
}

type DeleteSessionArgs struct {
	Req *kitex_gen.SessionDeleteRequest
}

func (p *DeleteSessionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteSessionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteSessionArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.SessionDeleteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteSessionArgs_Req_DEFAULT *kitex_gen.SessionDeleteRequest

func (p *DeleteSessionArgs) GetReq() *kitex_gen.SessionDeleteRequest {
	if !p.IsSetReq() {
		return DeleteSessionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteSessionArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteSessionResult struct {
	Success *kitex_gen.Response
}

var DeleteSessionResult_Success_DEFAULT *kitex_gen.Response

func (p *DeleteSessionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteSessionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteSessionResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteSessionResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return DeleteSessionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteSessionResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *DeleteSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.GetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserInfoArgs:
		success, err := handler.(kitex_gen.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *kitex_gen.GetUserRequest
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.GetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *kitex_gen.GetUserRequest

func (p *GetUserInfoArgs) GetReq() *kitex_gen.GetUserRequest {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserInfoResult struct {
	Success *kitex_gen.Response
}

var GetUserInfoResult_Success_DEFAULT *kitex_gen.Response

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_gen.UpdateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_gen.UserService).UpdateUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateUserInfoArgs:
		success, err := handler.(kitex_gen.UserService).UpdateUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateUserInfoArgs() interface{} {
	return &UpdateUserInfoArgs{}
}

func newUpdateUserInfoResult() interface{} {
	return &UpdateUserInfoResult{}
}

type UpdateUserInfoArgs struct {
	Req *kitex_gen.UpdateUserRequest
}

func (p *UpdateUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(kitex_gen.UpdateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserInfoArgs_Req_DEFAULT *kitex_gen.UpdateUserRequest

func (p *UpdateUserInfoArgs) GetReq() *kitex_gen.UpdateUserRequest {
	if !p.IsSetReq() {
		return UpdateUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUserInfoResult struct {
	Success *kitex_gen.Response
}

var UpdateUserInfoResult_Success_DEFAULT *kitex_gen.Response

func (p *UpdateUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserInfoResult) Unmarshal(in []byte) error {
	msg := new(kitex_gen.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserInfoResult) GetSuccess() *kitex_gen.Response {
	if !p.IsSetSuccess() {
		return UpdateUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_gen.Response)
}

func (p *UpdateUserInfoResult) IsSetSuccess() bool {
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

func (p *kClient) SendSMS(ctx context.Context, Req *kitex_gen.SMSRequest) (r *kitex_gen.Response, err error) {
	var _args SendSMSArgs
	_args.Req = Req
	var _result SendSMSResult
	if err = p.c.Call(ctx, "SendSMS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, Req *kitex_gen.RegRequest) (r *kitex_gen.Response, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *kitex_gen.LoginRequest) (r *kitex_gen.Response, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SessionAuth(ctx context.Context, Req *kitex_gen.SessionAuthRequest) (r *kitex_gen.Response, err error) {
	var _args SessionAuthArgs
	_args.Req = Req
	var _result SessionAuthResult
	if err = p.c.Call(ctx, "SessionAuth", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSessionInfo(ctx context.Context, Req *kitex_gen.SessionRequest) (r *kitex_gen.Response, err error) {
	var _args GetSessionInfoArgs
	_args.Req = Req
	var _result GetSessionInfoResult
	if err = p.c.Call(ctx, "GetSessionInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteSession(ctx context.Context, Req *kitex_gen.SessionDeleteRequest) (r *kitex_gen.Response, err error) {
	var _args DeleteSessionArgs
	_args.Req = Req
	var _result DeleteSessionResult
	if err = p.c.Call(ctx, "DeleteSession", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *kitex_gen.GetUserRequest) (r *kitex_gen.Response, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, Req *kitex_gen.UpdateUserRequest) (r *kitex_gen.Response, err error) {
	var _args UpdateUserInfoArgs
	_args.Req = Req
	var _result UpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
