// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: raftcmd.proto

package vearchpb

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

type OpType int32

const (
	OpType_CREATE OpType = 0
	OpType_DELETE OpType = 1
	OpType_BULK   OpType = 2
	OpType_GET    OpType = 3
	OpType_SEARCH OpType = 4
)

// Enum value maps for OpType.
var (
	OpType_name = map[int32]string{
		0: "CREATE",
		1: "DELETE",
		2: "BULK",
		3: "GET",
		4: "SEARCH",
	}
	OpType_value = map[string]int32{
		"CREATE": 0,
		"DELETE": 1,
		"BULK":   2,
		"GET":    3,
		"SEARCH": 4,
	}
)

func (x OpType) Enum() *OpType {
	p := new(OpType)
	*p = x
	return p
}

func (x OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_raftcmd_proto_enumTypes[0].Descriptor()
}

func (OpType) Type() protoreflect.EnumType {
	return &file_raftcmd_proto_enumTypes[0]
}

func (x OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpType.Descriptor instead.
func (OpType) EnumDescriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{0}
}

type CmdType int32

const (
	CmdType_WRITE       CmdType = 0
	CmdType_UPDATESPACE CmdType = 1
	CmdType_FLUSH       CmdType = 2
	CmdType_SEARCHDEL   CmdType = 3
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "WRITE",
		1: "UPDATESPACE",
		2: "FLUSH",
		3: "SEARCHDEL",
	}
	CmdType_value = map[string]int32{
		"WRITE":       0,
		"UPDATESPACE": 1,
		"FLUSH":       2,
		"SEARCHDEL":   3,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_raftcmd_proto_enumTypes[1].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_raftcmd_proto_enumTypes[1]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{1}
}

type PartitionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               OpType               `protobuf:"varint,1,opt,name=type,proto3,enum=OpType" json:"type,omitempty"`
	PartitionID        uint32               `protobuf:"varint,2,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	MessageID          string               `protobuf:"bytes,3,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Items              []*Item              `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	SearchRequest      *SearchRequest       `protobuf:"bytes,5,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	SearchResponse     *SearchResponse      `protobuf:"bytes,6,opt,name=search_response,json=searchResponse,proto3" json:"search_response,omitempty"`
	Data               []byte               `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Err                *Error               `protobuf:"bytes,8,opt,name=err,proto3" json:"err,omitempty"`
	SearchRequests     []*SearchRequest     `protobuf:"bytes,9,rep,name=search_requests,json=searchRequests,proto3" json:"search_requests,omitempty"`
	SearchResponses    []*SearchResponse    `protobuf:"bytes,10,rep,name=search_responses,json=searchResponses,proto3" json:"search_responses,omitempty"`
	DelNum             int32                `protobuf:"varint,11,opt,name=del_num,json=delNum,proto3" json:"del_num,omitempty"`
	DelByQueryResponse *DelByQueryeResponse `protobuf:"bytes,12,opt,name=del_by_query_response,json=delByQueryResponse,proto3" json:"del_by_query_response,omitempty"`
	IndexRequest       *IndexRequest        `protobuf:"bytes,13,opt,name=index_request,json=indexRequest,proto3" json:"index_request,omitempty"`
	IndexResponse      *IndexResponse       `protobuf:"bytes,14,opt,name=index_response,json=indexResponse,proto3" json:"index_response,omitempty"`
}

func (x *PartitionData) Reset() {
	*x = PartitionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftcmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionData) ProtoMessage() {}

func (x *PartitionData) ProtoReflect() protoreflect.Message {
	mi := &file_raftcmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionData.ProtoReflect.Descriptor instead.
func (*PartitionData) Descriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{0}
}

func (x *PartitionData) GetType() OpType {
	if x != nil {
		return x.Type
	}
	return OpType_CREATE
}

func (x *PartitionData) GetPartitionID() uint32 {
	if x != nil {
		return x.PartitionID
	}
	return 0
}

func (x *PartitionData) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *PartitionData) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PartitionData) GetSearchRequest() *SearchRequest {
	if x != nil {
		return x.SearchRequest
	}
	return nil
}

func (x *PartitionData) GetSearchResponse() *SearchResponse {
	if x != nil {
		return x.SearchResponse
	}
	return nil
}

func (x *PartitionData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PartitionData) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *PartitionData) GetSearchRequests() []*SearchRequest {
	if x != nil {
		return x.SearchRequests
	}
	return nil
}

func (x *PartitionData) GetSearchResponses() []*SearchResponse {
	if x != nil {
		return x.SearchResponses
	}
	return nil
}

func (x *PartitionData) GetDelNum() int32 {
	if x != nil {
		return x.DelNum
	}
	return 0
}

func (x *PartitionData) GetDelByQueryResponse() *DelByQueryeResponse {
	if x != nil {
		return x.DelByQueryResponse
	}
	return nil
}

func (x *PartitionData) GetIndexRequest() *IndexRequest {
	if x != nil {
		return x.IndexRequest
	}
	return nil
}

func (x *PartitionData) GetIndexResponse() *IndexResponse {
	if x != nil {
		return x.IndexResponse
	}
	return nil
}

// *********************** Raft *********************** //
type UpdateSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space   []byte `protobuf:"bytes,1,opt,name=Space,proto3" json:"Space,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UpdateSpace) Reset() {
	*x = UpdateSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftcmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpace) ProtoMessage() {}

