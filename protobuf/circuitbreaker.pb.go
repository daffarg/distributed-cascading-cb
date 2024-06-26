// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: circuitbreaker.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GeneralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method            string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url               string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Header            map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body              []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	RequiringEndpoint string            `protobuf:"bytes,5,opt,name=requiring_endpoint,json=requiringEndpoint,proto3" json:"requiring_endpoint,omitempty"`
	RequiringMethod   string            `protobuf:"bytes,6,opt,name=requiring_method,json=requiringMethod,proto3" json:"requiring_method,omitempty"`
}

func (x *GeneralRequest) Reset() {
	*x = GeneralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralRequest) ProtoMessage() {}

func (x *GeneralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralRequest.ProtoReflect.Descriptor instead.
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GeneralRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GeneralRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GeneralRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GeneralRequest) GetRequiringEndpoint() string {
	if x != nil {
		return x.RequiringEndpoint
	}
	return ""
}

func (x *GeneralRequest) GetRequiringMethod() string {
	if x != nil {
		return x.RequiringMethod
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url               string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header            map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequiringEndpoint string            `protobuf:"bytes,5,opt,name=requiring_endpoint,json=requiringEndpoint,proto3" json:"requiring_endpoint,omitempty"`
	RequiringMethod   string            `protobuf:"bytes,6,opt,name=requiring_method,json=requiringMethod,proto3" json:"requiring_method,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[1]
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
	return file_circuitbreaker_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetRequest) GetRequiringEndpoint() string {
	if x != nil {
		return x.RequiringEndpoint
	}
	return ""
}

func (x *GetRequest) GetRequiringMethod() string {
	if x != nil {
		return x.RequiringMethod
	}
	return ""
}

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url               string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header            map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body              []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	RequiringEndpoint string            `protobuf:"bytes,5,opt,name=requiring_endpoint,json=requiringEndpoint,proto3" json:"requiring_endpoint,omitempty"`
	RequiringMethod   string            `protobuf:"bytes,6,opt,name=requiring_method,json=requiringMethod,proto3" json:"requiring_method,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{2}
}

func (x *PostRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PostRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PostRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *PostRequest) GetRequiringEndpoint() string {
	if x != nil {
		return x.RequiringEndpoint
	}
	return ""
}

func (x *PostRequest) GetRequiringMethod() string {
	if x != nil {
		return x.RequiringMethod
	}
	return ""
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url               string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header            map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body              []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	RequiringEndpoint string            `protobuf:"bytes,5,opt,name=requiring_endpoint,json=requiringEndpoint,proto3" json:"requiring_endpoint,omitempty"`
	RequiringMethod   string            `protobuf:"bytes,6,opt,name=requiring_method,json=requiringMethod,proto3" json:"requiring_method,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{3}
}

func (x *PutRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PutRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PutRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *PutRequest) GetRequiringEndpoint() string {
	if x != nil {
		return x.RequiringEndpoint
	}
	return ""
}

func (x *PutRequest) GetRequiringMethod() string {
	if x != nil {
		return x.RequiringMethod
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url               string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Header            map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequiringEndpoint string            `protobuf:"bytes,5,opt,name=requiring_endpoint,json=requiringEndpoint,proto3" json:"requiring_endpoint,omitempty"`
	RequiringMethod   string            `protobuf:"bytes,6,opt,name=requiring_method,json=requiringMethod,proto3" json:"requiring_method,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DeleteRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DeleteRequest) GetRequiringEndpoint() string {
	if x != nil {
		return x.RequiringEndpoint
	}
	return ""
}

func (x *DeleteRequest) GetRequiringMethod() string {
	if x != nil {
		return x.RequiringMethod
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                    string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode                int32             `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Proto                     string            `protobuf:"bytes,3,opt,name=proto,proto3" json:"proto,omitempty"`
	ProtoMajor                int32             `protobuf:"varint,4,opt,name=proto_major,json=protoMajor,proto3" json:"proto_major,omitempty"`
	ProtoMinor                int32             `protobuf:"varint,5,opt,name=proto_minor,json=protoMinor,proto3" json:"proto_minor,omitempty"`
	Header                    map[string]string `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                      []byte            `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ContentLength             int64             `protobuf:"varint,8,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	IsFromAlternativeEndpoint bool              `protobuf:"varint,9,opt,name=is_from_alternative_endpoint,json=isFromAlternativeEndpoint,proto3" json:"is_from_alternative_endpoint,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_circuitbreaker_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *Response) GetProtoMajor() int32 {
	if x != nil {
		return x.ProtoMajor
	}
	return 0
}

func (x *Response) GetProtoMinor() int32 {
	if x != nil {
		return x.ProtoMinor
	}
	return 0
}

func (x *Response) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *Response) GetIsFromAlternativeEndpoint() bool {
	if x != nil {
		return x.IsFromAlternativeEndpoint
	}
	return false
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint  string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Timeout   uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circuitbreaker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_circuitbreaker_proto_msgTypes[6]
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
	return file_circuitbreaker_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Status) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Status) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_circuitbreaker_proto protoreflect.FileDescriptor

