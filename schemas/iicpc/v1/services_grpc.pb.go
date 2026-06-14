package v1
import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)
const _ = grpc.SupportPackageIsVersion9
const (
	Orchestrator_SubmitEngine_FullMethodName = "/iicpc.v1.Orchestrator/SubmitEngine"
	Orchestrator_GetRunStatus_FullMethodName = "/iicpc.v1.Orchestrator/GetRunStatus"
)
type OrchestratorClient interface {
	SubmitEngine(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	GetRunStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}
type orchestratorClient struct {
	cc grpc.ClientConnInterface
}
func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}
func (c *orchestratorClient) SubmitEngine(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, Orchestrator_SubmitEngine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *orchestratorClient) GetRunStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Orchestrator_GetRunStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
type OrchestratorServer interface {
	SubmitEngine(context.Context, *SubmitRequest) (*SubmitResponse, error)
	GetRunStatus(context.Context, *StatusRequest) (*StatusResponse, error)
}
type UnimplementedOrchestratorServer struct{}
func (UnimplementedOrchestratorServer) SubmitEngine(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SubmitEngine not implemented")
}
func (UnimplementedOrchestratorServer) GetRunStatus(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetRunStatus not implemented")
}
func (UnimplementedOrchestratorServer) testEmbeddedByValue() {}
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}
func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}
func _Orchestrator_SubmitEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).SubmitEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_SubmitEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).SubmitEngine(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Orchestrator_GetRunStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetRunStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_GetRunStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetRunStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iicpc.v1.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitEngine",
			Handler:    _Orchestrator_SubmitEngine_Handler,
		},
		{
			MethodName: "GetRunStatus",
			Handler:    _Orchestrator_GetRunStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iicpc/v1/services.proto",
}
const (
	SandboxDeployer_DeploySandbox_FullMethodName   = "/iicpc.v1.SandboxDeployer/DeploySandbox"
	SandboxDeployer_TeardownSandbox_FullMethodName = "/iicpc.v1.SandboxDeployer/TeardownSandbox"
)
type SandboxDeployerClient interface {
	DeploySandbox(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error)
	TeardownSandbox(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error)
}
type sandboxDeployerClient struct {
	cc grpc.ClientConnInterface
}
func NewSandboxDeployerClient(cc grpc.ClientConnInterface) SandboxDeployerClient {
	return &sandboxDeployerClient{cc}
}
func (c *sandboxDeployerClient) DeploySandbox(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeployResponse)
	err := c.cc.Invoke(ctx, SandboxDeployer_DeploySandbox_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *sandboxDeployerClient) TeardownSandbox(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeardownResponse)
	err := c.cc.Invoke(ctx, SandboxDeployer_TeardownSandbox_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
type SandboxDeployerServer interface {
	DeploySandbox(context.Context, *DeployRequest) (*DeployResponse, error)
	TeardownSandbox(context.Context, *TeardownRequest) (*TeardownResponse, error)
}
type UnimplementedSandboxDeployerServer struct{}
func (UnimplementedSandboxDeployerServer) DeploySandbox(context.Context, *DeployRequest) (*DeployResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DeploySandbox not implemented")
}
func (UnimplementedSandboxDeployerServer) TeardownSandbox(context.Context, *TeardownRequest) (*TeardownResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method TeardownSandbox not implemented")
}
func (UnimplementedSandboxDeployerServer) testEmbeddedByValue() {}
type UnsafeSandboxDeployerServer interface {
	mustEmbedUnimplementedSandboxDeployerServer()
}
func RegisterSandboxDeployerServer(s grpc.ServiceRegistrar, srv SandboxDeployerServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SandboxDeployer_ServiceDesc, srv)
}
func _SandboxDeployer_DeploySandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxDeployerServer).DeploySandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxDeployer_DeploySandbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxDeployerServer).DeploySandbox(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _SandboxDeployer_TeardownSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxDeployerServer).TeardownSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SandboxDeployer_TeardownSandbox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxDeployerServer).TeardownSandbox(ctx, req.(*TeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var SandboxDeployer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iicpc.v1.SandboxDeployer",
	HandlerType: (*SandboxDeployerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeploySandbox",
			Handler:    _SandboxDeployer_DeploySandbox_Handler,
		},
		{
			MethodName: "TeardownSandbox",
			Handler:    _SandboxDeployer_TeardownSandbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iicpc/v1/services.proto",
}
const (
	BotController_StartLoad_FullMethodName = "/iicpc.v1.BotController/StartLoad"
	BotController_StopLoad_FullMethodName  = "/iicpc.v1.BotController/StopLoad"
)
type BotControllerClient interface {
	StartLoad(ctx context.Context, in *StartLoadRequest, opts ...grpc.CallOption) (*StartLoadResponse, error)
	StopLoad(ctx context.Context, in *StopLoadRequest, opts ...grpc.CallOption) (*StopLoadResponse, error)
}
type botControllerClient struct {
	cc grpc.ClientConnInterface
}
func NewBotControllerClient(cc grpc.ClientConnInterface) BotControllerClient {
	return &botControllerClient{cc}
}
func (c *botControllerClient) StartLoad(ctx context.Context, in *StartLoadRequest, opts ...grpc.CallOption) (*StartLoadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartLoadResponse)
	err := c.cc.Invoke(ctx, BotController_StartLoad_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *botControllerClient) StopLoad(ctx context.Context, in *StopLoadRequest, opts ...grpc.CallOption) (*StopLoadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopLoadResponse)
	err := c.cc.Invoke(ctx, BotController_StopLoad_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
type BotControllerServer interface {
	StartLoad(context.Context, *StartLoadRequest) (*StartLoadResponse, error)
	StopLoad(context.Context, *StopLoadRequest) (*StopLoadResponse, error)
}
type UnimplementedBotControllerServer struct{}
func (UnimplementedBotControllerServer) StartLoad(context.Context, *StartLoadRequest) (*StartLoadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method StartLoad not implemented")
}
func (UnimplementedBotControllerServer) StopLoad(context.Context, *StopLoadRequest) (*StopLoadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method StopLoad not implemented")
}
func (UnimplementedBotControllerServer) testEmbeddedByValue() {}
type UnsafeBotControllerServer interface {
	mustEmbedUnimplementedBotControllerServer()
}
func RegisterBotControllerServer(s grpc.ServiceRegistrar, srv BotControllerServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BotController_ServiceDesc, srv)
}
func _BotController_StartLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotControllerServer).StartLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotController_StartLoad_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotControllerServer).StartLoad(ctx, req.(*StartLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _BotController_StopLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotControllerServer).StopLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotController_StopLoad_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotControllerServer).StopLoad(ctx, req.(*StopLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var BotController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iicpc.v1.BotController",
	HandlerType: (*BotControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartLoad",
			Handler:    _BotController_StartLoad_Handler,
		},
		{
			MethodName: "StopLoad",
			Handler:    _BotController_StopLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iicpc/v1/services.proto",
}
const (
	Ingester_StreamTelemetry_FullMethodName = "/iicpc.v1.Ingester/StreamTelemetry"
)
type IngesterClient interface {
	StreamTelemetry(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Telemetry, IngestResponse], error)
}
type ingesterClient struct {
	cc grpc.ClientConnInterface
}
func NewIngesterClient(cc grpc.ClientConnInterface) IngesterClient {
	return &ingesterClient{cc}
}
func (c *ingesterClient) StreamTelemetry(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Telemetry, IngestResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Ingester_ServiceDesc.Streams[0], Ingester_StreamTelemetry_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Telemetry, IngestResponse]{ClientStream: stream}
	return x, nil
}
type Ingester_StreamTelemetryClient = grpc.ClientStreamingClient[Telemetry, IngestResponse]
type IngesterServer interface {
	StreamTelemetry(grpc.ClientStreamingServer[Telemetry, IngestResponse]) error
}
type UnimplementedIngesterServer struct{}
func (UnimplementedIngesterServer) StreamTelemetry(grpc.ClientStreamingServer[Telemetry, IngestResponse]) error {
	return status.Error(codes.Unimplemented, "method StreamTelemetry not implemented")
}
func (UnimplementedIngesterServer) testEmbeddedByValue() {}
type UnsafeIngesterServer interface {
	mustEmbedUnimplementedIngesterServer()
}
func RegisterIngesterServer(s grpc.ServiceRegistrar, srv IngesterServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ingester_ServiceDesc, srv)
}
func _Ingester_StreamTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngesterServer).StreamTelemetry(&grpc.GenericServerStream[Telemetry, IngestResponse]{ServerStream: stream})
}
type Ingester_StreamTelemetryServer = grpc.ClientStreamingServer[Telemetry, IngestResponse]
var Ingester_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iicpc.v1.Ingester",
	HandlerType: (*IngesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTelemetry",
			Handler:       _Ingester_StreamTelemetry_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "iicpc/v1/services.proto",
}
const (
	Scoring_GetLiveLeaderboard_FullMethodName = "/iicpc.v1.Scoring/GetLiveLeaderboard"
)
type ScoringClient interface {
	GetLiveLeaderboard(ctx context.Context, in *LeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error)
}
type scoringClient struct {
	cc grpc.ClientConnInterface
}
func NewScoringClient(cc grpc.ClientConnInterface) ScoringClient {
	return &scoringClient{cc}
}
func (c *scoringClient) GetLiveLeaderboard(ctx context.Context, in *LeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, Scoring_GetLiveLeaderboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
type ScoringServer interface {
	GetLiveLeaderboard(context.Context, *LeaderboardRequest) (*LeaderboardResponse, error)
}
type UnimplementedScoringServer struct{}
func (UnimplementedScoringServer) GetLiveLeaderboard(context.Context, *LeaderboardRequest) (*LeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetLiveLeaderboard not implemented")
}
func (UnimplementedScoringServer) testEmbeddedByValue() {}
type UnsafeScoringServer interface {
	mustEmbedUnimplementedScoringServer()
}
func RegisterScoringServer(s grpc.ServiceRegistrar, srv ScoringServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Scoring_ServiceDesc, srv)
}
func _Scoring_GetLiveLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServer).GetLiveLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scoring_GetLiveLeaderboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServer).GetLiveLeaderboard(ctx, req.(*LeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var Scoring_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iicpc.v1.Scoring",
	HandlerType: (*ScoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLiveLeaderboard",
			Handler:    _Scoring_GetLiveLeaderboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iicpc/v1/services.proto",
}
