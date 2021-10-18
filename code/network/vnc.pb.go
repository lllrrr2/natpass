// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: vnc.proto

package network

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

type VncStatus int32

const (
	VncStatus_unset_st VncStatus = 0
	VncStatus_down     VncStatus = 1
	VncStatus_up       VncStatus = 2
)

// Enum value maps for VncStatus.
var (
	VncStatus_name = map[int32]string{
		0: "unset_st",
		1: "down",
		2: "up",
	}
	VncStatus_value = map[string]int32{
		"unset_st": 0,
		"down":     1,
		"up":       2,
	}
)

func (x VncStatus) Enum() *VncStatus {
	p := new(VncStatus)
	*p = x
	return p
}

func (x VncStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VncStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_vnc_proto_enumTypes[0].Descriptor()
}

func (VncStatus) Type() protoreflect.EnumType {
	return &file_vnc_proto_enumTypes[0]
}

func (x VncStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VncStatus.Descriptor instead.
func (VncStatus) EnumDescriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{0}
}

type VncImageEncoding int32

const (
	VncImage_raw  VncImageEncoding = 0
	VncImage_jpeg VncImageEncoding = 1
	VncImage_png  VncImageEncoding = 2
)

// Enum value maps for VncImageEncoding.
var (
	VncImageEncoding_name = map[int32]string{
		0: "raw",
		1: "jpeg",
		2: "png",
	}
	VncImageEncoding_value = map[string]int32{
		"raw":  0,
		"jpeg": 1,
		"png":  2,
	}
)

func (x VncImageEncoding) Enum() *VncImageEncoding {
	p := new(VncImageEncoding)
	*p = x
	return p
}

func (x VncImageEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VncImageEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_vnc_proto_enumTypes[1].Descriptor()
}

func (VncImageEncoding) Type() protoreflect.EnumType {
	return &file_vnc_proto_enumTypes[1]
}

func (x VncImageEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VncImageEncoding.Descriptor instead.
func (VncImageEncoding) EnumDescriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{1, 0}
}

type VncMouseButton int32

const (
	VncMouse_unset_btn VncMouseButton = 0
	VncMouse_left      VncMouseButton = 1
	VncMouse_middle    VncMouseButton = 2
	VncMouse_right     VncMouseButton = 3
)

// Enum value maps for VncMouseButton.
var (
	VncMouseButton_name = map[int32]string{
		0: "unset_btn",
		1: "left",
		2: "middle",
		3: "right",
	}
	VncMouseButton_value = map[string]int32{
		"unset_btn": 0,
		"left":      1,
		"middle":    2,
		"right":     3,
	}
)

func (x VncMouseButton) Enum() *VncMouseButton {
	p := new(VncMouseButton)
	*p = x
	return p
}

func (x VncMouseButton) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VncMouseButton) Descriptor() protoreflect.EnumDescriptor {
	return file_vnc_proto_enumTypes[2].Descriptor()
}

func (VncMouseButton) Type() protoreflect.EnumType {
	return &file_vnc_proto_enumTypes[2]
}

func (x VncMouseButton) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VncMouseButton.Descriptor instead.
func (VncMouseButton) EnumDescriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{2, 0}
}

type VncControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fps     uint32 `protobuf:"varint,1,opt,name=fps,proto3" json:"fps,omitempty"`
	Quality uint32 `protobuf:"varint,2,opt,name=quality,proto3" json:"quality,omitempty"` // image quality, percent
}

func (x *VncControl) Reset() {
	*x = VncControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vnc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VncControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VncControl) ProtoMessage() {}

func (x *VncControl) ProtoReflect() protoreflect.Message {
	mi := &file_vnc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VncControl.ProtoReflect.Descriptor instead.
func (*VncControl) Descriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{0}
}

func (x *VncControl) GetFps() uint32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *VncControl) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

type VncImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encode VncImageEncoding `protobuf:"varint,1,opt,name=encode,proto3,enum=network.VncImageEncoding" json:"encode,omitempty"`
	Data   []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *VncImage) Reset() {
	*x = VncImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vnc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VncImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VncImage) ProtoMessage() {}

func (x *VncImage) ProtoReflect() protoreflect.Message {
	mi := &file_vnc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VncImage.ProtoReflect.Descriptor instead.
func (*VncImage) Descriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{1}
}

func (x *VncImage) GetEncode() VncImageEncoding {
	if x != nil {
		return x.Encode
	}
	return VncImage_raw
}

