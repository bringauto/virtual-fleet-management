//*
// Detailed description can be found at Industrial Portal Protocol
// document located at our Google Disk
// https://drive.google.com/drive/folders/1ZE9VRs86QtP6GqTJBl6vRJLmkh1lTEc5

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: IndustrialPortalProtocol.proto

package ba_proto

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

type ConnectResponse_Type int32

const (
	ConnectResponse_OK ConnectResponse_Type = 0
	//*
	// If some car is already logged in under same company and name
	ConnectResponse_ALREADY_LOGGED ConnectResponse_Type = 1
)

// Enum value maps for ConnectResponse_Type.
var (
	ConnectResponse_Type_name = map[int32]string{
		0: "OK",
		1: "ALREADY_LOGGED",
	}
	ConnectResponse_Type_value = map[string]int32{
		"OK":             0,
		"ALREADY_LOGGED": 1,
	}
)

func (x ConnectResponse_Type) Enum() *ConnectResponse_Type {
	p := new(ConnectResponse_Type)
	*p = x
	return p
}

func (x ConnectResponse_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectResponse_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_IndustrialPortalProtocol_proto_enumTypes[0].Descriptor()
}

func (ConnectResponse_Type) Type() protoreflect.EnumType {
	return &file_IndustrialPortalProtocol_proto_enumTypes[0]
}

func (x ConnectResponse_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectResponse_Type.Descriptor instead.
func (ConnectResponse_Type) EnumDescriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{3, 0}
}

type Status_ServerError_Type int32

const (
	Status_ServerError_OK           Status_ServerError_Type = 0
	Status_ServerError_SERVER_ERROR Status_ServerError_Type = 1
)

// Enum value maps for Status_ServerError_Type.
var (
	Status_ServerError_Type_name = map[int32]string{
		0: "OK",
		1: "SERVER_ERROR",
	}
	Status_ServerError_Type_value = map[string]int32{
		"OK":           0,
		"SERVER_ERROR": 1,
	}
)

func (x Status_ServerError_Type) Enum() *Status_ServerError_Type {
	p := new(Status_ServerError_Type)
	*p = x
	return p
}

func (x Status_ServerError_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_ServerError_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_IndustrialPortalProtocol_proto_enumTypes[1].Descriptor()
}

func (Status_ServerError_Type) Type() protoreflect.EnumType {
	return &file_IndustrialPortalProtocol_proto_enumTypes[1]
}

func (x Status_ServerError_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_ServerError_Type.Descriptor instead.
func (Status_ServerError_Type) EnumDescriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{4, 0, 0}
}

type StatusResponse_Type int32

const (
	StatusResponse_OK StatusResponse_Type = 0
)

// Enum value maps for StatusResponse_Type.
var (
	StatusResponse_Type_name = map[int32]string{
		0: "OK",
	}
	StatusResponse_Type_value = map[string]int32{
		"OK": 0,
	}
)

func (x StatusResponse_Type) Enum() *StatusResponse_Type {
	p := new(StatusResponse_Type)
	*p = x
	return p
}

func (x StatusResponse_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusResponse_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_IndustrialPortalProtocol_proto_enumTypes[2].Descriptor()
}

func (StatusResponse_Type) Type() protoreflect.EnumType {
	return &file_IndustrialPortalProtocol_proto_enumTypes[2]
}

func (x StatusResponse_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusResponse_Type.Descriptor instead.
func (StatusResponse_Type) EnumDescriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{5, 0}
}

type CommandResponse_Type int32

const (
	CommandResponse_OK CommandResponse_Type = 0
)

// Enum value maps for CommandResponse_Type.
var (
	CommandResponse_Type_name = map[int32]string{
		0: "OK",
	}
	CommandResponse_Type_value = map[string]int32{
		"OK": 0,
	}
)

func (x CommandResponse_Type) Enum() *CommandResponse_Type {
	p := new(CommandResponse_Type)
	*p = x
	return p
}

func (x CommandResponse_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandResponse_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_IndustrialPortalProtocol_proto_enumTypes[3].Descriptor()
}

func (CommandResponse_Type) Type() protoreflect.EnumType {
	return &file_IndustrialPortalProtocol_proto_enumTypes[3]
}

func (x CommandResponse_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandResponse_Type.Descriptor instead.
func (CommandResponse_Type) EnumDescriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{7, 0}
}

