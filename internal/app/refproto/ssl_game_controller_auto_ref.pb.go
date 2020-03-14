// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_auto_ref.proto

package refproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// AutoRefRegistration is the first message that a client must send to the controller to identify itself
type AutoRefRegistration struct {
	// identifier is a unique name of the client
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature            *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AutoRefRegistration) Reset()         { *m = AutoRefRegistration{} }
func (m *AutoRefRegistration) String() string { return proto.CompactTextString(m) }
func (*AutoRefRegistration) ProtoMessage()    {}
func (*AutoRefRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{0}
}

func (m *AutoRefRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefRegistration.Unmarshal(m, b)
}
func (m *AutoRefRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefRegistration.Marshal(b, m, deterministic)
}
func (m *AutoRefRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefRegistration.Merge(m, src)
}
func (m *AutoRefRegistration) XXX_Size() int {
	return xxx_messageInfo_AutoRefRegistration.Size(m)
}
func (m *AutoRefRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefRegistration proto.InternalMessageInfo

func (m *AutoRefRegistration) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *AutoRefRegistration) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// AutoRefToController is the wrapper message for all subsequent messages from the autoRef to the controller
type AutoRefToController struct {
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// game_event is an optional event that the autoRef detected during the game
	GameEvent *GameEvent `protobuf:"bytes,2,opt,name=game_event,json=gameEvent" json:"game_event,omitempty"`
	// metadata on the current game state
	GameStateMetadata    *GameStateMetadata `protobuf:"bytes,5,opt,name=game_state_metadata,json=gameStateMetadata" json:"game_state_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AutoRefToController) Reset()         { *m = AutoRefToController{} }
func (m *AutoRefToController) String() string { return proto.CompactTextString(m) }
func (*AutoRefToController) ProtoMessage()    {}
func (*AutoRefToController) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{1}
}

func (m *AutoRefToController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefToController.Unmarshal(m, b)
}
func (m *AutoRefToController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefToController.Marshal(b, m, deterministic)
}
func (m *AutoRefToController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefToController.Merge(m, src)
}
func (m *AutoRefToController) XXX_Size() int {
	return xxx_messageInfo_AutoRefToController.Size(m)
}
func (m *AutoRefToController) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefToController.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefToController proto.InternalMessageInfo

func (m *AutoRefToController) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AutoRefToController) GetGameEvent() *GameEvent {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func (m *AutoRefToController) GetGameStateMetadata() *GameStateMetadata {
	if m != nil {
		return m.GameStateMetadata
	}
	return nil
}

type GameStateMetadata struct {
	// if current state is prepare_kickoff, prepare_penalty, stop or ball_placement:
	// true: all conditions are met to continue from the current state
	// false: at least one condition not met yet
	// not present: not supported by the autoRef
	// else:
	// undefined
	ReadyToContinue *bool `protobuf:"varint,1,opt,name=ready_to_continue,json=readyToContinue" json:"ready_to_continue,omitempty"`
	// number of yellow robots currently visible on the field
	NumberOfRobotsYellow *uint32 `protobuf:"varint,2,opt,name=number_of_robots_yellow,json=numberOfRobotsYellow" json:"number_of_robots_yellow,omitempty"`
	// number of blue robots currently visible on the field
	NumberOfRobotsBlue   *uint32  `protobuf:"varint,3,opt,name=number_of_robots_blue,json=numberOfRobotsBlue" json:"number_of_robots_blue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStateMetadata) Reset()         { *m = GameStateMetadata{} }
func (m *GameStateMetadata) String() string { return proto.CompactTextString(m) }
func (*GameStateMetadata) ProtoMessage()    {}
func (*GameStateMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2}
}

