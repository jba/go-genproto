// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto
// DO NOT EDIT!

package google_logging_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes a logs-based metric.  The value of the metric is the
// number of log entries that match a logs filter.
type LogMetric struct {
	// Required. The client-assigned metric identifier. Example:
	// `"severe_errors"`.  Metric identifiers are limited to 1000
	// characters and can include only the following characters: `A-Z`,
	// `a-z`, `0-9`, and the special characters `_-.,+!*',()%/\`.  The
	// forward-slash character (`/`) denotes a hierarchy of name pieces,
	// and it cannot be the first character of the name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this metric, which is used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// An [advanced logs filter](/logging/docs/view/advanced_filters).
	// Example: `"logName:syslog AND severity>=ERROR"`.
	Filter string `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (m *LogMetric) Reset()                    { *m = LogMetric{} }
func (m *LogMetric) String() string            { return proto.CompactTextString(m) }
func (*LogMetric) ProtoMessage()               {}
func (*LogMetric) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// The parameters to ListLogMetrics.
type ListLogMetricsRequest struct {
	// Required. The resource name containing the metrics.
	// Example: `"projects/my-project-id"`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If the `pageToken` parameter is supplied, then the next page of
	// results is retrieved.  The `pageToken` parameter must be set to the value
	// of the `nextPageToken` from the previous response.
	// The value of `parent` must be the same as in the previous request.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// You must check for presence of `nextPageToken` to determine if additional
	// results are available, which you can retrieve by passing the
	// `nextPageToken` value as the `pageToken` parameter in the next request.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListLogMetricsRequest) Reset()                    { *m = ListLogMetricsRequest{} }
func (m *ListLogMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsRequest) ProtoMessage()               {}
func (*ListLogMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// Result returned from ListLogMetrics.
type ListLogMetricsResponse struct {
	// A list of logs-based metrics.
	Metrics []*LogMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// included in the response.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogMetricsResponse) Reset()                    { *m = ListLogMetricsResponse{} }
func (m *ListLogMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsResponse) ProtoMessage()               {}
func (*ListLogMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ListLogMetricsResponse) GetMetrics() []*LogMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// The parameters to GetLogMetric.
type GetLogMetricRequest struct {
	// The resource name of the desired metric.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *GetLogMetricRequest) Reset()                    { *m = GetLogMetricRequest{} }
func (m *GetLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogMetricRequest) ProtoMessage()               {}
func (*GetLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// The parameters to CreateLogMetric.
type CreateLogMetricRequest struct {
	// The resource name of the project in which to create the metric.
	// Example: `"projects/my-project-id"`.
	//
	// The new metric must be provided in the request.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The new logs-based metric, which must not have an identifier that
	// already exists.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *CreateLogMetricRequest) Reset()                    { *m = CreateLogMetricRequest{} }
func (m *CreateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLogMetricRequest) ProtoMessage()               {}
func (*CreateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CreateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to UpdateLogMetric.
//
type UpdateLogMetricRequest struct {
	// The resource name of the metric to update.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	//
	// The updated metric must be provided in the request and have the
	// same identifier that is specified in `metricName`.
	// If the metric does not exist, it is created.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	// The updated metric, whose name must be the same as the
	// metric identifier in `metricName`. If `metricName` does not
	// exist, then a new metric is created.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *UpdateLogMetricRequest) Reset()                    { *m = UpdateLogMetricRequest{} }
func (m *UpdateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLogMetricRequest) ProtoMessage()               {}
func (*UpdateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to DeleteLogMetric.
type DeleteLogMetricRequest struct {
	// The resource name of the metric to delete.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *DeleteLogMetricRequest) Reset()                    { *m = DeleteLogMetricRequest{} }
func (m *DeleteLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogMetricRequest) ProtoMessage()               {}
func (*DeleteLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func init() {
	proto.RegisterType((*LogMetric)(nil), "google.logging.v2.LogMetric")
	proto.RegisterType((*ListLogMetricsRequest)(nil), "google.logging.v2.ListLogMetricsRequest")
	proto.RegisterType((*ListLogMetricsResponse)(nil), "google.logging.v2.ListLogMetricsResponse")
	proto.RegisterType((*GetLogMetricRequest)(nil), "google.logging.v2.GetLogMetricRequest")
	proto.RegisterType((*CreateLogMetricRequest)(nil), "google.logging.v2.CreateLogMetricRequest")
	proto.RegisterType((*UpdateLogMetricRequest)(nil), "google.logging.v2.UpdateLogMetricRequest")
	proto.RegisterType((*DeleteLogMetricRequest)(nil), "google.logging.v2.DeleteLogMetricRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MetricsServiceV2 service

type MetricsServiceV2Client interface {
	// Lists logs-based metrics.
	ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
}

type metricsServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceV2Client(cc *grpc.ClientConn) MetricsServiceV2Client {
	return &metricsServiceV2Client{cc}
}

func (c *metricsServiceV2Client) ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error) {
	out := new(ListLogMetricsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/ListLogMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/GetLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/CreateLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/UpdateLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/DeleteLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricsServiceV2 service

type MetricsServiceV2Server interface {
	// Lists logs-based metrics.
	ListLogMetrics(context.Context, *ListLogMetricsRequest) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(context.Context, *GetLogMetricRequest) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(context.Context, *CreateLogMetricRequest) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(context.Context, *UpdateLogMetricRequest) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(context.Context, *DeleteLogMetricRequest) (*google_protobuf4.Empty, error)
}

func RegisterMetricsServiceV2Server(s *grpc.Server, srv MetricsServiceV2Server) {
	s.RegisterService(&_MetricsServiceV2_serviceDesc, srv)
}

func _MetricsServiceV2_ListLogMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/ListLogMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, req.(*ListLogMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_GetLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/GetLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, req.(*GetLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_CreateLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/CreateLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, req.(*CreateLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_UpdateLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/UpdateLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, req.(*UpdateLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_DeleteLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/DeleteLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, req.(*DeleteLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.MetricsServiceV2",
	HandlerType: (*MetricsServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLogMetrics",
			Handler:    _MetricsServiceV2_ListLogMetrics_Handler,
		},
		{
			MethodName: "GetLogMetric",
			Handler:    _MetricsServiceV2_GetLogMetric_Handler,
		},
		{
			MethodName: "CreateLogMetric",
			Handler:    _MetricsServiceV2_CreateLogMetric_Handler,
		},
		{
			MethodName: "UpdateLogMetric",
			Handler:    _MetricsServiceV2_UpdateLogMetric_Handler,
		},
		{
			MethodName: "DeleteLogMetric",
			Handler:    _MetricsServiceV2_DeleteLogMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x8d, 0x15, 0xfa, 0x0a, 0x14, 0x8c, 0x16, 0x95, 0x00, 0x62, 0x8a, 0x50, 0x55,
	0x06, 0xc4, 0xa2, 0x1d, 0x93, 0x00, 0x71, 0xe1, 0xa7, 0x90, 0x06, 0x9a, 0x3a, 0x40, 0xe2, 0x54,
	0xa5, 0xd9, 0xab, 0x09, 0x6b, 0xe3, 0x10, 0xbb, 0xd5, 0x00, 0x71, 0xe1, 0xc6, 0x0d, 0x69, 0x07,
	0xfe, 0x30, 0xfe, 0x05, 0xae, 0xfc, 0x0f, 0x38, 0x4e, 0xd2, 0x95, 0xd6, 0xd0, 0x6e, 0x97, 0xc8,
	0x7e, 0xb6, 0xdf, 0xf7, 0xe3, 0xf7, 0xbe, 0x31, 0x3c, 0x65, 0x9c, 0xb3, 0x3e, 0x7a, 0x8c, 0xf7,
	0xfd, 0x88, 0x79, 0x3c, 0x61, 0x94, 0x61, 0x14, 0x27, 0x5c, 0x72, 0x9a, 0x2d, 0xf9, 0x71, 0x28,
	0x68, 0x9f, 0x33, 0x16, 0x46, 0x8c, 0x8e, 0x9a, 0xc5, 0xb0, 0x33, 0x40, 0x99, 0x84, 0x81, 0xf0,
	0xf4, 0x5e, 0x72, 0x3e, 0xcf, 0x93, 0xaf, 0x7a, 0xa3, 0xa6, 0xf3, 0x7c, 0xb1, 0xd4, 0xea, 0x43,
	0x05, 0x26, 0xa3, 0x30, 0xc0, 0x80, 0x47, 0xbd, 0x90, 0x51, 0x3f, 0x8a, 0xb8, 0xf4, 0x65, 0xc8,
	0xa3, 0x3c, 0xbb, 0xd3, 0x62, 0xa1, 0x7c, 0x37, 0xec, 0x7a, 0x01, 0x1f, 0xd0, 0x2c, 0x1d, 0xd5,
	0x0b, 0xdd, 0x61, 0x8f, 0xc6, 0xf2, 0x63, 0x8c, 0x82, 0xe2, 0x40, 0x0d, 0xb2, 0x6f, 0x76, 0xc8,
	0x7d, 0x0b, 0xe5, 0x2d, 0xce, 0x5e, 0x68, 0x4c, 0x42, 0xe0, 0x44, 0xe4, 0x0f, 0xb0, 0x66, 0xad,
	0x59, 0x8d, 0x72, 0x5b, 0x8f, 0xc9, 0x1a, 0x54, 0x76, 0x51, 0x04, 0x49, 0x18, 0xa7, 0x5a, 0xb5,
	0x25, 0xbd, 0x34, 0x19, 0x22, 0x36, 0x94, 0x7a, 0x61, 0x5f, 0x62, 0x52, 0x5b, 0xd6, 0x8b, 0xf9,
	0xcc, 0xdd, 0x83, 0xd5, 0xad, 0x50, 0xc8, 0x71, 0x7a, 0xd1, 0xc6, 0x0f, 0x43, 0x14, 0x32, 0x3d,
	0x10, 0xfb, 0x09, 0x46, 0x32, 0x17, 0xca, 0x67, 0xe4, 0x0a, 0x40, 0xec, 0x33, 0xec, 0x48, 0xbe,
	0x87, 0x85, 0x52, 0x39, 0x8d, 0xbc, 0x4a, 0x03, 0xe4, 0x12, 0xe8, 0x49, 0x47, 0x84, 0x9f, 0x50,
	0x4b, 0xad, 0xb4, 0x4f, 0xa5, 0x81, 0x1d, 0x35, 0x77, 0xf7, 0xc1, 0x9e, 0x16, 0x13, 0xb1, 0xaa,
	0x0d, 0x92, 0x4d, 0x38, 0x99, 0x77, 0x41, 0xc9, 0x2d, 0x37, 0x2a, 0xcd, 0xcb, 0xde, 0x4c, 0x1b,
	0xbc, 0xf1, 0xb9, 0x76, 0xb1, 0x99, 0xd4, 0xa1, 0x1a, 0xe1, 0xbe, 0xec, 0xcc, 0x20, 0x9d, 0x49,
	0xc3, 0xdb, 0x05, 0x96, 0xbb, 0x09, 0x17, 0x9e, 0xe1, 0xa1, 0x70, 0x71, 0xc9, 0xab, 0x50, 0xc9,
	0x32, 0x75, 0x26, 0x4a, 0x0a, 0x59, 0xe8, 0xa5, 0x8a, 0xb8, 0x3d, 0xb0, 0x1f, 0x25, 0xe8, 0x4b,
	0x9c, 0x39, 0xfa, 0xaf, 0xfa, 0x6c, 0x40, 0x29, 0x3b, 0xaf, 0x41, 0xe6, 0x5d, 0x24, 0xdf, 0xeb,
	0x72, 0xb0, 0x5f, 0xc7, 0xbb, 0x26, 0x9d, 0x79, 0x88, 0xc7, 0x14, 0xbc, 0x0b, 0xf6, 0x63, 0xec,
	0xe3, 0x31, 0x04, 0x9b, 0xbf, 0x57, 0xe0, 0x5c, 0xde, 0xbf, 0x9d, 0xcc, 0xed, 0x6f, 0x9a, 0xe4,
	0xc0, 0x82, 0xb3, 0x7f, 0xf7, 0x96, 0x34, 0x4c, 0x20, 0x26, 0xaf, 0x39, 0xd7, 0x17, 0xd8, 0x99,
	0x19, 0xc5, 0xbd, 0xf9, 0xf5, 0xe7, 0xaf, 0x83, 0xa5, 0x3a, 0xb9, 0xa6, 0x7e, 0xe0, 0x2e, 0x4a,
	0xff, 0x36, 0xfd, 0x9c, 0x15, 0xfe, 0x81, 0xfa, 0x55, 0xde, 0x63, 0x20, 0x05, 0x5d, 0xff, 0x42,
	0x0b, 0x7b, 0x7c, 0xb3, 0xe0, 0xf4, 0x64, 0xdf, 0x49, 0xdd, 0xa0, 0x64, 0x30, 0x86, 0xf3, 0xdf,
	0x22, 0xba, 0x2d, 0x0d, 0x71, 0x8b, 0xdc, 0x38, 0x84, 0x98, 0x28, 0xd9, 0x04, 0x49, 0x01, 0xa2,
	0x98, 0xc8, 0x77, 0x0b, 0xaa, 0x53, 0x5e, 0x22, 0xa6, 0x8b, 0x9b, 0xfd, 0x36, 0x87, 0x68, 0x43,
	0x13, 0x79, 0xee, 0x42, 0x65, 0xb9, 0x97, 0x9b, 0x80, 0xfc, 0x50, 0x48, 0x53, 0xb6, 0x33, 0x22,
	0x99, 0xad, 0x39, 0x07, 0xe9, 0xbe, 0x46, 0xba, 0xe3, 0x1c, 0xa5, 0x48, 0x63, 0x32, 0xd5, 0xb8,
	0xea, 0x94, 0x3f, 0x8d, 0x64, 0x66, 0x0f, 0x3b, 0x76, 0xb1, 0xb5, 0x78, 0x5b, 0xbd, 0x27, 0xe9,
	0x73, 0x5a, 0x34, 0x6e, 0xfd, 0x28, 0x4c, 0x0f, 0x2f, 0xc2, 0xaa, 0x7a, 0xad, 0x67, 0xc5, 0xb7,
	0xad, 0x6e, 0x49, 0xe7, 0x6f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x05, 0x6e, 0x8f, 0xcf, 0x7c,
	0x06, 0x00, 0x00,
}
