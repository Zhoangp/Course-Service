// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/course.proto

package pb

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

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	GetCourses(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error)
	GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error)
	GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error)
	Enrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error)
	GetCourseContent(ctx context.Context, in *GetCourseContentRequest, opts ...grpc.CallOption) (*GetCourseContentResponse, error)
	GetEnrollments(ctx context.Context, in *GetEnrollmentsRequest, opts ...grpc.CallOption) (*GetEnrollmentsResponse, error)
	GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error)
	PublishCourse(ctx context.Context, in *PublishCourseRequest, opts ...grpc.CallOption) (*PublishCourseResponse, error)
	GetCourseWithInstructor(ctx context.Context, in *GetCourseWithInstructorRequest, opts ...grpc.CallOption) (*GetCourseWithInstructorResponse, error)
	DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) GetCourses(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error) {
	out := new(GetCoursesResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error) {
	out := new(GetCourseResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error) {
	out := new(CreateCourseResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) Enrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error) {
	out := new(EnrollmentResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/Enrollment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourseContent(ctx context.Context, in *GetCourseContentRequest, opts ...grpc.CallOption) (*GetCourseContentResponse, error) {
	out := new(GetCourseContentResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetCourseContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetEnrollments(ctx context.Context, in *GetEnrollmentsRequest, opts ...grpc.CallOption) (*GetEnrollmentsResponse, error) {
	out := new(GetEnrollmentsResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetEnrollments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error) {
	out := new(GetPricesResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error) {
	out := new(UpdateCourseResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/UpdateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) PublishCourse(ctx context.Context, in *PublishCourseRequest, opts ...grpc.CallOption) (*PublishCourseResponse, error) {
	out := new(PublishCourseResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/PublishCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourseWithInstructor(ctx context.Context, in *GetCourseWithInstructorRequest, opts ...grpc.CallOption) (*GetCourseWithInstructorResponse, error) {
	out := new(GetCourseWithInstructorResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/GetCourseWithInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	out := new(DeleteCourseResponse)
	err := c.cc.Invoke(ctx, "/course.CourseService/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	GetCourses(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error)
	GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error)
	GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error)
	CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseResponse, error)
	Enrollment(context.Context, *EnrollmentRequest) (*EnrollmentResponse, error)
	GetCourseContent(context.Context, *GetCourseContentRequest) (*GetCourseContentResponse, error)
	GetEnrollments(context.Context, *GetEnrollmentsRequest) (*GetEnrollmentsResponse, error)
	GetPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseResponse, error)
	PublishCourse(context.Context, *PublishCourseRequest) (*PublishCourseResponse, error)
	GetCourseWithInstructor(context.Context, *GetCourseWithInstructorRequest) (*GetCourseWithInstructorResponse, error)
	DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) GetCourses(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseServiceServer) GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServiceServer) GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) Enrollment(context.Context, *EnrollmentRequest) (*EnrollmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enrollment not implemented")
}
func (UnimplementedCourseServiceServer) GetCourseContent(context.Context, *GetCourseContentRequest) (*GetCourseContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseContent not implemented")
}
func (UnimplementedCourseServiceServer) GetEnrollments(context.Context, *GetEnrollmentsRequest) (*GetEnrollmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnrollments not implemented")
}
func (UnimplementedCourseServiceServer) GetPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrices not implemented")
}
func (UnimplementedCourseServiceServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedCourseServiceServer) PublishCourse(context.Context, *PublishCourseRequest) (*PublishCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishCourse not implemented")
}
func (UnimplementedCourseServiceServer) GetCourseWithInstructor(context.Context, *GetCourseWithInstructorRequest) (*GetCourseWithInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseWithInstructor not implemented")
}
func (UnimplementedCourseServiceServer) DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourses(ctx, req.(*GetCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourse(ctx, req.(*GetCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetAllCategories(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_Enrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).Enrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/Enrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).Enrollment(ctx, req.(*EnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourseContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetCourseContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourseContent(ctx, req.(*GetCourseContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetEnrollments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnrollmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetEnrollments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetEnrollments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetEnrollments(ctx, req.(*GetEnrollmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetPrices(ctx, req.(*GetPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/UpdateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_PublishCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).PublishCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/PublishCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).PublishCourse(ctx, req.(*PublishCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourseWithInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseWithInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourseWithInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/GetCourseWithInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourseWithInstructor(ctx, req.(*GetCourseWithInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course.CourseService/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).DeleteCourse(ctx, req.(*DeleteCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourses",
			Handler:    _CourseService_GetCourses_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _CourseService_GetCourse_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _CourseService_GetAllCategories_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
		{
			MethodName: "Enrollment",
			Handler:    _CourseService_Enrollment_Handler,
		},
		{
			MethodName: "GetCourseContent",
			Handler:    _CourseService_GetCourseContent_Handler,
		},
		{
			MethodName: "GetEnrollments",
			Handler:    _CourseService_GetEnrollments_Handler,
		},
		{
			MethodName: "GetPrices",
			Handler:    _CourseService_GetPrices_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _CourseService_UpdateCourse_Handler,
		},
		{
			MethodName: "PublishCourse",
			Handler:    _CourseService_PublishCourse_Handler,
		},
		{
			MethodName: "GetCourseWithInstructor",
			Handler:    _CourseService_GetCourseWithInstructor_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _CourseService_DeleteCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/course.proto",
}
