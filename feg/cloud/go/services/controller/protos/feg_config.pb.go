// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg_config.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GyInitMethod int32

const (
	GyInitMethod_RESERVED    GyInitMethod = 0
	GyInitMethod_PER_SESSION GyInitMethod = 1
	GyInitMethod_PER_KEY     GyInitMethod = 2
)

var GyInitMethod_name = map[int32]string{
	0: "RESERVED",
	1: "PER_SESSION",
	2: "PER_KEY",
}
var GyInitMethod_value = map[string]int32{
	"RESERVED":    0,
	"PER_SESSION": 1,
	"PER_KEY":     2,
}

func (x GyInitMethod) String() string {
	return proto.EnumName(GyInitMethod_name, int32(x))
}
func (GyInitMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{0}
}

type DiamClientConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Retransmits          uint32   `protobuf:"varint,3,opt,name=retransmits,proto3" json:"retransmits,omitempty"`
	WatchdogInterval     uint32   `protobuf:"varint,4,opt,name=watchdog_interval,json=watchdogInterval,proto3" json:"watchdog_interval,omitempty"`
	RetryCount           uint32   `protobuf:"varint,5,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	LocalAddress         string   `protobuf:"bytes,6,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ProductName          string   `protobuf:"bytes,7,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Realm                string   `protobuf:"bytes,8,opt,name=realm,proto3" json:"realm,omitempty"`
	Host                 string   `protobuf:"bytes,9,opt,name=host,proto3" json:"host,omitempty"`
	DestRealm            string   `protobuf:"bytes,10,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	DestHost             string   `protobuf:"bytes,11,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamClientConfig) Reset()         { *m = DiamClientConfig{} }
func (m *DiamClientConfig) String() string { return proto.CompactTextString(m) }
func (*DiamClientConfig) ProtoMessage()    {}
func (*DiamClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{0}
}
func (m *DiamClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamClientConfig.Unmarshal(m, b)
}
func (m *DiamClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamClientConfig.Marshal(b, m, deterministic)
}
func (dst *DiamClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamClientConfig.Merge(dst, src)
}
func (m *DiamClientConfig) XXX_Size() int {
	return xxx_messageInfo_DiamClientConfig.Size(m)
}
func (m *DiamClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamClientConfig proto.InternalMessageInfo

func (m *DiamClientConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamClientConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamClientConfig) GetRetransmits() uint32 {
	if m != nil {
		return m.Retransmits
	}
	return 0
}

func (m *DiamClientConfig) GetWatchdogInterval() uint32 {
	if m != nil {
		return m.WatchdogInterval
	}
	return 0
}

func (m *DiamClientConfig) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *DiamClientConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamClientConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *DiamClientConfig) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *DiamClientConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DiamClientConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

