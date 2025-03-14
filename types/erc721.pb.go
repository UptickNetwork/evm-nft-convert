// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uptick/erc721/v1/erc721.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Owner enumerates the ownership of a ERC721 contract.
type Owner int32

const (
	// OWNER_UNSPECIFIED defines an invalid/undefined owner.
	OWNER_UNSPECIFIED Owner = 0
	// OWNER_MODULE erc721 is owned by the erc721 module account.
	OWNER_MODULE Owner = 1
	// EXTERNAL erc721 is owned by an external account.
	OWNER_EXTERNAL Owner = 2
)

var Owner_name = map[int32]string{
	0: "OWNER_UNSPECIFIED",
	1: "OWNER_MODULE",
	2: "OWNER_EXTERNAL",
}

var Owner_value = map[string]int32{
	"OWNER_UNSPECIFIED": 0,
	"OWNER_MODULE":      1,
	"OWNER_EXTERNAL":    2,
}

func (x Owner) String() string {
	return proto.EnumName(Owner_name, int32(x))
}

func (Owner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{0}
}

// TokenPair defines an instance that records a pairing consisting of a native
// Cosmos Coin and an ERC721 token address.
type TokenPair struct {
	// address of ERC721 contract token
	Erc721Address string `protobuf:"bytes,1,opt,name=erc721_address,json=erc721Address,proto3" json:"erc721_address,omitempty"`
	// cosmos nft class ID to be mapped to
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetErc721Address() string {
	if m != nil {
		return m.Erc721Address
	}
	return ""
}

func (m *TokenPair) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

// defines the unique id of nft asset
type UIDPair struct {
	// address of ERC721 contract token + tokenId
	Erc721Did string `protobuf:"bytes,1,opt,name=erc721_did,json=erc721Did,proto3" json:"erc721_did,omitempty"`
	// cosmos nft class ID to be mapped to + nftId
	ClassDid string `protobuf:"bytes,2,opt,name=class_did,json=classDid,proto3" json:"class_did,omitempty"`
}

func (m *UIDPair) Reset()         { *m = UIDPair{} }
func (m *UIDPair) String() string { return proto.CompactTextString(m) }
func (*UIDPair) ProtoMessage()    {}
func (*UIDPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4208f03f5270a65, []int{1}
}
func (m *UIDPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UIDPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UIDPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UIDPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIDPair.Merge(m, src)
}
func (m *UIDPair) XXX_Size() int {
	return m.Size()
}
func (m *UIDPair) XXX_DiscardUnknown() {
	xxx_messageInfo_UIDPair.DiscardUnknown(m)
}

var xxx_messageInfo_UIDPair proto.InternalMessageInfo

func (m *UIDPair) GetErc721Did() string {
	if m != nil {
		return m.Erc721Did
	}
	return ""
}

func (m *UIDPair) GetClassDid() string {
	if m != nil {
		return m.ClassDid
	}
	return ""
}

func init() {
	proto.RegisterEnum("uptick.erc721.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "uptick.erc721.v1.TokenPair")
	proto.RegisterType((*UIDPair)(nil), "uptick.erc721.v1.UIDPair")
}

func init() { proto.RegisterFile("uptick/erc721/v1/erc721.proto", fileDescriptor_e4208f03f5270a65) }

var fileDescriptor_e4208f03f5270a65 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2d, 0x28, 0xc9,
	0x4c, 0xce, 0xd6, 0x4f, 0x2d, 0x4a, 0x36, 0x37, 0x32, 0xd4, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x20, 0xd2, 0x7a, 0x50, 0xc1, 0x32, 0x43, 0x29, 0x91, 0xf4,
	0xfc, 0xf4, 0x7c, 0xb0, 0xa4, 0x3e, 0x88, 0x05, 0x51, 0xa7, 0x14, 0xcc, 0xc5, 0x19, 0x92, 0x9f,
	0x9d, 0x9a, 0x17, 0x90, 0x98, 0x59, 0x24, 0xa4, 0xca, 0xc5, 0x07, 0x51, 0x1f, 0x9f, 0x98, 0x92,
	0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x0b, 0x11, 0x75, 0x84,
	0x08, 0x0a, 0x49, 0x72, 0x71, 0x24, 0xe7, 0x24, 0x16, 0x17, 0xc7, 0x67, 0xa6, 0x48, 0x30, 0x81,
	0x15, 0xb0, 0x83, 0xf9, 0x9e, 0x29, 0x56, 0x2c, 0x2f, 0x16, 0xc8, 0x33, 0x2a, 0x79, 0x73, 0xb1,
	0x87, 0x7a, 0xba, 0x80, 0x8d, 0x94, 0xe5, 0xe2, 0x82, 0x1a, 0x99, 0x92, 0x99, 0x02, 0x35, 0x8e,
	0x13, 0x22, 0xe2, 0x92, 0x99, 0x22, 0x24, 0xcd, 0xc5, 0x09, 0x31, 0x2a, 0x05, 0x6e, 0x16, 0xc4,
	0x6c, 0x97, 0x4c, 0xa8, 0x61, 0x5a, 0x5e, 0x5c, 0xac, 0xfe, 0xe5, 0x79, 0xa9, 0x45, 0x42, 0xa2,
	0x5c, 0x82, 0xfe, 0xe1, 0x7e, 0xae, 0x41, 0xf1, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e,
	0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0x02, 0x5c, 0x3c, 0x10, 0x61, 0x5f, 0x7f, 0x97, 0x50, 0x1f,
	0x57, 0x01, 0x46, 0x21, 0x21, 0x2e, 0x3e, 0x88, 0x88, 0x6b, 0x44, 0x88, 0x6b, 0x90, 0x9f, 0xa3,
	0x8f, 0x00, 0x93, 0x14, 0x4b, 0xc7, 0x62, 0x39, 0x06, 0x27, 0x8f, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x0f, 0x05, 0x07, 0x9d, 0x5f, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0xb6, 0x7e, 0x6a, 0x59, 0xae,
	0x6e, 0x5e, 0x5a, 0x89, 0x6e, 0x72, 0x7e, 0x5e, 0x59, 0x6a, 0x51, 0x89, 0x7e, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0x38, 0xf8, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x8b, 0x85,
	0x67, 0x87, 0x01, 0x00, 0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Erc721Address != that1.Erc721Address {
		return false
	}
	if this.ClassId != that1.ClassId {
		return false
	}
	return true
}
func (this *UIDPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UIDPair)
	if !ok {
		that2, ok := that.(UIDPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Erc721Did != that1.Erc721Did {
		return false
	}
	if this.ClassDid != that1.ClassDid {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc721Address) > 0 {
		i -= len(m.Erc721Address)
		copy(dAtA[i:], m.Erc721Address)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Erc721Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UIDPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UIDPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UIDPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassDid) > 0 {
		i -= len(m.ClassDid)
		copy(dAtA[i:], m.ClassDid)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.ClassDid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc721Did) > 0 {
		i -= len(m.Erc721Did)
		copy(dAtA[i:], m.Erc721Did)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Erc721Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErc721(dAtA []byte, offset int, v uint64) int {
	offset -= sovErc721(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc721Address)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	return n
}

func (m *UIDPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc721Did)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.ClassDid)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	return n
}

func sovErc721(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErc721(x uint64) (n int) {
	return sovErc721(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func (m *UIDPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: UIDPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UIDPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func skipErc721(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
				return 0, ErrInvalidLengthErc721
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErc721
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErc721
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErc721        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErc721          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErc721 = fmt.Errorf("proto: unexpected end of group")
)
