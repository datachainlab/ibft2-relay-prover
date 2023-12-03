// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/ibft2/v1/ibft2.proto

package module

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClientState struct {
	ChainId         string       `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	IbcStoreAddress []byte       `protobuf:"bytes,2,opt,name=ibc_store_address,json=ibcStoreAddress,proto3" json:"ibc_store_address,omitempty"`
	LatestHeight    types.Height `protobuf:"bytes,3,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_368692294f11f9d7, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

type ConsensusState struct {
	Timestamp  uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Root       []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Validators [][]byte `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_368692294f11f9d7, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

type Header struct {
	BesuHeaderRlp     []byte       `protobuf:"bytes,1,opt,name=besu_header_rlp,json=besuHeaderRlp,proto3" json:"besu_header_rlp,omitempty"`
	Seals             [][]byte     `protobuf:"bytes,2,rep,name=seals,proto3" json:"seals,omitempty"`
	TrustedHeight     types.Height `protobuf:"bytes,3,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height"`
	AccountStateProof []byte       `protobuf:"bytes,4,opt,name=account_state_proof,json=accountStateProof,proto3" json:"account_state_proof,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_368692294f11f9d7, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.ibft2.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.ibft2.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.ibft2.v1.Header")
}

func init() {
	proto.RegisterFile("ibc/lightclients/ibft2/v1/ibft2.proto", fileDescriptor_368692294f11f9d7)
}