func (m *GameStateMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStateMetadata.Unmarshal(m, b)
}
func (m *GameStateMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStateMetadata.Marshal(b, m, deterministic)
}
func (m *GameStateMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStateMetadata.Merge(m, src)
}
func (m *GameStateMetadata) XXX_Size() int {
	return xxx_messageInfo_GameStateMetadata.Size(m)
}
func (m *GameStateMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStateMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GameStateMetadata proto.InternalMessageInfo

func (m *GameStateMetadata) GetReadyToContinue() bool {
	if m != nil && m.ReadyToContinue != nil {
		return *m.ReadyToContinue
	}
	return false
}

func (m *GameStateMetadata) GetNumberOfRobotsYellow() uint32 {
	if m != nil && m.NumberOfRobotsYellow != nil {
		return *m.NumberOfRobotsYellow
	}
	return 0
}

func (m *GameStateMetadata) GetNumberOfRobotsBlue() uint32 {
	if m != nil && m.NumberOfRobotsBlue != nil {
		return *m.NumberOfRobotsBlue
	}
	return 0
}

// ControllerToAutoRef is the wrapper message for all messages from controller to autoRef
type ControllerToAutoRef struct {
	// Types that are valid to be assigned to Msg:
	//	*ControllerToAutoRef_ControllerReply
	Msg                  isControllerToAutoRef_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ControllerToAutoRef) Reset()         { *m = ControllerToAutoRef{} }
func (m *ControllerToAutoRef) String() string { return proto.CompactTextString(m) }
func (*ControllerToAutoRef) ProtoMessage()    {}
func (*ControllerToAutoRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{3}
}

func (m *ControllerToAutoRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerToAutoRef.Unmarshal(m, b)
}
func (m *ControllerToAutoRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerToAutoRef.Marshal(b, m, deterministic)
}
func (m *ControllerToAutoRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerToAutoRef.Merge(m, src)
}
func (m *ControllerToAutoRef) XXX_Size() int {
	return xxx_messageInfo_ControllerToAutoRef.Size(m)
}
func (m *ControllerToAutoRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerToAutoRef.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerToAutoRef proto.InternalMessageInfo

type isControllerToAutoRef_Msg interface {
	isControllerToAutoRef_Msg()
}

type ControllerToAutoRef_ControllerReply struct {
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply,oneof"`
}

func (*ControllerToAutoRef_ControllerReply) isControllerToAutoRef_Msg() {}

func (m *ControllerToAutoRef) GetMsg() isControllerToAutoRef_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ControllerToAutoRef) GetControllerReply() *ControllerReply {
	if x, ok := m.GetMsg().(*ControllerToAutoRef_ControllerReply); ok {
		return x.ControllerReply
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ControllerToAutoRef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ControllerToAutoRef_ControllerReply)(nil),
	}
}

func init() {
	proto.RegisterType((*AutoRefRegistration)(nil), "AutoRefRegistration")
	proto.RegisterType((*AutoRefToController)(nil), "AutoRefToController")
	proto.RegisterType((*GameStateMetadata)(nil), "GameStateMetadata")
	proto.RegisterType((*ControllerToAutoRef)(nil), "ControllerToAutoRef")
}

func init() {
	proto.RegisterFile("ssl_game_controller_auto_ref.proto", fileDescriptor_2fbb3ba3bab9727c)
}

