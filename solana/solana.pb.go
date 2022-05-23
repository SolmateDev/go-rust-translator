// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/solana.proto

package solana

import (
	basic "github.com/SolmateDev/go-rust-translator/basic"
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

type ConfirmationLevel int32

const (
	ConfirmationLevel_PROCESSED ConfirmationLevel = 0
	ConfirmationLevel_CONFIRMED ConfirmationLevel = 1
	ConfirmationLevel_FINALIZED ConfirmationLevel = 2
)

// Enum value maps for ConfirmationLevel.
var (
	ConfirmationLevel_name = map[int32]string{
		0: "PROCESSED",
		1: "CONFIRMED",
		2: "FINALIZED",
	}
	ConfirmationLevel_value = map[string]int32{
		"PROCESSED": 0,
		"CONFIRMED": 1,
		"FINALIZED": 2,
	}
)

func (x ConfirmationLevel) Enum() *ConfirmationLevel {
	p := new(ConfirmationLevel)
	*p = x
	return p
}

func (x ConfirmationLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfirmationLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_solana_proto_enumTypes[0].Descriptor()
}

func (ConfirmationLevel) Type() protoreflect.EnumType {
	return &file_proto_solana_proto_enumTypes[0]
}

func (x ConfirmationLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfirmationLevel.Descriptor instead.
func (ConfirmationLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{0}
}

type SendBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx                [][]byte          `protobuf:"bytes,1,rep,name=tx,proto3" json:"tx,omitempty"`
	ConfirmationLevel ConfirmationLevel `protobuf:"varint,2,opt,name=confirmation_level,json=confirmationLevel,proto3,enum=solana.ConfirmationLevel" json:"confirmation_level,omitempty"`
}

func (x *SendBatchRequest) Reset() {
	*x = SendBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBatchRequest) ProtoMessage() {}

func (x *SendBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBatchRequest.ProtoReflect.Descriptor instead.
func (*SendBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{0}
}

func (x *SendBatchRequest) GetTx() [][]byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *SendBatchRequest) GetConfirmationLevel() ConfirmationLevel {
	if x != nil {
		return x.ConfirmationLevel
	}
	return ConfirmationLevel_PROCESSED
}

type SendBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature [][]byte `protobuf:"bytes,1,rep,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SendBatchResponse) Reset() {
	*x = SendBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBatchResponse) ProtoMessage() {}

func (x *SendBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBatchResponse.ProtoReflect.Descriptor instead.
func (*SendBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{1}
}

func (x *SendBatchResponse) GetSignature() [][]byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Genesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payer *basic.Keypair `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Mint  *basic.Keypair `protobuf:"bytes,2,opt,name=mint,proto3" json:"mint,omitempty"`
	Owner *basic.Pubkey  `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// u8
	Decimals []byte `protobuf:"bytes,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *Genesis) Reset() {
	*x = Genesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genesis) ProtoMessage() {}

func (x *Genesis) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genesis.ProtoReflect.Descriptor instead.
func (*Genesis) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{2}
}

func (x *Genesis) GetPayer() *basic.Keypair {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *Genesis) GetMint() *basic.Keypair {
	if x != nil {
		return x.Mint
	}
	return nil
}

func (x *Genesis) GetOwner() *basic.Pubkey {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Genesis) GetDecimals() []byte {
	if x != nil {
		return x.Decimals
	}
	return nil
}

type InitializeTokenAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//mint: Pubkey,
	Mint *basic.Pubkey `protobuf:"bytes,1,opt,name=mint,proto3" json:"mint,omitempty"`
	//owner_account: String,
	OwnerAccount *basic.Keypair `protobuf:"bytes,2,opt,name=owner_account,json=ownerAccount,proto3" json:"owner_account,omitempty"`
}

func (x *InitializeTokenAccount) Reset() {
	*x = InitializeTokenAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeTokenAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeTokenAccount) ProtoMessage() {}

func (x *InitializeTokenAccount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeTokenAccount.ProtoReflect.Descriptor instead.
func (*InitializeTokenAccount) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{3}
}

func (x *InitializeTokenAccount) GetMint() *basic.Pubkey {
	if x != nil {
		return x.Mint
	}
	return nil
}

func (x *InitializeTokenAccount) GetOwnerAccount() *basic.Keypair {
	if x != nil {
		return x.OwnerAccount
	}
	return nil
}

type Mint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//payer: String,
	Payer *basic.Keypair `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	//signer: String,
	Signer *basic.Keypair `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	//mint_pubkey: Pubkey,
	Mint *basic.Pubkey `protobuf:"bytes,3,opt,name=mint,proto3" json:"mint,omitempty"`
	// recipient: Option<Pubkey>,
	Recipient *basic.Pubkey `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	//quantity: u64,
	Quantity uint64 `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Mint) Reset() {
	*x = Mint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mint) ProtoMessage() {}

func (x *Mint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mint.ProtoReflect.Descriptor instead.
func (*Mint) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{4}
}

func (x *Mint) GetPayer() *basic.Keypair {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *Mint) GetSigner() *basic.Keypair {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *Mint) GetMint() *basic.Pubkey {
	if x != nil {
		return x.Mint
	}
	return nil
}

func (x *Mint) GetRecipient() *basic.Pubkey {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *Mint) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mint_pubkey: Pubkey,
	Mint *basic.Pubkey `protobuf:"bytes,1,opt,name=mint,proto3" json:"mint,omitempty"`
	//owner_pubkey: Pubkey,
	Owner *basic.Pubkey `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	//payer: String,
	Payer *basic.Keypair `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
}

func (x *CreateAccount) Reset() {
	*x = CreateAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solana_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccount) ProtoMessage() {}

