// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liderman/traderstack/strategy/v1/strategy.proto

package strategyv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Стратегия.
type Strategy struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StackId              string             `protobuf:"bytes,2,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	AccountId            string             `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	RunIntervalDuration  *duration.Duration `protobuf:"bytes,4,opt,name=run_interval_duration,json=runIntervalDuration,proto3" json:"run_interval_duration,omitempty"`
	Enabled              bool               `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Strategy) Reset()         { *m = Strategy{} }
func (m *Strategy) String() string { return proto.CompactTextString(m) }
func (*Strategy) ProtoMessage()    {}
func (*Strategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_af238c6140d96e48, []int{0}
}

func (m *Strategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strategy.Unmarshal(m, b)
}
func (m *Strategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strategy.Marshal(b, m, deterministic)
}
func (m *Strategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strategy.Merge(m, src)
}
func (m *Strategy) XXX_Size() int {
	return xxx_messageInfo_Strategy.Size(m)
}
func (m *Strategy) XXX_DiscardUnknown() {
	xxx_messageInfo_Strategy.DiscardUnknown(m)
}

var xxx_messageInfo_Strategy proto.InternalMessageInfo

func (m *Strategy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Strategy) GetStackId() string {
	if m != nil {
		return m.StackId
	}
	return ""
}

func (m *Strategy) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Strategy) GetRunIntervalDuration() *duration.Duration {
	if m != nil {
		return m.RunIntervalDuration
	}
	return nil
}

func (m *Strategy) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// Лог.
type StrategyLog struct {
	LogType              string               `protobuf:"bytes,1,opt,name=log_type,json=logType,proto3" json:"log_type,omitempty"`
	Message              string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StrategyLog) Reset()         { *m = StrategyLog{} }
func (m *StrategyLog) String() string { return proto.CompactTextString(m) }
func (*StrategyLog) ProtoMessage()    {}
func (*StrategyLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_af238c6140d96e48, []int{1}
}

func (m *StrategyLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyLog.Unmarshal(m, b)
}
func (m *StrategyLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyLog.Marshal(b, m, deterministic)
}
func (m *StrategyLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyLog.Merge(m, src)
}
func (m *StrategyLog) XXX_Size() int {
	return xxx_messageInfo_StrategyLog.Size(m)
}
func (m *StrategyLog) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyLog.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyLog proto.InternalMessageInfo

func (m *StrategyLog) GetLogType() string {
	if m != nil {
		return m.LogType
	}
	return ""
}

func (m *StrategyLog) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StrategyLog) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*Strategy)(nil), "liderman.traderstack.strategy.v1.Strategy")
	proto.RegisterType((*StrategyLog)(nil), "liderman.traderstack.strategy.v1.StrategyLog")
}

func init() {
	proto.RegisterFile("liderman/traderstack/strategy/v1/strategy.proto", fileDescriptor_af238c6140d96e48)
}

var fileDescriptor_af238c6140d96e48 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0x37, 0xdd, 0x96, 0xa1, 0x17, 0x11, 0xa1, 0x1b, 0xa8, 0x65, 0x78, 0xb1, 0xab,
	0x94, 0xce, 0x37, 0x18, 0xde, 0x0c, 0x26, 0x8c, 0xae, 0xec, 0x42, 0x06, 0x25, 0x6b, 0x62, 0x08,
	0xb6, 0x49, 0x49, 0xd3, 0xc2, 0x9e, 0xc2, 0x77, 0xf0, 0xd2, 0x77, 0xf0, 0x05, 0x7c, 0x2a, 0x69,
	0xd6, 0xcc, 0xa1, 0x17, 0xde, 0xf5, 0x9c, 0xf3, 0xff, 0x7f, 0xbf, 0x73, 0x02, 0x82, 0x8c, 0x13,
	0xaa, 0x72, 0x2c, 0x02, 0xad, 0x30, 0xa1, 0xaa, 0xd4, 0x38, 0x7d, 0x0d, 0x4a, 0xad, 0xb0, 0xa6,
	0x6c, 0x1f, 0xd4, 0xe1, 0xf1, 0x1b, 0x15, 0x4a, 0x6a, 0x09, 0x7d, 0x6b, 0x40, 0x27, 0x06, 0x74,
	0x14, 0xd5, 0xe1, 0xf8, 0x96, 0x49, 0xc9, 0x32, 0x1a, 0x18, 0xfd, 0xae, 0x7a, 0x09, 0x48, 0xa5,
	0xb0, 0xe6, 0x52, 0x1c, 0x12, 0xc6, 0x77, 0xbf, 0xe7, 0x9a, 0xe7, 0xb4, 0xd4, 0x38, 0x2f, 0x0e,
	0x82, 0xc9, 0xa7, 0x03, 0xfa, 0xeb, 0x36, 0x10, 0x5e, 0x02, 0x97, 0x13, 0xcf, 0xf1, 0x9d, 0xe9,
	0x20, 0x72, 0x39, 0x81, 0x23, 0xd0, 0x37, 0xbf, 0x4c, 0x38, 0xf1, 0x5c, 0xd3, 0xed, 0x99, 0x7a,
	0x41, 0xe0, 0x0d, 0x00, 0x38, 0x4d, 0x65, 0x25, 0x74, 0x33, 0xec, 0x98, 0xe1, 0xa0, 0xed, 0x2c,
	0x08, 0x7c, 0x02, 0xd7, 0xaa, 0x12, 0x09, 0x17, 0x9a, 0xaa, 0x1a, 0x67, 0x89, 0xc5, 0xf2, 0xba,
	0xbe, 0x33, 0x1d, 0xce, 0x46, 0xe8, 0xc0, 0x85, 0x2c, 0x17, 0x7a, 0x6c, 0x05, 0xd1, 0x95, 0xaa,
	0xc4, 0xa2, 0xb5, 0xd9, 0x26, 0xf4, 0x40, 0x8f, 0x0a, 0xbc, 0xcb, 0x28, 0xf1, 0xce, 0x7c, 0x67,
	0xda, 0x8f, 0x6c, 0x39, 0x51, 0x60, 0x68, 0xf1, 0x97, 0x92, 0x35, 0xc4, 0x99, 0x64, 0x89, 0xde,
	0x17, 0xb4, 0xdd, 0xa3, 0x97, 0x49, 0x16, 0xef, 0x0b, 0xda, 0x64, 0xe4, 0xb4, 0x2c, 0x31, 0xa3,
	0x76, 0x97, 0xb6, 0x84, 0x08, 0x74, 0x9b, 0xb3, 0x98, 0x2d, 0x86, 0xb3, 0xf1, 0x1f, 0xb6, 0xd8,
	0xde, 0x2c, 0x32, 0xba, 0xf9, 0x9b, 0x03, 0xee, 0x53, 0x99, 0xa3, 0xff, 0x5e, 0x67, 0x7e, 0x61,
	0xd1, 0x56, 0x4d, 0xd4, 0xca, 0x79, 0x06, 0x76, 0x5a, 0x87, 0xef, 0x6e, 0x67, 0x19, 0xaf, 0x3f,
	0x5c, 0x7f, 0x69, 0x53, 0xe2, 0x93, 0x14, 0x6b, 0x44, 0x9b, 0xf0, 0xeb, 0x47, 0xb2, 0x3d, 0x91,
	0x6c, 0xad, 0x64, 0xbb, 0x09, 0x77, 0xe7, 0x86, 0xf5, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x8a, 0xa9, 0x6b, 0x62, 0x02, 0x00, 0x00,
}
