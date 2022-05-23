// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/basic.proto

package basic

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

type TxStatus int32

const (
	TxStatus_PROCESSED TxStatus = 0
	TxStatus_CONFIRMED TxStatus = 1
	TxStatus_FINALIZED TxStatus = 2
)

// Enum value maps for TxStatus.
var (
	TxStatus_name = map[int32]string{
		0: "PROCESSED",
		1: "CONFIRMED",
		2: "FINALIZED",
	}
	TxStatus_value = map[string]int32{
		"PROCESSED": 0,
		"CONFIRMED": 1,
		"FINALIZED": 2,
	}
)

func (x TxStatus) Enum() *TxStatus {
	p := new(TxStatus)
	*p = x
	return p
}

func (x TxStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_basic_proto_enumTypes[0].Descriptor()
}

func (TxStatus) Type() protoreflect.EnumType {
	return &file_proto_basic_proto_enumTypes[0]
}

func (x TxStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxStatus.Descriptor instead.
func (TxStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{0}
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot  uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{0}
}

func (x *AccountInfo) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *AccountInfo) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SignedTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *SignedTx) Reset() {
	*x = SignedTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTx) ProtoMessage() {}

func (x *SignedTx) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTx.ProtoReflect.Descriptor instead.
func (*SignedTx) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{1}
}

func (x *SignedTx) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

type TxStatusWithSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TxStatus `protobuf:"varint,1,opt,name=status,proto3,enum=basic.TxStatus" json:"status,omitempty"`
	Slot   uint64   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *TxStatusWithSlot) Reset() {
	*x = TxStatusWithSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxStatusWithSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxStatusWithSlot) ProtoMessage() {}

func (x *TxStatusWithSlot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxStatusWithSlot.ProtoReflect.Descriptor instead.
func (*TxStatusWithSlot) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{2}
}

func (x *TxStatusWithSlot) GetStatus() TxStatus {
	if x != nil {
		return x.Status
	}
	return TxStatus_PROCESSED
}

func (x *TxStatusWithSlot) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{3}
}

type Pubkey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Pubkey) Reset() {
	*x = Pubkey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pubkey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pubkey) ProtoMessage() {}

func (x *Pubkey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pubkey.ProtoReflect.Descriptor instead.
func (*Pubkey) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{4}
}

func (x *Pubkey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Keypair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//	*Keypair_Seed
	//	*Keypair_PrivateKey
	Input isKeypair_Input `protobuf_oneof:"input"`
}

func (x *Keypair) Reset() {
	*x = Keypair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keypair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keypair) ProtoMessage() {}

func (x *Keypair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keypair.ProtoReflect.Descriptor instead.
func (*Keypair) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{5}
}

func (m *Keypair) GetInput() isKeypair_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *Keypair) GetSeed() []byte {
	if x, ok := x.GetInput().(*Keypair_Seed); ok {
		return x.Seed
	}
	return nil
}

func (x *Keypair) GetPrivateKey() []byte {
	if x, ok := x.GetInput().(*Keypair_PrivateKey); ok {
		return x.PrivateKey
	}
	return nil
}

type isKeypair_Input interface {
	isKeypair_Input()
}

type Keypair_Seed struct {
	Seed []byte `protobuf:"bytes,1,opt,name=seed,proto3,oneof"`
}

type Keypair_PrivateKey struct {
	PrivateKey []byte `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3,oneof"`
}

func (*Keypair_Seed) isKeypair_Input() {}

func (*Keypair_PrivateKey) isKeypair_Input() {}

type ProgramId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Pubkey `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProgramId) Reset() {
	*x = ProgramId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramId) ProtoMessage() {}

func (x *ProgramId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramId.ProtoReflect.Descriptor instead.
func (*ProgramId) Descriptor() ([]byte, []int) {
	return file_proto_basic_proto_rawDescGZIP(), []int{6}
}

func (x *ProgramId) GetId() *Pubkey {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_proto_basic_proto protoreflect.FileDescriptor

var file_proto_basic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0x37, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x78, 0x22,
	0x4f, 0x0a, 0x10, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x54, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x70, 0x61,
	0x69, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x2a, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x02, 0x69, 0x64,
	0x2a, 0x37, 0x0a, 0x08, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x02, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x6c, 0x6d, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x75, 0x73, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_proto_rawDescOnce sync.Once
	file_proto_basic_proto_rawDescData = file_proto_basic_proto_rawDesc
)

func file_proto_basic_proto_rawDescGZIP() []byte {
	file_proto_basic_proto_rawDescOnce.Do(func() {
		file_proto_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_proto_rawDescData)
	})
	return file_proto_basic_proto_rawDescData
}

var file_proto_basic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_basic_proto_goTypes = []interface{}{
	(TxStatus)(0),            // 0: basic.TxStatus
	(*AccountInfo)(nil),      // 1: basic.AccountInfo
	(*SignedTx)(nil),         // 2: basic.SignedTx
	(*TxStatusWithSlot)(nil), // 3: basic.TxStatusWithSlot
	(*Empty)(nil),            // 4: basic.Empty
	(*Pubkey)(nil),           // 5: basic.Pubkey
	(*Keypair)(nil),          // 6: basic.Keypair
	(*ProgramId)(nil),        // 7: basic.ProgramId
}
var file_proto_basic_proto_depIdxs = []int32{
	0, // 0: basic.TxStatusWithSlot.status:type_name -> basic.TxStatus
	5, // 1: basic.ProgramId.id:type_name -> basic.Pubkey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_basic_proto_init() }
func file_proto_basic_proto_init() {
	if File_proto_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_proto_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTx); i {
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
		file_proto_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxStatusWithSlot); i {
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
		file_proto_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_basic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pubkey); i {
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
		file_proto_basic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keypair); i {
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
		file_proto_basic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgramId); i {
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
	file_proto_basic_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Keypair_Seed)(nil),
		(*Keypair_PrivateKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_basic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_proto_goTypes,
		DependencyIndexes: file_proto_basic_proto_depIdxs,
		EnumInfos:         file_proto_basic_proto_enumTypes,
		MessageInfos:      file_proto_basic_proto_msgTypes,
	}.Build()
	File_proto_basic_proto = out.File
	file_proto_basic_proto_rawDesc = nil
	file_proto_basic_proto_goTypes = nil
	file_proto_basic_proto_depIdxs = nil
}
