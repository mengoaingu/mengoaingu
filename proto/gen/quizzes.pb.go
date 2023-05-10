// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: quizzes.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Level int32

const (
	Level_LEVEL_INVALID Level = 0
	Level_LEVEL_A1      Level = 1
	Level_LEVEL_A2      Level = 2
	Level_LEVEL_B1      Level = 3
	Level_LEVEL_B2      Level = 4
	Level_LEVEL_C1      Level = 5
	Level_LEVEL_C2      Level = 6
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "LEVEL_INVALID",
		1: "LEVEL_A1",
		2: "LEVEL_A2",
		3: "LEVEL_B1",
		4: "LEVEL_B2",
		5: "LEVEL_C1",
		6: "LEVEL_C2",
	}
	Level_value = map[string]int32{
		"LEVEL_INVALID": 0,
		"LEVEL_A1":      1,
		"LEVEL_A2":      2,
		"LEVEL_B1":      3,
		"LEVEL_B2":      4,
		"LEVEL_C1":      5,
		"LEVEL_C2":      6,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_quizzes_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_quizzes_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{0}
}

type QuestionType int32

const (
	QuestionType_QUESTION_TYPE_INVALID QuestionType = 0
	QuestionType_QUESTION_TYPE_PART1   QuestionType = 1
	QuestionType_QUESTION_TYPE_PART2   QuestionType = 2
	QuestionType_QUESTION_TYPE_PART3   QuestionType = 3
	QuestionType_QUESTION_TYPE_PART4   QuestionType = 4
	QuestionType_QUESTION_TYPE_PART5   QuestionType = 5
	QuestionType_QUESTION_TYPE_PART6   QuestionType = 6
	QuestionType_QUESTION_TYPE_PART7   QuestionType = 7
)

// Enum value maps for QuestionType.
var (
	QuestionType_name = map[int32]string{
		0: "QUESTION_TYPE_INVALID",
		1: "QUESTION_TYPE_PART1",
		2: "QUESTION_TYPE_PART2",
		3: "QUESTION_TYPE_PART3",
		4: "QUESTION_TYPE_PART4",
		5: "QUESTION_TYPE_PART5",
		6: "QUESTION_TYPE_PART6",
		7: "QUESTION_TYPE_PART7",
	}
	QuestionType_value = map[string]int32{
		"QUESTION_TYPE_INVALID": 0,
		"QUESTION_TYPE_PART1":   1,
		"QUESTION_TYPE_PART2":   2,
		"QUESTION_TYPE_PART3":   3,
		"QUESTION_TYPE_PART4":   4,
		"QUESTION_TYPE_PART5":   5,
		"QUESTION_TYPE_PART6":   6,
		"QUESTION_TYPE_PART7":   7,
	}
)

func (x QuestionType) Enum() *QuestionType {
	p := new(QuestionType)
	*p = x
	return p
}

func (x QuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_quizzes_proto_enumTypes[1].Descriptor()
}

func (QuestionType) Type() protoreflect.EnumType {
	return &file_quizzes_proto_enumTypes[1]
}

func (x QuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionType.Descriptor instead.
func (QuestionType) EnumDescriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{1}
}

type State int32

const (
	State_INVALID State = 0
	State_DOING   State = 1
	State_DONE    State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "INVALID",
		1: "DOING",
		2: "DONE",
	}
	State_value = map[string]int32{
		"INVALID": 0,
		"DOING":   1,
		"DONE":    2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_quizzes_proto_enumTypes[2].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_quizzes_proto_enumTypes[2]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{2}
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Thumbnail string `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Topic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Topic) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

type Quiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId     int64                 `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	IsFree      *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=is_free,json=isFree,proto3" json:"is_free,omitempty"`
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{1}
}

func (x *Quiz) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quiz) GetTopicId() int64 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *Quiz) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Quiz) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Quiz) GetIsFree() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsFree
	}
	return nil
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QuestionText string       `protobuf:"bytes,2,opt,name=question_text,json=questionText,proto3" json:"question_text,omitempty"`
	Answers      string       `protobuf:"bytes,3,opt,name=answers,proto3" json:"answers,omitempty"`
	Eplain       string       `protobuf:"bytes,4,opt,name=eplain,proto3" json:"eplain,omitempty"`
	Explain      string       `protobuf:"bytes,5,opt,name=explain,proto3" json:"explain,omitempty"`
	QuestionType QuestionType `protobuf:"varint,6,opt,name=question_type,json=questionType,proto3,enum=go.todoapp.proto.task.QuestionType" json:"question_type,omitempty"`
	Image        string       `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	Audio        string       `protobuf:"bytes,8,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{2}
}

func (x *Question) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetQuestionText() string {
	if x != nil {
		return x.QuestionText
	}
	return ""
}

func (x *Question) GetAnswers() string {
	if x != nil {
		return x.Answers
	}
	return ""
}

func (x *Question) GetEplain() string {
	if x != nil {
		return x.Eplain
	}
	return ""
}

func (x *Question) GetExplain() string {
	if x != nil {
		return x.Explain
	}
	return ""
}

func (x *Question) GetQuestionType() QuestionType {
	if x != nil {
		return x.QuestionType
	}
	return QuestionType_QUESTION_TYPE_INVALID
}

func (x *Question) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Question) GetAudio() string {
	if x != nil {
		return x.Audio
	}
	return ""
}

type QuizReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TopicId       int64                 `protobuf:"varint,3,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	QuizId        int64                 `protobuf:"varint,4,opt,name=quiz_id,json=quizId,proto3" json:"quiz_id,omitempty"`
	Percentage    string                `protobuf:"bytes,5,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Duration      int64                 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	TotalDuration int64                 `protobuf:"varint,7,opt,name=total_duration,json=totalDuration,proto3" json:"total_duration,omitempty"`
	Result        *wrapperspb.BoolValue `protobuf:"bytes,8,opt,name=result,proto3" json:"result,omitempty"`
	StartTime     int64                 `protobuf:"varint,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	TotalQuiz     int64                 `protobuf:"varint,10,opt,name=total_quiz,json=totalQuiz,proto3" json:"total_quiz,omitempty"`
	State         State                 `protobuf:"varint,11,opt,name=state,proto3,enum=go.todoapp.proto.task.State" json:"state,omitempty"`
	Answers       string                `protobuf:"bytes,12,opt,name=answers,proto3" json:"answers,omitempty"`
}

func (x *QuizReport) Reset() {
	*x = QuizReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quizzes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuizReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuizReport) ProtoMessage() {}

func (x *QuizReport) ProtoReflect() protoreflect.Message {
	mi := &file_quizzes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuizReport.ProtoReflect.Descriptor instead.
func (*QuizReport) Descriptor() ([]byte, []int) {
	return file_quizzes_proto_rawDescGZIP(), []int{3}
}

func (x *QuizReport) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QuizReport) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QuizReport) GetTopicId() int64 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *QuizReport) GetQuizId() int64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

func (x *QuizReport) GetPercentage() string {
	if x != nil {
		return x.Percentage
	}
	return ""
}

func (x *QuizReport) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *QuizReport) GetTotalDuration() int64 {
	if x != nil {
		return x.TotalDuration
	}
	return 0
}

func (x *QuizReport) GetResult() *wrapperspb.BoolValue {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *QuizReport) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *QuizReport) GetTotalQuiz() int64 {
	if x != nil {
		return x.TotalQuiz
	}
	return 0
}

func (x *QuizReport) GetState() State {
	if x != nil {
		return x.State
	}
	return State_INVALID
}

func (x *QuizReport) GetAnswers() string {
	if x != nil {
		return x.Answers
	}
	return ""
}

var File_quizzes_proto protoreflect.FileDescriptor

var file_quizzes_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x67, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73, 0x46, 0x72, 0x65,
	0x65, 0x22, 0x81, 0x02, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12,
	0x48, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x0a, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x75, 0x69, 0x7a,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x71, 0x75, 0x69, 0x7a, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x71, 0x75, 0x69, 0x7a, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x73, 0x2a, 0x6e, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x11, 0x0a,
	0x0d, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x31, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x32, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x42, 0x31, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x42, 0x32, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x43, 0x31, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x43, 0x32, 0x10, 0x06, 0x2a, 0xd8, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x31, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x32,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x33, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52,
	0x54, 0x34, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x35, 0x10, 0x05, 0x12, 0x17, 0x0a,
	0x13, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x36, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x37, 0x10, 0x07, 0x2a,
	0x29, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quizzes_proto_rawDescOnce sync.Once
	file_quizzes_proto_rawDescData = file_quizzes_proto_rawDesc
)

func file_quizzes_proto_rawDescGZIP() []byte {
	file_quizzes_proto_rawDescOnce.Do(func() {
		file_quizzes_proto_rawDescData = protoimpl.X.CompressGZIP(file_quizzes_proto_rawDescData)
	})
	return file_quizzes_proto_rawDescData
}

var file_quizzes_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_quizzes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_quizzes_proto_goTypes = []interface{}{
	(Level)(0),                   // 0: go.todoapp.proto.task.Level
	(QuestionType)(0),            // 1: go.todoapp.proto.task.QuestionType
	(State)(0),                   // 2: go.todoapp.proto.task.State
	(*Topic)(nil),                // 3: go.todoapp.proto.task.Topic
	(*Quiz)(nil),                 // 4: go.todoapp.proto.task.Quiz
	(*Question)(nil),             // 5: go.todoapp.proto.task.Question
	(*QuizReport)(nil),           // 6: go.todoapp.proto.task.QuizReport
	(*wrapperspb.BoolValue)(nil), // 7: google.protobuf.BoolValue
}
var file_quizzes_proto_depIdxs = []int32{
	7, // 0: go.todoapp.proto.task.Quiz.is_free:type_name -> google.protobuf.BoolValue
	1, // 1: go.todoapp.proto.task.Question.question_type:type_name -> go.todoapp.proto.task.QuestionType
	7, // 2: go.todoapp.proto.task.QuizReport.result:type_name -> google.protobuf.BoolValue
	2, // 3: go.todoapp.proto.task.QuizReport.state:type_name -> go.todoapp.proto.task.State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_quizzes_proto_init() }
func file_quizzes_proto_init() {
	if File_quizzes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quizzes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_quizzes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quiz); i {
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
		file_quizzes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_quizzes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuizReport); i {
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
			RawDescriptor: file_quizzes_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quizzes_proto_goTypes,
		DependencyIndexes: file_quizzes_proto_depIdxs,
		EnumInfos:         file_quizzes_proto_enumTypes,
		MessageInfos:      file_quizzes_proto_msgTypes,
	}.Build()
	File_quizzes_proto = out.File
	file_quizzes_proto_rawDesc = nil
	file_quizzes_proto_goTypes = nil
	file_quizzes_proto_depIdxs = nil
}