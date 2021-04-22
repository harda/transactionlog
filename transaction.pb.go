// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.8
// source: transaction.proto

package transactionlog

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

type TransactionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pcode      int32  `protobuf:"varint,2,opt,name=pcode,proto3" json:"pcode,omitempty"`
	Amount     int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Trace      int32  `protobuf:"varint,4,opt,name=trace,proto3" json:"trace,omitempty"`
	Localtime  string `protobuf:"bytes,5,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Localdate  string `protobuf:"bytes,6,opt,name=localdate,proto3" json:"localdate,omitempty"`
	Posentry   int32  `protobuf:"varint,7,opt,name=posentry,proto3" json:"posentry,omitempty"`
	Pan        string `protobuf:"bytes,8,opt,name=pan,proto3" json:"pan,omitempty"`
	Nii        int32  `protobuf:"varint,9,opt,name=nii,proto3" json:"nii,omitempty"`
	Poscon     int32  `protobuf:"varint,10,opt,name=poscon,proto3" json:"poscon,omitempty"`
	Refnum     string `protobuf:"bytes,11,opt,name=refnum,proto3" json:"refnum,omitempty"`
	Rc         string `protobuf:"bytes,12,opt,name=rc,proto3" json:"rc,omitempty"`
	Termid     string `protobuf:"bytes,13,opt,name=termid,proto3" json:"termid,omitempty"`
	Merchantid string `protobuf:"bytes,14,opt,name=merchantid,proto3" json:"merchantid,omitempty"`
	Bit59      string `protobuf:"bytes,15,opt,name=bit59,proto3" json:"bit59,omitempty"`
	Rawmsg     string `protobuf:"bytes,16,opt,name=rawmsg,proto3" json:"rawmsg,omitempty"`
}

func (x *TransactionLog) Reset() {
	*x = TransactionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLog) ProtoMessage() {}

func (x *TransactionLog) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLog.ProtoReflect.Descriptor instead.
func (*TransactionLog) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransactionLog) GetPcode() int32 {
	if x != nil {
		return x.Pcode
	}
	return 0
}

func (x *TransactionLog) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionLog) GetTrace() int32 {
	if x != nil {
		return x.Trace
	}
	return 0
}

func (x *TransactionLog) GetLocaltime() string {
	if x != nil {
		return x.Localtime
	}
	return ""
}

func (x *TransactionLog) GetLocaldate() string {
	if x != nil {
		return x.Localdate
	}
	return ""
}

func (x *TransactionLog) GetPosentry() int32 {
	if x != nil {
		return x.Posentry
	}
	return 0
}

func (x *TransactionLog) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *TransactionLog) GetNii() int32 {
	if x != nil {
		return x.Nii
	}
	return 0
}

func (x *TransactionLog) GetPoscon() int32 {
	if x != nil {
		return x.Poscon
	}
	return 0
}

func (x *TransactionLog) GetRefnum() string {
	if x != nil {
		return x.Refnum
	}
	return ""
}

func (x *TransactionLog) GetRc() string {
	if x != nil {
		return x.Rc
	}
	return ""
}

func (x *TransactionLog) GetTermid() string {
	if x != nil {
		return x.Termid
	}
	return ""
}

func (x *TransactionLog) GetMerchantid() string {
	if x != nil {
		return x.Merchantid
	}
	return ""
}

func (x *TransactionLog) GetBit59() string {
	if x != nil {
		return x.Bit59
	}
	return ""
}

func (x *TransactionLog) GetRawmsg() string {
	if x != nil {
		return x.Rawmsg
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x61, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x69, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x66, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x66, 0x6e, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x72, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x74, 0x35, 0x39, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x69, 0x74, 0x35, 0x39, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x6d, 0x73, 0x67, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x77, 0x6d, 0x73, 0x67, 0x22, 0x14, 0x0a, 0x02,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xc6, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x53, 0x61, 0x76,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x0f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67,
	0x1a, 0x03, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x0f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67,
	0x1a, 0x03, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x03, 0x2e, 0x49, 0x44, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x61, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_proto_goTypes = []interface{}{
	(*TransactionLog)(nil), // 0: TransactionLog
	(*ID)(nil),             // 1: ID
}
var file_transaction_proto_depIdxs = []int32{
	0, // 0: TransactionService.SaveTransactionLog:input_type -> TransactionLog
	1, // 1: TransactionService.SelectTransactionLog:input_type -> ID
	0, // 2: TransactionService.UpdateTransactionLog:input_type -> TransactionLog
	1, // 3: TransactionService.DeleteTransactionLog:input_type -> ID
	1, // 4: TransactionService.SaveTransactionLog:output_type -> ID
	0, // 5: TransactionService.SelectTransactionLog:output_type -> TransactionLog
	1, // 6: TransactionService.UpdateTransactionLog:output_type -> ID
	1, // 7: TransactionService.DeleteTransactionLog:output_type -> ID
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLog); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