var file_circuitbreaker_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02,
	0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3c, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2d, 0x0a, 0x12,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xed, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x39, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x81, 0x02, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf3, 0x01, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x3b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x8a, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12,
	0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x3f, 0x0a, 0x1c, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x6c,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69, 0x73, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x74,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x32, 0x9f, 0x02, 0x0a, 0x0e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x50, 0x75,
	0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_circuitbreaker_proto_rawDescOnce sync.Once
	file_circuitbreaker_proto_rawDescData = file_circuitbreaker_proto_rawDesc
)

func file_circuitbreaker_proto_rawDescGZIP() []byte {
	file_circuitbreaker_proto_rawDescOnce.Do(func() {
		file_circuitbreaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_circuitbreaker_proto_rawDescData)
	})
	return file_circuitbreaker_proto_rawDescData
}

var file_circuitbreaker_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_circuitbreaker_proto_goTypes = []interface{}{
	(*GeneralRequest)(nil), // 0: protobuf.GeneralRequest
	(*GetRequest)(nil),     // 1: protobuf.GetRequest
	(*PostRequest)(nil),    // 2: protobuf.PostRequest
	(*PutRequest)(nil),     // 3: protobuf.PutRequest
	(*DeleteRequest)(nil),  // 4: protobuf.DeleteRequest
	(*Response)(nil),       // 5: protobuf.Response
	(*Status)(nil),         // 6: protobuf.Status
	nil,                    // 7: protobuf.GeneralRequest.HeaderEntry
	nil,                    // 8: protobuf.GetRequest.HeaderEntry
	nil,                    // 9: protobuf.PostRequest.HeaderEntry
	nil,                    // 10: protobuf.PutRequest.HeaderEntry
	nil,                    // 11: protobuf.DeleteRequest.HeaderEntry
	nil,                    // 12: protobuf.Response.HeaderEntry
}
var file_circuitbreaker_proto_depIdxs = []int32{
	7,  // 0: protobuf.GeneralRequest.header:type_name -> protobuf.GeneralRequest.HeaderEntry
	8,  // 1: protobuf.GetRequest.header:type_name -> protobuf.GetRequest.HeaderEntry
	9,  // 2: protobuf.PostRequest.header:type_name -> protobuf.PostRequest.HeaderEntry
	10, // 3: protobuf.PutRequest.header:type_name -> protobuf.PutRequest.HeaderEntry
	11, // 4: protobuf.DeleteRequest.header:type_name -> protobuf.DeleteRequest.HeaderEntry
	12, // 5: protobuf.Response.header:type_name -> protobuf.Response.HeaderEntry
	0,  // 6: protobuf.CircuitBreaker.General:input_type -> protobuf.GeneralRequest
	1,  // 7: protobuf.CircuitBreaker.Get:input_type -> protobuf.GetRequest
	2,  // 8: protobuf.CircuitBreaker.Post:input_type -> protobuf.PostRequest
	3,  // 9: protobuf.CircuitBreaker.Put:input_type -> protobuf.PutRequest
	4,  // 10: protobuf.CircuitBreaker.Delete:input_type -> protobuf.DeleteRequest
	5,  // 11: protobuf.CircuitBreaker.General:output_type -> protobuf.Response
	5,  // 12: protobuf.CircuitBreaker.Get:output_type -> protobuf.Response
	5,  // 13: protobuf.CircuitBreaker.Post:output_type -> protobuf.Response
	5,  // 14: protobuf.CircuitBreaker.Put:output_type -> protobuf.Response
	5,  // 15: protobuf.CircuitBreaker.Delete:output_type -> protobuf.Response
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_circuitbreaker_proto_init() }
func file_circuitbreaker_proto_init() {
	if File_circuitbreaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_circuitbreaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralRequest); i {
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
		file_circuitbreaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_circuitbreaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequest); i {
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
		file_circuitbreaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_circuitbreaker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_circuitbreaker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_circuitbreaker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_circuitbreaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_circuitbreaker_proto_goTypes,
		DependencyIndexes: file_circuitbreaker_proto_depIdxs,
		MessageInfos:      file_circuitbreaker_proto_msgTypes,
	}.Build()
	File_circuitbreaker_proto = out.File
	file_circuitbreaker_proto_rawDesc = nil
	file_circuitbreaker_proto_goTypes = nil
	file_circuitbreaker_proto_depIdxs = nil
}
