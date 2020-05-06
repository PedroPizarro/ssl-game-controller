// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_engine_config.proto

package engine

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

type Config_Behavior int32

const (
	Config_BEHAVIOR_UNKNOWN         Config_Behavior = 0
	Config_BEHAVIOR_ACCEPT          Config_Behavior = 1
	Config_BEHAVIOR_ACCEPT_MAJORITY Config_Behavior = 2
	Config_BEHAVIOR_PROPOSE_ONLY    Config_Behavior = 3
	Config_BEHAVIOR_LOG             Config_Behavior = 4
	Config_BEHAVIOR_IGNORE          Config_Behavior = 5
)

var Config_Behavior_name = map[int32]string{
	0: "BEHAVIOR_UNKNOWN",
	1: "BEHAVIOR_ACCEPT",
	2: "BEHAVIOR_ACCEPT_MAJORITY",
	3: "BEHAVIOR_PROPOSE_ONLY",
	4: "BEHAVIOR_LOG",
	5: "BEHAVIOR_IGNORE",
}

var Config_Behavior_value = map[string]int32{
	"BEHAVIOR_UNKNOWN":         0,
	"BEHAVIOR_ACCEPT":          1,
	"BEHAVIOR_ACCEPT_MAJORITY": 2,
	"BEHAVIOR_PROPOSE_ONLY":    3,
	"BEHAVIOR_LOG":             4,
	"BEHAVIOR_IGNORE":          5,
}

func (x Config_Behavior) Enum() *Config_Behavior {
	p := new(Config_Behavior)
	*p = x
	return p
}

func (x Config_Behavior) String() string {
	return proto.EnumName(Config_Behavior_name, int32(x))
}

func (x *Config_Behavior) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Config_Behavior_value, data, "Config_Behavior")
	if err != nil {
		return err
	}
	*x = Config_Behavior(value)
	return nil
}

func (Config_Behavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_297fb0f3751cdb09, []int{0, 0}
}

type AutoRefConfig_Behavior int32

const (
	AutoRefConfig_BEHAVIOR_UNKNOWN AutoRefConfig_Behavior = 0
	AutoRefConfig_BEHAVIOR_ACCEPT  AutoRefConfig_Behavior = 1
	AutoRefConfig_BEHAVIOR_LOG     AutoRefConfig_Behavior = 2
	AutoRefConfig_BEHAVIOR_IGNORE  AutoRefConfig_Behavior = 3
)

var AutoRefConfig_Behavior_name = map[int32]string{
	0: "BEHAVIOR_UNKNOWN",
	1: "BEHAVIOR_ACCEPT",
	2: "BEHAVIOR_LOG",
	3: "BEHAVIOR_IGNORE",
}

var AutoRefConfig_Behavior_value = map[string]int32{
	"BEHAVIOR_UNKNOWN": 0,
	"BEHAVIOR_ACCEPT":  1,
	"BEHAVIOR_LOG":     2,
	"BEHAVIOR_IGNORE":  3,
}

func (x AutoRefConfig_Behavior) Enum() *AutoRefConfig_Behavior {
	p := new(AutoRefConfig_Behavior)
	*p = x
	return p
}

func (x AutoRefConfig_Behavior) String() string {
	return proto.EnumName(AutoRefConfig_Behavior_name, int32(x))
}

func (x *AutoRefConfig_Behavior) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AutoRefConfig_Behavior_value, data, "AutoRefConfig_Behavior")
	if err != nil {
		return err
	}
	*x = AutoRefConfig_Behavior(value)
	return nil
}

func (AutoRefConfig_Behavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_297fb0f3751cdb09, []int{1, 0}
}

