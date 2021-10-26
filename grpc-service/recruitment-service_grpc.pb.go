// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_service

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

// RecruitmentServiceClient is the client API for RecruitmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecruitmentServiceClient interface {
	// vacancies
	PostVacancy(ctx context.Context, in *PostVacancyRequest, opts ...grpc.CallOption) (*PostVacancyResponse, error)
	ShowVacancies(ctx context.Context, in *ShowVacanciesRequest, opts ...grpc.CallOption) (*ShowVacanciesResponse, error)
	SearchVacancies(ctx context.Context, in *SearchVacanciesRequest, opts ...grpc.CallOption) (*SearchVacanciesResponse, error)
	ChangeVacancyPosition(ctx context.Context, in *ChangeVacancyPositionRequest, opts ...grpc.CallOption) (*Empty, error)
	ApproveVacancy(ctx context.Context, in *ApproveVacancyRequest, opts ...grpc.CallOption) (*Empty, error)
	CloseVacancy(ctx context.Context, in *CloseVacancyRequest, opts ...grpc.CallOption) (*Empty, error)
	RejectVacancy(ctx context.Context, in *RejectVacancyRequest, opts ...grpc.CallOption) (*Empty, error)
	TakeInWorkVacancy(ctx context.Context, in *TakeInWorkVacancyRequest, opts ...grpc.CallOption) (*Empty, error)
	GetVacancy(ctx context.Context, in *GetVacancyRequest, opts ...grpc.CallOption) (*GetVacancyResponse, error)
	// recruitment
	ConsiderCandidate(ctx context.Context, in *ConsiderCandidateRequest, opts ...grpc.CallOption) (*ConsiderCandidateResponse, error)
	ConsiderCandidateAnotherVacancy(ctx context.Context, in *ConsiderCandidateAnotherVacancyRequest, opts ...grpc.CallOption) (*ConsiderCandidateAnotherVacancyResponse, error)
	AcceptRecruitmentStage(ctx context.Context, in *AcceptRecruitmentStageRequest, opts ...grpc.CallOption) (*Empty, error)
	DenyRecruitment(ctx context.Context, in *DenyRecruitmentRequest, opts ...grpc.CallOption) (*Empty, error)
	GetRecruitment(ctx context.Context, in *GetRecruitmentRequest, opts ...grpc.CallOption) (*GetRecruitmentResponse, error)
	ShowRecruitments(ctx context.Context, in *ShowRecruitmentRequest, opts ...grpc.CallOption) (*ShowRecruitmentsResponse, error)
}

type recruitmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecruitmentServiceClient(cc grpc.ClientConnInterface) RecruitmentServiceClient {
	return &recruitmentServiceClient{cc}
}

