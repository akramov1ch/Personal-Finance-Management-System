// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: incomeexpenses.proto

package grpc

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

type IncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	UserId   int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Category string  `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Date     string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *IncomeRequest) Reset() {
	*x = IncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeRequest) ProtoMessage() {}

func (x *IncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeRequest.ProtoReflect.Descriptor instead.
func (*IncomeRequest) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{0}
}

func (x *IncomeRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *IncomeRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IncomeRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *IncomeRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *IncomeRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type IncomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

func (x *IncomeResponse) Reset() {
	*x = IncomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeResponse) ProtoMessage() {}

func (x *IncomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeResponse.ProtoReflect.Descriptor instead.
func (*IncomeResponse) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{1}
}

func (x *IncomeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *IncomeResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type ExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	UserId   int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Category string  `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Date     string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ExpenseRequest) Reset() {
	*x = ExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseRequest) ProtoMessage() {}

func (x *ExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseRequest.ProtoReflect.Descriptor instead.
func (*ExpenseRequest) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{2}
}

func (x *ExpenseRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ExpenseRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ExpenseRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExpenseRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ExpenseRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

func (x *ExpenseResponse) Reset() {
	*x = ExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseResponse) ProtoMessage() {}

func (x *ExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseResponse.ProtoReflect.Descriptor instead.
func (*ExpenseResponse) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{3}
}

func (x *ExpenseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExpenseResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type GetTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetTransactionsRequest) Reset() {
	*x = GetTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsRequest) ProtoMessage() {}

func (x *GetTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string  `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	UserId        int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Type          string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"` // "income" or "expense"
	Amount        float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Category      string  `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Date          string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Transaction) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Transaction) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetTransactionsResponse) Reset() {
	*x = GetTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incomeexpenses_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse) ProtoMessage() {}

func (x *GetTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_incomeexpenses_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_incomeexpenses_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_incomeexpenses_proto protoreflect.FileDescriptor

var file_incomeexpenses_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xbc, 0x01,
	0x0a, 0x15, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_incomeexpenses_proto_rawDescOnce sync.Once
	file_incomeexpenses_proto_rawDescData = file_incomeexpenses_proto_rawDesc
)

func file_incomeexpenses_proto_rawDescGZIP() []byte {
	file_incomeexpenses_proto_rawDescOnce.Do(func() {
		file_incomeexpenses_proto_rawDescData = protoimpl.X.CompressGZIP(file_incomeexpenses_proto_rawDescData)
	})
	return file_incomeexpenses_proto_rawDescData
}

var file_incomeexpenses_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_incomeexpenses_proto_goTypes = []any{
	(*IncomeRequest)(nil),           // 0: IncomeRequest
	(*IncomeResponse)(nil),          // 1: IncomeResponse
	(*ExpenseRequest)(nil),          // 2: ExpenseRequest
	(*ExpenseResponse)(nil),         // 3: ExpenseResponse
	(*GetTransactionsRequest)(nil),  // 4: GetTransactionsRequest
	(*Transaction)(nil),             // 5: Transaction
	(*GetTransactionsResponse)(nil), // 6: GetTransactionsResponse
}
var file_incomeexpenses_proto_depIdxs = []int32{
	5, // 0: GetTransactionsResponse.transactions:type_name -> Transaction
	0, // 1: IncomeExpensesService.LogIncome:input_type -> IncomeRequest
	2, // 2: IncomeExpensesService.LogExpense:input_type -> ExpenseRequest
	4, // 3: IncomeExpensesService.GetTransactions:input_type -> GetTransactionsRequest
	1, // 4: IncomeExpensesService.LogIncome:output_type -> IncomeResponse
	3, // 5: IncomeExpensesService.LogExpense:output_type -> ExpenseResponse
	6, // 6: IncomeExpensesService.GetTransactions:output_type -> GetTransactionsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_incomeexpenses_proto_init() }
func file_incomeexpenses_proto_init() {
	if File_incomeexpenses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_incomeexpenses_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeRequest); i {
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
		file_incomeexpenses_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeResponse); i {
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
		file_incomeexpenses_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExpenseRequest); i {
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
		file_incomeexpenses_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExpenseResponse); i {
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
		file_incomeexpenses_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetTransactionsRequest); i {
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
		file_incomeexpenses_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Transaction); i {
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
		file_incomeexpenses_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetTransactionsResponse); i {
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
			RawDescriptor: file_incomeexpenses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_incomeexpenses_proto_goTypes,
		DependencyIndexes: file_incomeexpenses_proto_depIdxs,
		MessageInfos:      file_incomeexpenses_proto_msgTypes,
	}.Build()
	File_incomeexpenses_proto = out.File
	file_incomeexpenses_proto_rawDesc = nil
	file_incomeexpenses_proto_goTypes = nil
	file_incomeexpenses_proto_depIdxs = nil
}
