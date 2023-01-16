// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stoporders.proto

package investapi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//Направление сделки стоп-заявки.
type StopOrderDirection int32

const (
	StopOrderDirection_STOP_ORDER_DIRECTION_UNSPECIFIED StopOrderDirection = 0
	StopOrderDirection_STOP_ORDER_DIRECTION_BUY         StopOrderDirection = 1
	StopOrderDirection_STOP_ORDER_DIRECTION_SELL        StopOrderDirection = 2
)

var StopOrderDirection_name = map[int32]string{
	0: "STOP_ORDER_DIRECTION_UNSPECIFIED",
	1: "STOP_ORDER_DIRECTION_BUY",
	2: "STOP_ORDER_DIRECTION_SELL",
}

var StopOrderDirection_value = map[string]int32{
	"STOP_ORDER_DIRECTION_UNSPECIFIED": 0,
	"STOP_ORDER_DIRECTION_BUY":         1,
	"STOP_ORDER_DIRECTION_SELL":        2,
}

func (x StopOrderDirection) String() string {
	return proto.EnumName(StopOrderDirection_name, int32(x))
}

func (StopOrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{0}
}

//Тип экспирации стоп-заявке.
type StopOrderExpirationType int32

const (
	StopOrderExpirationType_STOP_ORDER_EXPIRATION_TYPE_UNSPECIFIED      StopOrderExpirationType = 0
	StopOrderExpirationType_STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_CANCEL StopOrderExpirationType = 1
	StopOrderExpirationType_STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_DATE   StopOrderExpirationType = 2
)

var StopOrderExpirationType_name = map[int32]string{
	0: "STOP_ORDER_EXPIRATION_TYPE_UNSPECIFIED",
	1: "STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_CANCEL",
	2: "STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_DATE",
}

var StopOrderExpirationType_value = map[string]int32{
	"STOP_ORDER_EXPIRATION_TYPE_UNSPECIFIED":      0,
	"STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_CANCEL": 1,
	"STOP_ORDER_EXPIRATION_TYPE_GOOD_TILL_DATE":   2,
}

func (x StopOrderExpirationType) String() string {
	return proto.EnumName(StopOrderExpirationType_name, int32(x))
}

func (StopOrderExpirationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{1}
}

//Тип стоп-заявки.
type StopOrderType int32

const (
	StopOrderType_STOP_ORDER_TYPE_UNSPECIFIED StopOrderType = 0
	StopOrderType_STOP_ORDER_TYPE_TAKE_PROFIT StopOrderType = 1
	StopOrderType_STOP_ORDER_TYPE_STOP_LOSS   StopOrderType = 2
	StopOrderType_STOP_ORDER_TYPE_STOP_LIMIT  StopOrderType = 3
)

var StopOrderType_name = map[int32]string{
	0: "STOP_ORDER_TYPE_UNSPECIFIED",
	1: "STOP_ORDER_TYPE_TAKE_PROFIT",
	2: "STOP_ORDER_TYPE_STOP_LOSS",
	3: "STOP_ORDER_TYPE_STOP_LIMIT",
}

var StopOrderType_value = map[string]int32{
	"STOP_ORDER_TYPE_UNSPECIFIED": 0,
	"STOP_ORDER_TYPE_TAKE_PROFIT": 1,
	"STOP_ORDER_TYPE_STOP_LOSS":   2,
	"STOP_ORDER_TYPE_STOP_LIMIT":  3,
}

func (x StopOrderType) String() string {
	return proto.EnumName(StopOrderType_name, int32(x))
}

func (StopOrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{2}
}

