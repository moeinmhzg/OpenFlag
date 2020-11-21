// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: pkg/proto/evaluation.proto

package evaluation

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityID      int64             `protobuf:"varint,1,opt,name=entityID,json=entity_id,proto3" json:"entityID,omitempty"`
	EntityType    string            `protobuf:"bytes,2,opt,name=entityType,json=entity_type,proto3" json:"entityType,omitempty"`
	EntityContext map[string]string `protobuf:"bytes,3,rep,name=entityContext,json=entity_context,proto3" json:"entityContext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetEntityID() int64 {
	if x != nil {
		return x.EntityID
	}
	return 0
}

func (x *Entity) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *Entity) GetEntityContext() map[string]string {
	if x != nil {
		return x.EntityContext
	}
	return nil
}

type EvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities          []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	Flags             []string  `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
	SaveContexts      bool      `protobuf:"varint,3,opt,name=save_contexts,proto3" json:"save_contexts,omitempty"`
	UseStoredContexts bool      `protobuf:"varint,4,opt,name=use_stored_contexts,proto3" json:"use_stored_contexts,omitempty"`
}

func (x *EvaluationRequest) Reset() {
	*x = EvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationRequest) ProtoMessage() {}

func (x *EvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationRequest.ProtoReflect.Descriptor instead.
func (*EvaluationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{1}
}

func (x *EvaluationRequest) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *EvaluationRequest) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *EvaluationRequest) GetSaveContexts() bool {
	if x != nil {
		return x.SaveContexts
	}
	return false
}

func (x *EvaluationRequest) GetUseStoredContexts() bool {
	if x != nil {
		return x.UseStoredContexts
	}
	return false
}

type EvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity      *Entity                          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Evaluations []*EvaluationResponse_Evaluation `protobuf:"bytes,2,rep,name=evaluations,proto3" json:"evaluations,omitempty"`
}

func (x *EvaluationResponse) Reset() {
	*x = EvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationResponse) ProtoMessage() {}

func (x *EvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationResponse.ProtoReflect.Descriptor instead.
func (*EvaluationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{2}
}

func (x *EvaluationResponse) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *EvaluationResponse) GetEvaluations() []*EvaluationResponse_Evaluation {
	if x != nil {
		return x.Evaluations
	}
	return nil
}

type EvaluationResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*EvaluationResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *EvaluationResponseList) Reset() {
	*x = EvaluationResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationResponseList) ProtoMessage() {}

func (x *EvaluationResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationResponseList.ProtoReflect.Descriptor instead.
func (*EvaluationResponseList) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{3}
}

func (x *EvaluationResponseList) GetList() []*EvaluationResponse {
	if x != nil {
		return x.List
	}
	return nil
}

type EvaluationResponse_Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VariantKey        string `protobuf:"bytes,1,opt,name=variant_key,proto3" json:"variant_key,omitempty"`
	VariantAttachment []byte `protobuf:"bytes,2,opt,name=variant_attachment,proto3" json:"variant_attachment,omitempty"`
}

func (x *EvaluationResponse_Variant) Reset() {
	*x = EvaluationResponse_Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationResponse_Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationResponse_Variant) ProtoMessage() {}

func (x *EvaluationResponse_Variant) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationResponse_Variant.ProtoReflect.Descriptor instead.
func (*EvaluationResponse_Variant) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EvaluationResponse_Variant) GetVariantKey() string {
	if x != nil {
		return x.VariantKey
	}
	return ""
}

func (x *EvaluationResponse_Variant) GetVariantAttachment() []byte {
	if x != nil {
		return x.VariantAttachment
	}
	return nil
}

type EvaluationResponse_Evaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag    string                      `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Variant *EvaluationResponse_Variant `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
}

func (x *EvaluationResponse_Evaluation) Reset() {
	*x = EvaluationResponse_Evaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_evaluation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationResponse_Evaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationResponse_Evaluation) ProtoMessage() {}

func (x *EvaluationResponse_Evaluation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_evaluation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationResponse_Evaluation.ProtoReflect.Descriptor instead.
func (*EvaluationResponse_Evaluation) Descriptor() ([]byte, []int) {
	return file_pkg_proto_evaluation_proto_rawDescGZIP(), []int{2, 1}
}

func (x *EvaluationResponse_Evaluation) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *EvaluationResponse_Evaluation) GetVariant() *EvaluationResponse_Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

var File_pkg_proto_evaluation_proto protoreflect.FileDescriptor