type Config struct {
	GameEventBehavior    map[string]Config_Behavior `protobuf:"bytes,3,rep,name=game_event_behavior,json=gameEventBehavior" json:"game_event_behavior,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=Config_Behavior"`
	AutoRefConfigs       map[string]*AutoRefConfig  `protobuf:"bytes,4,rep,name=auto_ref_configs,json=autoRefConfigs" json:"auto_ref_configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_297fb0f3751cdb09, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetGameEventBehavior() map[string]Config_Behavior {
	if m != nil {
		return m.GameEventBehavior
	}
	return nil
}

func (m *Config) GetAutoRefConfigs() map[string]*AutoRefConfig {
	if m != nil {
		return m.AutoRefConfigs
	}
	return nil
}

type AutoRefConfig struct {
	GameEventBehavior    map[string]AutoRefConfig_Behavior `protobuf:"bytes,1,rep,name=game_event_behavior,json=gameEventBehavior" json:"game_event_behavior,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=AutoRefConfig_Behavior"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AutoRefConfig) Reset()         { *m = AutoRefConfig{} }
func (m *AutoRefConfig) String() string { return proto.CompactTextString(m) }
func (*AutoRefConfig) ProtoMessage()    {}
func (*AutoRefConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_297fb0f3751cdb09, []int{1}
}