func (x *UpdateSpace) ProtoReflect() protoreflect.Message {
	mi := &file_raftcmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpace.ProtoReflect.Descriptor instead.
func (*UpdateSpace) Descriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSpace) GetSpace() []byte {
	if x != nil {
		return x.Space
	}
	return nil
}

func (x *UpdateSpace) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DocCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    OpType   `protobuf:"varint,1,opt,name=type,proto3,enum=OpType" json:"type,omitempty"`
	Version int64    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Slot    uint32   `protobuf:"varint,5,opt,name=slot,proto3" json:"slot,omitempty"`
	Doc     []byte   `protobuf:"bytes,7,opt,name=doc,proto3" json:"doc,omitempty"`
	Docs    [][]byte `protobuf:"bytes,8,rep,name=docs,proto3" json:"docs,omitempty"`
}

func (x *DocCmd) Reset() {
	*x = DocCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftcmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocCmd) ProtoMessage() {}

func (x *DocCmd) ProtoReflect() protoreflect.Message {
	mi := &file_raftcmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocCmd.ProtoReflect.Descriptor instead.
func (*DocCmd) Descriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{2}
}

func (x *DocCmd) GetType() OpType {
	if x != nil {
		return x.Type
	}
	return OpType_CREATE
}

func (x *DocCmd) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DocCmd) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *DocCmd) GetDoc() []byte {
	if x != nil {
		return x.Doc
	}
	return nil
}

func (x *DocCmd) GetDocs() [][]byte {
	if x != nil {
		return x.Docs
	}
	return nil
}

type RaftCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          CmdType         `protobuf:"varint,1,opt,name=type,proto3,enum=CmdType" json:"type,omitempty"`
	WriteCommand  *DocCmd         `protobuf:"bytes,2,opt,name=write_command,json=writeCommand,proto3" json:"write_command,omitempty"`
	UpdateSpace   *UpdateSpace    `protobuf:"bytes,3,opt,name=update_space,json=updateSpace,proto3" json:"update_space,omitempty"`
	SearchDelReq  *SearchRequest  `protobuf:"bytes,4,opt,name=search_del_req,json=searchDelReq,proto3" json:"search_del_req,omitempty"`
	SearchDelResp *SearchResponse `protobuf:"bytes,5,opt,name=search_del_resp,json=searchDelResp,proto3" json:"search_del_resp,omitempty"`
}

func (x *RaftCommand) Reset() {
	*x = RaftCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftcmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftCommand) ProtoMessage() {}

func (x *RaftCommand) ProtoReflect() protoreflect.Message {
	mi := &file_raftcmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftCommand.ProtoReflect.Descriptor instead.
func (*RaftCommand) Descriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{3}
}

func (x *RaftCommand) GetType() CmdType {
	if x != nil {
		return x.Type
	}
	return CmdType_WRITE
}

func (x *RaftCommand) GetWriteCommand() *DocCmd {
	if x != nil {
		return x.WriteCommand
	}
	return nil
}

func (x *RaftCommand) GetUpdateSpace() *UpdateSpace {
	if x != nil {
		return x.UpdateSpace
	}
	return nil
}

func (x *RaftCommand) GetSearchDelReq() *SearchRequest {
	if x != nil {
		return x.SearchDelReq
	}
	return nil
}

func (x *RaftCommand) GetSearchDelResp() *SearchResponse {
	if x != nil {
		return x.SearchDelResp
	}
	return nil
}

type SnapData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SnapData) Reset() {
	*x = SnapData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftcmd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapData) ProtoMessage() {}

func (x *SnapData) ProtoReflect() protoreflect.Message {
	mi := &file_raftcmd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapData.ProtoReflect.Descriptor instead.
func (*SnapData) Descriptor() ([]byte, []int) {
	return file_raftcmd_proto_rawDescGZIP(), []int{4}
}

func (x *SnapData) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SnapData) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_raftcmd_proto protoreflect.FileDescriptor

