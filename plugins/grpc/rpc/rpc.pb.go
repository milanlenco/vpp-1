// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import interfaces1 "github.com/ligato/vpp-agent/plugins/linuxv2/model/interfaces"
import l31 "github.com/ligato/vpp-agent/plugins/linuxv2/model/l3"
import acl "github.com/ligato/vpp-agent/plugins/vppv2/model/acl"
import interfaces "github.com/ligato/vpp-agent/plugins/vppv2/model/interfaces"
import ipsec "github.com/ligato/vpp-agent/plugins/vppv2/model/ipsec"
import l2 "github.com/ligato/vpp-agent/plugins/vppv2/model/l2"
import l3 "github.com/ligato/vpp-agent/plugins/vppv2/model/l3"
import nat "github.com/ligato/vpp-agent/plugins/vppv2/model/nat"
import punt "github.com/ligato/vpp-agent/plugins/vppv2/model/punt"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Data request is an inventory of supported data types with one or multiple
// items of every type. Universal type for every data change/resync request
type DataRequest struct {
	AccessLists          []*acl.Acl                      `protobuf:"bytes,10,rep,name=AccessLists" json:"AccessLists,omitempty"`
	Interfaces           []*interfaces.Interface         `protobuf:"bytes,20,rep,name=Interfaces" json:"Interfaces,omitempty"`
	BridgeDomains        []*l2.BridgeDomain              `protobuf:"bytes,40,rep,name=BridgeDomains" json:"BridgeDomains,omitempty"`
	FIBs                 []*l2.FIBEntry                  `protobuf:"bytes,41,rep,name=FIBs" json:"FIBs,omitempty"`
	XCons                []*l2.XConnectPair              `protobuf:"bytes,42,rep,name=XCons" json:"XCons,omitempty"`
	StaticRoutes         []*l3.StaticRoute               `protobuf:"bytes,50,rep,name=StaticRoutes" json:"StaticRoutes,omitempty"`
	ArpEntries           []*l3.ARPEntry                  `protobuf:"bytes,51,rep,name=ArpEntries" json:"ArpEntries,omitempty"`
	ProxyArp             *l3.ProxyARP                    `protobuf:"bytes,52,opt,name=ProxyArp" json:"ProxyArp,omitempty"`
	IPScanNeighbor       *l3.IPScanNeighbor              `protobuf:"bytes,53,opt,name=IPScanNeighbor" json:"IPScanNeighbor,omitempty"`
	SPDs                 []*ipsec.SecurityPolicyDatabase `protobuf:"bytes,60,rep,name=SPDs" json:"SPDs,omitempty"`
	SAs                  []*ipsec.SecurityAssociation    `protobuf:"bytes,61,rep,name=SAs" json:"SAs,omitempty"`
	IPRedirectPunts      []*punt.IpRedirect              `protobuf:"bytes,65,rep,name=IPRedirectPunts" json:"IPRedirectPunts,omitempty"`
	ToHostPunts          []*punt.ToHost                  `protobuf:"bytes,66,rep,name=ToHostPunts" json:"ToHostPunts,omitempty"`
	NatGlobal            *nat.Nat44Global                `protobuf:"bytes,71,opt,name=NatGlobal" json:"NatGlobal,omitempty"`
	DNATs                []*nat.DNat44                   `protobuf:"bytes,72,rep,name=DNATs" json:"DNATs,omitempty"`
	LinuxInterfaces      []*interfaces1.Interface        `protobuf:"bytes,80,rep,name=LinuxInterfaces" json:"LinuxInterfaces,omitempty"`
	LinuxArpEntries      []*l31.StaticARPEntry           `protobuf:"bytes,90,rep,name=LinuxArpEntries" json:"LinuxArpEntries,omitempty"`
	LinuxRoutes          []*l31.StaticRoute              `protobuf:"bytes,91,rep,name=LinuxRoutes" json:"LinuxRoutes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ad67b345f2565a7d, []int{0}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (dst *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(dst, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetAccessLists() []*acl.Acl {
	if m != nil {
		return m.AccessLists
	}
	return nil
}

func (m *DataRequest) GetInterfaces() []*interfaces.Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *DataRequest) GetBridgeDomains() []*l2.BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *DataRequest) GetFIBs() []*l2.FIBEntry {
	if m != nil {
		return m.FIBs
	}
	return nil
}

func (m *DataRequest) GetXCons() []*l2.XConnectPair {
	if m != nil {
		return m.XCons
	}
	return nil
}

func (m *DataRequest) GetStaticRoutes() []*l3.StaticRoute {
	if m != nil {
		return m.StaticRoutes
	}
	return nil
}

func (m *DataRequest) GetArpEntries() []*l3.ARPEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

func (m *DataRequest) GetProxyArp() *l3.ProxyARP {
	if m != nil {
		return m.ProxyArp
	}
	return nil
}

func (m *DataRequest) GetIPScanNeighbor() *l3.IPScanNeighbor {
	if m != nil {
		return m.IPScanNeighbor
	}
	return nil
}

func (m *DataRequest) GetSPDs() []*ipsec.SecurityPolicyDatabase {
	if m != nil {
		return m.SPDs
	}
	return nil
}

func (m *DataRequest) GetSAs() []*ipsec.SecurityAssociation {
	if m != nil {
		return m.SAs
	}
	return nil
}

func (m *DataRequest) GetIPRedirectPunts() []*punt.IpRedirect {
	if m != nil {
		return m.IPRedirectPunts
	}
	return nil
}

func (m *DataRequest) GetToHostPunts() []*punt.ToHost {
	if m != nil {
		return m.ToHostPunts
	}
	return nil
}

func (m *DataRequest) GetNatGlobal() *nat.Nat44Global {
	if m != nil {
		return m.NatGlobal
	}
	return nil
}

func (m *DataRequest) GetDNATs() []*nat.DNat44 {
	if m != nil {
		return m.DNATs
	}
	return nil
}

func (m *DataRequest) GetLinuxInterfaces() []*interfaces1.Interface {
	if m != nil {
		return m.LinuxInterfaces
	}
	return nil
}

func (m *DataRequest) GetLinuxArpEntries() []*l31.StaticARPEntry {
	if m != nil {
		return m.LinuxArpEntries
	}
	return nil
}

func (m *DataRequest) GetLinuxRoutes() []*l31.StaticRoute {
	if m != nil {
		return m.LinuxRoutes
	}
	return nil
}

// Response to data change 'put'
type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ad67b345f2565a7d, []int{1}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (dst *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(dst, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

// Response to data change 'del'
type DelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelResponse) Reset()         { *m = DelResponse{} }
func (m *DelResponse) String() string { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()    {}
func (*DelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ad67b345f2565a7d, []int{2}
}
func (m *DelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelResponse.Unmarshal(m, b)
}
func (m *DelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelResponse.Marshal(b, m, deterministic)
}
func (dst *DelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelResponse.Merge(dst, src)
}
func (m *DelResponse) XXX_Size() int {
	return xxx_messageInfo_DelResponse.Size(m)
}
func (m *DelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelResponse proto.InternalMessageInfo

// Response to data resync
type ResyncResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResyncResponse) Reset()         { *m = ResyncResponse{} }
func (m *ResyncResponse) String() string { return proto.CompactTextString(m) }
func (*ResyncResponse) ProtoMessage()    {}
func (*ResyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ad67b345f2565a7d, []int{3}
}
func (m *ResyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResyncResponse.Unmarshal(m, b)
}
func (m *ResyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResyncResponse.Marshal(b, m, deterministic)
}
func (dst *ResyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResyncResponse.Merge(dst, src)
}
func (m *ResyncResponse) XXX_Size() int {
	return xxx_messageInfo_ResyncResponse.Size(m)
}
func (m *ResyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResyncResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DataRequest)(nil), "rpc.DataRequest")
	proto.RegisterType((*PutResponse)(nil), "rpc.PutResponse")
	proto.RegisterType((*DelResponse)(nil), "rpc.DelResponse")
	proto.RegisterType((*ResyncResponse)(nil), "rpc.ResyncResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataChangeService service

type DataChangeServiceClient interface {
	// Creates or updates one or multiple configuration items
	Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type dataChangeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataChangeServiceClient(cc *grpc.ClientConn) DataChangeServiceClient {
	return &dataChangeServiceClient{cc}
}

func (c *dataChangeServiceClient) Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataChangeService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataChangeServiceClient) Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataChangeService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataChangeService service

type DataChangeServiceServer interface {
	// Creates or updates one or multiple configuration items
	Put(context.Context, *DataRequest) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(context.Context, *DataRequest) (*DelResponse, error)
}

func RegisterDataChangeServiceServer(s *grpc.Server, srv DataChangeServiceServer) {
	s.RegisterService(&_DataChangeService_serviceDesc, srv)
}

func _DataChangeService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Put(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataChangeService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Del(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataChangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataChangeService",
	HandlerType: (*DataChangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _DataChangeService_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _DataChangeService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for DataResyncService service

type DataResyncServiceClient interface {
	// Calls vpp-agent resync
	Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error)
}

type dataResyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataResyncServiceClient(cc *grpc.ClientConn) DataResyncServiceClient {
	return &dataResyncServiceClient{cc}
}

func (c *dataResyncServiceClient) Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error) {
	out := new(ResyncResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataResyncService/Resync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataResyncService service

type DataResyncServiceServer interface {
	// Calls vpp-agent resync
	Resync(context.Context, *DataRequest) (*ResyncResponse, error)
}

func RegisterDataResyncServiceServer(s *grpc.Server, srv DataResyncServiceServer) {
	s.RegisterService(&_DataResyncService_serviceDesc, srv)
}

func _DataResyncService_Resync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataResyncServiceServer).Resync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataResyncService/Resync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataResyncServiceServer).Resync(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataResyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataResyncService",
	HandlerType: (*DataResyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resync",
			Handler:    _DataResyncService_Resync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_ad67b345f2565a7d) }

var fileDescriptor_rpc_ad67b345f2565a7d = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x4f, 0x1b, 0x39,
	0x14, 0xd6, 0x2a, 0xc0, 0x82, 0xc3, 0x25, 0xeb, 0xdd, 0x95, 0x46, 0x59, 0xad, 0x44, 0x79, 0xa8,
	0x00, 0xd1, 0x19, 0x35, 0x03, 0xad, 0x44, 0x41, 0xd5, 0x84, 0xe1, 0x92, 0x8a, 0x46, 0x23, 0x87,
	0x87, 0xaa, 0x7d, 0xa8, 0x1c, 0xc7, 0x04, 0x4b, 0xc6, 0x9e, 0xda, 0x1e, 0x44, 0x7e, 0x4d, 0xff,
	0x6a, 0x65, 0x7b, 0x12, 0x26, 0x11, 0xad, 0x9a, 0x3c, 0x4c, 0x94, 0x39, 0xdf, 0xc5, 0x47, 0xc7,
	0xdf, 0x49, 0xc0, 0x9a, 0xca, 0x49, 0x98, 0x2b, 0x69, 0x24, 0xac, 0xa9, 0x9c, 0x34, 0x4f, 0x87,
	0xcc, 0xdc, 0x15, 0xfd, 0x90, 0xc8, 0xfb, 0x88, 0xb3, 0x21, 0x36, 0x32, 0x7a, 0xc8, 0xf3, 0x57,
	0x78, 0x48, 0x85, 0x89, 0x72, 0x5e, 0x0c, 0x99, 0xd0, 0xb6, 0xf2, 0xd0, 0x8a, 0xee, 0xe5, 0x80,
	0xf2, 0x08, 0x13, 0xf7, 0x78, 0x8f, 0xe6, 0x87, 0x79, 0xe5, 0x4c, 0x18, 0xaa, 0x6e, 0x31, 0xa1,
	0xfa, 0xe9, 0x6b, 0xe9, 0xf5, 0x6e, 0x5e, 0x2f, 0xde, 0x8a, 0xfa, 0x83, 0x52, 0x7c, 0xb2, 0x80,
	0xf8, 0x96, 0xf5, 0x4b, 0x75, 0xb2, 0x80, 0xfa, 0x91, 0x48, 0x21, 0x28, 0x31, 0x0b, 0x37, 0x10,
	0x47, 0x58, 0xe5, 0xa5, 0xfa, 0x7c, 0x01, 0xb5, 0x36, 0xd8, 0x30, 0xf2, 0x55, 0xc9, 0xc2, 0x8c,
	0x47, 0x38, 0xf7, 0x6d, 0x0a, 0x6c, 0xec, 0x53, 0xca, 0xdf, 0xcf, 0x2b, 0xcf, 0x0b, 0x0b, 0x14,
	0xc2, 0x2c, 0x3a, 0x47, 0x96, 0x6b, 0x4a, 0xfc, 0x67, 0x69, 0xf1, 0xf1, 0x77, 0x2c, 0x38, 0x13,
	0xc5, 0xe3, 0x2f, 0x33, 0xa5, 0xe7, 0x99, 0xc8, 0xb4, 0x1d, 0x8f, 0x23, 0x1e, 0x7b, 0xf9, 0xce,
	0xf7, 0x3f, 0x41, 0x3d, 0xc5, 0x06, 0x23, 0xfa, 0xad, 0xa0, 0xda, 0xc0, 0x7d, 0x50, 0x4f, 0x08,
	0xa1, 0x5a, 0x5f, 0x33, 0x6d, 0x74, 0x00, 0xb6, 0x6b, 0xbb, 0xf5, 0xd6, 0x6a, 0x68, 0x17, 0x22,
	0x21, 0x1c, 0x55, 0x41, 0x78, 0x04, 0x40, 0x67, 0xd2, 0x4e, 0xf0, 0x8f, 0xa3, 0xfe, 0x1b, 0x56,
	0x3a, 0x9c, 0xa0, 0xa8, 0x42, 0x84, 0x6f, 0xc0, 0x46, 0x5b, 0xb1, 0xc1, 0x90, 0xa6, 0xf2, 0x1e,
	0x33, 0xa1, 0x83, 0x5d, 0xa7, 0x6c, 0x84, 0xbc, 0x15, 0x56, 0x01, 0x34, 0x4d, 0x83, 0xdb, 0x60,
	0xe9, 0xa2, 0xd3, 0xd6, 0xc1, 0x9e, 0xa3, 0xaf, 0x5b, 0xfa, 0x45, 0xa7, 0x7d, 0x2e, 0x8c, 0x1a,
	0x21, 0x87, 0xc0, 0x97, 0x60, 0xf9, 0xd3, 0x99, 0x14, 0x3a, 0xd8, 0x7f, 0x72, 0xb4, 0x05, 0x9b,
	0xe2, 0x0c, 0x33, 0x85, 0x3c, 0x0c, 0x63, 0xb0, 0xde, 0x73, 0xd9, 0x42, 0x36, 0x5a, 0x3a, 0x68,
	0x39, 0xfa, 0x56, 0xc8, 0xe3, 0xb0, 0x52, 0x47, 0x53, 0x24, 0x78, 0x00, 0x40, 0xa2, 0x72, 0x7b,
	0x1c, 0xa3, 0x3a, 0x88, 0xc7, 0x4d, 0xc4, 0x61, 0x82, 0x32, 0xdf, 0x44, 0x05, 0x87, 0xbb, 0x60,
	0x35, 0x53, 0xf2, 0x71, 0x94, 0xa8, 0x3c, 0x38, 0xdc, 0xfe, 0x63, 0xcc, 0xf5, 0x35, 0x94, 0xa1,
	0x09, 0x0a, 0x8f, 0xc1, 0x66, 0x27, 0xeb, 0x11, 0x2c, 0xba, 0x94, 0x0d, 0xef, 0xfa, 0x52, 0x05,
	0x47, 0x8e, 0x0f, 0x2d, 0x7f, 0x1a, 0x41, 0x33, 0x4c, 0xf8, 0x1a, 0x2c, 0xf5, 0xb2, 0x54, 0x07,
	0x27, 0xae, 0x9b, 0xff, 0x43, 0x9f, 0xb3, 0x1e, 0x25, 0x85, 0x62, 0x66, 0x94, 0x49, 0xce, 0xc8,
	0xc8, 0xde, 0x6e, 0x1f, 0x6b, 0x8a, 0x1c, 0x15, 0x1e, 0x80, 0x5a, 0x2f, 0xd1, 0xc1, 0xa9, 0x53,
	0x34, 0x67, 0x14, 0x89, 0xd6, 0x92, 0x30, 0x6c, 0x98, 0x14, 0xc8, 0xd2, 0xe0, 0x31, 0xd8, 0xea,
	0x64, 0x88, 0x0e, 0x98, 0xb2, 0x23, 0x2c, 0x84, 0xd1, 0x41, 0x52, 0xce, 0xd6, 0x6d, 0x45, 0x27,
	0x1f, 0x83, 0x68, 0x96, 0x08, 0x43, 0x50, 0xbf, 0x91, 0x57, 0x52, 0x97, 0xba, 0x76, 0x39, 0x31,
	0xa7, 0xf3, 0x00, 0xaa, 0x12, 0x60, 0x08, 0xd6, 0xba, 0xd8, 0x5c, 0x72, 0xd9, 0xc7, 0x3c, 0xb8,
	0x74, 0x33, 0x68, 0x84, 0x76, 0x77, 0xbb, 0xd8, 0x1c, 0x1e, 0xfa, 0x3a, 0x7a, 0xa2, 0xc0, 0x17,
	0x60, 0x39, 0xed, 0x26, 0x37, 0x3a, 0xb8, 0x72, 0xce, 0x75, 0xc7, 0x4d, 0x1d, 0x19, 0x79, 0x04,
	0x9e, 0x83, 0xad, 0x6b, 0x1b, 0xfd, 0x4a, 0x4c, 0x33, 0x47, 0xfe, 0x2f, 0x74, 0x2b, 0xf1, 0x7c,
	0x58, 0x67, 0x35, 0xb0, 0x5d, 0xda, 0x54, 0xee, 0xff, 0xb3, 0xb3, 0x09, 0x4a, 0x9b, 0x49, 0x70,
	0x26, 0x59, 0x98, 0x15, 0xc0, 0xb7, 0xa0, 0xee, 0x4a, 0x65, 0xe4, 0xbe, 0x94, 0xdb, 0x32, 0xa3,
	0xf7, 0xc1, 0xab, 0x32, 0x77, 0x36, 0x40, 0x3d, 0x2b, 0x0c, 0xa2, 0x3a, 0x97, 0x42, 0x53, 0xfb,
	0x9a, 0x52, 0x3e, 0x79, 0x6d, 0x80, 0x4d, 0x44, 0xf5, 0x48, 0x90, 0x71, 0xa5, 0xc5, 0xc0, 0x5f,
	0xf6, 0xca, 0xcf, 0xee, 0xb0, 0x18, 0xd2, 0x1e, 0x55, 0x0f, 0x8c, 0x50, 0xb8, 0x07, 0x6a, 0x59,
	0x61, 0x60, 0x23, 0xb4, 0xff, 0x8e, 0x95, 0x7d, 0x6f, 0xfa, 0x4a, 0xe5, 0x00, 0x4b, 0x4d, 0x29,
	0xff, 0x29, 0xb5, 0x72, 0x78, 0x2b, 0xf5, 0x47, 0xf9, 0x06, 0xc6, 0x47, 0x45, 0x60, 0xc5, 0x17,
	0x9e, 0xb1, 0xf8, 0xdb, 0x55, 0xa6, 0x1b, 0xee, 0xaf, 0xb8, 0x5f, 0xa2, 0xf8, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x9d, 0x03, 0xd1, 0xba, 0x07, 0x00, 0x00,
}
