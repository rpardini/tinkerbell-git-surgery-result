// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: template/template.proto

package template

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[0]
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
	return file_template_template_proto_rawDescGZIP(), []int{0}
}

type WorkflowTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Data      string                 `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WorkflowTemplate) Reset() {
	*x = WorkflowTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTemplate) ProtoMessage() {}

func (x *WorkflowTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTemplate.ProtoReflect.Descriptor instead.
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowTemplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *WorkflowTemplate) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FilterBy:
	//	*ListRequest_Name
	FilterBy isListRequest_FilterBy `protobuf_oneof:"filter_by"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{4}
}

func (m *ListRequest) GetFilterBy() isListRequest_FilterBy {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (x *ListRequest) GetName() string {
	if x, ok := x.GetFilterBy().(*ListRequest_Name); ok {
		return x.Name
	}
	return ""
}

type isListRequest_FilterBy interface {
	isListRequest_FilterBy()
}

type ListRequest_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

func (*ListRequest_Name) isListRequest_FilterBy() {}

var File_template_template_proto protoreflect.FileDescriptor

var file_template_template_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x81, 0x02,
	0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x30, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x32, 0x9a, 0x06, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9f,
	0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x36,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x97, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62,
	0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x37, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x30, 0x01, 0x12, 0x81, 0x01, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x31, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_template_proto_rawDescOnce sync.Once
	file_template_template_proto_rawDescData = file_template_template_proto_rawDesc
)

func file_template_template_proto_rawDescGZIP() []byte {
	file_template_template_proto_rawDescOnce.Do(func() {
		file_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_template_proto_rawDescData)
	})
	return file_template_template_proto_rawDescData
}

var file_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_template_template_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: github.com.tinkerbell.tink.protos.template.Empty
	(*WorkflowTemplate)(nil),      // 1: github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	(*CreateResponse)(nil),        // 2: github.com.tinkerbell.tink.protos.template.CreateResponse
	(*GetRequest)(nil),            // 3: github.com.tinkerbell.tink.protos.template.GetRequest
	(*ListRequest)(nil),           // 4: github.com.tinkerbell.tink.protos.template.ListRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_template_template_proto_depIdxs = []int32{
	5, // 0: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.updated_at:type_name -> google.protobuf.Timestamp
	5, // 2: github.com.tinkerbell.tink.protos.template.WorkflowTemplate.deleted_at:type_name -> google.protobuf.Timestamp
	1, // 3: github.com.tinkerbell.tink.protos.template.TemplateService.CreateTemplate:input_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	3, // 4: github.com.tinkerbell.tink.protos.template.TemplateService.GetTemplate:input_type -> github.com.tinkerbell.tink.protos.template.GetRequest
	3, // 5: github.com.tinkerbell.tink.protos.template.TemplateService.DeleteTemplate:input_type -> github.com.tinkerbell.tink.protos.template.GetRequest
	4, // 6: github.com.tinkerbell.tink.protos.template.TemplateService.ListTemplates:input_type -> github.com.tinkerbell.tink.protos.template.ListRequest
	1, // 7: github.com.tinkerbell.tink.protos.template.TemplateService.UpdateTemplate:input_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	2, // 8: github.com.tinkerbell.tink.protos.template.TemplateService.CreateTemplate:output_type -> github.com.tinkerbell.tink.protos.template.CreateResponse
	1, // 9: github.com.tinkerbell.tink.protos.template.TemplateService.GetTemplate:output_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	0, // 10: github.com.tinkerbell.tink.protos.template.TemplateService.DeleteTemplate:output_type -> github.com.tinkerbell.tink.protos.template.Empty
	1, // 11: github.com.tinkerbell.tink.protos.template.TemplateService.ListTemplates:output_type -> github.com.tinkerbell.tink.protos.template.WorkflowTemplate
	0, // 12: github.com.tinkerbell.tink.protos.template.TemplateService.UpdateTemplate:output_type -> github.com.tinkerbell.tink.protos.template.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_template_template_proto_init() }
func file_template_template_proto_init() {
	if File_template_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_template_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowTemplate); i {
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
		file_template_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_template_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_template_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
	file_template_template_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ListRequest_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_template_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_template_proto_goTypes,
		DependencyIndexes: file_template_template_proto_depIdxs,
		MessageInfos:      file_template_template_proto_msgTypes,
	}.Build()
	File_template_template_proto = out.File
	file_template_template_proto_rawDesc = nil
	file_template_template_proto_goTypes = nil
	file_template_template_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateServiceClient interface {
	CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error)
	GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error)
	ListTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TemplateService_ListTemplatesClient, error)
	UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) ListTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TemplateService_ListTemplatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TemplateService_serviceDesc.Streams[0], "/github.com.tinkerbell.tink.protos.template.TemplateService/ListTemplates", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateServiceListTemplatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TemplateService_ListTemplatesClient interface {
	Recv() (*WorkflowTemplate, error)
	grpc.ClientStream
}

type templateServiceListTemplatesClient struct {
	grpc.ClientStream
}

func (x *templateServiceListTemplatesClient) Recv() (*WorkflowTemplate, error) {
	m := new(WorkflowTemplate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *templateServiceClient) UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
type TemplateServiceServer interface {
	CreateTemplate(context.Context, *WorkflowTemplate) (*CreateResponse, error)
	GetTemplate(context.Context, *GetRequest) (*WorkflowTemplate, error)
	DeleteTemplate(context.Context, *GetRequest) (*Empty, error)
	ListTemplates(*ListRequest, TemplateService_ListTemplatesServer) error
	UpdateTemplate(context.Context, *WorkflowTemplate) (*Empty, error)
}

// UnimplementedTemplateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct {
}

func (*UnimplementedTemplateServiceServer) CreateTemplate(context.Context, *WorkflowTemplate) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (*UnimplementedTemplateServiceServer) GetTemplate(context.Context, *GetRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (*UnimplementedTemplateServiceServer) DeleteTemplate(context.Context, *GetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (*UnimplementedTemplateServiceServer) ListTemplates(*ListRequest, TemplateService_ListTemplatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}
func (*UnimplementedTemplateServiceServer) UpdateTemplate(context.Context, *WorkflowTemplate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}

func RegisterTemplateServiceServer(s *grpc.Server, srv TemplateServiceServer) {
	s.RegisterService(&_TemplateService_serviceDesc, srv)
}

func _TemplateService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).DeleteTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_ListTemplates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemplateServiceServer).ListTemplates(m, &templateServiceListTemplatesServer{stream})
}

type TemplateService_ListTemplatesServer interface {
	Send(*WorkflowTemplate) error
	grpc.ServerStream
}

type templateServiceListTemplatesServer struct {
	grpc.ServerStream
}

func (x *templateServiceListTemplatesServer) Send(m *WorkflowTemplate) error {
	return x.ServerStream.SendMsg(m)
}

func _TemplateService_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).UpdateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.tinkerbell.tink.protos.template.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _TemplateService_CreateTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _TemplateService_GetTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _TemplateService_DeleteTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _TemplateService_UpdateTemplate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTemplates",
			Handler:       _TemplateService_ListTemplates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "template/template.proto",
}
