// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: report_action_status_request.proto

package proto

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

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

// The various state a workflow can be
type StateType int32

const (
	// Unspecified is the default state of a workflow. It means that the state of
	// the workflow is not known.
	StateType_UNSPECIFIED StateType = 0
	// A workflow is in pending state when it is waiting for the hardware to pick
	// it up and start the execution.
	StateType_PENDING StateType = 1
	// A workflow is in a running state when the tink-worker started the
	// exeuction of it, and it is currently in execution inside the device
	// itself.
	StateType_RUNNING StateType = 2
	// Failed is a final state. Something wrong happened during the execution of
	// the workflow inside the target. Have a look at the logs to see if you can
	// spot what is going on.
	StateType_FAILED StateType = 3
	// Timeout is final state, almost like FAILED but it communicate to you that
	// an action or the overall workflow reached the specified timeout.
	StateType_TIMEOUT StateType = 4
	// This is the state we all deserve. The execution of the workflow is over
	// and everything is just fine. Sit down, and enjoy your great work.
	StateType_SUCCESS StateType = 5
)

// Enum value maps for StateType.
var (
	StateType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "PENDING",
		2: "RUNNING",
		3: "FAILED",
		4: "TIMEOUT",
		5: "SUCCESS",
	}
	StateType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"PENDING":     1,
		"RUNNING":     2,
		"FAILED":      3,
		"TIMEOUT":     4,
		"SUCCESS":     5,
	}
)

func (x StateType) Enum() *StateType {
	p := new(StateType)
	*p = x
	return p
}

func (x StateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateType) Descriptor() protoreflect.EnumDescriptor {
	return file_report_action_status_request_proto_enumTypes[0].Descriptor()
}

func (StateType) Type() protoreflect.EnumType {
	return &file_report_action_status_request_proto_enumTypes[0]
}

