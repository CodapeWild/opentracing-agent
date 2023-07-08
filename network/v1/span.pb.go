// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: idl/proto/v1/span.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpanLocation int32

const (
	SpanLocation_Unknown SpanLocation = 0
	SpanLocation_Entry   SpanLocation = 1
	SpanLocation_Local   SpanLocation = 2
	SpanLocation_Exit    SpanLocation = 3
)

// Enum value maps for SpanLocation.
var (
	SpanLocation_name = map[int32]string{
		0: "Unknown",
		1: "Entry",
		2: "Local",
		3: "Exit",
	}
	SpanLocation_value = map[string]int32{
		"Unknown": 0,
		"Entry":   1,
		"Local":   2,
		"Exit":    3,
	}
)

func (x SpanLocation) Enum() *SpanLocation {
	p := new(SpanLocation)
	*p = x
	return p
}

func (x SpanLocation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanLocation) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_proto_v1_span_proto_enumTypes[0].Descriptor()
}

func (SpanLocation) Type() protoreflect.EnumType {
	return &file_idl_proto_v1_span_proto_enumTypes[0]
}

func (x SpanLocation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanLocation.Descriptor instead.
func (SpanLocation) EnumDescriptor() ([]byte, []int) {
	return file_idl_proto_v1_span_proto_rawDescGZIP(), []int{0}
}

type SpanStatus int32

const (
	SpanStatus_Undefined SpanStatus = 0
	SpanStatus_OK        SpanStatus = 1
	SpanStatus_Error     SpanStatus = 2
	SpanStatus_Recover   SpanStatus = 3
	SpanStatus_Fatal     SpanStatus = 4
)

// Enum value maps for SpanStatus.
var (
	SpanStatus_name = map[int32]string{
		0: "Undefined",
		1: "OK",
		2: "Error",
		3: "Recover",
		4: "Fatal",
	}
	SpanStatus_value = map[string]int32{
		"Undefined": 0,
		"OK":        1,
		"Error":     2,
		"Recover":   3,
		"Fatal":     4,
	}
)

func (x SpanStatus) Enum() *SpanStatus {
	p := new(SpanStatus)
	*p = x
	return p
}

func (x SpanStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_proto_v1_span_proto_enumTypes[1].Descriptor()
}

func (SpanStatus) Type() protoreflect.EnumType {
	return &file_idl_proto_v1_span_proto_enumTypes[1]
}

func (x SpanStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanStatus.Descriptor instead.
func (SpanStatus) EnumDescriptor() ([]byte, []int) {
	return file_idl_proto_v1_span_proto_rawDescGZIP(), []int{1}
}

type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID   uint64              `protobuf:"varint,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	ParentID  uint64              `protobuf:"varint,2,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	SpanID    uint64              `protobuf:"varint,3,opt,name=SpanID,proto3" json:"SpanID,omitempty"`
	Service   string              `protobuf:"bytes,4,opt,name=Service,proto3" json:"Service,omitempty"`
	Resource  string              `protobuf:"bytes,5,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Operation string              `protobuf:"bytes,6,opt,name=Operation,proto3" json:"Operation,omitempty"`
	Location  SpanLocation        `protobuf:"varint,7,opt,name=Location,proto3,enum=opentracing.v1.SpanLocation" json:"Location,omitempty"`
	Status    SpanStatus          `protobuf:"varint,8,opt,name=Status,proto3,enum=opentracing.v1.SpanStatus" json:"Status,omitempty"`
	Meta      map[string]string   `protobuf:"bytes,9,rep,name=Meta,proto3" json:"Meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metrics   map[string]*Numeric `protobuf:"bytes,10,rep,name=Metrics,proto3" json:"Metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Start     uint64              `protobuf:"varint,11,opt,name=Start,proto3" json:"Start,omitempty"`
	Duration  uint64              `protobuf:"varint,12,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Source    *SpanVendor         `protobuf:"bytes,13,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_v1_span_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_v1_span_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_idl_proto_v1_span_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetTraceID() uint64 {
	if x != nil {
		return x.TraceID
	}
	return 0
}

func (x *Span) GetParentID() uint64 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *Span) GetSpanID() uint64 {
	if x != nil {
		return x.SpanID
	}
	return 0
}

func (x *Span) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Span) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Span) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Span) GetLocation() SpanLocation {
	if x != nil {
		return x.Location
	}
	return SpanLocation_Unknown
}

func (x *Span) GetStatus() SpanStatus {
	if x != nil {
		return x.Status
	}
	return SpanStatus_Undefined
}

