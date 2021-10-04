// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package course

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

// CoursesClient is the client API for Courses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursesClient interface {
	// Sends a greeting
	AddCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CoursesReply, error)
}

type coursesClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursesClient(cc grpc.ClientConnInterface) CoursesClient {
	return &coursesClient{cc}
}

func (c *coursesClient) AddCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CoursesReply, error) {
	out := new(CoursesReply)
	err := c.cc.Invoke(ctx, "/main.Courses/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursesServer is the server API for Courses service.
// All implementations must embed UnimplementedCoursesServer
// for forward compatibility
type CoursesServer interface {
	// Sends a greeting
	AddCourse(context.Context, *CourseRequest) (*CoursesReply, error)
	mustEmbedUnimplementedCoursesServer()
}

// UnimplementedCoursesServer must be embedded to have forward compatible implementations.
type UnimplementedCoursesServer struct {
}

func (UnimplementedCoursesServer) AddCourse(context.Context, *CourseRequest) (*CoursesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedCoursesServer) mustEmbedUnimplementedCoursesServer() {}

// UnsafeCoursesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursesServer will
// result in compilation errors.
type UnsafeCoursesServer interface {
	mustEmbedUnimplementedCoursesServer()
}

func RegisterCoursesServer(s grpc.ServiceRegistrar, srv CoursesServer) {
	s.RegisterService(&Courses_ServiceDesc, srv)
}

func _Courses_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Courses/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).AddCourse(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Courses_ServiceDesc is the grpc.ServiceDesc for Courses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Courses",
	HandlerType: (*CoursesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCourse",
			Handler:    _Courses_AddCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courses.proto",
}
