// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/grpc/proto3/network.proto

package proto3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_tendermint_tendermint_p2p "github.com/tendermint/tendermint/p2p"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty1) Reset()         { *m = Empty1{} }
func (m *Empty1) String() string { return proto.CompactTextString(m) }
func (*Empty1) ProtoMessage()    {}
func (*Empty1) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_8d833252a1306f2b, []int{0}
}
func (m *Empty1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Empty1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty1.Merge(dst, src)
}
func (m *Empty1) XXX_Size() int {
	return m.Size()
}
func (m *Empty1) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty1.DiscardUnknown(m)
}

var xxx_messageInfo_Empty1 proto.InternalMessageInfo

func (*Empty1) XXX_MessageName() string {
	return "proto3.Empty1"
}

type Peer struct {
	NodeInfo             github_com_tendermint_tendermint_p2p.DefaultNodeInfo `protobuf:"bytes,1,opt,name=NodeInfo,proto3,customtype=github.com/tendermint/tendermint/p2p.DefaultNodeInfo" json:"NodeInfo"`
	IsOutbound           bool                                                 `protobuf:"varint,2,opt,name=IsOutbound,proto3" json:"IsOutbound,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_8d833252a1306f2b, []int{1}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetIsOutbound() bool {
	if m != nil {
		return m.IsOutbound
	}
	return false
}

func (*Peer) XXX_MessageName() string {
	return "proto3.Peer"
}

type PeerResponse struct {
	Peer                 []*Peer  `protobuf:"bytes,1,rep,name=Peer" json:"Peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_8d833252a1306f2b, []int{2}
}
func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(dst, src)
}
func (m *PeerResponse) XXX_Size() int {
	return m.Size()
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetPeer() []*Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (*PeerResponse) XXX_MessageName() string {
	return "proto3.PeerResponse"
}

type NetInfoResponse struct {
	Listening            bool     `protobuf:"varint,1,opt,name=Listening,proto3" json:"Listening,omitempty"`
	Listeners            []string `protobuf:"bytes,2,rep,name=Listeners" json:"Listeners,omitempty"`
	Peers                []*Peer  `protobuf:"bytes,3,rep,name=Peers" json:"Peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetInfoResponse) Reset()         { *m = NetInfoResponse{} }
func (m *NetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetInfoResponse) ProtoMessage()    {}
func (*NetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_8d833252a1306f2b, []int{3}
}
func (m *NetInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetInfoResponse.Merge(dst, src)
}
func (m *NetInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *NetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetInfoResponse proto.InternalMessageInfo

func (m *NetInfoResponse) GetListening() bool {
	if m != nil {
		return m.Listening
	}
	return false
}

func (m *NetInfoResponse) GetListeners() []string {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *NetInfoResponse) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (*NetInfoResponse) XXX_MessageName() string {
	return "proto3.NetInfoResponse"
}
func init() {
	proto.RegisterType((*Empty1)(nil), "proto3.Empty1")
	proto.RegisterType((*Peer)(nil), "proto3.Peer")
	proto.RegisterType((*PeerResponse)(nil), "proto3.PeerResponse")
	proto.RegisterType((*NetInfoResponse)(nil), "proto3.NetInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkClient interface {
	GetNetworkInfo(ctx context.Context, in *Empty1, opts ...grpc.CallOption) (*NetInfoResponse, error)
	GetPeers(ctx context.Context, in *Empty1, opts ...grpc.CallOption) (*PeerResponse, error)
}

type networkClient struct {
	cc *grpc.ClientConn
}

func NewNetworkClient(cc *grpc.ClientConn) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) GetNetworkInfo(ctx context.Context, in *Empty1, opts ...grpc.CallOption) (*NetInfoResponse, error) {
	out := new(NetInfoResponse)
	err := c.cc.Invoke(ctx, "/proto3.Network/GetNetworkInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetPeers(ctx context.Context, in *Empty1, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/proto3.Network/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
type NetworkServer interface {
	GetNetworkInfo(context.Context, *Empty1) (*NetInfoResponse, error)
	GetPeers(context.Context, *Empty1) (*PeerResponse, error)
}

func RegisterNetworkServer(s *grpc.Server, srv NetworkServer) {
	s.RegisterService(&_Network_serviceDesc, srv)
}

func _Network_GetNetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetNetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto3.Network/GetNetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetNetworkInfo(ctx, req.(*Empty1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto3.Network/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetPeers(ctx, req.(*Empty1))
	}
	return interceptor(ctx, in, info, handler)
}

var _Network_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto3.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkInfo",
			Handler:    _Network_GetNetworkInfo_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _Network_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/grpc/proto3/network.proto",
}

func (m *Empty1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintNetwork(dAtA, i, uint64(m.NodeInfo.Size()))
	n1, err := m.NodeInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.IsOutbound {
		dAtA[i] = 0x10
		i++
		if m.IsOutbound {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PeerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Peer) > 0 {
		for _, msg := range m.Peer {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNetwork(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NetInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Listening {
		dAtA[i] = 0x8
		i++
		if m.Listening {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Listeners) > 0 {
		for _, s := range m.Listeners {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Peers) > 0 {
		for _, msg := range m.Peers {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintNetwork(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNetwork(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NodeInfo.Size()
	n += 1 + l + sovNetwork(uint64(l))
	if m.IsOutbound {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PeerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Peer) > 0 {
		for _, e := range m.Peer {
			l = e.Size()
			n += 1 + l + sovNetwork(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NetInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Listening {
		n += 2
	}
	if len(m.Listeners) > 0 {
		for _, s := range m.Listeners {
			l = len(s)
			n += 1 + l + sovNetwork(uint64(l))
		}
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovNetwork(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetwork(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetwork(x uint64) (n int) {
	return sovNetwork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOutbound", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsOutbound = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peer = append(m.Peer, &Peer{})
			if err := m.Peer[len(m.Peer)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Listening", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Listening = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Listeners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Listeners = append(m.Listeners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetwork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNetwork
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetwork
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetwork(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetwork = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetwork   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("rpc/grpc/proto3/network.proto", fileDescriptor_network_8d833252a1306f2b)
}

var fileDescriptor_network_8d833252a1306f2b = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x65, 0x41, 0xb1, 0x8c, 0x04, 0x93, 0xc6, 0xc4, 0x86, 0xe0, 0xd2, 0xf4, 0xd4, 0x8b, 0x05,
	0xc1, 0x8b, 0x89, 0x27, 0xa2, 0x21, 0x24, 0x06, 0x4d, 0x4f, 0x5e, 0x2d, 0x0c, 0xb5, 0x51, 0x76,
	0x6b, 0xbb, 0xd5, 0x78, 0xf3, 0x67, 0x78, 0xf4, 0xa7, 0x78, 0xe4, 0xe8, 0xd9, 0x03, 0x31, 0xf4,
	0x8f, 0x18, 0x76, 0xa9, 0xd4, 0x8f, 0x4b, 0x33, 0xef, 0xcd, 0xbc, 0x37, 0xaf, 0xb3, 0xb0, 0x1f,
	0x85, 0xa3, 0x96, 0xbf, 0xfc, 0x84, 0x11, 0x17, 0xbc, 0xdb, 0x62, 0x28, 0x1e, 0x79, 0x74, 0xeb,
	0x48, 0xa8, 0x97, 0x15, 0x5b, 0x3f, 0xf0, 0x03, 0x71, 0x93, 0x78, 0xce, 0x88, 0x4f, 0x5b, 0x3e,
	0xf7, 0xb9, 0x9a, 0xf6, 0x92, 0x89, 0x44, 0x12, 0xc8, 0x4a, 0xc9, 0x2c, 0x0d, 0xca, 0x67, 0xd3,
	0x50, 0x3c, 0x1d, 0x5a, 0xcf, 0x04, 0x36, 0x2e, 0x11, 0x23, 0xfd, 0x0a, 0xb4, 0x21, 0x1f, 0xe3,
	0x80, 0x4d, 0xb8, 0x41, 0x4c, 0x62, 0x57, 0x7b, 0x27, 0xb3, 0x79, 0xb3, 0xf0, 0x31, 0x6f, 0x1e,
	0xe5, 0xbc, 0x05, 0xb2, 0x31, 0x46, 0xd3, 0x80, 0x89, 0x7c, 0x19, 0x76, 0x42, 0xe7, 0x14, 0x27,
	0xd7, 0xc9, 0x9d, 0xc8, 0x3c, 0xdc, 0x6f, 0x37, 0x9d, 0x02, 0x0c, 0xe2, 0x8b, 0x44, 0x78, 0x3c,
	0x61, 0x63, 0xa3, 0x68, 0x12, 0x5b, 0x73, 0x73, 0x8c, 0xd5, 0x86, 0xea, 0x32, 0x81, 0x8b, 0x71,
	0xc8, 0x59, 0x8c, 0xba, 0xa9, 0x12, 0x19, 0xc4, 0x2c, 0xd9, 0xdb, 0x9d, 0xaa, 0x8a, 0xdc, 0x75,
	0xe4, 0x8c, 0xec, 0x58, 0xf7, 0xb0, 0x33, 0x44, 0x21, 0xd7, 0x64, 0xa2, 0x06, 0x54, 0xce, 0x83,
	0x58, 0x20, 0x0b, 0x98, 0x2f, 0xf3, 0x6b, 0xee, 0x9a, 0x58, 0x77, 0x31, 0x8a, 0x8d, 0xa2, 0x59,
	0xb2, 0x2b, 0xee, 0x9a, 0xd0, 0x2d, 0xd8, 0x5c, 0xda, 0xc6, 0x46, 0xe9, 0x9f, 0x8d, 0xaa, 0xd5,
	0x79, 0x80, 0xad, 0xa1, 0xba, 0xbc, 0x7e, 0x0c, 0xb5, 0x3e, 0x8a, 0x15, 0x92, 0x7f, 0x58, 0xcb,
	0x14, 0xea, 0xa8, 0xf5, 0xbd, 0x0c, 0xff, 0x4e, 0xd9, 0x06, 0xad, 0x8f, 0x42, 0x3a, 0xfe, 0x11,
	0xed, 0xfe, 0x58, 0xbb, 0x52, 0xf4, 0x1a, 0xb3, 0x05, 0x25, 0xef, 0x0b, 0x4a, 0x3e, 0x17, 0x94,
	0xbc, 0xa4, 0xb4, 0xf0, 0x9a, 0xd2, 0xc2, 0x5b, 0x4a, 0xc9, 0x2c, 0xa5, 0xc4, 0x5b, 0x3d, 0xff,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x7e, 0xd9, 0xca, 0x26, 0x02, 0x00, 0x00,
}