func (m *DiamClientConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

type DiamServerConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	LocalAddress         string   `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	DestHost             string   `protobuf:"bytes,4,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	DestRealm            string   `protobuf:"bytes,5,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamServerConfig) Reset()         { *m = DiamServerConfig{} }
func (m *DiamServerConfig) String() string { return proto.CompactTextString(m) }
func (*DiamServerConfig) ProtoMessage()    {}
func (*DiamServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{1}
}
func (m *DiamServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamServerConfig.Unmarshal(m, b)
}
func (m *DiamServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamServerConfig.Marshal(b, m, deterministic)
}
func (dst *DiamServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamServerConfig.Merge(dst, src)
}
func (m *DiamServerConfig) XXX_Size() int {
	return xxx_messageInfo_DiamServerConfig.Size(m)
}
func (m *DiamServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamServerConfig proto.InternalMessageInfo

func (m *DiamServerConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamServerConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamServerConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamServerConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

func (m *DiamServerConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

type S6AConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *S6AConfig) Reset()         { *m = S6AConfig{} }
func (m *S6AConfig) String() string { return proto.CompactTextString(m) }
func (*S6AConfig) ProtoMessage()    {}
func (*S6AConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{2}
}
func (m *S6AConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S6AConfig.Unmarshal(m, b)
}
func (m *S6AConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S6AConfig.Marshal(b, m, deterministic)
}
func (dst *S6AConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S6AConfig.Merge(dst, src)
}
func (m *S6AConfig) XXX_Size() int {
	return xxx_messageInfo_S6AConfig.Size(m)
}
func (m *S6AConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_S6AConfig.DiscardUnknown(m)
}

var xxx_messageInfo_S6AConfig proto.InternalMessageInfo

func (m *S6AConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

type GxConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GxConfig) Reset()         { *m = GxConfig{} }
func (m *GxConfig) String() string { return proto.CompactTextString(m) }
func (*GxConfig) ProtoMessage()    {}
func (*GxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{3}
}
func (m *GxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GxConfig.Unmarshal(m, b)
}
func (m *GxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GxConfig.Marshal(b, m, deterministic)
}
func (dst *GxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GxConfig.Merge(dst, src)
}
func (m *GxConfig) XXX_Size() int {
	return xxx_messageInfo_GxConfig.Size(m)
}
func (m *GxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GxConfig proto.InternalMessageInfo

func (m *GxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

type GyConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	InitMethod           GyInitMethod      `protobuf:"varint,2,opt,name=init_method,json=initMethod,proto3,enum=feg.GyInitMethod" json:"init_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GyConfig) Reset()         { *m = GyConfig{} }
func (m *GyConfig) String() string { return proto.CompactTextString(m) }
func (*GyConfig) ProtoMessage()    {}
func (*GyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{4}
}
func (m *GyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GyConfig.Unmarshal(m, b)
}
func (m *GyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GyConfig.Marshal(b, m, deterministic)
}
func (dst *GyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GyConfig.Merge(dst, src)
}
func (m *GyConfig) XXX_Size() int {
	return xxx_messageInfo_GyConfig.Size(m)
}
func (m *GyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GyConfig proto.InternalMessageInfo

func (m *GyConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *GyConfig) GetInitMethod() GyInitMethod {
	if m != nil {
		return m.InitMethod
	}
	return GyInitMethod_RESERVED
}

type SwxConfig struct {
	Server *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// After auth, verify Non-3GPP IP Access enabled
	VerifyAuthorization  bool     `protobuf:"varint,2,opt,name=verify_authorization,json=verifyAuthorization,proto3" json:"verify_authorization,omitempty"`
	CacheTTLSeconds      uint32   `protobuf:"varint,3,opt,name=CacheTTLSeconds,proto3" json:"CacheTTLSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwxConfig) Reset()         { *m = SwxConfig{} }
func (m *SwxConfig) String() string { return proto.CompactTextString(m) }
func (*SwxConfig) ProtoMessage()    {}
func (*SwxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{5}
}
func (m *SwxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwxConfig.Unmarshal(m, b)
}
func (m *SwxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwxConfig.Marshal(b, m, deterministic)
}
func (dst *SwxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwxConfig.Merge(dst, src)
}
func (m *SwxConfig) XXX_Size() int {
	return xxx_messageInfo_SwxConfig.Size(m)
}
func (m *SwxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SwxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SwxConfig proto.InternalMessageInfo

func (m *SwxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *SwxConfig) GetVerifyAuthorization() bool {
	if m != nil {
		return m.VerifyAuthorization
	}
	return false
}

func (m *SwxConfig) GetCacheTTLSeconds() uint32 {
	if m != nil {
		return m.CacheTTLSeconds
	}
	return 0
}

type HSSConfig struct {
	Server *DiamServerConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Operator configuration field for LTE.
	LteAuthOp []byte `protobuf:"bytes,2,opt,name=lte_auth_op,json=lteAuthOp,proto3" json:"lte_auth_op,omitempty"`
	// Authentication management field for LTE.
	LteAuthAmf []byte `protobuf:"bytes,3,opt,name=lte_auth_amf,json=lteAuthAmf,proto3" json:"lte_auth_amf,omitempty"`
	// Maps from IMSI to SubscriptionProfile.
	SubProfiles map[string]*HSSConfig_SubscriptionProfile `protobuf:"bytes,4,rep,name=sub_profiles,json=subProfiles,proto3" json:"sub_profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If an IMSI if not found in sub_profiles, the default profile is used instead.
	DefaultSubProfile    *HSSConfig_SubscriptionProfile `protobuf:"bytes,5,opt,name=default_sub_profile,json=defaultSubProfile,proto3" json:"default_sub_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *HSSConfig) Reset()         { *m = HSSConfig{} }
func (m *HSSConfig) String() string { return proto.CompactTextString(m) }
func (*HSSConfig) ProtoMessage()    {}
func (*HSSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{6}
}
func (m *HSSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig.Unmarshal(m, b)
}
func (m *HSSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig.Marshal(b, m, deterministic)
}
func (dst *HSSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig.Merge(dst, src)
}
func (m *HSSConfig) XXX_Size() int {
	return xxx_messageInfo_HSSConfig.Size(m)
}
func (m *HSSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig proto.InternalMessageInfo

func (m *HSSConfig) GetServer() *DiamServerConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *HSSConfig) GetLteAuthOp() []byte {
	if m != nil {
		return m.LteAuthOp
	}
	return nil
}

func (m *HSSConfig) GetLteAuthAmf() []byte {
	if m != nil {
		return m.LteAuthAmf
	}
	return nil
}

func (m *HSSConfig) GetSubProfiles() map[string]*HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.SubProfiles
	}
	return nil
}

func (m *HSSConfig) GetDefaultSubProfile() *HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.DefaultSubProfile
	}
	return nil
}

type HSSConfig_SubscriptionProfile struct {
	// Maximum uplink bit rate (AMBR-UL)
	MaxUlBitRate uint64 `protobuf:"varint,1,opt,name=max_ul_bit_rate,json=maxUlBitRate,proto3" json:"max_ul_bit_rate,omitempty"`
	// Maximum downlink bit rate (AMBR-DL)
	MaxDlBitRate         uint64   `protobuf:"varint,2,opt,name=max_dl_bit_rate,json=maxDlBitRate,proto3" json:"max_dl_bit_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSSConfig_SubscriptionProfile) Reset()         { *m = HSSConfig_SubscriptionProfile{} }
func (m *HSSConfig_SubscriptionProfile) String() string { return proto.CompactTextString(m) }
func (*HSSConfig_SubscriptionProfile) ProtoMessage()    {}
func (*HSSConfig_SubscriptionProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{6, 0}
}
func (m *HSSConfig_SubscriptionProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Unmarshal(m, b)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Marshal(b, m, deterministic)
}
func (dst *HSSConfig_SubscriptionProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.Merge(dst, src)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Size() int {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Size(m)
}
func (m *HSSConfig_SubscriptionProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig_SubscriptionProfile proto.InternalMessageInfo

func (m *HSSConfig_SubscriptionProfile) GetMaxUlBitRate() uint64 {
	if m != nil {
		return m.MaxUlBitRate
	}
	return 0
}

func (m *HSSConfig_SubscriptionProfile) GetMaxDlBitRate() uint64 {
	if m != nil {
		return m.MaxDlBitRate
	}
	return 0
}

type HealthConfig struct {
	// Services the health service is responsible for tracking
	HealthServices []string `protobuf:"bytes,1,rep,name=health_services,json=healthServices,proto3" json:"health_services,omitempty"`
	// Frequency of FeG health manager updates to the cloud
	UpdateIntervalSecs uint32 `protobuf:"varint,2,opt,name=update_interval_secs,json=updateIntervalSecs,proto3" json:"update_interval_secs,omitempty"`
	// Period to disable connection creation when requested to from cloud
	CloudDisablePeriodSecs uint32 `protobuf:"varint,3,opt,name=cloud_disable_period_secs,json=cloudDisablePeriodSecs,proto3" json:"cloud_disable_period_secs,omitempty"`
	// Period to disable connection creation when locally determined
	LocalDisablePeriodSecs uint32 `protobuf:"varint,4,opt,name=local_disable_period_secs,json=localDisablePeriodSecs,proto3" json:"local_disable_period_secs,omitempty"`
	// The number of consecutive health update failures before locally disabling
	UpdateFailureThreshold uint32 `protobuf:"varint,5,opt,name=update_failure_threshold,json=updateFailureThreshold,proto3" json:"update_failure_threshold,omitempty"`
	// Percentage of request failures considered to be unhealthy
	RequestFailureThreshold float32 `protobuf:"fixed32,6,opt,name=request_failure_threshold,json=requestFailureThreshold,proto3" json:"request_failure_threshold,omitempty"`
	// Minimum number of requests necessary to consider a metrics interval valid
	MinimumRequestThreshold uint32 `protobuf:"varint,7,opt,name=minimum_request_threshold,json=minimumRequestThreshold,proto3" json:"minimum_request_threshold,omitempty"`
	// Cpu utilization healthy threshold
	CpuUtilizationThreshold float32 `protobuf:"fixed32,8,opt,name=cpu_utilization_threshold,json=cpuUtilizationThreshold,proto3" json:"cpu_utilization_threshold,omitempty"`
	// Available memory healthy threshold
	MemoryAvailableThreshold float32  `protobuf:"fixed32,9,opt,name=memory_available_threshold,json=memoryAvailableThreshold,proto3" json:"memory_available_threshold,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *HealthConfig) Reset()         { *m = HealthConfig{} }
func (m *HealthConfig) String() string { return proto.CompactTextString(m) }
func (*HealthConfig) ProtoMessage()    {}
func (*HealthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{7}
}
func (m *HealthConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthConfig.Unmarshal(m, b)
}
func (m *HealthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthConfig.Marshal(b, m, deterministic)
}
func (dst *HealthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthConfig.Merge(dst, src)
}
func (m *HealthConfig) XXX_Size() int {
	return xxx_messageInfo_HealthConfig.Size(m)
}
func (m *HealthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HealthConfig proto.InternalMessageInfo

func (m *HealthConfig) GetHealthServices() []string {
	if m != nil {
		return m.HealthServices
	}
	return nil
}

func (m *HealthConfig) GetUpdateIntervalSecs() uint32 {
	if m != nil {
		return m.UpdateIntervalSecs
	}
	return 0
}

func (m *HealthConfig) GetCloudDisablePeriodSecs() uint32 {
	if m != nil {
		return m.CloudDisablePeriodSecs
	}
	return 0
}

func (m *HealthConfig) GetLocalDisablePeriodSecs() uint32 {
	if m != nil {
		return m.LocalDisablePeriodSecs
	}
	return 0
}

func (m *HealthConfig) GetUpdateFailureThreshold() uint32 {
	if m != nil {
		return m.UpdateFailureThreshold
	}
	return 0
}

func (m *HealthConfig) GetRequestFailureThreshold() float32 {
	if m != nil {
		return m.RequestFailureThreshold
	}
	return 0
}

func (m *HealthConfig) GetMinimumRequestThreshold() uint32 {
	if m != nil {
		return m.MinimumRequestThreshold
	}
	return 0
}

func (m *HealthConfig) GetCpuUtilizationThreshold() float32 {
	if m != nil {
		return m.CpuUtilizationThreshold
	}
	return 0
}

func (m *HealthConfig) GetMemoryAvailableThreshold() float32 {
	if m != nil {
		return m.MemoryAvailableThreshold
	}
	return 0
}

type EapAkaConfig struct {
	Timeout              *EapAkaConfig_Timeouts `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	PlmnIds              []string               `protobuf:"bytes,2,rep,name=PlmnIds,proto3" json:"PlmnIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EapAkaConfig) Reset()         { *m = EapAkaConfig{} }
func (m *EapAkaConfig) String() string { return proto.CompactTextString(m) }
func (*EapAkaConfig) ProtoMessage()    {}
func (*EapAkaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{8}
}
func (m *EapAkaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapAkaConfig.Unmarshal(m, b)
}
func (m *EapAkaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapAkaConfig.Marshal(b, m, deterministic)
}
func (dst *EapAkaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapAkaConfig.Merge(dst, src)
}
func (m *EapAkaConfig) XXX_Size() int {
	return xxx_messageInfo_EapAkaConfig.Size(m)
}
func (m *EapAkaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EapAkaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EapAkaConfig proto.InternalMessageInfo

func (m *EapAkaConfig) GetTimeout() *EapAkaConfig_Timeouts {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *EapAkaConfig) GetPlmnIds() []string {
	if m != nil {
		return m.PlmnIds
	}
	return nil
}

type EapAkaConfig_Timeouts struct {
	ChallengeMs            uint32   `protobuf:"varint,1,opt,name=ChallengeMs,proto3" json:"ChallengeMs,omitempty"`
	ErrorNotificationMs    uint32   `protobuf:"varint,2,opt,name=ErrorNotificationMs,proto3" json:"ErrorNotificationMs,omitempty"`
	SessionMs              uint32   `protobuf:"varint,3,opt,name=SessionMs,proto3" json:"SessionMs,omitempty"`
	SessionAuthenticatedMs uint32   `protobuf:"varint,4,opt,name=SessionAuthenticatedMs,proto3" json:"SessionAuthenticatedMs,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *EapAkaConfig_Timeouts) Reset()         { *m = EapAkaConfig_Timeouts{} }
func (m *EapAkaConfig_Timeouts) String() string { return proto.CompactTextString(m) }
func (*EapAkaConfig_Timeouts) ProtoMessage()    {}
func (*EapAkaConfig_Timeouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{8, 0}
}
func (m *EapAkaConfig_Timeouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Unmarshal(m, b)
}
func (m *EapAkaConfig_Timeouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Marshal(b, m, deterministic)
}
func (dst *EapAkaConfig_Timeouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapAkaConfig_Timeouts.Merge(dst, src)
}
func (m *EapAkaConfig_Timeouts) XXX_Size() int {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Size(m)
}
func (m *EapAkaConfig_Timeouts) XXX_DiscardUnknown() {
	xxx_messageInfo_EapAkaConfig_Timeouts.DiscardUnknown(m)
}

var xxx_messageInfo_EapAkaConfig_Timeouts proto.InternalMessageInfo

func (m *EapAkaConfig_Timeouts) GetChallengeMs() uint32 {
	if m != nil {
		return m.ChallengeMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetErrorNotificationMs() uint32 {
	if m != nil {
		return m.ErrorNotificationMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetSessionMs() uint32 {
	if m != nil {
		return m.SessionMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetSessionAuthenticatedMs() uint32 {
	if m != nil {
		return m.SessionAuthenticatedMs
	}
	return 0
}

type Config struct {
	// FeG config params
	S6A                  *S6AConfig    `protobuf:"bytes,4,opt,name=s6a,proto3" json:"s6a,omitempty"`
	Gx                   *GxConfig     `protobuf:"bytes,5,opt,name=gx,proto3" json:"gx,omitempty"`
	Gy                   *GyConfig     `protobuf:"bytes,6,opt,name=gy,proto3" json:"gy,omitempty"`
	ServedNetworkIds     []string      `protobuf:"bytes,7,rep,name=served_network_ids,json=servedNetworkIds,proto3" json:"served_network_ids,omitempty"`
	Hss                  *HSSConfig    `protobuf:"bytes,8,opt,name=hss,proto3" json:"hss,omitempty"`
	Swx                  *SwxConfig    `protobuf:"bytes,9,opt,name=swx,proto3" json:"swx,omitempty"`
	Health               *HealthConfig `protobuf:"bytes,10,opt,name=health,proto3" json:"health,omitempty"`
	EapAka               *EapAkaConfig `protobuf:"bytes,11,opt,name=eap_aka,json=eapAka,proto3" json:"eap_aka,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_400ee177e7de03cd, []int{9}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetS6A() *S6AConfig {
	if m != nil {
		return m.S6A
	}
	return nil
}

func (m *Config) GetGx() *GxConfig {
	if m != nil {
		return m.Gx
	}
	return nil
}

func (m *Config) GetGy() *GyConfig {
	if m != nil {
		return m.Gy
	}
	return nil
}

func (m *Config) GetServedNetworkIds() []string {
	if m != nil {
		return m.ServedNetworkIds
	}
	return nil
}

func (m *Config) GetHss() *HSSConfig {
	if m != nil {
		return m.Hss
	}
	return nil
}

func (m *Config) GetSwx() *SwxConfig {
	if m != nil {
		return m.Swx
	}
	return nil
}

func (m *Config) GetHealth() *HealthConfig {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *Config) GetEapAka() *EapAkaConfig {
	if m != nil {
		return m.EapAka
	}
	return nil
}

func init() {
	proto.RegisterType((*DiamClientConfig)(nil), "feg.DiamClientConfig")
	proto.RegisterType((*DiamServerConfig)(nil), "feg.DiamServerConfig")
	proto.RegisterType((*S6AConfig)(nil), "feg.S6aConfig")
	proto.RegisterType((*GxConfig)(nil), "feg.GxConfig")
	proto.RegisterType((*GyConfig)(nil), "feg.GyConfig")
	proto.RegisterType((*SwxConfig)(nil), "feg.SwxConfig")
	proto.RegisterType((*HSSConfig)(nil), "feg.HSSConfig")
	proto.RegisterMapType((map[string]*HSSConfig_SubscriptionProfile)(nil), "feg.HSSConfig.SubProfilesEntry")
	proto.RegisterType((*HSSConfig_SubscriptionProfile)(nil), "feg.HSSConfig.SubscriptionProfile")
	proto.RegisterType((*HealthConfig)(nil), "feg.HealthConfig")
	proto.RegisterType((*EapAkaConfig)(nil), "feg.EapAkaConfig")
	proto.RegisterType((*EapAkaConfig_Timeouts)(nil), "feg.EapAkaConfig.Timeouts")
	proto.RegisterType((*Config)(nil), "feg.Config")
	proto.RegisterEnum("feg.GyInitMethod", GyInitMethod_name, GyInitMethod_value)
}

func init() { proto.RegisterFile("feg_config.proto", fileDescriptor_feg_config_400ee177e7de03cd) }

var fileDescriptor_feg_config_400ee177e7de03cd = []byte{
	// 1152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xfe, 0xd9, 0x4e, 0x62, 0xfb, 0xac, 0xd3, 0xb8, 0x93, 0xfe, 0xd2, 0x6d, 0xa0, 0xd4, 0x18,
	0x21, 0x42, 0x81, 0xa8, 0x04, 0x54, 0xb5, 0x11, 0x37, 0x69, 0x62, 0xda, 0x08, 0x92, 0x46, 0xb3,
	0x29, 0x12, 0xdc, 0x8c, 0x26, 0xbb, 0x63, 0x7b, 0x94, 0xfd, 0x63, 0x66, 0x66, 0x93, 0x98, 0x57,
	0xe0, 0x1a, 0xf1, 0x0c, 0x5c, 0x22, 0xf1, 0x3a, 0xdc, 0xf2, 0x1c, 0x68, 0xcf, 0xcc, 0xae, 0xdd,
	0xc4, 0x12, 0x28, 0x5c, 0xed, 0xce, 0xf9, 0xbe, 0x6f, 0xe6, 0x9b, 0xb3, 0x67, 0xce, 0x2c, 0x74,
	0x87, 0x62, 0xc4, 0xc2, 0x2c, 0x1d, 0xca, 0xd1, 0xf6, 0x44, 0x65, 0x26, 0x23, 0x8d, 0xa1, 0x18,
	0xf5, 0xff, 0xaa, 0x43, 0xf7, 0x40, 0xf2, 0x64, 0x3f, 0x96, 0x22, 0x35, 0xfb, 0x88, 0x93, 0x4d,
	0x68, 0x21, 0x25, 0xcc, 0x62, 0xbf, 0xd6, 0xab, 0x6d, 0xb5, 0x69, 0x35, 0x26, 0x3e, 0x34, 0x79,
	0x14, 0x29, 0xa1, 0xb5, 0x5f, 0x47, 0xa8, 0x1c, 0x92, 0x1e, 0x78, 0x4a, 0x18, 0xc5, 0x53, 0x9d,
	0x48, 0xa3, 0xfd, 0x46, 0xaf, 0xb6, 0xb5, 0x4a, 0xe7, 0x43, 0xe4, 0x13, 0xb8, 0x7b, 0xc9, 0x4d,
	0x38, 0x8e, 0xb2, 0x11, 0x93, 0xa9, 0x11, 0xea, 0x82, 0xc7, 0xfe, 0x12, 0xf2, 0xba, 0x25, 0x70,
	0xe8, 0xe2, 0xe4, 0x91, 0x9d, 0x6e, 0xca, 0xc2, 0x2c, 0x4f, 0x8d, 0xbf, 0x8c, 0x34, 0xc0, 0xd0,
	0x7e, 0x11, 0x21, 0x1f, 0xc0, 0x6a, 0x9c, 0x85, 0x3c, 0x66, 0xa5, 0x9f, 0x15, 0xf4, 0xd3, 0xc1,
	0xe0, 0x9e, 0x33, 0xf5, 0x3e, 0x74, 0x26, 0x2a, 0x8b, 0xf2, 0xd0, 0xb0, 0x94, 0x27, 0xc2, 0x6f,
	0x22, 0xc7, 0x73, 0xb1, 0x63, 0x9e, 0x08, 0x72, 0x0f, 0x96, 0x95, 0xe0, 0x71, 0xe2, 0xb7, 0x10,
	0xb3, 0x03, 0x42, 0x60, 0x69, 0x9c, 0x69, 0xe3, 0xb7, 0x31, 0x88, 0xef, 0xe4, 0x21, 0x40, 0x24,
	0xb4, 0x61, 0x96, 0x0e, 0x88, 0xb4, 0x8b, 0x08, 0x45, 0xc9, 0x3b, 0x80, 0x03, 0x86, 0x3a, 0xcf,
	0xe6, 0xad, 0x08, 0xbc, 0xca, 0xb4, 0xe9, 0xff, 0x56, 0xb3, 0x89, 0x0e, 0x84, 0xba, 0x10, 0xea,
	0x3f, 0x25, 0xfa, 0xc6, 0xc6, 0x1b, 0x0b, 0x36, 0xfe, 0x96, 0x99, 0xa5, 0xb7, 0xcd, 0x5c, 0xdb,
	0xc8, 0xf2, 0xb5, 0x8d, 0xf4, 0x77, 0xa1, 0x1d, 0x3c, 0xe5, 0xce, 0xe3, 0x67, 0xb0, 0xa2, 0xd1,
	0x33, 0x3a, 0xf4, 0x76, 0xfe, 0xbf, 0x3d, 0x14, 0xa3, 0xed, 0xeb, 0x35, 0x43, 0x1d, 0xa9, 0xff,
	0x1c, 0x5a, 0x2f, 0xaf, 0x6e, 0x27, 0x4d, 0xa0, 0xf5, 0x72, 0x7a, 0x2b, 0x29, 0xd9, 0x01, 0x4f,
	0xa6, 0xd2, 0xb0, 0x44, 0x98, 0x71, 0x16, 0x61, 0xc2, 0xee, 0xec, 0xdc, 0x45, 0xcd, 0xcb, 0xe9,
	0x61, 0x2a, 0xcd, 0x11, 0x02, 0x14, 0x64, 0xf5, 0xde, 0xff, 0xb5, 0x06, 0xed, 0xe0, 0xf2, 0x76,
	0x5e, 0xc9, 0xe7, 0x70, 0xef, 0x42, 0x28, 0x39, 0x9c, 0x32, 0x9e, 0x9b, 0x71, 0xa6, 0xe4, 0x4f,
	0xdc, 0xc8, 0x2c, 0xc5, 0x95, 0x5b, 0x74, 0xdd, 0x62, 0x7b, 0xf3, 0x10, 0xd9, 0x82, 0xb5, 0x7d,
	0x1e, 0x8e, 0xc5, 0xe9, 0xe9, 0xb7, 0x81, 0x08, 0xb3, 0x34, 0x2a, 0xcf, 0xc8, 0xf5, 0x70, 0xff,
	0xcf, 0x06, 0xb4, 0x5f, 0x05, 0xc1, 0x3f, 0x3a, 0x9b, 0xaf, 0xa5, 0xca, 0xd9, 0x7b, 0xe0, 0xc5,
	0x46, 0xa0, 0x2d, 0x96, 0x4d, 0xd0, 0x50, 0x87, 0xb6, 0x63, 0x23, 0x0a, 0x37, 0xaf, 0x27, 0xa4,
	0x07, 0x9d, 0x0a, 0xe7, 0xc9, 0x10, 0x3d, 0x74, 0x28, 0x38, 0xc2, 0x5e, 0x32, 0x24, 0x2f, 0xa0,
	0xa3, 0xf3, 0x33, 0x36, 0x51, 0xd9, 0x50, 0xc6, 0x42, 0xfb, 0x4b, 0xbd, 0xc6, 0x96, 0xb7, 0xf3,
	0x08, 0x97, 0xad, 0x6c, 0x6d, 0x07, 0xf9, 0xd9, 0x89, 0x63, 0x0c, 0x52, 0xa3, 0xa6, 0xd4, 0xd3,
	0xb3, 0x08, 0xa1, 0xb0, 0x1e, 0x89, 0x21, 0xcf, 0x63, 0xc3, 0xe6, 0xe6, 0xc2, 0x52, 0xf3, 0x76,
	0xfa, 0x37, 0xa7, 0xd2, 0xa1, 0x92, 0x93, 0x22, 0x4d, 0x6e, 0x06, 0x7a, 0xd7, 0xc9, 0x67, 0xcb,
	0x6c, 0x86, 0xb0, 0xbe, 0x80, 0x49, 0x3e, 0x84, 0xb5, 0x84, 0x5f, 0xb1, 0x3c, 0x66, 0x67, 0xd2,
	0x30, 0xc5, 0x8d, 0xc0, 0x44, 0x2d, 0xd1, 0x4e, 0xc2, 0xaf, 0xde, 0xc4, 0x2f, 0xa4, 0xa1, 0xdc,
	0x54, 0xb4, 0x68, 0x8e, 0x56, 0xaf, 0x68, 0x07, 0x25, 0x6d, 0xf3, 0x0c, 0xba, 0xd7, 0x77, 0x46,
	0xba, 0xd0, 0x38, 0x17, 0x53, 0x77, 0x42, 0x8b, 0x57, 0xf2, 0x0c, 0x96, 0x2f, 0x78, 0x9c, 0xdb,
	0x29, 0xfe, 0xdd, 0x86, 0xac, 0x60, 0xb7, 0xfe, 0xac, 0xd6, 0xff, 0x79, 0x09, 0x3a, 0xaf, 0x04,
	0x8f, 0xcd, 0xd8, 0x7d, 0xe2, 0x8f, 0x60, 0x6d, 0x8c, 0x63, 0x56, 0x7c, 0x44, 0x19, 0x0a, 0xed,
	0xd7, 0x7a, 0x8d, 0xad, 0x36, 0xbd, 0x63, 0xc3, 0x81, 0x8b, 0x92, 0x27, 0x70, 0x2f, 0x9f, 0x44,
	0xdc, 0x88, 0xaa, 0x7f, 0x32, 0x2d, 0x42, 0xdb, 0x21, 0x56, 0x29, 0xb1, 0x58, 0xd9, 0x42, 0x03,
	0x11, 0x6a, 0xf2, 0x1c, 0x1e, 0x84, 0x71, 0x96, 0x47, 0x2c, 0x92, 0x9a, 0x9f, 0xc5, 0x82, 0x4d,
	0x84, 0x92, 0x59, 0x64, 0x65, 0xb6, 0xfe, 0x36, 0x90, 0x70, 0x60, 0xf1, 0x13, 0x84, 0x4b, 0xa9,
	0xed, 0x33, 0x8b, 0xa4, 0xb6, 0x6d, 0x6f, 0x20, 0xe1, 0xa6, 0xf4, 0x19, 0xf8, 0xce, 0xe7, 0x90,
	0xcb, 0x38, 0x57, 0x82, 0x99, 0xb1, 0x12, 0x7a, 0x9c, 0xc5, 0x91, 0xeb, 0xe4, 0x1b, 0x16, 0xff,
	0xda, 0xc2, 0xa7, 0x25, 0x4a, 0x76, 0xe1, 0x81, 0x12, 0x3f, 0xe6, 0x45, 0x77, 0xba, 0x29, 0x2d,
	0x3a, 0x7c, 0x9d, 0xde, 0x77, 0x84, 0x45, 0xda, 0x44, 0xa6, 0x32, 0xc9, 0x13, 0x56, 0xce, 0x31,
	0xd3, 0x36, 0x71, 0xd9, 0xfb, 0x8e, 0x40, 0x2d, 0xfe, 0x96, 0x36, 0x9c, 0xe4, 0x2c, 0x37, 0x32,
	0x76, 0x07, 0x76, 0x4e, 0xdb, 0xb2, 0xeb, 0x86, 0x93, 0xfc, 0xcd, 0x0c, 0x9f, 0x69, 0xbf, 0x82,
	0xcd, 0x44, 0x24, 0x99, 0x9a, 0x32, 0x7e, 0xc1, 0x65, 0x8c, 0xb9, 0x9a, 0x89, 0xdb, 0x28, 0xf6,
	0x2d, 0x63, 0xaf, 0x24, 0x54, 0xea, 0xfe, 0x2f, 0x75, 0xe8, 0x0c, 0xf8, 0x64, 0xef, 0xbc, 0xec,
	0xb8, 0x5f, 0x42, 0xd3, 0xc8, 0x44, 0x64, 0xb9, 0x71, 0x27, 0x7e, 0x13, 0xcb, 0x6b, 0x9e, 0xb3,
	0x7d, 0x6a, 0x09, 0x9a, 0x96, 0xd4, 0xe2, 0xbe, 0x38, 0x89, 0x93, 0xf4, 0x30, 0x2a, 0xaa, 0xa1,
	0xa8, 0x9d, 0x72, 0xb8, 0xf9, 0x47, 0x0d, 0x5a, 0x25, 0xbf, 0xb8, 0xa5, 0xf7, 0xc7, 0x3c, 0x8e,
	0x45, 0x3a, 0x12, 0x47, 0x1a, 0x17, 0x58, 0xa5, 0xf3, 0x21, 0xf2, 0x04, 0xd6, 0x07, 0x4a, 0x65,
	0xea, 0x38, 0x33, 0x72, 0x28, 0x43, 0xdc, 0xeb, 0x51, 0x59, 0x62, 0x8b, 0x20, 0xf2, 0x2e, 0xb4,
	0x03, 0xa1, 0xb5, 0xe5, 0xd9, 0x9a, 0x9a, 0x05, 0xc8, 0x53, 0xd8, 0x70, 0x83, 0xa2, 0xc1, 0x88,
	0xd4, 0x14, 0x42, 0x11, 0x1d, 0x55, 0x35, 0xb4, 0x18, 0xed, 0xff, 0x5e, 0x87, 0x15, 0x97, 0x91,
	0x1e, 0x34, 0xf4, 0x53, 0x8e, 0x7c, 0x6f, 0xe7, 0x0e, 0x66, 0xa3, 0xba, 0xa0, 0x68, 0x01, 0x91,
	0x87, 0x50, 0x1f, 0x5d, 0xb9, 0xf6, 0xb2, 0x6a, 0xfb, 0xbe, 0xeb, 0xec, 0xb4, 0x3e, 0xba, 0x42,
	0x78, 0x8a, 0xe5, 0x53, 0xc1, 0xd3, 0x0a, 0x9e, 0x92, 0x4f, 0x81, 0x60, 0xf7, 0x8c, 0x58, 0x2a,
	0xcc, 0x65, 0xa6, 0xce, 0x99, 0x8c, 0xb4, 0xdf, 0xc4, 0x34, 0x76, 0x2d, 0x72, 0x6c, 0x81, 0xc3,
	0xa8, 0x48, 0x61, 0x63, 0xac, 0x35, 0x16, 0x45, 0xe9, 0xa6, 0x3a, 0xfa, 0xb4, 0x80, 0xd0, 0xef,
	0xe5, 0x15, 0x7e, 0xf9, 0xca, 0x6f, 0x79, 0xd3, 0xd0, 0x02, 0x22, 0x1f, 0xc3, 0x8a, 0x3d, 0xda,
	0xf8, 0x1b, 0xe1, 0xb9, 0xbb, 0x6a, 0xbe, 0x29, 0x50, 0x47, 0x20, 0x8f, 0xa1, 0x29, 0xf8, 0x84,
	0xf1, 0x73, 0x8e, 0x3f, 0x15, 0x25, 0x77, 0xbe, 0x1c, 0xe8, 0x8a, 0xc0, 0xd1, 0xe3, 0x5d, 0xe8,
	0xcc, 0xdf, 0x77, 0xa4, 0x03, 0x2d, 0x3a, 0x08, 0x06, 0xf4, 0xbb, 0xc1, 0x41, 0xf7, 0x7f, 0x64,
	0x0d, 0xbc, 0x93, 0x01, 0x65, 0xc1, 0x20, 0x08, 0x0e, 0x5f, 0x1f, 0x77, 0x6b, 0xc4, 0x83, 0x66,
	0x11, 0xf8, 0x66, 0xf0, 0x7d, 0xb7, 0xfe, 0xa2, 0xf5, 0xc3, 0x0a, 0xfe, 0x7c, 0xe8, 0x33, 0xfb,
	0xfc, 0xe2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x29, 0x92, 0xa6, 0x34, 0x0a, 0x00, 0x00,
}
