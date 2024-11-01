// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: strategy.proto

package pbstrategy

import (
	group "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/group"
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

// strategy source resource reference: pkg/dal/table/strategy.go
type Strategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *StrategySpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *StrategyState      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Attachment *StrategyAttachment `protobuf:"bytes,4,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *Revision           `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
	App        *AppSpec            `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *Strategy) Reset() {
	*x = Strategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strategy) ProtoMessage() {}

func (x *Strategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strategy.ProtoReflect.Descriptor instead.
func (*Strategy) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *Strategy) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Strategy) GetSpec() *StrategySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Strategy) GetStatus() *StrategyState {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Strategy) GetAttachment() *StrategyAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Strategy) GetRevision() *Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

func (x *Strategy) GetApp() *AppSpec {
	if x != nil {
		return x.App
	}
	return nil
}

// ReleaseSpec source resource reference: pkg/dal/table/release.go
type StrategySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReleaseId         uint32 `protobuf:"varint,2,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	AsDefault         bool   `protobuf:"varint,3,opt,name=as_default,json=asDefault,proto3" json:"as_default,omitempty"`
	Scope             *Scope `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	Namespace         string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PublishType       string `protobuf:"bytes,6,opt,name=publish_type,json=publishType,proto3" json:"publish_type,omitempty"`
	PublishTime       string `protobuf:"bytes,7,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	PublishStatus     string `protobuf:"bytes,8,opt,name=publish_status,json=publishStatus,proto3" json:"publish_status,omitempty"`
	RejectReason      string `protobuf:"bytes,9,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
	Approver          string `protobuf:"bytes,10,opt,name=approver,proto3" json:"approver,omitempty"`
	ApproverProgress  string `protobuf:"bytes,11,opt,name=approver_progress,json=approverProgress,proto3" json:"approver_progress,omitempty"`
	Memo              string `protobuf:"bytes,12,opt,name=memo,proto3" json:"memo,omitempty"`
	FinalApprovalTime string `protobuf:"bytes,13,opt,name=final_approval_time,json=finalApprovalTime,proto3" json:"final_approval_time,omitempty"`
}

func (x *StrategySpec) Reset() {
	*x = StrategySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySpec) ProtoMessage() {}

func (x *StrategySpec) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySpec.ProtoReflect.Descriptor instead.
func (*StrategySpec) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *StrategySpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StrategySpec) GetReleaseId() uint32 {
	if x != nil {
		return x.ReleaseId
	}
	return 0
}

func (x *StrategySpec) GetAsDefault() bool {
	if x != nil {
		return x.AsDefault
	}
	return false
}

func (x *StrategySpec) GetScope() *Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *StrategySpec) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *StrategySpec) GetPublishType() string {
	if x != nil {
		return x.PublishType
	}
	return ""
}

func (x *StrategySpec) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *StrategySpec) GetPublishStatus() string {
	if x != nil {
		return x.PublishStatus
	}
	return ""
}

func (x *StrategySpec) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

func (x *StrategySpec) GetApprover() string {
	if x != nil {
		return x.Approver
	}
	return ""
}

func (x *StrategySpec) GetApproverProgress() string {
	if x != nil {
		return x.ApproverProgress
	}
	return ""
}

func (x *StrategySpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *StrategySpec) GetFinalApprovalTime() string {
	if x != nil {
		return x.FinalApprovalTime
	}
	return ""
}

// Scope defines the scope
type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*group.Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{2}
}

func (x *Scope) GetGroups() []*group.Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

// StrategyState defines the strategy's state
type StrategyState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubState string `protobuf:"bytes,1,opt,name=pub_state,json=pubState,proto3" json:"pub_state,omitempty"`
}

func (x *StrategyState) Reset() {
	*x = StrategyState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyState) ProtoMessage() {}

func (x *StrategyState) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyState.ProtoReflect.Descriptor instead.
func (*StrategyState) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{3}
}

func (x *StrategyState) GetPubState() string {
	if x != nil {
		return x.PubState
	}
	return ""
}

// StrategyAttachment defines the strategy's attachment
type StrategyAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId         uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId         uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	StrategySetId uint32 `protobuf:"varint,3,opt,name=strategy_set_id,json=strategySetId,proto3" json:"strategy_set_id,omitempty"`
}

func (x *StrategyAttachment) Reset() {
	*x = StrategyAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyAttachment) ProtoMessage() {}

func (x *StrategyAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyAttachment.ProtoReflect.Descriptor instead.
func (*StrategyAttachment) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{4}
}

func (x *StrategyAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *StrategyAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *StrategyAttachment) GetStrategySetId() uint32 {
	if x != nil {
		return x.StrategySetId
	}
	return 0
}

// // Revision defines the strategy's revision
type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Reviser   string `protobuf:"bytes,2,opt,name=reviser,proto3" json:"reviser,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{5}
}

func (x *Revision) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Revision) GetReviser() string {
	if x != nil {
		return x.Reviser
	}
	return ""
}