var file_pkg_proto_evaluation_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a,
	0x40, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x22, 0xce, 0x02, 0x0a, 0x12, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0b, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x5b, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x62, 0x0a, 0x0a, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x12, 0x40, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x16, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x32, 0x5d, 0x0a, 0x0a, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_evaluation_proto_rawDescOnce sync.Once
	file_pkg_proto_evaluation_proto_rawDescData = file_pkg_proto_evaluation_proto_rawDesc
)

func file_pkg_proto_evaluation_proto_rawDescGZIP() []byte {
	file_pkg_proto_evaluation_proto_rawDescOnce.Do(func() {
		file_pkg_proto_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_evaluation_proto_rawDescData)
	})
	return file_pkg_proto_evaluation_proto_rawDescData
}

var file_pkg_proto_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_evaluation_proto_goTypes = []interface{}{
	(*Entity)(nil),                        // 0: evaluation.Entity
	(*EvaluationRequest)(nil),             // 1: evaluation.EvaluationRequest
	(*EvaluationResponse)(nil),            // 2: evaluation.EvaluationResponse
	(*EvaluationResponseList)(nil),        // 3: evaluation.EvaluationResponseList
	nil,                                   // 4: evaluation.Entity.EntityContextEntry
	(*EvaluationResponse_Variant)(nil),    // 5: evaluation.EvaluationResponse.Variant
	(*EvaluationResponse_Evaluation)(nil), // 6: evaluation.EvaluationResponse.Evaluation
}
var file_pkg_proto_evaluation_proto_depIdxs = []int32{
	4, // 0: evaluation.Entity.entityContext:type_name -> evaluation.Entity.EntityContextEntry
	0, // 1: evaluation.EvaluationRequest.entities:type_name -> evaluation.Entity
	0, // 2: evaluation.EvaluationResponse.entity:type_name -> evaluation.Entity
	6, // 3: evaluation.EvaluationResponse.evaluations:type_name -> evaluation.EvaluationResponse.Evaluation
	2, // 4: evaluation.EvaluationResponseList.list:type_name -> evaluation.EvaluationResponse
	5, // 5: evaluation.EvaluationResponse.Evaluation.variant:type_name -> evaluation.EvaluationResponse.Variant
	1, // 6: evaluation.Evaluation.Evaluate:input_type -> evaluation.EvaluationRequest
	3, // 7: evaluation.Evaluation.Evaluate:output_type -> evaluation.EvaluationResponseList
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_proto_evaluation_proto_init() }
func file_pkg_proto_evaluation_proto_init() {
	if File_pkg_proto_evaluation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_evaluation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_pkg_proto_evaluation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationRequest); i {
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
		file_pkg_proto_evaluation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationResponse); i {
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
		file_pkg_proto_evaluation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationResponseList); i {
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
		file_pkg_proto_evaluation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationResponse_Variant); i {
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
		file_pkg_proto_evaluation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationResponse_Evaluation); i {
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
			RawDescriptor: file_pkg_proto_evaluation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_evaluation_proto_goTypes,
		DependencyIndexes: file_pkg_proto_evaluation_proto_depIdxs,
		MessageInfos:      file_pkg_proto_evaluation_proto_msgTypes,
	}.Build()
	File_pkg_proto_evaluation_proto = out.File
	file_pkg_proto_evaluation_proto_rawDesc = nil
	file_pkg_proto_evaluation_proto_goTypes = nil
	file_pkg_proto_evaluation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EvaluationClient is the client API for Evaluation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluationClient interface {
	Evaluate(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponseList, error)
}

type evaluationClient struct {
	cc grpc.ClientConnInterface
}

func NewEvaluationClient(cc grpc.ClientConnInterface) EvaluationClient {
	return &evaluationClient{cc}
}

func (c *evaluationClient) Evaluate(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponseList, error) {
	out := new(EvaluationResponseList)
	err := c.cc.Invoke(ctx, "/evaluation.Evaluation/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluationServer is the server API for Evaluation service.
type EvaluationServer interface {
	Evaluate(context.Context, *EvaluationRequest) (*EvaluationResponseList, error)
}

// UnimplementedEvaluationServer can be embedded to have forward compatible implementations.
type UnimplementedEvaluationServer struct {
}

func (*UnimplementedEvaluationServer) Evaluate(context.Context, *EvaluationRequest) (*EvaluationResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterEvaluationServer(s *grpc.Server, srv EvaluationServer) {
	s.RegisterService(&_Evaluation_serviceDesc, srv)
}

func _Evaluation_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluationServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evaluation.Evaluation/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluationServer).Evaluate(ctx, req.(*EvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Evaluation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evaluation.Evaluation",
	HandlerType: (*EvaluationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _Evaluation_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/evaluation.proto",
}