var file_raftcmd_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x66, 0x74, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x35,
	0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x37, 0x0a,
	0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x47, 0x0a, 0x15, 0x64,
	0x65, 0x6c, 0x5f, 0x62, 0x79, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x44, 0x65, 0x6c,
	0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x12, 0x64, 0x65, 0x6c, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79,
	0x0a, 0x06, 0x44, 0x6f, 0x63, 0x43, 0x6d, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x64, 0x6f, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x0b, 0x52, 0x61,
	0x66, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x44, 0x6f, 0x63, 0x43, 0x6d, 0x64, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x0f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x32, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x3f, 0x0a, 0x06, 0x4f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x55, 0x4c, 0x4b, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x04, 0x2a, 0x3f, 0x0a, 0x07, 0x43, 0x6d,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x55, 0x53, 0x48, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x44, 0x45, 0x4c, 0x10, 0x03, 0x42, 0x0e, 0x48, 0x01, 0x5a,
	0x0a, 0x2e, 0x2f, 0x76, 0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_raftcmd_proto_rawDescOnce sync.Once
	file_raftcmd_proto_rawDescData = file_raftcmd_proto_rawDesc
)

func file_raftcmd_proto_rawDescGZIP() []byte {
	file_raftcmd_proto_rawDescOnce.Do(func() {
		file_raftcmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_raftcmd_proto_rawDescData)
	})
	return file_raftcmd_proto_rawDescData
}

var file_raftcmd_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_raftcmd_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_raftcmd_proto_goTypes = []interface{}{
	(OpType)(0),                 // 0: OpType
	(CmdType)(0),                // 1: CmdType
	(*PartitionData)(nil),       // 2: PartitionData
	(*UpdateSpace)(nil),         // 3: UpdateSpace
	(*DocCmd)(nil),              // 4: DocCmd
	(*RaftCommand)(nil),         // 5: RaftCommand
	(*SnapData)(nil),            // 6: SnapData
	(*Item)(nil),                // 7: Item
	(*SearchRequest)(nil),       // 8: SearchRequest
	(*SearchResponse)(nil),      // 9: SearchResponse
	(*Error)(nil),               // 10: Error
	(*DelByQueryeResponse)(nil), // 11: DelByQueryeResponse
	(*IndexRequest)(nil),        // 12: IndexRequest
	(*IndexResponse)(nil),       // 13: IndexResponse
}
var file_raftcmd_proto_depIdxs = []int32{
	0,  // 0: PartitionData.type:type_name -> OpType
	7,  // 1: PartitionData.items:type_name -> Item
	8,  // 2: PartitionData.search_request:type_name -> SearchRequest
	9,  // 3: PartitionData.search_response:type_name -> SearchResponse
	10, // 4: PartitionData.err:type_name -> Error
	8,  // 5: PartitionData.search_requests:type_name -> SearchRequest
	9,  // 6: PartitionData.search_responses:type_name -> SearchResponse
	11, // 7: PartitionData.del_by_query_response:type_name -> DelByQueryeResponse
	12, // 8: PartitionData.index_request:type_name -> IndexRequest
	13, // 9: PartitionData.index_response:type_name -> IndexResponse
	0,  // 10: DocCmd.type:type_name -> OpType
	1,  // 11: RaftCommand.type:type_name -> CmdType
	4,  // 12: RaftCommand.write_command:type_name -> DocCmd
	3,  // 13: RaftCommand.update_space:type_name -> UpdateSpace
	8,  // 14: RaftCommand.search_del_req:type_name -> SearchRequest
	9,  // 15: RaftCommand.search_del_resp:type_name -> SearchResponse
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_raftcmd_proto_init() }
func file_raftcmd_proto_init() {
	if File_raftcmd_proto != nil {
		return
	}
	file_errors_proto_init()
	file_data_model_proto_init()
	file_router_grpc_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_raftcmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartitionData); i {
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
		file_raftcmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSpace); i {
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
		file_raftcmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocCmd); i {
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
		file_raftcmd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftCommand); i {
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
		file_raftcmd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapData); i {
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
			RawDescriptor: file_raftcmd_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_raftcmd_proto_goTypes,
		DependencyIndexes: file_raftcmd_proto_depIdxs,
		EnumInfos:         file_raftcmd_proto_enumTypes,
		MessageInfos:      file_raftcmd_proto_msgTypes,
	}.Build()
	File_raftcmd_proto = out.File
	file_raftcmd_proto_rawDesc = nil
	file_raftcmd_proto_goTypes = nil
	file_raftcmd_proto_depIdxs = nil
}