var fileDescriptor_368692294f11f9d7 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x4d, 0xe8, 0x52, 0xa8, 0xbb, 0xdb, 0xaa, 0xa6, 0x87, 0x74, 0x85, 0xb2, 0xab, 0x95, 0x40,
	0x2b, 0xa4, 0x26, 0xda, 0xf2, 0x0b, 0x68, 0x85, 0x28, 0x27, 0x50, 0x7a, 0xe3, 0x12, 0xf9, 0x6b,
	0x13, 0x4b, 0xde, 0x38, 0xb2, 0x27, 0x91, 0xf8, 0x17, 0xdc, 0xf9, 0x31, 0x5c, 0xf7, 0xd8, 0x23,
	0x27, 0x04, 0xbb, 0x7f, 0x04, 0xd9, 0x8e, 0xa0, 0xe2, 0xc6, 0xed, 0xcd, 0x7b, 0x6f, 0xec, 0x79,
	0xf6, 0xa0, 0x17, 0x92, 0xb2, 0x5c, 0xc9, 0xaa, 0x06, 0xa6, 0xa4, 0x68, 0xc0, 0xe6, 0x92, 0xae,
	0xe1, 0x2a, 0xef, 0x57, 0x01, 0x64, 0xad, 0xd1, 0xa0, 0xf1, 0x85, 0xa4, 0x2c, 0x7b, 0x68, 0xcb,
	0x82, 0xda, 0xaf, 0xa6, 0x33, 0x77, 0x02, 0xd3, 0x46, 0xe4, 0x41, 0x72, 0xad, 0x01, 0x85, 0xde,
	0xe9, 0xac, 0xd2, 0xba, 0x52, 0x22, 0xf7, 0x15, 0xed, 0xd6, 0x39, 0xc8, 0x8d, 0xb0, 0x40, 0x36,
	0xed, 0x60, 0x48, 0xff, 0x35, 0xf0, 0xce, 0x10, 0x90, 0xba, 0x19, 0xf4, 0xf3, 0x4a, 0x57, 0xda,
	0xc3, 0xdc, 0xa1, 0xc0, 0x2e, 0xbe, 0xc6, 0xe8, 0xf8, 0xc6, 0xdf, 0x73, 0x07, 0x04, 0x04, 0xbe,
	0x40, 0x4f, 0x59, 0x4d, 0x64, 0x53, 0x4a, 0x9e, 0xc4, 0xf3, 0x78, 0x79, 0x54, 0x3c, 0xf1, 0xf5,
	0x7b, 0x8e, 0x5f, 0xa1, 0x33, 0x49, 0x59, 0x69, 0x41, 0x1b, 0x51, 0x12, 0xce, 0x8d, 0xb0, 0x36,
	0x79, 0x34, 0x8f, 0x97, 0xe3, 0xe2, 0x54, 0x52, 0x76, 0xe7, 0xf8, 0x37, 0x81, 0xc6, 0x6f, 0xd1,
	0x44, 0x11, 0x10, 0x16, 0xca, 0x5a, 0xb8, 0xbc, 0xc9, 0xc1, 0x3c, 0x5e, 0x1e, 0x5f, 0x4d, 0x33,
	0xf7, 0x02, 0x2e, 0x66, 0x36, 0x84, 0xeb, 0x57, 0xd9, 0xad, 0x77, 0x5c, 0x8f, 0xb6, 0x3f, 0x66,
	0x51, 0x31, 0x0e, 0x6d, 0x81, 0x5b, 0x50, 0x74, 0x72, 0xa3, 0x1b, 0x2b, 0x1a, 0xdb, 0xd9, 0x30,
	0xdf, 0x73, 0x74, 0xf4, 0x27, 0xb8, 0x1f, 0x70, 0x54, 0xfc, 0x25, 0x30, 0x46, 0x23, 0xa3, 0x35,
	0x0c, 0x53, 0x79, 0x8c, 0x53, 0x84, 0x7a, 0xa2, 0x24, 0x27, 0xa0, 0x8d, 0x4d, 0x0e, 0xe6, 0x07,
	0xcb, 0x71, 0xf1, 0x80, 0x59, 0x7c, 0x8b, 0xd1, 0xe1, 0xad, 0x20, 0x5c, 0x18, 0xfc, 0x12, 0x9d,
	0x52, 0x61, 0xbb, 0xb2, 0xf6, 0x65, 0x69, 0x54, 0xb8, 0x62, 0x5c, 0x4c, 0x1c, 0x1d, 0x4c, 0x85,
	0x6a, 0xf1, 0x39, 0x7a, 0x6c, 0x05, 0x51, 0x2e, 0xbd, 0x3b, 0x2d, 0x14, 0xf8, 0x1d, 0x3a, 0x01,
	0xd3, 0x59, 0x10, 0xfc, 0x7f, 0x43, 0x4f, 0x86, 0xbe, 0x40, 0xe2, 0x0c, 0x3d, 0x23, 0x8c, 0xe9,
	0xae, 0x81, 0xd2, 0xba, 0xd0, 0x65, 0x6b, 0xb4, 0x5e, 0x27, 0x23, 0x3f, 0xca, 0xd9, 0x20, 0xf9,
	0xe7, 0xf8, 0xe8, 0x84, 0xeb, 0x0f, 0xdb, 0x5f, 0x69, 0xb4, 0xdd, 0xa5, 0xf1, 0xfd, 0x2e, 0x8d,
	0x7f, 0xee, 0xd2, 0xf8, 0xcb, 0x3e, 0x8d, 0xee, 0xf7, 0x69, 0xf4, 0x7d, 0x9f, 0x46, 0x9f, 0x56,
	0x95, 0x84, 0xba, 0xa3, 0x19, 0xd3, 0x9b, 0x9c, 0x13, 0x20, 0xfe, 0x3b, 0x15, 0xa1, 0x61, 0x33,
	0x2f, 0x8d, 0x50, 0xe4, 0xf3, 0x65, 0x6b, 0x74, 0x2f, 0x4c, 0xbe, 0xd1, 0xbc, 0x53, 0x82, 0x1e,
	0xfa, 0xdd, 0x78, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x14, 0xc7, 0x49, 0xeb, 0xd7, 0x02, 0x00,
	0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIbft2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.IbcStoreAddress) > 0 {
		i -= len(m.IbcStoreAddress)
		copy(dAtA[i:], m.IbcStoreAddress)
		i = encodeVarintIbft2(dAtA, i, uint64(len(m.IbcStoreAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIbft2(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintIbft2(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintIbft2(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintIbft2(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountStateProof) > 0 {
		i -= len(m.AccountStateProof)
		copy(dAtA[i:], m.AccountStateProof)
		i = encodeVarintIbft2(dAtA, i, uint64(len(m.AccountStateProof)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TrustedHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIbft2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Seals) > 0 {
		for iNdEx := len(m.Seals) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seals[iNdEx])
			copy(dAtA[i:], m.Seals[iNdEx])
			i = encodeVarintIbft2(dAtA, i, uint64(len(m.Seals[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BesuHeaderRlp) > 0 {
		i -= len(m.BesuHeaderRlp)
		copy(dAtA[i:], m.BesuHeaderRlp)
		i = encodeVarintIbft2(dAtA, i, uint64(len(m.BesuHeaderRlp)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbft2(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbft2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIbft2(uint64(l))
	}
	l = len(m.IbcStoreAddress)
	if l > 0 {
		n += 1 + l + sovIbft2(uint64(l))
	}
	l = m.LatestHeight.Size()
	n += 1 + l + sovIbft2(uint64(l))
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovIbft2(uint64(m.Timestamp))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovIbft2(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, b := range m.Validators {
			l = len(b)
			n += 1 + l + sovIbft2(uint64(l))
		}
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BesuHeaderRlp)
	if l > 0 {
		n += 1 + l + sovIbft2(uint64(l))
	}
	if len(m.Seals) > 0 {
		for _, b := range m.Seals {
			l = len(b)
			n += 1 + l + sovIbft2(uint64(l))
		}
	}
	l = m.TrustedHeight.Size()
	n += 1 + l + sovIbft2(uint64(l))
	l = len(m.AccountStateProof)
	if l > 0 {
		n += 1 + l + sovIbft2(uint64(l))
	}
	return n
}

func sovIbft2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbft2(x uint64) (n int) {
	return sovIbft2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbft2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStoreAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcStoreAddress = append(m.IbcStoreAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.IbcStoreAddress == nil {
				m.IbcStoreAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbft2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbft2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbft2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, make([]byte, postIndex-iNdEx))
			copy(m.Validators[len(m.Validators)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbft2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbft2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbft2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BesuHeaderRlp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BesuHeaderRlp = append(m.BesuHeaderRlp[:0], dAtA[iNdEx:postIndex]...)
			if m.BesuHeaderRlp == nil {
				m.BesuHeaderRlp = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seals", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seals = append(m.Seals, make([]byte, postIndex-iNdEx))
			copy(m.Seals[len(m.Seals)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TrustedHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountStateProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbft2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbft2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountStateProof = append(m.AccountStateProof[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountStateProof == nil {
				m.AccountStateProof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbft2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbft2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIbft2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbft2
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
					return 0, ErrIntOverflowIbft2
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIbft2
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
			if length < 0 {
				return 0, ErrInvalidLengthIbft2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbft2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbft2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbft2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbft2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbft2 = fmt.Errorf("proto: unexpected end of group")
)
