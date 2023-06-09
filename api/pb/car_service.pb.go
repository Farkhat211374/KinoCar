// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: car_service.proto

package pb

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

type CreateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type CreateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"` // will have a blog id
}

func (x *CreateCarResponse) Reset() {
	*x = CreateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarResponse) ProtoMessage() {}

func (x *CreateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarResponse.ProtoReflect.Descriptor instead.
func (*CreateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type ReadCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *ReadCarRequest) Reset() {
	*x = ReadCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarRequest) ProtoMessage() {}

func (x *ReadCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarRequest.ProtoReflect.Descriptor instead.
func (*ReadCarRequest) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReadCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ReadCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ReadCarResponse) Reset() {
	*x = ReadCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarResponse) ProtoMessage() {}

func (x *ReadCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarResponse.ProtoReflect.Descriptor instead.
func (*ReadCarResponse) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarResponse) Reset() {
	*x = UpdateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarResponse) ProtoMessage() {}

func (x *UpdateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarResponse.ProtoReflect.Descriptor instead.
func (*UpdateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type DeleteCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarRequest) Reset() {
	*x = DeleteCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarRequest) ProtoMessage() {}

func (x *DeleteCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type DeleteCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarResponse) Reset() {
	*x = DeleteCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarResponse) ProtoMessage() {}

func (x *DeleteCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarResponse.ProtoReflect.Descriptor instead.
func (*DeleteCarResponse) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCarResponse) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ListCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCarRequest) Reset() {
	*x = ListCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarRequest) ProtoMessage() {}

func (x *ListCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarRequest.ProtoReflect.Descriptor instead.
func (*ListCarRequest) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{8}
}

type ListCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ListCarResponse) Reset() {
	*x = ListCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarResponse) ProtoMessage() {}

func (x *ListCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarResponse.ProtoReflect.Descriptor instead.
func (*ListCarResponse) Descriptor() ([]byte, []int) {
	return file_car_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

var File_car_service_proto protoreflect.FileDescriptor

var file_car_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x27, 0x0a, 0x0e, 0x52, 0x65, 0x61,
	0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72,
	0x22, 0x2d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22,
	0x2e, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22,
	0x29, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x63,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x32, 0xa4, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x1c, 0x5a,
	0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x72, 0x6b,
	0x68, 0x61, 0x74, 0x2f, 0x4b, 0x69, 0x6e, 0x6f, 0x43, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_car_service_proto_rawDescOnce sync.Once
	file_car_service_proto_rawDescData = file_car_service_proto_rawDesc
)

func file_car_service_proto_rawDescGZIP() []byte {
	file_car_service_proto_rawDescOnce.Do(func() {
		file_car_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_service_proto_rawDescData)
	})
	return file_car_service_proto_rawDescData
}

var file_car_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_car_service_proto_goTypes = []interface{}{
	(*CreateCarRequest)(nil),  // 0: pb.CreateCarRequest
	(*CreateCarResponse)(nil), // 1: pb.CreateCarResponse
	(*ReadCarRequest)(nil),    // 2: pb.ReadCarRequest
	(*ReadCarResponse)(nil),   // 3: pb.ReadCarResponse
	(*UpdateCarRequest)(nil),  // 4: pb.UpdateCarRequest
	(*UpdateCarResponse)(nil), // 5: pb.UpdateCarResponse
	(*DeleteCarRequest)(nil),  // 6: pb.DeleteCarRequest
	(*DeleteCarResponse)(nil), // 7: pb.DeleteCarResponse
	(*ListCarRequest)(nil),    // 8: pb.ListCarRequest
	(*ListCarResponse)(nil),   // 9: pb.ListCarResponse
	(*Car)(nil),               // 10: pb.Car
}
var file_car_service_proto_depIdxs = []int32{
	10, // 0: pb.CreateCarRequest.car:type_name -> pb.Car
	10, // 1: pb.CreateCarResponse.car:type_name -> pb.Car
	10, // 2: pb.ReadCarResponse.car:type_name -> pb.Car
	10, // 3: pb.UpdateCarRequest.car:type_name -> pb.Car
	10, // 4: pb.UpdateCarResponse.car:type_name -> pb.Car
	10, // 5: pb.ListCarResponse.car:type_name -> pb.Car
	0,  // 6: pb.CarService.CreateCar:input_type -> pb.CreateCarRequest
	2,  // 7: pb.CarService.ReadCar:input_type -> pb.ReadCarRequest
	4,  // 8: pb.CarService.UpdateCar:input_type -> pb.UpdateCarRequest
	6,  // 9: pb.CarService.DeleteCar:input_type -> pb.DeleteCarRequest
	8,  // 10: pb.CarService.ListCar:input_type -> pb.ListCarRequest
	1,  // 11: pb.CarService.CreateCar:output_type -> pb.CreateCarResponse
	3,  // 12: pb.CarService.ReadCar:output_type -> pb.ReadCarResponse
	5,  // 13: pb.CarService.UpdateCar:output_type -> pb.UpdateCarResponse
	7,  // 14: pb.CarService.DeleteCar:output_type -> pb.DeleteCarResponse
	9,  // 15: pb.CarService.ListCar:output_type -> pb.ListCarResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_car_service_proto_init() }
func file_car_service_proto_init() {
	if File_car_service_proto != nil {
		return
	}
	file_car_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_car_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarRequest); i {
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
		file_car_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarResponse); i {
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
		file_car_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarRequest); i {
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
		file_car_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarResponse); i {
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
		file_car_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_car_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarResponse); i {
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
		file_car_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarRequest); i {
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
		file_car_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarResponse); i {
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
		file_car_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarRequest); i {
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
		file_car_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarResponse); i {
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
			RawDescriptor: file_car_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_service_proto_goTypes,
		DependencyIndexes: file_car_service_proto_depIdxs,
		MessageInfos:      file_car_service_proto_msgTypes,
	}.Build()
	File_car_service_proto = out.File
	file_car_service_proto_rawDesc = nil
	file_car_service_proto_goTypes = nil
	file_car_service_proto_depIdxs = nil
}
