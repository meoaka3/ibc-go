// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/controller/v1/controller.proto

package types

import (
	fmt "fmt"
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

// Params defines the set of on-chain interchain accounts parameters.
// The following parameters may be used to disable the controller submodule.
type Params struct {
	// controller_enabled enables or disables the controller submodule.
	ControllerEnabled bool `protobuf:"varint,1,opt,name=controller_enabled,json=controllerEnabled,proto3" json:"controller_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_177fd0fec5eb3400, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetControllerEnabled() bool {
	if m != nil {
		return m.ControllerEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "ibc.applications.interchain_accounts.controller.v1.Params")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/controller/v1/controller.proto", fileDescriptor_177fd0fec5eb3400)
}

var fileDescriptor_177fd0fec5eb3400 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xa5, 0x42, 0xd9, 0xc8, 0xc4, 0x64, 0x21, 0x26, 0x96, 0xf8, 0xd4, 0x32, 0x64,
	0x07, 0xb1, 0x57, 0x8c, 0x2c, 0x95, 0x7d, 0xb5, 0xda, 0x43, 0x8e, 0x2f, 0xf2, 0x39, 0x91, 0x78,
	0x0b, 0x1e, 0x8b, 0xb1, 0x23, 0x23, 0x4a, 0x5e, 0x04, 0x01, 0x43, 0x32, 0x64, 0x3c, 0x7d, 0xba,
	0x4f, 0xff, 0x57, 0x3e, 0x91, 0x43, 0xb0, 0x5d, 0x17, 0x08, 0x6d, 0x26, 0x8e, 0x02, 0x14, 0xb3,
	0x4f, 0x78, 0xb6, 0x14, 0x0f, 0x16, 0x91, 0xfb, 0x98, 0x05, 0x90, 0x63, 0x4e, 0x1c, 0x82, 0x4f,
	0x30, 0x6c, 0x17, 0x97, 0xe9, 0x12, 0x67, 0xae, 0x76, 0xe4, 0xd0, 0x2c, 0x25, 0x66, 0x45, 0x62,
	0x16, 0x6f, 0xc3, 0xf6, 0xae, 0x29, 0x37, 0x7b, 0x9b, 0x6c, 0x2b, 0x55, 0x5d, 0x56, 0x33, 0x3a,
	0xf8, 0x68, 0x5d, 0xf0, 0xc7, 0x1b, 0x75, 0xab, 0xee, 0xaf, 0x5e, 0xae, 0x67, 0xf2, 0xfc, 0x0f,
	0x1e, 0xdf, 0x3e, 0x47, 0xad, 0x2e, 0xa3, 0x56, 0xdf, 0xa3, 0x56, 0x1f, 0x93, 0x2e, 0x2e, 0x93,
	0x2e, 0xbe, 0x26, 0x5d, 0xbc, 0xee, 0x4f, 0x94, 0xcf, 0xbd, 0x33, 0xc8, 0x2d, 0x20, 0x4b, 0xcb,
	0x02, 0xe4, 0xb0, 0x3e, 0x31, 0x0c, 0x0d, 0xb4, 0x7c, 0xec, 0x83, 0x97, 0xdf, 0x56, 0x81, 0x5d,
	0x53, 0xcf, 0x0b, 0xeb, 0xb5, 0xcc, 0xfc, 0xde, 0x79, 0x71, 0x9b, 0xbf, 0xbe, 0x87, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xe6, 0x80, 0x41, 0x26, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ControllerEnabled {
		i--
		if m.ControllerEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintController(dAtA []byte, offset int, v uint64) int {
	offset -= sovController(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ControllerEnabled {
		n += 2
	}
	return n
}

func sovController(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozController(x uint64) (n int) {
	return sovController(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControllerEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ControllerEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
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
func skipController(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowController
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
					return 0, ErrIntOverflowController
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
					return 0, ErrIntOverflowController
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
				return 0, ErrInvalidLengthController
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupController
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthController
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthController        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowController          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupController = fmt.Errorf("proto: unexpected end of group")
)