func (x *Revision) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Revision) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// AuditStrategy strategy relate audit
type AuditStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublishType       string `protobuf:"bytes,1,opt,name=publish_type,json=publishType,proto3" json:"publish_type,omitempty"`
	PublishTime       string `protobuf:"bytes,2,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	PublishStatus     string `protobuf:"bytes,3,opt,name=publish_status,json=publishStatus,proto3" json:"publish_status,omitempty"`
	RejectReason      string `protobuf:"bytes,4,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
	Approver          string `protobuf:"bytes,5,opt,name=approver,proto3" json:"approver,omitempty"`
	ApproverProgress  string `protobuf:"bytes,6,opt,name=approver_progress,json=approverProgress,proto3" json:"approver_progress,omitempty"`
	UpdatedAt         string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Reviser           string `protobuf:"bytes,8,opt,name=reviser,proto3" json:"reviser,omitempty"`
	ReleaseId         uint32 `protobuf:"varint,9,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	Scope             *Scope `protobuf:"bytes,10,opt,name=scope,proto3" json:"scope,omitempty"`
	Creator           string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
	FinalApprovalTime string `protobuf:"bytes,12,opt,name=final_approval_time,json=finalApprovalTime,proto3" json:"final_approval_time,omitempty"`
}

func (x *AuditStrategy) Reset() {
	*x = AuditStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditStrategy) ProtoMessage() {}

func (x *AuditStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditStrategy.ProtoReflect.Descriptor instead.
func (*AuditStrategy) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{6}
}

func (x *AuditStrategy) GetPublishType() string {
	if x != nil {
		return x.PublishType
	}
	return ""
}

func (x *AuditStrategy) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *AuditStrategy) GetPublishStatus() string {
	if x != nil {
		return x.PublishStatus
	}
	return ""
}

func (x *AuditStrategy) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

func (x *AuditStrategy) GetApprover() string {
	if x != nil {
		return x.Approver
	}
	return ""
}

func (x *AuditStrategy) GetApproverProgress() string {
	if x != nil {
		return x.ApproverProgress
	}
	return ""
}

func (x *AuditStrategy) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AuditStrategy) GetReviser() string {
	if x != nil {
		return x.Reviser
	}
	return ""
}

func (x *AuditStrategy) GetReleaseId() uint32 {
	if x != nil {
		return x.ReleaseId
	}
	return 0
}

func (x *AuditStrategy) GetScope() *Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *AuditStrategy) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *AuditStrategy) GetFinalApprovalTime() string {
	if x != nil {
		return x.FinalApprovalTime
	}
	return ""
}

// AppSpec app relate
type AppSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *AppSpec) Reset() {
	*x = AppSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSpec) ProtoMessage() {}

func (x *AppSpec) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppSpec.ProtoReflect.Descriptor instead.
func (*AppSpec) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{7}
}

func (x *AppSpec) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_strategy_proto protoreflect.FileDescriptor

var file_strategy_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x1a, 0x23, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x02, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0xc6, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x2f, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x6a, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb5, 0x03, 0x0a, 0x0d, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75,
	0x65, 0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x63, 0x73, 0x2f, 0x62, 0x63, 0x73,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x62, 0x73,
	0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x3b, 0x70, 0x62,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strategy_proto_rawDescOnce sync.Once
	file_strategy_proto_rawDescData = file_strategy_proto_rawDesc
)

func file_strategy_proto_rawDescGZIP() []byte {
	file_strategy_proto_rawDescOnce.Do(func() {
		file_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_proto_rawDescData)
	})
	return file_strategy_proto_rawDescData
}

var file_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strategy_proto_goTypes = []interface{}{
	(*Strategy)(nil),           // 0: pbstrategy.Strategy
	(*StrategySpec)(nil),       // 1: pbstrategy.StrategySpec
	(*Scope)(nil),              // 2: pbstrategy.Scope
	(*StrategyState)(nil),      // 3: pbstrategy.StrategyState
	(*StrategyAttachment)(nil), // 4: pbstrategy.StrategyAttachment
	(*Revision)(nil),           // 5: pbstrategy.Revision
	(*AuditStrategy)(nil),      // 6: pbstrategy.AuditStrategy
	(*AppSpec)(nil),            // 7: pbstrategy.AppSpec
	(*group.Group)(nil),        // 8: pbgroup.Group
}
var file_strategy_proto_depIdxs = []int32{
	1, // 0: pbstrategy.Strategy.spec:type_name -> pbstrategy.StrategySpec
	3, // 1: pbstrategy.Strategy.status:type_name -> pbstrategy.StrategyState
	4, // 2: pbstrategy.Strategy.attachment:type_name -> pbstrategy.StrategyAttachment
	5, // 3: pbstrategy.Strategy.revision:type_name -> pbstrategy.Revision
	7, // 4: pbstrategy.Strategy.app:type_name -> pbstrategy.AppSpec
	2, // 5: pbstrategy.StrategySpec.scope:type_name -> pbstrategy.Scope
	8, // 6: pbstrategy.Scope.groups:type_name -> pbgroup.Group
	2, // 7: pbstrategy.AuditStrategy.scope:type_name -> pbstrategy.Scope
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_strategy_proto_init() }
func file_strategy_proto_init() {
	if File_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strategy); i {
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
		file_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySpec); i {
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
		file_strategy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
		file_strategy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyState); i {
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
		file_strategy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyAttachment); i {
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
		file_strategy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
		file_strategy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditStrategy); i {
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
		file_strategy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppSpec); i {
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
			RawDescriptor: file_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_strategy_proto_goTypes,
		DependencyIndexes: file_strategy_proto_depIdxs,
		MessageInfos:      file_strategy_proto_msgTypes,
	}.Build()
	File_strategy_proto = out.File
	file_strategy_proto_rawDesc = nil
	file_strategy_proto_goTypes = nil
	file_strategy_proto_depIdxs = nil
}