//*
// Special message which contains other IndustrialPortal messages
// Every message of this type can contain only one of the IndustrialPortal messages
// From Server to Client only.
type MessageIndustrialPortal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*MessageIndustrialPortal_ConnectReponse
	//	*MessageIndustrialPortal_StatusResponse
	//	*MessageIndustrialPortal_Command
	Type isMessageIndustrialPortal_Type `protobuf_oneof:"Type"`
}

func (x *MessageIndustrialPortal) Reset() {
	*x = MessageIndustrialPortal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIndustrialPortal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIndustrialPortal) ProtoMessage() {}

func (x *MessageIndustrialPortal) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIndustrialPortal.ProtoReflect.Descriptor instead.
func (*MessageIndustrialPortal) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{0}
}

func (m *MessageIndustrialPortal) GetType() isMessageIndustrialPortal_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *MessageIndustrialPortal) GetConnectReponse() *ConnectResponse {
	if x, ok := x.GetType().(*MessageIndustrialPortal_ConnectReponse); ok {
		return x.ConnectReponse
	}
	return nil
}

func (x *MessageIndustrialPortal) GetStatusResponse() *StatusResponse {
	if x, ok := x.GetType().(*MessageIndustrialPortal_StatusResponse); ok {
		return x.StatusResponse
	}
	return nil
}

func (x *MessageIndustrialPortal) GetCommand() *Command {
	if x, ok := x.GetType().(*MessageIndustrialPortal_Command); ok {
		return x.Command
	}
	return nil
}

type isMessageIndustrialPortal_Type interface {
	isMessageIndustrialPortal_Type()
}

type MessageIndustrialPortal_ConnectReponse struct {
	ConnectReponse *ConnectResponse `protobuf:"bytes,1,opt,name=connectReponse,proto3,oneof"`
}

type MessageIndustrialPortal_StatusResponse struct {
	StatusResponse *StatusResponse `protobuf:"bytes,2,opt,name=statusResponse,proto3,oneof"`
}

type MessageIndustrialPortal_Command struct {
	Command *Command `protobuf:"bytes,3,opt,name=command,proto3,oneof"`
}

func (*MessageIndustrialPortal_ConnectReponse) isMessageIndustrialPortal_Type() {}

func (*MessageIndustrialPortal_StatusResponse) isMessageIndustrialPortal_Type() {}

func (*MessageIndustrialPortal_Command) isMessageIndustrialPortal_Type() {}

//*
// Special message which contains other Daemon messages
// Every message of this type can contain only one of the Daemon messages
// From Client to Server only.
type MessageDaemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*MessageDaemon_Connect
	//	*MessageDaemon_Status
	//	*MessageDaemon_CommandResponse
	Type isMessageDaemon_Type `protobuf_oneof:"Type"`
}

func (x *MessageDaemon) Reset() {
	*x = MessageDaemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDaemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDaemon) ProtoMessage() {}

func (x *MessageDaemon) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDaemon.ProtoReflect.Descriptor instead.
func (*MessageDaemon) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{1}
}

func (m *MessageDaemon) GetType() isMessageDaemon_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *MessageDaemon) GetConnect() *Connect {
	if x, ok := x.GetType().(*MessageDaemon_Connect); ok {
		return x.Connect
	}
	return nil
}

func (x *MessageDaemon) GetStatus() *Status {
	if x, ok := x.GetType().(*MessageDaemon_Status); ok {
		return x.Status
	}
	return nil
}

func (x *MessageDaemon) GetCommandResponse() *CommandResponse {
	if x, ok := x.GetType().(*MessageDaemon_CommandResponse); ok {
		return x.CommandResponse
	}
	return nil
}

type isMessageDaemon_Type interface {
	isMessageDaemon_Type()
}

type MessageDaemon_Connect struct {
	Connect *Connect `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type MessageDaemon_Status struct {
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

type MessageDaemon_CommandResponse struct {
	CommandResponse *CommandResponse `protobuf:"bytes,3,opt,name=commandResponse,proto3,oneof"`
}

func (*MessageDaemon_Connect) isMessageDaemon_Type() {}

func (*MessageDaemon_Status) isMessageDaemon_Type() {}

func (*MessageDaemon_CommandResponse) isMessageDaemon_Type() {}

//*
// Connect message information
// First message in new communication.
type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// Company name
	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	//
	// Car name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//
	// sessionId of the car
	// - generated before sending first message
	// - stays same in each session
	// - is for check if server communicates with the same car
	SessionId string `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{2}
}