var fileDescriptor_2fbb3ba3bab9727c = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0xdb, 0x40,
	0x10, 0x85, 0x2b, 0x5b, 0x06, 0x7b, 0x4a, 0xb1, 0xbc, 0x76, 0xa9, 0xe8, 0xa1, 0x18, 0x9d, 0xdc,
	0x1e, 0x0c, 0x2d, 0xf4, 0xd8, 0x43, 0x5d, 0x4a, 0x82, 0x21, 0x04, 0xd6, 0xbe, 0x84, 0x1c, 0x96,
	0xb5, 0x3d, 0x12, 0x82, 0xd5, 0x8e, 0x59, 0xad, 0x12, 0xfc, 0x93, 0xf2, 0x17, 0xf2, 0xeb, 0x82,
	0x56, 0xb2, 0x64, 0xc7, 0x21, 0x37, 0xf1, 0xbd, 0xf7, 0x76, 0x66, 0x1e, 0x82, 0x28, 0xcf, 0x95,
	0x48, 0x64, 0x86, 0x62, 0x4b, 0xda, 0x1a, 0x52, 0x0a, 0x8d, 0x90, 0x85, 0x25, 0x61, 0x30, 0x9e,
	0xef, 0x0d, 0x59, 0xfa, 0x3a, 0x7d, 0xcb, 0xb3, 0xa5, 0x2c, 0x23, 0x5d, 0x3b, 0x26, 0x8d, 0x03,
	0x1f, 0x50, 0xdb, 0x8a, 0x46, 0x02, 0xc6, 0x7f, 0x0b, 0x4b, 0x1c, 0x63, 0x8e, 0x49, 0x9a, 0x5b,
	0x23, 0x6d, 0x4a, 0x9a, 0x7d, 0x03, 0x48, 0x77, 0xa8, 0x6d, 0x1a, 0xa7, 0x68, 0x42, 0x6f, 0xda,
	0x99, 0x0d, 0xf8, 0x09, 0x61, 0x33, 0x18, 0xe4, 0x69, 0xa2, 0xa5, 0x2d, 0x0c, 0x86, 0x9d, 0xa9,
	0x37, 0xfb, 0xf8, 0x0b, 0xe6, 0xab, 0x23, 0xe1, 0xad, 0x18, 0x3d, 0x7b, 0xcd, 0x84, 0x35, 0xfd,
	0x6b, 0x76, 0x3b, 0x7f, 0xc1, 0x7b, 0xe7, 0x05, 0xf6, 0x1d, 0xa0, 0x5d, 0xbb, 0x19, 0x76, 0x25,
	0x33, 0xfc, 0x5f, 0x12, 0x3e, 0x48, 0x8e, 0x9f, 0x6c, 0x01, 0x63, 0x67, 0xcd, 0xad, 0xb4, 0x28,
	0x32, 0xb4, 0x72, 0x27, 0xad, 0x0c, 0x7b, 0x2e, 0xc3, 0x5c, 0x66, 0x55, 0x4a, 0x37, 0xb5, 0xc2,
	0x47, 0xc9, 0x6b, 0xb4, 0xf4, 0xfb, 0xdd, 0xc0, 0x5f, 0xfa, 0x7d, 0x3f, 0xe8, 0x45, 0x4f, 0x1e,
	0x8c, 0x2e, 0x42, 0xec, 0x07, 0x8c, 0x0c, 0xca, 0xdd, 0x41, 0x58, 0x72, 0x6d, 0xa7, 0xba, 0xa8,
	0x4e, 0xe8, 0xf3, 0xa1, 0x13, 0xaa, 0x43, 0x4b, 0xcc, 0x7e, 0xc3, 0x17, 0x5d, 0x64, 0x1b, 0x34,
	0x82, 0x62, 0x61, 0x68, 0x43, 0x36, 0x17, 0x07, 0x54, 0x8a, 0x1e, 0xdd, 0x25, 0x9f, 0xf8, 0xa4,
	0x92, 0x6f, 0x63, 0xee, 0xc4, 0x3b, 0xa7, 0xb1, 0x9f, 0xf0, 0xf9, 0x22, 0xb6, 0x51, 0x05, 0x86,
	0x5d, 0x17, 0x62, 0xe7, 0xa1, 0x85, 0x2a, 0x30, 0xba, 0x87, 0x71, 0x5b, 0xef, 0x9a, 0xea, 0xce,
	0xd9, 0x1f, 0x08, 0x4e, 0xfe, 0x08, 0x83, 0x7b, 0x75, 0xa8, 0xeb, 0x0e, 0xe6, 0xad, 0x9f, 0x97,
	0xfc, 0xfa, 0x03, 0x1f, 0x6e, 0xcf, 0xd1, 0xa2, 0x07, 0xdd, 0x2c, 0x4f, 0x5e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x49, 0xad, 0x39, 0x85, 0x83, 0x02, 0x00, 0x00,
}