// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/remind/remind.proto

package remind_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for RemindServer service

type RemindServerService interface {
	AddUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*EmptyResponse, error)
	AddBatchUnread(ctx context.Context, in *RemindBatchRequest, opts ...client.CallOption) (*EmptyResponse, error)
	DeleteUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*EmptyResponse, error)
	CheckUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*CheckResponse, error)
}

type remindServerService struct {
	c    client.Client
	name string
}

func NewRemindServerService(name string, c client.Client) RemindServerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "remind"
	}
	return &remindServerService{
		c:    c,
		name: name,
	}
}

func (c *remindServerService) AddUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "RemindServer.AddUnread", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindServerService) AddBatchUnread(ctx context.Context, in *RemindBatchRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "RemindServer.AddBatchUnread", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindServerService) DeleteUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "RemindServer.DeleteUnread", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remindServerService) CheckUnread(ctx context.Context, in *RemindInfo, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.name, "RemindServer.CheckUnread", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RemindServer service

type RemindServerHandler interface {
	AddUnread(context.Context, *RemindInfo, *EmptyResponse) error
	AddBatchUnread(context.Context, *RemindBatchRequest, *EmptyResponse) error
	DeleteUnread(context.Context, *RemindInfo, *EmptyResponse) error
	CheckUnread(context.Context, *RemindInfo, *CheckResponse) error
}

func RegisterRemindServerHandler(s server.Server, hdlr RemindServerHandler, opts ...server.HandlerOption) error {
	type remindServer interface {
		AddUnread(ctx context.Context, in *RemindInfo, out *EmptyResponse) error
		AddBatchUnread(ctx context.Context, in *RemindBatchRequest, out *EmptyResponse) error
		DeleteUnread(ctx context.Context, in *RemindInfo, out *EmptyResponse) error
		CheckUnread(ctx context.Context, in *RemindInfo, out *CheckResponse) error
	}
	type RemindServer struct {
		remindServer
	}
	h := &remindServerHandler{hdlr}
	return s.Handle(s.NewHandler(&RemindServer{h}, opts...))
}

type remindServerHandler struct {
	RemindServerHandler
}

func (h *remindServerHandler) AddUnread(ctx context.Context, in *RemindInfo, out *EmptyResponse) error {
	return h.RemindServerHandler.AddUnread(ctx, in, out)
}

func (h *remindServerHandler) AddBatchUnread(ctx context.Context, in *RemindBatchRequest, out *EmptyResponse) error {
	return h.RemindServerHandler.AddBatchUnread(ctx, in, out)
}

func (h *remindServerHandler) DeleteUnread(ctx context.Context, in *RemindInfo, out *EmptyResponse) error {
	return h.RemindServerHandler.DeleteUnread(ctx, in, out)
}

func (h *remindServerHandler) CheckUnread(ctx context.Context, in *RemindInfo, out *CheckResponse) error {
	return h.RemindServerHandler.CheckUnread(ctx, in, out)
}