func (x *CreateAccount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccount.ProtoReflect.Descriptor instead.
func (*CreateAccount) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAccount) GetMint() *basic.Pubkey {
	if x != nil {
		return x.Mint
	}
	return nil
}

func (x *CreateAccount) GetOwner() *basic.Pubkey {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *CreateAccount) GetPayer() *basic.Keypair {
	if x != nil {
		return x.Payer
	}
	return nil
}

var File_proto_solana_proto protoreflect.FileDescriptor

var file_proto_solana_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x1a, 0x11, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x02, 0x74, 0x78, 0x12, 0x48, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x31, 0x0a,
	0x11, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x94, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x24, 0x0a, 0x05,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72,
	0x52, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x22, 0x70, 0x0a, 0x16, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x04,
	0x6d, 0x69, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x52, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x04, 0x4d, 0x69,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6d,
	0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x7d, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65, 0x79,
	0x70, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x2a, 0x40, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x02, 0x32, 0xb0, 0x02,
	0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x78, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x52, 0x75, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x0f, 0x2e, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x1a, 0x0c, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x19, 0x52,
	0x75, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e,
	0x61, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x52, 0x75,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15,
	0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4b, 0x65,
	0x79, 0x70, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x69,
	0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x4d, 0x69, 0x6e, 0x74,
	0x1a, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x6f, 0x6c, 0x6d, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x75, 0x73,
	0x74, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_solana_proto_rawDescOnce sync.Once
	file_proto_solana_proto_rawDescData = file_proto_solana_proto_rawDesc
)

func file_proto_solana_proto_rawDescGZIP() []byte {
	file_proto_solana_proto_rawDescOnce.Do(func() {
		file_proto_solana_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_solana_proto_rawDescData)
	})
	return file_proto_solana_proto_rawDescData
}

var file_proto_solana_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_solana_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_solana_proto_goTypes = []interface{}{
	(ConfirmationLevel)(0),         // 0: solana.ConfirmationLevel
	(*SendBatchRequest)(nil),       // 1: solana.SendBatchRequest
	(*SendBatchResponse)(nil),      // 2: solana.SendBatchResponse
	(*Genesis)(nil),                // 3: solana.Genesis
	(*InitializeTokenAccount)(nil), // 4: solana.InitializeTokenAccount
	(*Mint)(nil),                   // 5: solana.Mint
	(*CreateAccount)(nil),          // 6: solana.CreateAccount
	(*basic.Keypair)(nil),          // 7: basic.Keypair
	(*basic.Pubkey)(nil),           // 8: basic.Pubkey
	(*basic.Empty)(nil),            // 9: basic.Empty
}
var file_proto_solana_proto_depIdxs = []int32{
	0,  // 0: solana.SendBatchRequest.confirmation_level:type_name -> solana.ConfirmationLevel
	7,  // 1: solana.Genesis.payer:type_name -> basic.Keypair
	7,  // 2: solana.Genesis.mint:type_name -> basic.Keypair
	8,  // 3: solana.Genesis.owner:type_name -> basic.Pubkey
	8,  // 4: solana.InitializeTokenAccount.mint:type_name -> basic.Pubkey
	7,  // 5: solana.InitializeTokenAccount.owner_account:type_name -> basic.Keypair
	7,  // 6: solana.Mint.payer:type_name -> basic.Keypair
	7,  // 7: solana.Mint.signer:type_name -> basic.Keypair
	8,  // 8: solana.Mint.mint:type_name -> basic.Pubkey
	8,  // 9: solana.Mint.recipient:type_name -> basic.Pubkey
	8,  // 10: solana.CreateAccount.mint:type_name -> basic.Pubkey
	8,  // 11: solana.CreateAccount.owner:type_name -> basic.Pubkey
	7,  // 12: solana.CreateAccount.payer:type_name -> basic.Keypair
	1,  // 13: solana.Broadcast.SendTx:input_type -> solana.SendBatchRequest
	3,  // 14: solana.Broadcast.RunGenesis:input_type -> solana.Genesis
	4,  // 15: solana.Broadcast.RunInitializeTokenAccount:input_type -> solana.InitializeTokenAccount
	6,  // 16: solana.Broadcast.RunCreateAccount:input_type -> solana.CreateAccount
	5,  // 17: solana.Broadcast.RunMint:input_type -> solana.Mint
	2,  // 18: solana.Broadcast.SendTx:output_type -> solana.SendBatchResponse
	9,  // 19: solana.Broadcast.RunGenesis:output_type -> basic.Empty
	7,  // 20: solana.Broadcast.RunInitializeTokenAccount:output_type -> basic.Keypair
	7,  // 21: solana.Broadcast.RunCreateAccount:output_type -> basic.Keypair
	9,  // 22: solana.Broadcast.RunMint:output_type -> basic.Empty
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_solana_proto_init() }
func file_proto_solana_proto_init() {
	if File_proto_solana_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_solana_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBatchRequest); i {
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
		file_proto_solana_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBatchResponse); i {
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
		file_proto_solana_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genesis); i {
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
		file_proto_solana_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeTokenAccount); i {
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
		file_proto_solana_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mint); i {
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
		file_proto_solana_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccount); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_solana_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_solana_proto_goTypes,
		DependencyIndexes: file_proto_solana_proto_depIdxs,
		EnumInfos:         file_proto_solana_proto_enumTypes,
		MessageInfos:      file_proto_solana_proto_msgTypes,
	}.Build()
	File_proto_solana_proto = out.File
	file_proto_solana_proto_rawDesc = nil
	file_proto_solana_proto_goTypes = nil
	file_proto_solana_proto_depIdxs = nil
}