func (x *VncImage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type VncMouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type VncStatus      `protobuf:"varint,1,opt,name=type,proto3,enum=network.VncStatus" json:"type,omitempty"`
	Btn  VncMouseButton `protobuf:"varint,2,opt,name=btn,proto3,enum=network.VncMouseButton" json:"btn,omitempty"`
	X    uint32         `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y    uint32         `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *VncMouse) Reset() {
	*x = VncMouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vnc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VncMouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VncMouse) ProtoMessage() {}

func (x *VncMouse) ProtoReflect() protoreflect.Message {
	mi := &file_vnc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VncMouse.ProtoReflect.Descriptor instead.
func (*VncMouse) Descriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{2}
}

func (x *VncMouse) GetType() VncStatus {
	if x != nil {
		return x.Type
	}
	return VncStatus_unset_st
}

func (x *VncMouse) GetBtn() VncMouseButton {
	if x != nil {
		return x.Btn
	}
	return VncMouse_unset_btn
}

func (x *VncMouse) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *VncMouse) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type VncKeyboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type VncStatus `protobuf:"varint,1,opt,name=type,proto3,enum=network.VncStatus" json:"type,omitempty"`
	// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VncKeyboard) Reset() {
	*x = VncKeyboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vnc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VncKeyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VncKeyboard) ProtoMessage() {}

func (x *VncKeyboard) ProtoReflect() protoreflect.Message {
	mi := &file_vnc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VncKeyboard.ProtoReflect.Descriptor instead.
func (*VncKeyboard) Descriptor() ([]byte, []int) {
	return file_vnc_proto_rawDescGZIP(), []int{3}
}

func (x *VncKeyboard) GetType() VncStatus {
	if x != nil {
		return x.Type
	}
	return VncStatus_unset_st
}

func (x *VncKeyboard) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_vnc_proto protoreflect.FileDescriptor

var file_vnc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x76, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x22, 0x39, 0x0a, 0x0b, 0x76, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22,
	0x7c, 0x0a, 0x09, 0x76, 0x6e, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x06,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x07, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x6a, 0x70,
	0x65, 0x67, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x70, 0x6e, 0x67, 0x10, 0x02, 0x22, 0xb7, 0x01,
	0x0a, 0x09, 0x76, 0x6e, 0x63, 0x5f, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x62, 0x74, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e, 0x63, 0x5f,
	0x6d, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x03, 0x62, 0x74,
	0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x38, 0x0a,
	0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x75, 0x6e, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x74, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x10, 0x03, 0x22, 0x49, 0x0a, 0x0c, 0x76, 0x6e, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x76, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x2a, 0x2c, 0x0a, 0x0a, 0x76, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x75, 0x70, 0x10, 0x02,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vnc_proto_rawDescOnce sync.Once
	file_vnc_proto_rawDescData = file_vnc_proto_rawDesc
)

func file_vnc_proto_rawDescGZIP() []byte {
	file_vnc_proto_rawDescOnce.Do(func() {
		file_vnc_proto_rawDescData = protoimpl.X.CompressGZIP(file_vnc_proto_rawDescData)
	})
	return file_vnc_proto_rawDescData
}

var file_vnc_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_vnc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vnc_proto_goTypes = []interface{}{
	(VncStatus)(0),        // 0: network.vnc_status
	(VncImageEncoding)(0), // 1: network.vnc_image.encoding
	(VncMouseButton)(0),   // 2: network.vnc_mouse.button
	(*VncControl)(nil),    // 3: network.vnc_control
	(*VncImage)(nil),      // 4: network.vnc_image
	(*VncMouse)(nil),      // 5: network.vnc_mouse
	(*VncKeyboard)(nil),   // 6: network.vnc_keyboard
}
var file_vnc_proto_depIdxs = []int32{
	1, // 0: network.vnc_image.encode:type_name -> network.vnc_image.encoding
	0, // 1: network.vnc_mouse.type:type_name -> network.vnc_status
	2, // 2: network.vnc_mouse.btn:type_name -> network.vnc_mouse.button
	0, // 3: network.vnc_keyboard.type:type_name -> network.vnc_status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vnc_proto_init() }
func file_vnc_proto_init() {
	if File_vnc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vnc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VncControl); i {
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
		file_vnc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VncImage); i {
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
		file_vnc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VncMouse); i {
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
		file_vnc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VncKeyboard); i {
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
			RawDescriptor: file_vnc_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vnc_proto_goTypes,
		DependencyIndexes: file_vnc_proto_depIdxs,
		EnumInfos:         file_vnc_proto_enumTypes,
		MessageInfos:      file_vnc_proto_msgTypes,
	}.Build()
	File_vnc_proto = out.File
	file_vnc_proto_rawDesc = nil
	file_vnc_proto_goTypes = nil
	file_vnc_proto_depIdxs = nil
}
