package v1
import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)
const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
type SubmitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContestantId  string                 `protobuf:"bytes,1,opt,name=contestant_id,json=contestantId,proto3" json:"contestant_id,omitempty"`
	ArtifactUrl   string                 `protobuf:"bytes,2,opt,name=artifact_url,json=artifactUrl,proto3" json:"artifact_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*SubmitRequest) ProtoMessage() {}
func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{0}
}
func (x *SubmitRequest) GetContestantId() string {
	if x != nil {
		return x.ContestantId
	}
	return ""
}
func (x *SubmitRequest) GetArtifactUrl() string {
	if x != nil {
		return x.ArtifactUrl
	}
	return ""
}
type SubmitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*SubmitResponse) ProtoMessage() {}
func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{1}
}
func (x *SubmitResponse) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
type StatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StatusRequest) ProtoMessage() {}
func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{2}
}
func (x *StatusRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
type StatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StatusResponse) ProtoMessage() {}
func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{3}
}
func (x *StatusResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}
type DeployRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	ImageRef      string                 `protobuf:"bytes,2,opt,name=image_ref,json=imageRef,proto3" json:"image_ref,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DeployRequest) ProtoMessage() {}
func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{4}
}
func (x *DeployRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *DeployRequest) GetImageRef() string {
	if x != nil {
		return x.ImageRef
	}
	return ""
}
type DeployResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	PodIp         string                 `protobuf:"bytes,2,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DeployResponse) ProtoMessage() {}
func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{5}
}
func (x *DeployResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}
func (x *DeployResponse) GetPodIp() string {
	if x != nil {
		return x.PodIp
	}
	return ""
}
type TeardownRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *TeardownRequest) Reset() {
	*x = TeardownRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *TeardownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*TeardownRequest) ProtoMessage() {}
func (x *TeardownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*TeardownRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{6}
}
func (x *TeardownRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
type TeardownResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *TeardownResponse) Reset() {
	*x = TeardownResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *TeardownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*TeardownResponse) ProtoMessage() {}
func (x *TeardownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*TeardownResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{7}
}
func (x *TeardownResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}
type StartLoadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	TargetIp      string                 `protobuf:"bytes,2,opt,name=target_ip,json=targetIp,proto3" json:"target_ip,omitempty"`
	DurationSec   int32                  `protobuf:"varint,3,opt,name=duration_sec,json=durationSec,proto3" json:"duration_sec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StartLoadRequest) Reset() {
	*x = StartLoadRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StartLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StartLoadRequest) ProtoMessage() {}
func (x *StartLoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StartLoadRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{8}
}
func (x *StartLoadRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *StartLoadRequest) GetTargetIp() string {
	if x != nil {
		return x.TargetIp
	}
	return ""
}
func (x *StartLoadRequest) GetDurationSec() int32 {
	if x != nil {
		return x.DurationSec
	}
	return 0
}
type StartLoadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StartLoadResponse) Reset() {
	*x = StartLoadResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StartLoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StartLoadResponse) ProtoMessage() {}
func (x *StartLoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StartLoadResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{9}
}
func (x *StartLoadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}
type StopLoadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StopLoadRequest) Reset() {
	*x = StopLoadRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StopLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StopLoadRequest) ProtoMessage() {}
func (x *StopLoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StopLoadRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{10}
}
func (x *StopLoadRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
type StopLoadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StopLoadResponse) Reset() {
	*x = StopLoadResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StopLoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StopLoadResponse) ProtoMessage() {}
func (x *StopLoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StopLoadResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{11}
}
func (x *StopLoadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}
type IngestResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *IngestResponse) Reset() {
	*x = IngestResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *IngestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*IngestResponse) ProtoMessage() {}
func (x *IngestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*IngestResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{12}
}
func (x *IngestResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}
type LeaderboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaderboardRequest) Reset() {
	*x = LeaderboardRequest{}
	mi := &file_iicpc_v1_services_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaderboardRequest) ProtoMessage() {}