func (m *AutoRefConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefConfig.Unmarshal(m, b)
}
func (m *AutoRefConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefConfig.Marshal(b, m, deterministic)
}
func (m *AutoRefConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefConfig.Merge(m, src)
}
func (m *AutoRefConfig) XXX_Size() int {
	return xxx_messageInfo_AutoRefConfig.Size(m)
}
func (m *AutoRefConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefConfig proto.InternalMessageInfo

func (m *AutoRefConfig) GetGameEventBehavior() map[string]AutoRefConfig_Behavior {
	if m != nil {
		return m.GameEventBehavior
	}
	return nil
}

func init() {
	proto.RegisterEnum("Config_Behavior", Config_Behavior_name, Config_Behavior_value)
	proto.RegisterEnum("AutoRefConfig_Behavior", AutoRefConfig_Behavior_name, AutoRefConfig_Behavior_value)
	proto.RegisterType((*Config)(nil), "Config")
	proto.RegisterMapType((map[string]*AutoRefConfig)(nil), "Config.AutoRefConfigsEntry")
	proto.RegisterMapType((map[string]Config_Behavior)(nil), "Config.GameEventBehaviorEntry")
	proto.RegisterType((*AutoRefConfig)(nil), "AutoRefConfig")
	proto.RegisterMapType((map[string]AutoRefConfig_Behavior)(nil), "AutoRefConfig.GameEventBehaviorEntry")
}

func init() {
	proto.RegisterFile("ssl_gc_engine_config.proto", fileDescriptor_297fb0f3751cdb09)
}

var fileDescriptor_297fb0f3751cdb09 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x8e, 0x93, 0x40,
	0x14, 0xc6, 0x05, 0x76, 0x8d, 0xce, 0x6a, 0x1d, 0xa7, 0xfe, 0xc1, 0x6a, 0x4c, 0xd3, 0xa8, 0xe9,
	0x0d, 0x90, 0xf4, 0xca, 0x78, 0x61, 0xa4, 0x84, 0xd4, 0x6a, 0x65, 0xea, 0x74, 0x77, 0xcd, 0x1a,
	0xcd, 0x64, 0x4a, 0xa6, 0x2c, 0x91, 0xce, 0x10, 0x18, 0x9a, 0xec, 0x83, 0xf8, 0x04, 0xbe, 0x99,
	0x4f, 0x62, 0x5a, 0xb6, 0xa4, 0x10, 0x2e, 0xd4, 0x3b, 0xf8, 0xce, 0xe1, 0xf7, 0x9d, 0x73, 0xbe,
	0x00, 0x7a, 0x79, 0x9e, 0xd0, 0x28, 0xa4, 0x5c, 0x44, 0xb1, 0xe0, 0x34, 0x94, 0x62, 0x15, 0x47,
	0x76, 0x9a, 0x49, 0x25, 0x07, 0xbf, 0x0d, 0x70, 0xd3, 0xdb, 0x09, 0x28, 0x00, 0xdd, 0x88, 0xad,
	0x39, 0xe5, 0x1b, 0x2e, 0x14, 0x5d, 0xf2, 0x4b, 0xb6, 0x89, 0x65, 0x66, 0x1a, 0x7d, 0x63, 0x78,
	0x32, 0x7a, 0x6e, 0x97, 0x5d, 0xf6, 0x84, 0xad, 0xb9, 0xbf, 0xed, 0x18, 0x5f, 0x37, 0xf8, 0x42,
	0x65, 0x57, 0xe4, 0x7e, 0xd4, 0xd4, 0x91, 0x0f, 0x20, 0x2b, 0x94, 0xa4, 0x19, 0x5f, 0x5d, 0x7b,
	0xe6, 0xe6, 0xd1, 0x0e, 0xf6, 0x74, 0x0f, 0x73, 0x0b, 0x25, 0x09, 0x5f, 0x95, 0x6f, 0x79, 0x49,
	0xea, 0xb0, 0x9a, 0xd8, 0x3b, 0x07, 0x8f, 0xda, 0x3d, 0x11, 0x04, 0xc6, 0x0f, 0x7e, 0x65, 0x6a,
	0x7d, 0x6d, 0x78, 0x9b, 0x6c, 0x1f, 0xd1, 0x2b, 0x70, 0xbc, 0x61, 0x49, 0xc1, 0x4d, 0xbd, 0xaf,
	0x0d, 0x3b, 0x23, 0xb8, 0xf7, 0xd9, 0x7f, 0x47, 0xca, 0xf2, 0x1b, 0xfd, 0xb5, 0xd6, 0xfb, 0x0c,
	0xba, 0x2d, 0xf6, 0x2d, 0xd0, 0x17, 0x87, 0xd0, 0x93, 0x51, 0xa7, 0x3e, 0xf5, 0x01, 0x72, 0xf0,
	0x53, 0x03, 0xb7, 0xaa, 0xf5, 0x1f, 0x00, 0x38, 0xf6, 0xdf, 0xbb, 0xe7, 0x53, 0x4c, 0xe8, 0x59,
	0xf0, 0x31, 0xc0, 0x5f, 0x02, 0x78, 0x03, 0x75, 0xc1, 0xbd, 0x4a, 0x75, 0x3d, 0xcf, 0x9f, 0x9f,
	0x42, 0x0d, 0x3d, 0x03, 0x66, 0x43, 0xa4, 0x9f, 0xdc, 0x0f, 0x98, 0x4c, 0x4f, 0x2f, 0xa0, 0x8e,
	0x9e, 0x80, 0x87, 0x55, 0x75, 0x4e, 0xf0, 0x1c, 0x2f, 0x7c, 0x8a, 0x83, 0xd9, 0x05, 0x34, 0x10,
	0x04, 0x77, 0xaa, 0xd2, 0x0c, 0x4f, 0xe0, 0x51, 0x8d, 0x3f, 0x9d, 0x04, 0x98, 0xf8, 0xf0, 0x78,
	0xf0, 0x4b, 0x07, 0x77, 0x6b, 0x43, 0xa3, 0xb3, 0xf6, 0xac, 0xb5, 0x5d, 0x3c, 0x2f, 0xeb, 0x1b,
	0xfe, 0x7d, 0xe4, 0xbd, 0xef, 0xff, 0x90, 0x95, 0x55, 0xcf, 0xea, 0x71, 0xc3, 0xb4, 0x25, 0xb2,
	0xc1, 0xb7, 0xff, 0x3b, 0x6f, 0xf3, 0x4a, 0x7a, 0xdb, 0x95, 0x8c, 0xf1, 0xbb, 0xaf, 0x6f, 0xa3,
	0x58, 0x5d, 0x16, 0x4b, 0x3b, 0x94, 0x6b, 0x87, 0xc8, 0xa5, 0xf4, 0x8a, 0xd4, 0x5a, 0x2c, 0x66,
	0x4e, 0x9e, 0x27, 0xd6, 0x76, 0x59, 0x2b, 0x94, 0x42, 0x65, 0x32, 0x49, 0x78, 0xe6, 0xc4, 0x42,
	0xf1, 0x4c, 0xb0, 0xc4, 0x61, 0x69, 0xea, 0x94, 0x7f, 0xd6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x20, 0xaf, 0x5b, 0x50, 0x69, 0x03, 0x00, 0x00,
}