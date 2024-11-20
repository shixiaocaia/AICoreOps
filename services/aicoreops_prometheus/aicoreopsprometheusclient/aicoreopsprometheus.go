// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_prometheus.proto

package aicoreopsprometheusclient

import (
	"context"

	"aicoreops_prometheus/aicoreops_prometheus"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = aicoreops_prometheus.Request
	Response = aicoreops_prometheus.Response

	AicoreopsPrometheus interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAicoreopsPrometheus struct {
		cli zrpc.Client
	}
)

func NewAicoreopsPrometheus(cli zrpc.Client) AicoreopsPrometheus {
	return &defaultAicoreopsPrometheus{
		cli: cli,
	}
}

func (m *defaultAicoreopsPrometheus) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := aicoreops_prometheus.NewAicoreopsPrometheusClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
