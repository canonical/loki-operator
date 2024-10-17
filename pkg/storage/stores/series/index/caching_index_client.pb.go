// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/storage/stores/series/index/caching_index_client.proto

package index

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type CacheEntry struct {
	Column Bytes `protobuf:"bytes,1,opt,name=Column,proto3,customtype=Bytes" json:"Column"`
	Value  Bytes `protobuf:"bytes,2,opt,name=Value,proto3,customtype=Bytes" json:"Value"`
}

func (m *CacheEntry) Reset()      { *m = CacheEntry{} }
func (*CacheEntry) ProtoMessage() {}
func (*CacheEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3c8533c7598e01, []int{0}
}
func (m *CacheEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheEntry.Merge(m, src)
}
func (m *CacheEntry) XXX_Size() int {
	return m.Size()
}
func (m *CacheEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CacheEntry proto.InternalMessageInfo

type ReadBatch struct {
	Entries []CacheEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries"`
	Key     string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The time at which the key expires.
	Expiry int64 `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// The number of entries; used for cardinality limiting.
	// entries will be empty when this is set.
	Cardinality int32 `protobuf:"varint,4,opt,name=cardinality,proto3" json:"cardinality,omitempty"`
}

func (m *ReadBatch) Reset()      { *m = ReadBatch{} }
func (*ReadBatch) ProtoMessage() {}
func (*ReadBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3c8533c7598e01, []int{1}
}
func (m *ReadBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBatch.Merge(m, src)
}
func (m *ReadBatch) XXX_Size() int {
	return m.Size()
}
func (m *ReadBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBatch proto.InternalMessageInfo

func (m *ReadBatch) GetEntries() []CacheEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ReadBatch) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReadBatch) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *ReadBatch) GetCardinality() int32 {
	if m != nil {
		return m.Cardinality
	}
	return 0
}

func init() {
	proto.RegisterType((*CacheEntry)(nil), "index.CacheEntry")
	proto.RegisterType((*ReadBatch)(nil), "index.ReadBatch")
}

func init() {
	proto.RegisterFile("pkg/storage/stores/series/index/caching_index_client.proto", fileDescriptor_1a3c8533c7598e01)
}

var fileDescriptor_1a3c8533c7598e01 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xbe, 0xfb, 0xa5, 0xe9, 0x8f, 0x5e, 0x15, 0xf4, 0x10, 0x09, 0x0e, 0x6f, 0x43, 0x45, 0xc8,
	0xd4, 0xe0, 0x9f, 0xc9, 0x31, 0xc5, 0x2f, 0x90, 0x41, 0xdc, 0xca, 0x99, 0xbe, 0xa4, 0x47, 0xe3,
	0xa5, 0x24, 0x57, 0x68, 0x36, 0x67, 0x27, 0x3f, 0x86, 0x1f, 0xa5, 0x63, 0xc7, 0xe2, 0x50, 0xec,
	0x75, 0x71, 0xec, 0x47, 0x90, 0x5c, 0x2b, 0x3a, 0x38, 0xdd, 0xf3, 0x8f, 0x7b, 0x1e, 0x5e, 0x76,
	0x3b, 0x19, 0xa7, 0x61, 0xa9, 0xf3, 0x42, 0xa4, 0x68, 0x5f, 0x2c, 0xc3, 0x12, 0x0b, 0x89, 0x65,
	0x28, 0xd5, 0x10, 0x67, 0x61, 0x22, 0x92, 0x91, 0x54, 0xe9, 0xc0, 0xb2, 0x41, 0x92, 0x49, 0x54,
	0xba, 0x37, 0x29, 0x72, 0x9d, 0x73, 0xd7, 0x6a, 0x67, 0x27, 0x69, 0x9e, 0xe6, 0x56, 0x09, 0x6b,
	0xb4, 0x33, 0xbb, 0x0f, 0x8c, 0xf5, 0x45, 0x32, 0xc2, 0x3b, 0xa5, 0x8b, 0x8a, 0x5f, 0xb0, 0x66,
	0x3f, 0xcf, 0xa6, 0x4f, 0xca, 0xa3, 0x3e, 0x0d, 0x0e, 0xa2, 0xc3, 0xf9, 0xaa, 0x43, 0xde, 0x57,
	0x1d, 0x37, 0xaa, 0x34, 0x96, 0xf1, 0xde, 0xe4, 0xe7, 0xcc, 0xbd, 0x17, 0xd9, 0x14, 0xbd, 0x7f,
	0x7f, 0xa5, 0x76, 0x5e, 0xf7, 0x85, 0xb2, 0x56, 0x8c, 0x62, 0x18, 0x09, 0x9d, 0x8c, 0xf8, 0x25,
	0xfb, 0x8f, 0x4a, 0xd7, 0x83, 0x3d, 0xea, 0x3b, 0x41, 0xfb, 0xea, 0xb8, 0x67, 0x67, 0xf5, 0x7e,
	0xda, 0xa3, 0x46, 0xfd, 0x4f, 0xfc, 0x9d, 0xe3, 0x47, 0xcc, 0x19, 0x63, 0x65, 0x3b, 0x5a, 0x71,
	0x0d, 0xf9, 0x29, 0x6b, 0xe2, 0x6c, 0x22, 0x8b, 0xca, 0x73, 0x7c, 0x1a, 0x38, 0xf1, 0x9e, 0x71,
	0x9f, 0xb5, 0x13, 0x51, 0x0c, 0xa5, 0x12, 0x99, 0xd4, 0x95, 0xd7, 0xf0, 0x69, 0xe0, 0xc6, 0xbf,
	0xa5, 0xe8, 0x66, 0xb1, 0x06, 0xb2, 0x5c, 0x03, 0xd9, 0xae, 0x81, 0x3e, 0x1b, 0xa0, 0x6f, 0x06,
	0xe8, 0xdc, 0x00, 0x5d, 0x18, 0xa0, 0x1f, 0x06, 0xe8, 0xa7, 0x01, 0xb2, 0x35, 0x40, 0x5f, 0x37,
	0x40, 0x16, 0x1b, 0x20, 0xcb, 0x0d, 0x90, 0xc7, 0xa6, 0xbd, 0xd1, 0xf5, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xad, 0xdc, 0x48, 0x36, 0x7e, 0x01, 0x00, 0x00,
}

func (this *CacheEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CacheEntry)
	if !ok {
		that2, ok := that.(CacheEntry)
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
	if !this.Column.Equal(that1.Column) {
		return false
	}
	if !this.Value.Equal(that1.Value) {
		return false
	}
	return true
}
func (this *ReadBatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReadBatch)
	if !ok {
		that2, ok := that.(ReadBatch)
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
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(&that1.Entries[i]) {
			return false
		}
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Expiry != that1.Expiry {
		return false
	}
	if this.Cardinality != that1.Cardinality {
		return false
	}
	return true
}
func (this *CacheEntry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&index.CacheEntry{")
	s = append(s, "Column: "+fmt.Sprintf("%#v", this.Column)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReadBatch) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&index.ReadBatch{")
	if this.Entries != nil {
		vs := make([]*CacheEntry, len(this.Entries))
		for i := range vs {
			vs[i] = &this.Entries[i]
		}
		s = append(s, "Entries: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Expiry: "+fmt.Sprintf("%#v", this.Expiry)+",\n")
	s = append(s, "Cardinality: "+fmt.Sprintf("%#v", this.Cardinality)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCachingIndexClient(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CacheEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCachingIndexClient(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Column.Size()
		i -= size
		if _, err := m.Column.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCachingIndexClient(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ReadBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cardinality != 0 {
		i = encodeVarintCachingIndexClient(dAtA, i, uint64(m.Cardinality))
		i--
		dAtA[i] = 0x20
	}
	if m.Expiry != 0 {
		i = encodeVarintCachingIndexClient(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCachingIndexClient(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCachingIndexClient(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCachingIndexClient(dAtA []byte, offset int, v uint64) int {
	offset -= sovCachingIndexClient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CacheEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Column.Size()
	n += 1 + l + sovCachingIndexClient(uint64(l))
	l = m.Value.Size()
	n += 1 + l + sovCachingIndexClient(uint64(l))
	return n
}

func (m *ReadBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovCachingIndexClient(uint64(l))
		}
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCachingIndexClient(uint64(l))
	}
	if m.Expiry != 0 {
		n += 1 + sovCachingIndexClient(uint64(m.Expiry))
	}
	if m.Cardinality != 0 {
		n += 1 + sovCachingIndexClient(uint64(m.Cardinality))
	}
	return n
}

func sovCachingIndexClient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCachingIndexClient(x uint64) (n int) {
	return sovCachingIndexClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CacheEntry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CacheEntry{`,
		`Column:` + fmt.Sprintf("%v", this.Column) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReadBatch) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEntries := "[]CacheEntry{"
	for _, f := range this.Entries {
		repeatedStringForEntries += strings.Replace(strings.Replace(f.String(), "CacheEntry", "CacheEntry", 1), `&`, ``, 1) + ","
	}
	repeatedStringForEntries += "}"
	s := strings.Join([]string{`&ReadBatch{`,
		`Entries:` + repeatedStringForEntries + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Expiry:` + fmt.Sprintf("%v", this.Expiry) + `,`,
		`Cardinality:` + fmt.Sprintf("%v", this.Cardinality) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCachingIndexClient(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CacheEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCachingIndexClient
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
			return fmt.Errorf("proto: CacheEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Column", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
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
				return ErrInvalidLengthCachingIndexClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Column.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
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
				return ErrInvalidLengthCachingIndexClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCachingIndexClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCachingIndexClient
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
func (m *ReadBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCachingIndexClient
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
			return fmt.Errorf("proto: ReadBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
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
				return ErrInvalidLengthCachingIndexClient
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, CacheEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
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
				return ErrInvalidLengthCachingIndexClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cardinality", wireType)
			}
			m.Cardinality = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCachingIndexClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cardinality |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCachingIndexClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCachingIndexClient
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCachingIndexClient
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
func skipCachingIndexClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCachingIndexClient
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
					return 0, ErrIntOverflowCachingIndexClient
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
					return 0, ErrIntOverflowCachingIndexClient
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
				return 0, ErrInvalidLengthCachingIndexClient
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCachingIndexClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCachingIndexClient
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
				next, err := skipCachingIndexClient(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCachingIndexClient
				}
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
	ErrInvalidLengthCachingIndexClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCachingIndexClient   = fmt.Errorf("proto: integer overflow")
)
