// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: booking.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookingService_CreateBooking_FullMethodName        = "/booking.BookingService/CreateBooking"
	BookingService_GetBookingByUser_FullMethodName     = "/booking.BookingService/GetBookingByUser"
	BookingService_GetBookingsBySection_FullMethodName = "/booking.BookingService/GetBookingsBySection"
	BookingService_RemoveBookingByUser_FullMethodName  = "/booking.BookingService/RemoveBookingByUser"
	BookingService_ModifySeatByUser_FullMethodName     = "/booking.BookingService/ModifySeatByUser"
)

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	// CreateBooking books a ticket for the user
	CreateBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	// GetBookingByUser returns a booking for the user
	GetBookingByUser(ctx context.Context, in *GetBookingByUserRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	// GetBookingsBySection returns list of bookings for the section
	GetBookingsBySection(ctx context.Context, in *GetBookingsBySectionRequest, opts ...grpc.CallOption) (*BookingListResponse, error)
	// RemoveBookingByUser removes a booking for the user
	RemoveBookingByUser(ctx context.Context, in *RemoveBookingByUserRequest, opts ...grpc.CallOption) (*RemoveBookingResponse, error)
	// ModifySeatByUser modifies the seat for the user
	ModifySeatByUser(ctx context.Context, in *SeatModificationRequest, opts ...grpc.CallOption) (*SeatModificationResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CreateBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingByUser(ctx context.Context, in *GetBookingByUserRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, BookingService_GetBookingByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingsBySection(ctx context.Context, in *GetBookingsBySectionRequest, opts ...grpc.CallOption) (*BookingListResponse, error) {
	out := new(BookingListResponse)
	err := c.cc.Invoke(ctx, BookingService_GetBookingsBySection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) RemoveBookingByUser(ctx context.Context, in *RemoveBookingByUserRequest, opts ...grpc.CallOption) (*RemoveBookingResponse, error) {
	out := new(RemoveBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_RemoveBookingByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ModifySeatByUser(ctx context.Context, in *SeatModificationRequest, opts ...grpc.CallOption) (*SeatModificationResponse, error) {
	out := new(SeatModificationResponse)
	err := c.cc.Invoke(ctx, BookingService_ModifySeatByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations should embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	// CreateBooking books a ticket for the user
	CreateBooking(context.Context, *BookingRequest) (*BookingResponse, error)
	// GetBookingByUser returns a booking for the user
	GetBookingByUser(context.Context, *GetBookingByUserRequest) (*BookingResponse, error)
	// GetBookingsBySection returns list of bookings for the section
	GetBookingsBySection(context.Context, *GetBookingsBySectionRequest) (*BookingListResponse, error)
	// RemoveBookingByUser removes a booking for the user
	RemoveBookingByUser(context.Context, *RemoveBookingByUserRequest) (*RemoveBookingResponse, error)
	// ModifySeatByUser modifies the seat for the user
	ModifySeatByUser(context.Context, *SeatModificationRequest) (*SeatModificationResponse, error)
}

// UnimplementedBookingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *BookingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingByUser(context.Context, *GetBookingByUserRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingByUser not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingsBySection(context.Context, *GetBookingsBySectionRequest) (*BookingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingsBySection not implemented")
}
func (UnimplementedBookingServiceServer) RemoveBookingByUser(context.Context, *RemoveBookingByUserRequest) (*RemoveBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBookingByUser not implemented")
}
func (UnimplementedBookingServiceServer) ModifySeatByUser(context.Context, *SeatModificationRequest) (*SeatModificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeatByUser not implemented")
}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetBookingByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingByUser(ctx, req.(*GetBookingByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingsBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingsBySectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingsBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetBookingsBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingsBySection(ctx, req.(*GetBookingsBySectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_RemoveBookingByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBookingByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).RemoveBookingByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_RemoveBookingByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).RemoveBookingByUser(ctx, req.(*RemoveBookingByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ModifySeatByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatModificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ModifySeatByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_ModifySeatByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ModifySeatByUser(ctx, req.(*SeatModificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetBookingByUser",
			Handler:    _BookingService_GetBookingByUser_Handler,
		},
		{
			MethodName: "GetBookingsBySection",
			Handler:    _BookingService_GetBookingsBySection_Handler,
		},
		{
			MethodName: "RemoveBookingByUser",
			Handler:    _BookingService_RemoveBookingByUser_Handler,
		},
		{
			MethodName: "ModifySeatByUser",
			Handler:    _BookingService_ModifySeatByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