func (x *Connect) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Connect) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Connect) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// ConnectResponse information
// Response only to Connect message
type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      ConnectResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=IndustrialPortal.ConnectResponse_Type" json:"type,omitempty"`
	SessionId string               `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectResponse) GetType() ConnectResponse_Type {
	if x != nil {
		return x.Type
	}
	return ConnectResponse_OK
}

func (x *ConnectResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// Car Status information with error type
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// CarStatus information
	// - full description is in the CarStateProtocol.proto
	CarStatus *CarStatus          `protobuf:"bytes,1,opt,name=carStatus,proto3" json:"carStatus,omitempty"`
	Server    *Status_ServerError `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	SessionId string              `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetCarStatus() *CarStatus {
	if x != nil {
		return x.CarStatus
	}
	return nil
}

func (x *Status) GetServer() *Status_ServerError {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Status) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// StatusResponse information
// Response only to Status message
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      StatusResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=IndustrialPortal.StatusResponse_Type" json:"type,omitempty"`
	SessionId string              `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetType() StatusResponse_Type {
	if x != nil {
		return x.Type
	}
	return StatusResponse_OK
}

func (x *StatusResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// Command message information
// - contains command for a car
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// CarCommand information
	// - full description is in CarStateProtocol.proto
	CarCommand *CarCommand `protobuf:"bytes,1,opt,name=carCommand,proto3" json:"carCommand,omitempty"`
	SessionId  string      `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{6}
}

func (x *Command) GetCarCommand() *CarCommand {
	if x != nil {
		return x.CarCommand
	}
	return nil
}

