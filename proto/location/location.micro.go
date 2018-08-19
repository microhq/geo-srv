// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/microhq/geo-srv/proto/location/location.proto

/*
Package location is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/geo-srv/proto/location/location.proto

It has these top-level messages:
	ReadRequest
	ReadResponse
	SaveRequest
	SaveResponse
	SearchRequest
	SearchResponse
*/
package location

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/microhq/geo-srv/proto"

import (
	context "context"
	api "github.com/micro/go-api"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Location service

type LocationService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type locationService struct {
	c    client.Client
	name string
}

func NewLocationService(name string, c client.Client) LocationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "location"
	}
	return &locationService{
		c:    c,
		name: name,
	}
}

func (c *locationService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Save", in)
	out := new(SaveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Location service

type LocationHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Save(context.Context, *SaveRequest, *SaveResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterLocationHandler(s server.Server, hdlr LocationHandler, opts ...server.HandlerOption) error {
	type location interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type Location struct {
		location
	}
	h := &locationHandler{hdlr}
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
}

type locationHandler struct {
	LocationHandler
}

func (h *locationHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.LocationHandler.Read(ctx, in, out)
}

func (h *locationHandler) Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error {
	return h.LocationHandler.Save(ctx, in, out)
}

func (h *locationHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.LocationHandler.Search(ctx, in, out)
}
