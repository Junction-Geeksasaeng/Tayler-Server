// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paper/paper.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Paper struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id        uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Host      string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	PaperName string `protobuf:"bytes,4,opt,name=paperName,proto3" json:"paperName,omitempty"`
	Owner     string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Price     string `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *Paper) Reset()         { *m = Paper{} }
func (m *Paper) String() string { return proto.CompactTextString(m) }
func (*Paper) ProtoMessage()    {}
func (*Paper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebf187adf54427ce, []int{0}
}
func (m *Paper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Paper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Paper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Paper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paper.Merge(m, src)
}
func (m *Paper) XXX_Size() int {
	return m.Size()
}
func (m *Paper) XXX_DiscardUnknown() {
	xxx_messageInfo_Paper.DiscardUnknown(m)
}

var xxx_messageInfo_Paper proto.InternalMessageInfo

func (m *Paper) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Paper) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Paper) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Paper) GetPaperName() string {
	if m != nil {
		return m.PaperName
	}
	return ""
}

func (m *Paper) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Paper) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func init() {
	proto.RegisterType((*Paper)(nil), "paper.paper.Paper")
}

func init() { proto.RegisterFile("paper/paper.proto", fileDescriptor_ebf187adf54427ce) }

var fileDescriptor_ebf187adf54427ce = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0x2c, 0x48,
	0x2d, 0xd2, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x50, 0x0e, 0x88, 0x54,
	0xea, 0x67, 0xe4, 0x62, 0x0d, 0x00, 0xb1, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b,
	0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0x8c,
	0xfc, 0xe2, 0x12, 0x09, 0x66, 0xb0, 0x32, 0x30, 0x5b, 0x48, 0x86, 0x8b, 0x13, 0x6c, 0xa0, 0x5f,
	0x62, 0x6e, 0xaa, 0x04, 0x0b, 0x58, 0x02, 0x21, 0x20, 0x24, 0xc2, 0xc5, 0x9a, 0x5f, 0x9e, 0x97,
	0x5a, 0x24, 0xc1, 0x0a, 0x96, 0x81, 0x70, 0x40, 0xa2, 0x05, 0x45, 0x99, 0xc9, 0xa9, 0x12, 0x6c,
	0x10, 0x51, 0x30, 0xc7, 0x49, 0xf7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x84, 0x21, 0x7e, 0xa9, 0x80, 0xf8, 0x46, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec,
	0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x39, 0x9a, 0xa8, 0xe9, 0x00, 0x00, 0x00,
}

func (m *Paper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Paper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Paper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintPaper(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintPaper(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PaperName) > 0 {
		i -= len(m.PaperName)
		copy(dAtA[i:], m.PaperName)
		i = encodeVarintPaper(dAtA, i, uint64(len(m.PaperName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintPaper(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintPaper(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPaper(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaper(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Paper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPaper(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPaper(uint64(m.Id))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovPaper(uint64(l))
	}
	l = len(m.PaperName)
	if l > 0 {
		n += 1 + l + sovPaper(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovPaper(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovPaper(uint64(l))
	}
	return n
}

func sovPaper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaper(x uint64) (n int) {
	return sovPaper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Paper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaper
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
			return fmt.Errorf("proto: Paper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Paper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
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
				return ErrInvalidLengthPaper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
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
				return ErrInvalidLengthPaper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaperName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
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
				return ErrInvalidLengthPaper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaperName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
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
				return ErrInvalidLengthPaper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaper
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
				return ErrInvalidLengthPaper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPaper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaper
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
func skipPaper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaper
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
					return 0, ErrIntOverflowPaper
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
					return 0, ErrIntOverflowPaper
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
				return 0, ErrInvalidLengthPaper
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaper
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaper
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaper        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaper          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaper = fmt.Errorf("proto: unexpected end of group")
)