func (x *Command) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// CommandResponse information
// Response only to Command message
type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      CommandResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=IndustrialPortal.CommandResponse_Type" json:"type,omitempty"`
	SessionId string               `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{7}
}

func (x *CommandResponse) GetType() CommandResponse_Type {
	if x != nil {
		return x.Type
	}
	return CommandResponse_OK
}

func (x *CommandResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

//*
// Server error information
// If the car fins out that the server is down,
// it starts to save stops that car finnished.
// When the server is back again,
// it sends ServerError::Type SERVER_ERROR
// with finished stops, so the server can mark them as done
type Status_ServerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Status_ServerError_Type `protobuf:"varint,1,opt,name=type,proto3,enum=IndustrialPortal.Status_ServerError_Type" json:"type,omitempty"`
	//*
	// All stops, which were finished, when the server has been down
	Stops []*Stop `protobuf:"bytes,2,rep,name=stops,proto3" json:"stops,omitempty"`
}

func (x *Status_ServerError) Reset() {
	*x = Status_ServerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IndustrialPortalProtocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_ServerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_ServerError) ProtoMessage() {}

func (x *Status_ServerError) ProtoReflect() protoreflect.Message {
	mi := &file_IndustrialPortalProtocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_ServerError.ProtoReflect.Descriptor instead.
func (*Status_ServerError) Descriptor() ([]byte, []int) {
	return file_IndustrialPortalProtocol_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Status_ServerError) GetType() Status_ServerError_Type {
	if x != nil {
		return x.Type
	}
	return Status_ServerError_OK
}

func (x *Status_ServerError) GetStops() []*Stop {
	if x != nil {
		return x.Stops
	}
	return nil
}

var File_IndustrialPortalProtocol_proto protoreflect.FileDescriptor

var file_IndustrialPortalProtocol_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x1a, 0x16, 0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x17, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c,
	0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd1,
	0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x55, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x10, 0x01, 0x22, 0xbe, 0x02, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x61, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x61, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x63, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x9c, 0x01,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x79, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x22, 0x65, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7b,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2e,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IndustrialPortalProtocol_proto_rawDescOnce sync.Once
	file_IndustrialPortalProtocol_proto_rawDescData = file_IndustrialPortalProtocol_proto_rawDesc
)

func file_IndustrialPortalProtocol_proto_rawDescGZIP() []byte {
	file_IndustrialPortalProtocol_proto_rawDescOnce.Do(func() {
		file_IndustrialPortalProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_IndustrialPortalProtocol_proto_rawDescData)
	})
	return file_IndustrialPortalProtocol_proto_rawDescData
}

var file_IndustrialPortalProtocol_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_IndustrialPortalProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_IndustrialPortalProtocol_proto_goTypes = []interface{}{
	(ConnectResponse_Type)(0),       // 0: IndustrialPortal.ConnectResponse.Type
	(Status_ServerError_Type)(0),    // 1: IndustrialPortal.Status.ServerError.Type
	(StatusResponse_Type)(0),        // 2: IndustrialPortal.StatusResponse.Type
	(CommandResponse_Type)(0),       // 3: IndustrialPortal.CommandResponse.Type
	(*MessageIndustrialPortal)(nil), // 4: IndustrialPortal.MessageIndustrialPortal
	(*MessageDaemon)(nil),           // 5: IndustrialPortal.MessageDaemon
	(*Connect)(nil),                 // 6: IndustrialPortal.Connect
	(*ConnectResponse)(nil),         // 7: IndustrialPortal.ConnectResponse
	(*Status)(nil),                  // 8: IndustrialPortal.Status
	(*StatusResponse)(nil),          // 9: IndustrialPortal.StatusResponse
	(*Command)(nil),                 // 10: IndustrialPortal.Command
	(*CommandResponse)(nil),         // 11: IndustrialPortal.CommandResponse
	(*Status_ServerError)(nil),      // 12: IndustrialPortal.Status.ServerError
	(*CarStatus)(nil),               // 13: CarStateProtocol.CarStatus
	(*CarCommand)(nil),              // 14: CarStateProtocol.CarCommand
	(*Stop)(nil),                    // 15: CarStateProtocol.Stop
}
var file_IndustrialPortalProtocol_proto_depIdxs = []int32{
	7,  // 0: IndustrialPortal.MessageIndustrialPortal.connectReponse:type_name -> IndustrialPortal.ConnectResponse
	9,  // 1: IndustrialPortal.MessageIndustrialPortal.statusResponse:type_name -> IndustrialPortal.StatusResponse
	10, // 2: IndustrialPortal.MessageIndustrialPortal.command:type_name -> IndustrialPortal.Command
	6,  // 3: IndustrialPortal.MessageDaemon.connect:type_name -> IndustrialPortal.Connect
	8,  // 4: IndustrialPortal.MessageDaemon.status:type_name -> IndustrialPortal.Status
	11, // 5: IndustrialPortal.MessageDaemon.commandResponse:type_name -> IndustrialPortal.CommandResponse
	0,  // 6: IndustrialPortal.ConnectResponse.type:type_name -> IndustrialPortal.ConnectResponse.Type
	13, // 7: IndustrialPortal.Status.carStatus:type_name -> CarStateProtocol.CarStatus
	12, // 8: IndustrialPortal.Status.server:type_name -> IndustrialPortal.Status.ServerError
	2,  // 9: IndustrialPortal.StatusResponse.type:type_name -> IndustrialPortal.StatusResponse.Type
	14, // 10: IndustrialPortal.Command.carCommand:type_name -> CarStateProtocol.CarCommand
	3,  // 11: IndustrialPortal.CommandResponse.type:type_name -> IndustrialPortal.CommandResponse.Type
	1,  // 12: IndustrialPortal.Status.ServerError.type:type_name -> IndustrialPortal.Status.ServerError.Type
	15, // 13: IndustrialPortal.Status.ServerError.stops:type_name -> CarStateProtocol.Stop
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_IndustrialPortalProtocol_proto_init() }
func file_IndustrialPortalProtocol_proto_init() {
	if File_IndustrialPortalProtocol_proto != nil {
		return
	}
	file_CarStateProtocol_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IndustrialPortalProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageIndustrialPortal); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDaemon); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_IndustrialPortalProtocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_ServerError); i {
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
	file_IndustrialPortalProtocol_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MessageIndustrialPortal_ConnectReponse)(nil),
		(*MessageIndustrialPortal_StatusResponse)(nil),
		(*MessageIndustrialPortal_Command)(nil),
	}
	file_IndustrialPortalProtocol_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MessageDaemon_Connect)(nil),
		(*MessageDaemon_Status)(nil),
		(*MessageDaemon_CommandResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_IndustrialPortalProtocol_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IndustrialPortalProtocol_proto_goTypes,
		DependencyIndexes: file_IndustrialPortalProtocol_proto_depIdxs,
		EnumInfos:         file_IndustrialPortalProtocol_proto_enumTypes,
		MessageInfos:      file_IndustrialPortalProtocol_proto_msgTypes,
	}.Build()
	File_IndustrialPortalProtocol_proto = out.File
	file_IndustrialPortalProtocol_proto_rawDesc = nil
	file_IndustrialPortalProtocol_proto_goTypes = nil
	file_IndustrialPortalProtocol_proto_depIdxs = nil
}