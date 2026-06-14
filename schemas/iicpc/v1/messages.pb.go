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
type Order struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RunId          string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	SequenceNumber int64                  `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	OrderType      string                 `protobuf:"bytes,3,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Side           string                 `protobuf:"bytes,4,opt,name=side,proto3" json:"side,omitempty"`
	Quantity       int32                  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price          int64                  `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	IntendedSendTs int64                  `protobuf:"varint,7,opt,name=intended_send_ts,json=intendedSendTs,proto3" json:"intended_send_ts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}
func (x *Order) Reset() {
	*x = Order{}
	mi := &file_iicpc_v1_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Order) ProtoMessage() {}
func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Order) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_messages_proto_rawDescGZIP(), []int{0}
}
func (x *Order) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *Order) GetSequenceNumber() int64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}
func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}
func (x *Order) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}
func (x *Order) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}
func (x *Order) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}
func (x *Order) GetIntendedSendTs() int64 {
	if x != nil {
		return x.IntendedSendTs
	}
	return 0
}
type Ack struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RunId          string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	SequenceNumber int64                  `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	Status         string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	AckTs          int64                  `protobuf:"varint,4,opt,name=ack_ts,json=ackTs,proto3" json:"ack_ts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}
func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_iicpc_v1_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Ack) ProtoMessage() {}
func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Ack) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_messages_proto_rawDescGZIP(), []int{1}
}
func (x *Ack) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *Ack) GetSequenceNumber() int64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}
func (x *Ack) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}
func (x *Ack) GetAckTs() int64 {
	if x != nil {
		return x.AckTs
	}
	return 0
}
type Telemetry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	SendTs        int64                  `protobuf:"varint,2,opt,name=send_ts,json=sendTs,proto3" json:"send_ts,omitempty"`
	AckTs         int64                  `protobuf:"varint,3,opt,name=ack_ts,json=ackTs,proto3" json:"ack_ts,omitempty"`
	LatencyUs     int32                  `protobuf:"varint,4,opt,name=latency_us,json=latencyUs,proto3" json:"latency_us,omitempty"`
	ErrorCode     string                 `protobuf:"bytes,5,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *Telemetry) Reset() {
	*x = Telemetry{}
	mi := &file_iicpc_v1_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Telemetry) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Telemetry) ProtoMessage() {}
func (x *Telemetry) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_messages_proto_rawDescGZIP(), []int{2}
}
func (x *Telemetry) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *Telemetry) GetSendTs() int64 {
	if x != nil {
		return x.SendTs
	}
	return 0
}
func (x *Telemetry) GetAckTs() int64 {
	if x != nil {
		return x.AckTs
	}
	return 0
}
func (x *Telemetry) GetLatencyUs() int32 {
	if x != nil {
		return x.LatencyUs
	}
	return 0
}
func (x *Telemetry) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}
type RunEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	State         string                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *RunEvent) Reset() {
	*x = RunEvent{}
	mi := &file_iicpc_v1_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *RunEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*RunEvent) ProtoMessage() {}
func (x *RunEvent) ProtoReflect() protoreflect.Message {
	mi := &file_iicpc_v1_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*RunEvent) Descriptor() ([]byte, []int) {
	return file_iicpc_v1_messages_proto_rawDescGZIP(), []int{3}
}
func (x *RunEvent) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}
func (x *RunEvent) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}
var File_iicpc_v1_messages_proto protoreflect.FileDescriptor
const file_iicpc_v1_messages_proto_rawDesc = "" +
	"\n" +
	"\x17iicpc/v1/messages.proto\x12\biicpc.v1\"\xd6\x01\n" +
	"\x05Order\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12'\n" +
	"\x0fsequence_number\x18\x02 \x01(\x03R\x0esequenceNumber\x12\x1d\n" +
	"\n" +
	"order_type\x18\x03 \x01(\tR\torderType\x12\x12\n" +
	"\x04side\x18\x04 \x01(\tR\x04side\x12\x1a\n" +
	"\bquantity\x18\x05 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x06 \x01(\x03R\x05price\x12(\n" +
	"\x10intended_send_ts\x18\a \x01(\x03R\x0eintendedSendTs\"t\n" +
	"\x03Ack\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12'\n" +
	"\x0fsequence_number\x18\x02 \x01(\x03R\x0esequenceNumber\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\x12\x15\n" +
	"\x06ack_ts\x18\x04 \x01(\x03R\x05ackTs\"\x90\x01\n" +
	"\tTelemetry\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x17\n" +
	"\asend_ts\x18\x02 \x01(\x03R\x06sendTs\x12\x15\n" +
	"\x06ack_ts\x18\x03 \x01(\x03R\x05ackTs\x12\x1d\n" +
	"\n" +
	"latency_us\x18\x04 \x01(\x05R\tlatencyUs\x12\x1d\n" +
	"\n" +
	"error_code\x18\x05 \x01(\tR\terrorCode\"7\n" +
	"\bRunEvent\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x14\n" +
	"\x05state\x18\x02 \x01(\tR\x05stateB;Z9github.com/devrajdas/iicpc-platform-2026/schemas/iicpc/v1b\x06proto3"
var (
	file_iicpc_v1_messages_proto_rawDescOnce sync.Once
	file_iicpc_v1_messages_proto_rawDescData []byte
)
func file_iicpc_v1_messages_proto_rawDescGZIP() []byte {
	file_iicpc_v1_messages_proto_rawDescOnce.Do(func() {
		file_iicpc_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_iicpc_v1_messages_proto_rawDesc), len(file_iicpc_v1_messages_proto_rawDesc)))
	})
	return file_iicpc_v1_messages_proto_rawDescData
}
var file_iicpc_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_iicpc_v1_messages_proto_goTypes = []any{
	(*Order)(nil),
	(*Ack)(nil),
	(*Telemetry)(nil),
	(*RunEvent)(nil),
}
var file_iicpc_v1_messages_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}
func init() { file_iicpc_v1_messages_proto_init() }
func file_iicpc_v1_messages_proto_init() {
	if File_iicpc_v1_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_iicpc_v1_messages_proto_rawDesc), len(file_iicpc_v1_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iicpc_v1_messages_proto_goTypes,
		DependencyIndexes: file_iicpc_v1_messages_proto_depIdxs,
		MessageInfos:      file_iicpc_v1_messages_proto_msgTypes,
	}.Build()
	File_iicpc_v1_messages_proto = out.File
	file_iicpc_v1_messages_proto_goTypes = nil
	file_iicpc_v1_messages_proto_depIdxs = nil
}