func (x *LeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{13}
}
func (x *LeaderboardRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}
type LeaderboardResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	RankedContestants []string               `protobuf:"bytes,1,rep,name=ranked_contestants,json=rankedContestants,proto3" json:"ranked_contestants,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}
func (x *LeaderboardResponse) Reset() {
	*x = LeaderboardResponse{}
	mi := &file_iicpc_v1_services_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaderboardResponse) ProtoMessage() {}
func (x *LeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_services_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_services_proto_rawDescGZIP(), []int{14}
}
func (x *LeaderboardResponse) GetRankedContestants() []string {
	if x != nil {
		return x.RankedContestants
	}
	return nil
}
var File_iicpc_v1_services_proto protoreflect.FileDescriptor
const file_iicpc_v1_services_proto_rawDesc = "" +
	"\n" +
	"\x17iicpc/v1/services.proto\x12\biicpc.v1\x1a\x17iicpc/v1/messages.proto\"W\n" +
	"\rSubmitRequest\x12#\n" +
	"\rcontestant_id\x18\x01 \x01(\tR\fcontestantId\x12!\n" +
	"\fartifact_url\x18\x02 \x01(\tR\vartifactUrl\"'\n" +
	"\x0eSubmitResponse\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\"&\n" +
	"\rStatusRequest\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\"&\n" +
	"\x0eStatusResponse\x12\x14\n" +
	"\x05state\x18\x01 \x01(\tR\x05state\"C\n" +
	"\rDeployRequest\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x1b\n" +
	"\timage_ref\x18\x02 \x01(\tR\bimageRef\"A\n" +
	"\x0eDeployResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x15\n" +
	"\x06pod_ip\x18\x02 \x01(\tR\x05podIp\"(\n" +
	"\x0fTeardownRequest\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\",\n" +
	"\x10TeardownResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"i\n" +
	"\x10StartLoadRequest\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x1b\n" +
	"\ttarget_ip\x18\x02 \x01(\tR\btargetIp\x12!\n" +
	"\fduration_sec\x18\x03 \x01(\x05R\vdurationSec\"-\n" +
	"\x11StartLoadResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"(\n" +
	"\x0fStopLoadRequest\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\",\n" +
	"\x10StopLoadResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"*\n" +
	"\x0eIngestResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"*\n" +
	"\x12LeaderboardRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\"D\n" +
	"\x13LeaderboardResponse\x12-\n" +
	"\x12ranked_contestants\x18\x01 \x03(\tR\x11rankedContestants2\x94\x01\n" +
	"\fOrchestrator\x12A\n" +
	"\fSubmitEngine\x12\x17.iicpc.v1.SubmitRequest\x1a\x18.iicpc.v1.SubmitResponse\x12A\n" +
	"\fGetRunStatus\x12\x17.iicpc.v1.StatusRequest\x1a\x18.iicpc.v1.StatusResponse2\x9f\x01\n" +
	"\x0fSandboxDeployer\x12B\n" +
	"\rDeploySandbox\x12\x17.iicpc.v1.DeployRequest\x1a\x18.iicpc.v1.DeployResponse\x12H\n" +
	"\x0fTeardownSandbox\x12\x19.iicpc.v1.TeardownRequest\x1a\x1a.iicpc.v1.TeardownResponse2\x98\x01\n" +
	"\rBotController\x12D\n" +
	"\tStartLoad\x12\x1a.iicpc.v1.StartLoadRequest\x1a\x1b.iicpc.v1.StartLoadResponse\x12A\n" +
	"\bStopLoad\x12\x19.iicpc.v1.StopLoadRequest\x1a\x1a.iicpc.v1.StopLoadResponse2N\n" +
	"\bIngester\x12B\n" +
	"\x0fStreamTelemetry\x12\x13.iicpc.v1.Telemetry\x1a\x18.iicpc.v1.IngestResponse(\x012\\\n" +
	"\aScoring\x12Q\n" +
	"\x12GetLiveLeaderboard\x12\x1c.iicpc.v1.LeaderboardRequest\x1a\x1d.iicpc.v1.LeaderboardResponseB;Z9github.com/devrajdas/iicpc-platform-2026/schemas/iicpc/v1b\x06proto3"
var (
	file_iicpc_v1_services_proto_rawDescOnce sync.Once
	file_iicpc_v1_services_proto_rawDescData []byte
)
func file_iicpc_v1_services_proto_rawDescGZIP() []byte {
	file_iicpc_v1_services_proto_rawDescOnce.Do(func() {
		file_iicpc_v1_services_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_iicpc_v1_services_proto_rawDesc), len(file_iicpc_v1_services_proto_rawDesc)))
	})
	return file_iicpc_v1_services_proto_rawDescData
}
var file_iicpc_v1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_iicpc_v1_services_proto_goTypes = []any{
	(*SubmitRequest)(nil),
	(*SubmitResponse)(nil),
	(*StatusRequest)(nil),
	(*StatusResponse)(nil),
	(*DeployRequest)(nil),
	(*DeployResponse)(nil),
	(*TeardownRequest)(nil),
	(*TeardownResponse)(nil),
	(*StartLoadRequest)(nil),
	(*StartLoadResponse)(nil),
	(*StopLoadRequest)(nil),
	(*StopLoadResponse)(nil),
	(*IngestResponse)(nil),
	(*LeaderboardRequest)(nil),
	(*LeaderboardResponse)(nil),
	(*Telemetry)(nil),
}
var file_iicpc_v1_services_proto_depIdxs = []int32{
	0,
	2,
	4,
	6,
	8,
	10,
	15,
	13,
	1,
	3,
	5,
	7,
	9,
	11,
	12,
	14,
	8,
	0,
	0,
	0,
	0,
}
func init() { file_iicpc_v1_services_proto_init() }
func file_iicpc_v1_services_proto_init() {
	if File_iicpc_v1_services_proto != nil {
		return
	}
	file_iicpc_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_iicpc_v1_services_proto_rawDesc), len(file_iicpc_v1_services_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_iicpc_v1_services_proto_goTypes,
		DependencyIndexes: file_iicpc_v1_services_proto_depIdxs,
		MessageInfos:      file_iicpc_v1_services_proto_msgTypes,
	}.Build()
	File_iicpc_v1_services_proto = out.File
	file_iicpc_v1_services_proto_goTypes = nil
	file_iicpc_v1_services_proto_depIdxs = nil
}