//Запрос выставления стоп-заявки.
type PostStopOrderRequest struct {
	Figi                 string                  `protobuf:"bytes,1,opt,name=figi,proto3" json:"figi,omitempty"`
	Quantity             int64                   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                *Quotation              `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	StopPrice            *Quotation              `protobuf:"bytes,4,opt,name=stop_price,json=stopPrice,proto3" json:"stop_price,omitempty"`
	Direction            StopOrderDirection      `protobuf:"varint,5,opt,name=direction,proto3,enum=tinkoff.public.invest.api.contract.v1.StopOrderDirection" json:"direction,omitempty"`
	AccountId            string                  `protobuf:"bytes,6,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ExpirationType       StopOrderExpirationType `protobuf:"varint,7,opt,name=expiration_type,json=expirationType,proto3,enum=tinkoff.public.invest.api.contract.v1.StopOrderExpirationType" json:"expiration_type,omitempty"`
	StopOrderType        StopOrderType           `protobuf:"varint,8,opt,name=stop_order_type,json=stopOrderType,proto3,enum=tinkoff.public.invest.api.contract.v1.StopOrderType" json:"stop_order_type,omitempty"`
	ExpireDate           *timestamp.Timestamp    `protobuf:"bytes,9,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PostStopOrderRequest) Reset()         { *m = PostStopOrderRequest{} }
func (m *PostStopOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PostStopOrderRequest) ProtoMessage()    {}
func (*PostStopOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{0}
}

func (m *PostStopOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostStopOrderRequest.Unmarshal(m, b)
}
func (m *PostStopOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostStopOrderRequest.Marshal(b, m, deterministic)
}
func (m *PostStopOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostStopOrderRequest.Merge(m, src)
}
func (m *PostStopOrderRequest) XXX_Size() int {
	return xxx_messageInfo_PostStopOrderRequest.Size(m)
}
func (m *PostStopOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostStopOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostStopOrderRequest proto.InternalMessageInfo

func (m *PostStopOrderRequest) GetFigi() string {
	if m != nil {
		return m.Figi
	}
	return ""
}

func (m *PostStopOrderRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *PostStopOrderRequest) GetPrice() *Quotation {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *PostStopOrderRequest) GetStopPrice() *Quotation {
	if m != nil {
		return m.StopPrice
	}
	return nil
}

func (m *PostStopOrderRequest) GetDirection() StopOrderDirection {
	if m != nil {
		return m.Direction
	}
	return StopOrderDirection_STOP_ORDER_DIRECTION_UNSPECIFIED
}

func (m *PostStopOrderRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PostStopOrderRequest) GetExpirationType() StopOrderExpirationType {
	if m != nil {
		return m.ExpirationType
	}
	return StopOrderExpirationType_STOP_ORDER_EXPIRATION_TYPE_UNSPECIFIED
}

func (m *PostStopOrderRequest) GetStopOrderType() StopOrderType {
	if m != nil {
		return m.StopOrderType
	}
	return StopOrderType_STOP_ORDER_TYPE_UNSPECIFIED
}

func (m *PostStopOrderRequest) GetExpireDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireDate
	}
	return nil
}

//Результат выставления стоп-заявки.
type PostStopOrderResponse struct {
	StopOrderId          string   `protobuf:"bytes,1,opt,name=stop_order_id,json=stopOrderId,proto3" json:"stop_order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostStopOrderResponse) Reset()         { *m = PostStopOrderResponse{} }
func (m *PostStopOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PostStopOrderResponse) ProtoMessage()    {}
func (*PostStopOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{1}
}

func (m *PostStopOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostStopOrderResponse.Unmarshal(m, b)
}
func (m *PostStopOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostStopOrderResponse.Marshal(b, m, deterministic)
}
func (m *PostStopOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostStopOrderResponse.Merge(m, src)
}
func (m *PostStopOrderResponse) XXX_Size() int {
	return xxx_messageInfo_PostStopOrderResponse.Size(m)
}
func (m *PostStopOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostStopOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostStopOrderResponse proto.InternalMessageInfo

func (m *PostStopOrderResponse) GetStopOrderId() string {
	if m != nil {
		return m.StopOrderId
	}
	return ""
}

//Запрос получения списка активных стоп-заявок.
type GetStopOrdersRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStopOrdersRequest) Reset()         { *m = GetStopOrdersRequest{} }
func (m *GetStopOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*GetStopOrdersRequest) ProtoMessage()    {}
func (*GetStopOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{2}
}

func (m *GetStopOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStopOrdersRequest.Unmarshal(m, b)
}
func (m *GetStopOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStopOrdersRequest.Marshal(b, m, deterministic)
}
func (m *GetStopOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStopOrdersRequest.Merge(m, src)
}
func (m *GetStopOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_GetStopOrdersRequest.Size(m)
}
func (m *GetStopOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStopOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStopOrdersRequest proto.InternalMessageInfo

func (m *GetStopOrdersRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

//Список активных стоп-заявок.
type GetStopOrdersResponse struct {
	StopOrders           []*StopOrder `protobuf:"bytes,1,rep,name=stop_orders,json=stopOrders,proto3" json:"stop_orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetStopOrdersResponse) Reset()         { *m = GetStopOrdersResponse{} }
func (m *GetStopOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*GetStopOrdersResponse) ProtoMessage()    {}
func (*GetStopOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{3}
}

func (m *GetStopOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStopOrdersResponse.Unmarshal(m, b)
}
func (m *GetStopOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStopOrdersResponse.Marshal(b, m, deterministic)
}
func (m *GetStopOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStopOrdersResponse.Merge(m, src)
}
func (m *GetStopOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_GetStopOrdersResponse.Size(m)
}
func (m *GetStopOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStopOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStopOrdersResponse proto.InternalMessageInfo

func (m *GetStopOrdersResponse) GetStopOrders() []*StopOrder {
	if m != nil {
		return m.StopOrders
	}
	return nil
}

//Запрос отмены выставленной стоп-заявки.
type CancelStopOrderRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	StopOrderId          string   `protobuf:"bytes,2,opt,name=stop_order_id,json=stopOrderId,proto3" json:"stop_order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelStopOrderRequest) Reset()         { *m = CancelStopOrderRequest{} }
func (m *CancelStopOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CancelStopOrderRequest) ProtoMessage()    {}
func (*CancelStopOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{4}
}

func (m *CancelStopOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelStopOrderRequest.Unmarshal(m, b)
}
func (m *CancelStopOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelStopOrderRequest.Marshal(b, m, deterministic)
}
func (m *CancelStopOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelStopOrderRequest.Merge(m, src)
}
func (m *CancelStopOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CancelStopOrderRequest.Size(m)
}
func (m *CancelStopOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelStopOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelStopOrderRequest proto.InternalMessageInfo

func (m *CancelStopOrderRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *CancelStopOrderRequest) GetStopOrderId() string {
	if m != nil {
		return m.StopOrderId
	}
	return ""
}

//Результат отмены выставленной стоп-заявки.
type CancelStopOrderResponse struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CancelStopOrderResponse) Reset()         { *m = CancelStopOrderResponse{} }
func (m *CancelStopOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CancelStopOrderResponse) ProtoMessage()    {}
func (*CancelStopOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{5}
}

func (m *CancelStopOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelStopOrderResponse.Unmarshal(m, b)
}
func (m *CancelStopOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelStopOrderResponse.Marshal(b, m, deterministic)
}
func (m *CancelStopOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelStopOrderResponse.Merge(m, src)
}
func (m *CancelStopOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CancelStopOrderResponse.Size(m)
}
func (m *CancelStopOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelStopOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelStopOrderResponse proto.InternalMessageInfo

func (m *CancelStopOrderResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

//Информация о стоп-заявке.
type StopOrder struct {
	StopOrderId          string               `protobuf:"bytes,1,opt,name=stop_order_id,json=stopOrderId,proto3" json:"stop_order_id,omitempty"`
	LotsRequested        int64                `protobuf:"varint,2,opt,name=lots_requested,json=lotsRequested,proto3" json:"lots_requested,omitempty"`
	Figi                 string               `protobuf:"bytes,3,opt,name=figi,proto3" json:"figi,omitempty"`
	Direction            StopOrderDirection   `protobuf:"varint,4,opt,name=direction,proto3,enum=tinkoff.public.invest.api.contract.v1.StopOrderDirection" json:"direction,omitempty"`
	Currency             string               `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	OrderType            StopOrderType        `protobuf:"varint,6,opt,name=order_type,json=orderType,proto3,enum=tinkoff.public.invest.api.contract.v1.StopOrderType" json:"order_type,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	ActivationDateTime   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=activation_date_time,json=activationDateTime,proto3" json:"activation_date_time,omitempty"`
	ExpirationTime       *timestamp.Timestamp `protobuf:"bytes,9,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	Price                *MoneyValue          `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	StopPrice            *MoneyValue          `protobuf:"bytes,11,opt,name=stop_price,json=stopPrice,proto3" json:"stop_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StopOrder) Reset()         { *m = StopOrder{} }
func (m *StopOrder) String() string { return proto.CompactTextString(m) }
func (*StopOrder) ProtoMessage()    {}
func (*StopOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_66aa9cbcf0b3ec31, []int{6}
}

func (m *StopOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopOrder.Unmarshal(m, b)
}
func (m *StopOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopOrder.Marshal(b, m, deterministic)
}
func (m *StopOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopOrder.Merge(m, src)
}
func (m *StopOrder) XXX_Size() int {
	return xxx_messageInfo_StopOrder.Size(m)
}
func (m *StopOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_StopOrder.DiscardUnknown(m)
}

var xxx_messageInfo_StopOrder proto.InternalMessageInfo

func (m *StopOrder) GetStopOrderId() string {
	if m != nil {
		return m.StopOrderId
	}
	return ""
}

func (m *StopOrder) GetLotsRequested() int64 {
	if m != nil {
		return m.LotsRequested
	}
	return 0
}

func (m *StopOrder) GetFigi() string {
	if m != nil {
		return m.Figi
	}
	return ""
}

func (m *StopOrder) GetDirection() StopOrderDirection {
	if m != nil {
		return m.Direction
	}
	return StopOrderDirection_STOP_ORDER_DIRECTION_UNSPECIFIED
}

func (m *StopOrder) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *StopOrder) GetOrderType() StopOrderType {
	if m != nil {
		return m.OrderType
	}
	return StopOrderType_STOP_ORDER_TYPE_UNSPECIFIED
}

func (m *StopOrder) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *StopOrder) GetActivationDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.ActivationDateTime
	}
	return nil
}

func (m *StopOrder) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *StopOrder) GetPrice() *MoneyValue {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *StopOrder) GetStopPrice() *MoneyValue {
	if m != nil {
		return m.StopPrice
	}
	return nil
}

func init() {
	proto.RegisterEnum("tinkoff.public.invest.api.contract.v1.StopOrderDirection", StopOrderDirection_name, StopOrderDirection_value)
	proto.RegisterEnum("tinkoff.public.invest.api.contract.v1.StopOrderExpirationType", StopOrderExpirationType_name, StopOrderExpirationType_value)
	proto.RegisterEnum("tinkoff.public.invest.api.contract.v1.StopOrderType", StopOrderType_name, StopOrderType_value)
	proto.RegisterType((*PostStopOrderRequest)(nil), "tinkoff.public.invest.api.contract.v1.PostStopOrderRequest")
	proto.RegisterType((*PostStopOrderResponse)(nil), "tinkoff.public.invest.api.contract.v1.PostStopOrderResponse")
	proto.RegisterType((*GetStopOrdersRequest)(nil), "tinkoff.public.invest.api.contract.v1.GetStopOrdersRequest")
	proto.RegisterType((*GetStopOrdersResponse)(nil), "tinkoff.public.invest.api.contract.v1.GetStopOrdersResponse")
	proto.RegisterType((*CancelStopOrderRequest)(nil), "tinkoff.public.invest.api.contract.v1.CancelStopOrderRequest")
	proto.RegisterType((*CancelStopOrderResponse)(nil), "tinkoff.public.invest.api.contract.v1.CancelStopOrderResponse")
	proto.RegisterType((*StopOrder)(nil), "tinkoff.public.invest.api.contract.v1.StopOrder")
}

func init() {
	proto.RegisterFile("stoporders.proto", fileDescriptor_66aa9cbcf0b3ec31)
}

var fileDescriptor_66aa9cbcf0b3ec31 = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x49, 0xda, 0x6d, 0x4e, 0x36, 0x6d, 0x76, 0xd4, 0x65, 0x4d, 0xd8, 0x65, 0xa3, 0x88,
	0x45, 0xa5, 0x08, 0x97, 0x16, 0xb8, 0x40, 0x85, 0x95, 0xb2, 0x89, 0x5b, 0x59, 0x64, 0x6b, 0xaf,
	0xe3, 0x2d, 0x2c, 0x54, 0xb2, 0x5c, 0x7b, 0x5a, 0x0d, 0xa4, 0x1e, 0xaf, 0x3d, 0x2e, 0xe4, 0x15,
	0xb8, 0xda, 0x37, 0x40, 0x42, 0xe2, 0x86, 0xa7, 0xe0, 0x9a, 0x57, 0xe0, 0x65, 0x90, 0xc7, 0x8e,
	0x7f, 0x12, 0xd3, 0xba, 0x85, 0xbb, 0xcc, 0xf8, 0x7c, 0xdf, 0x39, 0x73, 0xe6, 0x3b, 0xdf, 0x04,
	0x3a, 0x01, 0xa3, 0x1e, 0xf5, 0x1d, 0xec, 0x07, 0x92, 0xe7, 0x53, 0x46, 0xd1, 0x13, 0x46, 0xdc,
	0x1f, 0xe9, 0xd9, 0x99, 0xe4, 0x85, 0xa7, 0x53, 0x62, 0x4b, 0xc4, 0xbd, 0xc4, 0x01, 0x93, 0x2c,
	0x8f, 0x48, 0x36, 0x75, 0x99, 0x6f, 0xd9, 0x4c, 0xba, 0xdc, 0xed, 0x3e, 0x3e, 0xa7, 0xf4, 0x7c,
	0x8a, 0x77, 0x38, 0xe8, 0x34, 0x3c, 0xdb, 0x61, 0xe4, 0x02, 0x07, 0xcc, 0xba, 0xf0, 0x62, 0x9e,
	0xee, 0x5d, 0x9b, 0x5e, 0x5c, 0x50, 0x37, 0x5e, 0xf5, 0xff, 0x6e, 0xc0, 0xa6, 0x46, 0x03, 0x36,
	0x61, 0xd4, 0x53, 0xa3, 0x74, 0x3a, 0x7e, 0x1d, 0xe2, 0x80, 0x21, 0x04, 0x8d, 0x33, 0x72, 0x4e,
	0x44, 0xa1, 0x27, 0x6c, 0x35, 0x75, 0xfe, 0x1b, 0x75, 0x61, 0xed, 0x75, 0x68, 0xb9, 0x8c, 0xb0,
	0x99, 0x58, 0xeb, 0x09, 0x5b, 0x75, 0x3d, 0x5d, 0xa3, 0x03, 0x58, 0xf1, 0x7c, 0x62, 0x63, 0xb1,
	0xde, 0x13, 0xb6, 0x5a, 0x7b, 0x9f, 0x48, 0x95, 0xca, 0x95, 0x5e, 0x84, 0x94, 0x59, 0x8c, 0x50,
	0x57, 0x8f, 0xe1, 0x48, 0x05, 0x88, 0x8e, 0x6e, 0xc6, 0x64, 0x8d, 0x5b, 0x92, 0x35, 0x23, 0x0e,
	0x8d, 0x13, 0x7e, 0x03, 0x4d, 0x87, 0xf8, 0xd8, 0x8e, 0xf6, 0xc5, 0x95, 0x9e, 0xb0, 0xb5, 0xbe,
	0xf7, 0x45, 0x45, 0xbe, 0xb4, 0x29, 0xa3, 0x39, 0x81, 0x9e, 0x71, 0xa1, 0x47, 0x00, 0x96, 0x6d,
	0xd3, 0xd0, 0x65, 0x26, 0x71, 0xc4, 0x55, 0xde, 0xa7, 0x66, 0xb2, 0xa3, 0x38, 0xe8, 0x1c, 0x36,
	0xf0, 0xcf, 0x1e, 0xf1, 0x79, 0x41, 0x26, 0x9b, 0x79, 0x58, 0xbc, 0xc3, 0xb3, 0x3f, 0xbd, 0x69,
	0x76, 0x39, 0xa5, 0x31, 0x66, 0x1e, 0xd6, 0xd7, 0x71, 0x61, 0x8d, 0x4e, 0x60, 0x83, 0x77, 0x8c,
	0xab, 0x25, 0x4e, 0xb4, 0xc6, 0x13, 0x7d, 0x76, 0xd3, 0x44, 0x9c, 0xbe, 0x1d, 0xe4, 0x97, 0x68,
	0x1f, 0x5a, 0x3c, 0x1f, 0x36, 0x1d, 0x8b, 0x61, 0xb1, 0xc9, 0x2f, 0xa4, 0x2b, 0xc5, 0x2a, 0x93,
	0xe6, 0x2a, 0x93, 0x8c, 0xb9, 0xca, 0x74, 0x88, 0xc3, 0x47, 0x16, 0xc3, 0xfd, 0x7d, 0xb8, 0xbf,
	0x20, 0xae, 0xc0, 0xa3, 0x6e, 0x80, 0x51, 0x1f, 0xda, 0xb9, 0x9a, 0x89, 0x93, 0xc8, 0xac, 0x95,
	0xe6, 0x56, 0x9c, 0xfe, 0xe7, 0xb0, 0x79, 0x88, 0x33, 0x6c, 0x30, 0x57, 0x66, 0xb1, 0xef, 0xc2,
	0x42, 0xdf, 0xfb, 0x3f, 0xc0, 0xfd, 0x05, 0x58, 0x92, 0xf3, 0x05, 0xb4, 0xb2, 0x9c, 0x81, 0x28,
	0xf4, 0xea, 0x37, 0x90, 0x56, 0x76, 0x04, 0x48, 0x6b, 0x0c, 0xfa, 0xdf, 0xc3, 0xdb, 0x43, 0xcb,
	0xb5, 0xf1, 0x74, 0x69, 0x7c, 0xae, 0x2e, 0x72, 0xf9, 0xfc, 0xb5, 0xe5, 0xf3, 0x2b, 0xf0, 0x60,
	0x89, 0x3c, 0x39, 0x8a, 0x04, 0x8d, 0x68, 0xac, 0x39, 0xef, 0xd5, 0xb7, 0xc1, 0xe3, 0xfa, 0xbf,
	0xae, 0x40, 0x33, 0x65, 0xa9, 0xd2, 0x7c, 0xf4, 0x04, 0xd6, 0xa7, 0x94, 0x05, 0xa6, 0x1f, 0x9f,
	0x07, 0x3b, 0xc9, 0xc0, 0xb7, 0xa3, 0x5d, 0x7d, 0xbe, 0x99, 0xba, 0x44, 0x3d, 0xe7, 0x12, 0x85,
	0x81, 0x6b, 0xfc, 0x8f, 0x03, 0xd7, 0x85, 0x35, 0x3b, 0xf4, 0x7d, 0xec, 0xda, 0x33, 0x3e, 0xc8,
	0x4d, 0x3d, 0x5d, 0xa3, 0x09, 0x40, 0x4e, 0xff, 0xab, 0xff, 0x41, 0xff, 0x4d, 0x9a, 0xd7, 0xbe,
	0xed, 0x63, 0x8b, 0x25, 0xda, 0xbf, 0x73, 0xbd, 0xf6, 0xe3, 0xf0, 0x48, 0xfb, 0x68, 0x0c, 0x9b,
	0x96, 0xcd, 0xc8, 0x65, 0x3c, 0xff, 0x11, 0x81, 0xc9, 0xef, 0x6c, 0xed, 0x5a, 0x16, 0x94, 0xe1,
	0x22, 0xa6, 0xe8, 0x03, 0x1a, 0x16, 0xdd, 0x24, 0x22, 0xba, 0x7e, 0x14, 0xf3, 0x4e, 0x11, 0x91,
	0x1c, 0xce, 0x3d, 0x1a, 0x38, 0x74, 0xb7, 0x62, 0x7f, 0x9e, 0x53, 0x17, 0xcf, 0x8e, 0xad, 0x69,
	0x88, 0xe7, 0x26, 0xad, 0x15, 0x4c, 0xba, 0x75, 0x5b, 0xb6, 0xcc, 0xa5, 0xb7, 0x7f, 0x02, 0xb4,
	0x7c, 0xf9, 0xe8, 0x7d, 0xe8, 0x4d, 0x0c, 0x55, 0x33, 0x55, 0x7d, 0x24, 0xeb, 0xe6, 0x48, 0xd1,
	0xe5, 0xa1, 0xa1, 0xa8, 0x47, 0xe6, 0xcb, 0xa3, 0x89, 0x26, 0x0f, 0x95, 0x03, 0x45, 0x1e, 0x75,
	0xde, 0x42, 0x0f, 0x41, 0x2c, 0x8d, 0x7a, 0xf6, 0xf2, 0x55, 0x47, 0x40, 0x8f, 0xe0, 0x9d, 0xd2,
	0xaf, 0x13, 0x79, 0x3c, 0xee, 0xd4, 0xb6, 0x7f, 0x17, 0xe0, 0xc1, 0xbf, 0x38, 0x2d, 0xda, 0x86,
	0x0f, 0x72, 0x50, 0xf9, 0x5b, 0x4d, 0xd1, 0x07, 0x1c, 0x6b, 0xbc, 0xd2, 0xe4, 0x85, 0x22, 0x76,
	0xe0, 0xa3, 0x2b, 0x62, 0x0f, 0x55, 0x75, 0x64, 0x1a, 0xca, 0x78, 0x6c, 0x0e, 0x07, 0x47, 0x43,
	0x79, 0xdc, 0x11, 0xd0, 0xc7, 0xf0, 0x61, 0x25, 0xc0, 0x68, 0x60, 0xc8, 0x9d, 0xda, 0xf6, 0x1b,
	0x01, 0xda, 0x05, 0xa1, 0xa2, 0xc7, 0xf0, 0x6e, 0x8e, 0xa0, 0xa4, 0xa4, 0x92, 0x00, 0x63, 0xf0,
	0xb5, 0x6c, 0x6a, 0xba, 0x7a, 0xa0, 0x18, 0x4b, 0xad, 0xe1, 0x01, 0x7c, 0x3d, 0x56, 0x27, 0x93,
	0x4e, 0x0d, 0xbd, 0x07, 0xdd, 0xf2, 0xcf, 0xca, 0x73, 0xc5, 0xe8, 0xd4, 0xf7, 0xfe, 0xac, 0xc3,
	0xbd, 0xcc, 0x67, 0x27, 0xd8, 0xbf, 0x8c, 0xb4, 0xf1, 0x8b, 0x00, 0xed, 0x82, 0xe9, 0xa3, 0xfd,
	0x8a, 0xca, 0x28, 0xfb, 0x1f, 0xd2, 0xfd, 0xf2, 0x76, 0xe0, 0xc4, 0x28, 0xa3, 0x62, 0x0a, 0xaf,
	0x41, 0xe5, 0x62, 0xca, 0x9e, 0x9e, 0xca, 0xc5, 0x94, 0x3f, 0x40, 0x6f, 0x04, 0xd8, 0x58, 0x70,
	0x74, 0xf4, 0x55, 0x45, 0xc6, 0xf2, 0x67, 0xa6, 0xfb, 0xf4, 0xb6, 0xf0, 0xb8, 0xa4, 0x67, 0x16,
	0x3c, 0xf4, 0xc3, 0x8c, 0x83, 0x2c, 0xe0, 0x34, 0xe1, 0xbb, 0xbb, 0xd2, 0xce, 0x7e, 0x4c, 0x6a,
	0x79, 0xe4, 0xb7, 0xda, 0x8a, 0xa1, 0x0c, 0x34, 0xe5, 0x8f, 0xda, 0xa6, 0x91, 0x60, 0x14, 0xfe,
	0x6d, 0xe0, 0x11, 0xe9, 0x78, 0xf7, 0xaf, 0xda, 0xbd, 0x64, 0xfb, 0x24, 0xde, 0x3e, 0x39, 0xde,
	0x3d, 0x5d, 0xe5, 0xc6, 0xf4, 0xe9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x44, 0xbd, 0x37,
	0xd2, 0x0a, 0x00, 0x00,
}