func (c *recruitmentServiceClient) PostVacancy(ctx context.Context, in *PostVacancyRequest, opts ...grpc.CallOption) (*PostVacancyResponse, error) {
	out := new(PostVacancyResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/PostVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ShowVacancies(ctx context.Context, in *ShowVacanciesRequest, opts ...grpc.CallOption) (*ShowVacanciesResponse, error) {
	out := new(ShowVacanciesResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ShowVacancies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) SearchVacancies(ctx context.Context, in *SearchVacanciesRequest, opts ...grpc.CallOption) (*SearchVacanciesResponse, error) {
	out := new(SearchVacanciesResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/SearchVacancies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ChangeVacancyPosition(ctx context.Context, in *ChangeVacancyPositionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ChangeVacancyPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ApproveVacancy(ctx context.Context, in *ApproveVacancyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ApproveVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) CloseVacancy(ctx context.Context, in *CloseVacancyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/CloseVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) RejectVacancy(ctx context.Context, in *RejectVacancyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/RejectVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) TakeInWorkVacancy(ctx context.Context, in *TakeInWorkVacancyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/TakeInWorkVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) GetVacancy(ctx context.Context, in *GetVacancyRequest, opts ...grpc.CallOption) (*GetVacancyResponse, error) {
	out := new(GetVacancyResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/GetVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ConsiderCandidate(ctx context.Context, in *ConsiderCandidateRequest, opts ...grpc.CallOption) (*ConsiderCandidateResponse, error) {
	out := new(ConsiderCandidateResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ConsiderCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ConsiderCandidateAnotherVacancy(ctx context.Context, in *ConsiderCandidateAnotherVacancyRequest, opts ...grpc.CallOption) (*ConsiderCandidateAnotherVacancyResponse, error) {
	out := new(ConsiderCandidateAnotherVacancyResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ConsiderCandidateAnotherVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) AcceptRecruitmentStage(ctx context.Context, in *AcceptRecruitmentStageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/AcceptRecruitmentStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) DenyRecruitment(ctx context.Context, in *DenyRecruitmentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/RecruitmentService/DenyRecruitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) GetRecruitment(ctx context.Context, in *GetRecruitmentRequest, opts ...grpc.CallOption) (*GetRecruitmentResponse, error) {
	out := new(GetRecruitmentResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/GetRecruitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitmentServiceClient) ShowRecruitments(ctx context.Context, in *ShowRecruitmentRequest, opts ...grpc.CallOption) (*ShowRecruitmentsResponse, error) {
	out := new(ShowRecruitmentsResponse)
	err := c.cc.Invoke(ctx, "/RecruitmentService/ShowRecruitments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecruitmentServiceServer is the server API for RecruitmentService service.
// All implementations must embed UnimplementedRecruitmentServiceServer
// for forward compatibility
type RecruitmentServiceServer interface {
	// vacancies
	PostVacancy(context.Context, *PostVacancyRequest) (*PostVacancyResponse, error)
	ShowVacancies(context.Context, *ShowVacanciesRequest) (*ShowVacanciesResponse, error)
	SearchVacancies(context.Context, *SearchVacanciesRequest) (*SearchVacanciesResponse, error)
	ChangeVacancyPosition(context.Context, *ChangeVacancyPositionRequest) (*Empty, error)
	ApproveVacancy(context.Context, *ApproveVacancyRequest) (*Empty, error)
	CloseVacancy(context.Context, *CloseVacancyRequest) (*Empty, error)
	RejectVacancy(context.Context, *RejectVacancyRequest) (*Empty, error)
	TakeInWorkVacancy(context.Context, *TakeInWorkVacancyRequest) (*Empty, error)
	GetVacancy(context.Context, *GetVacancyRequest) (*GetVacancyResponse, error)
	// recruitment
	ConsiderCandidate(context.Context, *ConsiderCandidateRequest) (*ConsiderCandidateResponse, error)
	ConsiderCandidateAnotherVacancy(context.Context, *ConsiderCandidateAnotherVacancyRequest) (*ConsiderCandidateAnotherVacancyResponse, error)
	AcceptRecruitmentStage(context.Context, *AcceptRecruitmentStageRequest) (*Empty, error)
	DenyRecruitment(context.Context, *DenyRecruitmentRequest) (*Empty, error)
	GetRecruitment(context.Context, *GetRecruitmentRequest) (*GetRecruitmentResponse, error)
	ShowRecruitments(context.Context, *ShowRecruitmentRequest) (*ShowRecruitmentsResponse, error)
	mustEmbedUnimplementedRecruitmentServiceServer()
}

// UnimplementedRecruitmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecruitmentServiceServer struct {
}

func (UnimplementedRecruitmentServiceServer) PostVacancy(context.Context, *PostVacancyRequest) (*PostVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) ShowVacancies(context.Context, *ShowVacanciesRequest) (*ShowVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowVacancies not implemented")
}
func (UnimplementedRecruitmentServiceServer) SearchVacancies(context.Context, *SearchVacanciesRequest) (*SearchVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVacancies not implemented")
}
func (UnimplementedRecruitmentServiceServer) ChangeVacancyPosition(context.Context, *ChangeVacancyPositionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeVacancyPosition not implemented")
}
func (UnimplementedRecruitmentServiceServer) ApproveVacancy(context.Context, *ApproveVacancyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) CloseVacancy(context.Context, *CloseVacancyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) RejectVacancy(context.Context, *RejectVacancyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) TakeInWorkVacancy(context.Context, *TakeInWorkVacancyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeInWorkVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) GetVacancy(context.Context, *GetVacancyRequest) (*GetVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) ConsiderCandidate(context.Context, *ConsiderCandidateRequest) (*ConsiderCandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsiderCandidate not implemented")
}
func (UnimplementedRecruitmentServiceServer) ConsiderCandidateAnotherVacancy(context.Context, *ConsiderCandidateAnotherVacancyRequest) (*ConsiderCandidateAnotherVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsiderCandidateAnotherVacancy not implemented")
}
func (UnimplementedRecruitmentServiceServer) AcceptRecruitmentStage(context.Context, *AcceptRecruitmentStageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRecruitmentStage not implemented")
}
func (UnimplementedRecruitmentServiceServer) DenyRecruitment(context.Context, *DenyRecruitmentRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyRecruitment not implemented")
}
func (UnimplementedRecruitmentServiceServer) GetRecruitment(context.Context, *GetRecruitmentRequest) (*GetRecruitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecruitment not implemented")
}
func (UnimplementedRecruitmentServiceServer) ShowRecruitments(context.Context, *ShowRecruitmentRequest) (*ShowRecruitmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowRecruitments not implemented")
}
func (UnimplementedRecruitmentServiceServer) mustEmbedUnimplementedRecruitmentServiceServer() {}

// UnsafeRecruitmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecruitmentServiceServer will
// result in compilation errors.
type UnsafeRecruitmentServiceServer interface {
	mustEmbedUnimplementedRecruitmentServiceServer()
}

func RegisterRecruitmentServiceServer(s grpc.ServiceRegistrar, srv RecruitmentServiceServer) {
	s.RegisterService(&RecruitmentService_ServiceDesc, srv)
}

func _RecruitmentService_PostVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).PostVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/PostVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).PostVacancy(ctx, req.(*PostVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ShowVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ShowVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ShowVacancies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ShowVacancies(ctx, req.(*ShowVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_SearchVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).SearchVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/SearchVacancies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).SearchVacancies(ctx, req.(*SearchVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ChangeVacancyPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeVacancyPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ChangeVacancyPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ChangeVacancyPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ChangeVacancyPosition(ctx, req.(*ChangeVacancyPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ApproveVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ApproveVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ApproveVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ApproveVacancy(ctx, req.(*ApproveVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_CloseVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).CloseVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/CloseVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).CloseVacancy(ctx, req.(*CloseVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_RejectVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).RejectVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/RejectVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).RejectVacancy(ctx, req.(*RejectVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_TakeInWorkVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeInWorkVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).TakeInWorkVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/TakeInWorkVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).TakeInWorkVacancy(ctx, req.(*TakeInWorkVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_GetVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).GetVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/GetVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).GetVacancy(ctx, req.(*GetVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ConsiderCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsiderCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ConsiderCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ConsiderCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ConsiderCandidate(ctx, req.(*ConsiderCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ConsiderCandidateAnotherVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsiderCandidateAnotherVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ConsiderCandidateAnotherVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ConsiderCandidateAnotherVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ConsiderCandidateAnotherVacancy(ctx, req.(*ConsiderCandidateAnotherVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_AcceptRecruitmentStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRecruitmentStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).AcceptRecruitmentStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/AcceptRecruitmentStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).AcceptRecruitmentStage(ctx, req.(*AcceptRecruitmentStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_DenyRecruitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyRecruitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).DenyRecruitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/DenyRecruitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).DenyRecruitment(ctx, req.(*DenyRecruitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_GetRecruitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecruitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).GetRecruitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/GetRecruitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).GetRecruitment(ctx, req.(*GetRecruitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitmentService_ShowRecruitments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRecruitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitmentServiceServer).ShowRecruitments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitmentService/ShowRecruitments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitmentServiceServer).ShowRecruitments(ctx, req.(*ShowRecruitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecruitmentService_ServiceDesc is the grpc.ServiceDesc for RecruitmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecruitmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RecruitmentService",
	HandlerType: (*RecruitmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostVacancy",
			Handler:    _RecruitmentService_PostVacancy_Handler,
		},
		{
			MethodName: "ShowVacancies",
			Handler:    _RecruitmentService_ShowVacancies_Handler,
		},
		{
			MethodName: "SearchVacancies",
			Handler:    _RecruitmentService_SearchVacancies_Handler,
		},
		{
			MethodName: "ChangeVacancyPosition",
			Handler:    _RecruitmentService_ChangeVacancyPosition_Handler,
		},
		{
			MethodName: "ApproveVacancy",
			Handler:    _RecruitmentService_ApproveVacancy_Handler,
		},
		{
			MethodName: "CloseVacancy",
			Handler:    _RecruitmentService_CloseVacancy_Handler,
		},
		{
			MethodName: "RejectVacancy",
			Handler:    _RecruitmentService_RejectVacancy_Handler,
		},
		{
			MethodName: "TakeInWorkVacancy",
			Handler:    _RecruitmentService_TakeInWorkVacancy_Handler,
		},
		{
			MethodName: "GetVacancy",
			Handler:    _RecruitmentService_GetVacancy_Handler,
		},
		{
			MethodName: "ConsiderCandidate",
			Handler:    _RecruitmentService_ConsiderCandidate_Handler,
		},
		{
			MethodName: "ConsiderCandidateAnotherVacancy",
			Handler:    _RecruitmentService_ConsiderCandidateAnotherVacancy_Handler,
		},
		{
			MethodName: "AcceptRecruitmentStage",
			Handler:    _RecruitmentService_AcceptRecruitmentStage_Handler,
		},
		{
			MethodName: "DenyRecruitment",
			Handler:    _RecruitmentService_DenyRecruitment_Handler,
		},
		{
			MethodName: "GetRecruitment",
			Handler:    _RecruitmentService_GetRecruitment_Handler,
		},
		{
			MethodName: "ShowRecruitments",
			Handler:    _RecruitmentService_ShowRecruitments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-service/recruitment-service.proto",
}