func (x StateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateType.Descriptor instead.
func (StateType) EnumDescriptor() ([]byte, []int) {
	return file_report_action_status_request_proto_rawDescGZIP(), []int{0}
}

// ActionStatusRequest is the state of a single Workflow Action
type ActionStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The workflow id
	WorkflowId *string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId" json:"workflow_id,omitempty"`
	// The worker id
	WorkerId *string `protobuf:"bytes,2,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	// The name of the task this action is part of
	TaskId *string `protobuf:"bytes,3,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	// The action id
	ActionId *string `protobuf:"bytes,4,opt,name=action_id,json=actionId" json:"action_id,omitempty"`
	// The name of the action
	ActionName *string `protobuf:"bytes,5,opt,name=action_name,json=actionName" json:"action_name,omitempty"`
	// The state of the action. Those are the same described for workflow as
	// well. pending, running, successful and so on.
	ActionState *StateType `protobuf:"varint,6,opt,name=action_state,json=actionState,enum=proto.StateType" json:"action_state,omitempty"`
	// This is the time when the action started the execution
	ExecutionStart *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=execution_start,json=executionStart" json:"execution_start,omitempty"`
	// This is the time when the action stopped the execution
	ExecutionStop *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=execution_stop,json=executionStop" json:"execution_stop,omitempty"`
	// The execution duration time for the action
	ExecutionDuration *string `protobuf:"bytes,9,opt,name=execution_duration,json=executionDuration" json:"execution_duration,omitempty"`
	// The message returned from the action.
	Message       *ActionMessage `protobuf:"bytes,10,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionStatusRequest) Reset() {
	*x = ActionStatusRequest{}
	mi := &file_report_action_status_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionStatusRequest) ProtoMessage() {}

func (x *ActionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_action_status_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionStatusRequest.ProtoReflect.Descriptor instead.
func (*ActionStatusRequest) Descriptor() ([]byte, []int) {
	return file_report_action_status_request_proto_rawDescGZIP(), []int{0}
}

func (x *ActionStatusRequest) GetWorkflowId() string {
	if x != nil && x.WorkflowId != nil {
		return *x.WorkflowId
	}
	return ""
}

func (x *ActionStatusRequest) GetWorkerId() string {
	if x != nil && x.WorkerId != nil {
		return *x.WorkerId
	}
	return ""
}

func (x *ActionStatusRequest) GetTaskId() string {
	if x != nil && x.TaskId != nil {
		return *x.TaskId
	}
	return ""
}

func (x *ActionStatusRequest) GetActionId() string {
	if x != nil && x.ActionId != nil {
		return *x.ActionId
	}
	return ""
}

func (x *ActionStatusRequest) GetActionName() string {
	if x != nil && x.ActionName != nil {
		return *x.ActionName
	}
	return ""
}

func (x *ActionStatusRequest) GetActionState() StateType {
	if x != nil && x.ActionState != nil {
		return *x.ActionState
	}
	return StateType_UNSPECIFIED
}

func (x *ActionStatusRequest) GetExecutionStart() *timestamppb.Timestamp {
	if x != nil {
		return x.ExecutionStart
	}
	return nil
}

func (x *ActionStatusRequest) GetExecutionStop() *timestamppb.Timestamp {
	if x != nil {
		return x.ExecutionStop
	}
	return nil
}

func (x *ActionStatusRequest) GetExecutionDuration() string {
	if x != nil && x.ExecutionDuration != nil {
		return *x.ExecutionDuration
	}
	return ""
}

func (x *ActionStatusRequest) GetMessage() *ActionMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

// ActionMessage to report the status of a single action, it's an object so it can be extended
type ActionMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Message is the human readable message that can be used to describe the status of the action
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionMessage) Reset() {
	*x = ActionMessage{}
	mi := &file_report_action_status_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMessage) ProtoMessage() {}

func (x *ActionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_report_action_status_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMessage.ProtoReflect.Descriptor instead.
func (*ActionMessage) Descriptor() ([]byte, []int) {
	return file_report_action_status_request_proto_rawDescGZIP(), []int{1}
}

func (x *ActionMessage) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_report_action_status_request_proto protoreflect.FileDescriptor

var file_report_action_status_request_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x03, 0x0a,
	0x13, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43,
	0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2a, 0x5c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x05, 0x42, 0x8b,
	0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x1e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x08, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_report_action_status_request_proto_rawDescOnce sync.Once
	file_report_action_status_request_proto_rawDescData []byte
)

func file_report_action_status_request_proto_rawDescGZIP() []byte {
	file_report_action_status_request_proto_rawDescOnce.Do(func() {
		file_report_action_status_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_report_action_status_request_proto_rawDesc), len(file_report_action_status_request_proto_rawDesc)))
	})
	return file_report_action_status_request_proto_rawDescData
}

var file_report_action_status_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_report_action_status_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_report_action_status_request_proto_goTypes = []any{
	(StateType)(0),                // 0: proto.StateType
	(*ActionStatusRequest)(nil),   // 1: proto.ActionStatusRequest
	(*ActionMessage)(nil),         // 2: proto.ActionMessage
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_report_action_status_request_proto_depIdxs = []int32{
	0, // 0: proto.ActionStatusRequest.action_state:type_name -> proto.StateType
	3, // 1: proto.ActionStatusRequest.execution_start:type_name -> google.protobuf.Timestamp
	3, // 2: proto.ActionStatusRequest.execution_stop:type_name -> google.protobuf.Timestamp
	2, // 3: proto.ActionStatusRequest.message:type_name -> proto.ActionMessage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_report_action_status_request_proto_init() }
func file_report_action_status_request_proto_init() {
	if File_report_action_status_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_report_action_status_request_proto_rawDesc), len(file_report_action_status_request_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_report_action_status_request_proto_goTypes,
		DependencyIndexes: file_report_action_status_request_proto_depIdxs,
		EnumInfos:         file_report_action_status_request_proto_enumTypes,
		MessageInfos:      file_report_action_status_request_proto_msgTypes,
	}.Build()
	File_report_action_status_request_proto = out.File
	file_report_action_status_request_proto_goTypes = nil
	file_report_action_status_request_proto_depIdxs = nil
}
