// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/proto/service.proto

package ride_order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TrackOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Latitude  float32                `protobuf:"fixed32,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32                `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *TrackOrderRequest) Reset() {
	*x = TrackOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackOrderRequest) ProtoMessage() {}

func (x *TrackOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackOrderRequest.ProtoReflect.Descriptor instead.
func (*TrackOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *TrackOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackOrderRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TrackOrderRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *TrackOrderRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type TrackOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TrackOrderResponse) Reset() {
	*x = TrackOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackOrderResponse) ProtoMessage() {}

func (x *TrackOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackOrderResponse.ProtoReflect.Descriptor instead.
func (*TrackOrderResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_service_proto_rawDescGZIP(), []int{1}
}

type GetTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTrackRequest) Reset() {
	*x = GetTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackRequest) ProtoMessage() {}

func (x *GetTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackRequest.ProtoReflect.Descriptor instead.
func (*GetTrackRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrackRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Track []*TrackItem `protobuf:"bytes,1,rep,name=track,proto3" json:"track,omitempty"`
}

func (x *GetTrackResponse) Reset() {
	*x = GetTrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackResponse) ProtoMessage() {}

func (x *GetTrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackResponse.ProtoReflect.Descriptor instead.
func (*GetTrackResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrackResponse) GetTrack() []*TrackItem {
	if x != nil {
		return x.Track
	}
	return nil
}

type TrackItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Latitude  float32                `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32                `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *TrackItem) Reset() {
	*x = TrackItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackItem) ProtoMessage() {}

func (x *TrackItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackItem.ProtoReflect.Descriptor instead.
func (*TrackItem) Descriptor() ([]byte, []int) {
	return file_api_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *TrackItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TrackItem) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *TrackItem) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_api_proto_service_proto protoreflect.FileDescriptor

var file_api_proto_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x22, 0x7f, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x32, 0x74, 0x0a, 0x04, 0x52, 0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_service_proto_rawDescOnce sync.Once
	file_api_proto_service_proto_rawDescData = file_api_proto_service_proto_rawDesc
)

func file_api_proto_service_proto_rawDescGZIP() []byte {
	file_api_proto_service_proto_rawDescOnce.Do(func() {
		file_api_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_service_proto_rawDescData)
	})
	return file_api_proto_service_proto_rawDescData
}

var file_api_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_service_proto_goTypes = []interface{}{
	(*TrackOrderRequest)(nil),     // 0: TrackOrderRequest
	(*TrackOrderResponse)(nil),    // 1: TrackOrderResponse
	(*GetTrackRequest)(nil),       // 2: GetTrackRequest
	(*GetTrackResponse)(nil),      // 3: GetTrackResponse
	(*TrackItem)(nil),             // 4: TrackItem
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_proto_service_proto_depIdxs = []int32{
	5, // 0: TrackOrderRequest.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: GetTrackResponse.track:type_name -> TrackItem
	5, // 2: TrackItem.createdAt:type_name -> google.protobuf.Timestamp
	0, // 3: Ride.TrackOrder:input_type -> TrackOrderRequest
	2, // 4: Ride.GetTrack:input_type -> GetTrackRequest
	1, // 5: Ride.TrackOrder:output_type -> TrackOrderResponse
	3, // 6: Ride.GetTrack:output_type -> GetTrackResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_service_proto_init() }
func file_api_proto_service_proto_init() {
	if File_api_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackOrderRequest); i {
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
		file_api_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackOrderResponse); i {
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
		file_api_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackRequest); i {
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
		file_api_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackResponse); i {
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
		file_api_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackItem); i {
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
			RawDescriptor: file_api_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_service_proto_goTypes,
		DependencyIndexes: file_api_proto_service_proto_depIdxs,
		MessageInfos:      file_api_proto_service_proto_msgTypes,
	}.Build()
	File_api_proto_service_proto = out.File
	file_api_proto_service_proto_rawDesc = nil
	file_api_proto_service_proto_goTypes = nil
	file_api_proto_service_proto_depIdxs = nil
}