func (x *Span) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Span) GetMetrics() map[string]*Numeric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Span) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Span) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Span) GetSource() *SpanVendor {
	if x != nil {
		return x.Source
	}
	return nil
}

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trace []*Span `protobuf:"bytes,1,rep,name=Trace,proto3" json:"Trace,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_v1_span_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_v1_span_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_idl_proto_v1_span_proto_rawDescGZIP(), []int{1}
}

func (x *Trace) GetTrace() []*Span {
	if x != nil {
		return x.Trace
	}
	return nil
}

var File_idl_proto_v1_span_proto protoreflect.FileDescriptor

var file_idl_proto_v1_span_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x04, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61,
	0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a,
	0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x33, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e,
	0x52, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2a, 0x3b, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x6e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x78,
	0x69, 0x74, 0x10, 0x03, 0x2a, 0x46, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x10, 0x04, 0x32, 0x97, 0x01, 0x0a,
	0x13, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x15, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x1a, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x70,
	0x61, 0x6e, 0x12, 0x14, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x1a, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x64, 0x61, 0x70, 0x65, 0x57, 0x69, 0x6c, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_proto_v1_span_proto_rawDescOnce sync.Once
	file_idl_proto_v1_span_proto_rawDescData = file_idl_proto_v1_span_proto_rawDesc
)

func file_idl_proto_v1_span_proto_rawDescGZIP() []byte {
	file_idl_proto_v1_span_proto_rawDescOnce.Do(func() {
		file_idl_proto_v1_span_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_proto_v1_span_proto_rawDescData)
	})
	return file_idl_proto_v1_span_proto_rawDescData
}

var file_idl_proto_v1_span_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_idl_proto_v1_span_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_proto_v1_span_proto_goTypes = []interface{}{
	(SpanLocation)(0),  // 0: opentracing.v1.SpanLocation
	(SpanStatus)(0),    // 1: opentracing.v1.SpanStatus
	(*Span)(nil),       // 2: opentracing.v1.Span
	(*Trace)(nil),      // 3: opentracing.v1.Trace
	nil,                // 4: opentracing.v1.Span.MetaEntry
	nil,                // 5: opentracing.v1.Span.MetricsEntry
	(*SpanVendor)(nil), // 6: opentracing.v1.SpanVendor
	(*Numeric)(nil),    // 7: opentracing.v1.Numeric
	(*Response)(nil),   // 8: opentracing.v1.Response
}
var file_idl_proto_v1_span_proto_depIdxs = []int32{
	0, // 0: opentracing.v1.Span.Location:type_name -> opentracing.v1.SpanLocation
	1, // 1: opentracing.v1.Span.Status:type_name -> opentracing.v1.SpanStatus
	4, // 2: opentracing.v1.Span.Meta:type_name -> opentracing.v1.Span.MetaEntry
	5, // 3: opentracing.v1.Span.Metrics:type_name -> opentracing.v1.Span.MetricsEntry
	6, // 4: opentracing.v1.Span.Source:type_name -> opentracing.v1.SpanVendor
	2, // 5: opentracing.v1.Trace.Trace:type_name -> opentracing.v1.Span
	7, // 6: opentracing.v1.Span.MetricsEntry.value:type_name -> opentracing.v1.Numeric
	3, // 7: opentracing.v1.TracesReportService.SendTrace:input_type -> opentracing.v1.Trace
	2, // 8: opentracing.v1.TracesReportService.SendSpan:input_type -> opentracing.v1.Span
	8, // 9: opentracing.v1.TracesReportService.SendTrace:output_type -> opentracing.v1.Response
	8, // 10: opentracing.v1.TracesReportService.SendSpan:output_type -> opentracing.v1.Response
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_idl_proto_v1_span_proto_init() }
func file_idl_proto_v1_span_proto_init() {
	if File_idl_proto_v1_span_proto != nil {
		return
	}
	file_idl_proto_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_idl_proto_v1_span_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_proto_v1_span_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_proto_v1_span_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_proto_v1_span_proto_goTypes,
		DependencyIndexes: file_idl_proto_v1_span_proto_depIdxs,
		EnumInfos:         file_idl_proto_v1_span_proto_enumTypes,
		MessageInfos:      file_idl_proto_v1_span_proto_msgTypes,
	}.Build()
	File_idl_proto_v1_span_proto = out.File
	file_idl_proto_v1_span_proto_rawDesc = nil
	file_idl_proto_v1_span_proto_goTypes = nil
	file_idl_proto_v1_span_proto_depIdxs = nil
}
