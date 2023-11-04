// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: proto/quiz.proto

package proto

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

type GetQuestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQuestionsRequest) Reset() {
	*x = GetQuestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsRequest) ProtoMessage() {}

func (x *GetQuestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{0}
}

type GetQuestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *GetQuestionsResponse) Reset() {
	*x = GetQuestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsResponse) ProtoMessage() {}

func (x *GetQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text    string    `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Options []*Option `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[2]
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
	return file_proto_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *Question) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{3}
}

func (x *Option) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Option) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId     int32 `protobuf:"varint,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
	ChosenOptionId int32 `protobuf:"varint,2,opt,name=chosenOptionId,proto3" json:"chosenOptionId,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{4}
}

func (x *Answer) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answer) GetChosenOptionId() int32 {
	if x != nil {
		return x.ChosenOptionId
	}
	return 0
}

type RegisterAnswersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*Answer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *RegisterAnswersRequest) Reset() {
	*x = RegisterAnswersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAnswersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAnswersRequest) ProtoMessage() {}

func (x *RegisterAnswersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAnswersRequest.ProtoReflect.Descriptor instead.
func (*RegisterAnswersRequest) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterAnswersRequest) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type RegisterAnswersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizRating       int32             `protobuf:"varint,1,opt,name=quizRating,proto3" json:"quizRating,omitempty"`
	QuesitonsResults []*QuestionResult `protobuf:"bytes,2,rep,name=quesitonsResults,proto3" json:"quesitonsResults,omitempty"`
	Statistics       *Statistics       `protobuf:"bytes,3,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *RegisterAnswersResponse) Reset() {
	*x = RegisterAnswersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAnswersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAnswersResponse) ProtoMessage() {}

func (x *RegisterAnswersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAnswersResponse.ProtoReflect.Descriptor instead.
func (*RegisterAnswersResponse) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterAnswersResponse) GetQuizRating() int32 {
	if x != nil {
		return x.QuizRating
	}
	return 0
}

func (x *RegisterAnswersResponse) GetQuesitonsResults() []*QuestionResult {
	if x != nil {
		return x.QuesitonsResults
	}
	return nil
}

func (x *RegisterAnswersResponse) GetStatistics() *Statistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

type QuestionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId      int32 `protobuf:"varint,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
	ChosenOptionId  int32 `protobuf:"varint,2,opt,name=chosenOptionId,proto3" json:"chosenOptionId,omitempty"`
	CorrectOptionId int32 `protobuf:"varint,3,opt,name=correctOptionId,proto3" json:"correctOptionId,omitempty"`
	IsCorrect       bool  `protobuf:"varint,4,opt,name=isCorrect,proto3" json:"isCorrect,omitempty"`
}

func (x *QuestionResult) Reset() {
	*x = QuestionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionResult) ProtoMessage() {}

func (x *QuestionResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionResult.ProtoReflect.Descriptor instead.
func (*QuestionResult) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{7}
}

func (x *QuestionResult) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionResult) GetChosenOptionId() int32 {
	if x != nil {
		return x.ChosenOptionId
	}
	return 0
}

func (x *QuestionResult) GetCorrectOptionId() int32 {
	if x != nil {
		return x.CorrectOptionId
	}
	return 0
}

func (x *QuestionResult) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

type Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BetterThan float32 `protobuf:"fixed32,3,opt,name=betterThan,proto3" json:"betterThan,omitempty"`
}

func (x *Statistics) Reset() {
	*x = Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quiz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistics) ProtoMessage() {}

func (x *Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistics.ProtoReflect.Descriptor instead.
func (*Statistics) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{8}
}

func (x *Statistics) GetBetterThan() float32 {
	if x != nil {
		return x.BetterThan
	}
	return 0
}

var File_proto_quiz_proto protoreflect.FileDescriptor

var file_proto_quiz_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x71, 0x75, 0x69, 0x7a, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2c, 0x0a,
	0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x50, 0x0a, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63,
	0x68, 0x6f, 0x73, 0x65, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22,
	0xad, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x71,
	0x75, 0x69, 0x7a, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x71, 0x75, 0x69, 0x7a, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x10, 0x71,
	0x75, 0x65, 0x73, 0x69, 0x74, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x10, 0x71, 0x75, 0x65,
	0x73, 0x69, 0x74, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22,
	0xa0, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x73,
	0x65, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x22, 0x2c, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x65, 0x74, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x62, 0x65, 0x74, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e,
	0x32, 0xa8, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x19, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x75,
	0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quiz_proto_rawDescOnce sync.Once
	file_proto_quiz_proto_rawDescData = file_proto_quiz_proto_rawDesc
)

func file_proto_quiz_proto_rawDescGZIP() []byte {
	file_proto_quiz_proto_rawDescOnce.Do(func() {
		file_proto_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quiz_proto_rawDescData)
	})
	return file_proto_quiz_proto_rawDescData
}

var file_proto_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_quiz_proto_goTypes = []interface{}{
	(*GetQuestionsRequest)(nil),     // 0: quiz.GetQuestionsRequest
	(*GetQuestionsResponse)(nil),    // 1: quiz.GetQuestionsResponse
	(*Question)(nil),                // 2: quiz.Question
	(*Option)(nil),                  // 3: quiz.Option
	(*Answer)(nil),                  // 4: quiz.Answer
	(*RegisterAnswersRequest)(nil),  // 5: quiz.RegisterAnswersRequest
	(*RegisterAnswersResponse)(nil), // 6: quiz.RegisterAnswersResponse
	(*QuestionResult)(nil),          // 7: quiz.QuestionResult
	(*Statistics)(nil),              // 8: quiz.Statistics
}
var file_proto_quiz_proto_depIdxs = []int32{
	2, // 0: quiz.GetQuestionsResponse.questions:type_name -> quiz.Question
	3, // 1: quiz.Question.options:type_name -> quiz.Option
	4, // 2: quiz.RegisterAnswersRequest.answers:type_name -> quiz.Answer
	7, // 3: quiz.RegisterAnswersResponse.quesitonsResults:type_name -> quiz.QuestionResult
	8, // 4: quiz.RegisterAnswersResponse.statistics:type_name -> quiz.Statistics
	0, // 5: quiz.QuizService.GetQuestions:input_type -> quiz.GetQuestionsRequest
	5, // 6: quiz.QuizService.RegisterAnswers:input_type -> quiz.RegisterAnswersRequest
	1, // 7: quiz.QuizService.GetQuestions:output_type -> quiz.GetQuestionsResponse
	6, // 8: quiz.QuizService.RegisterAnswers:output_type -> quiz.RegisterAnswersResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_quiz_proto_init() }
func file_proto_quiz_proto_init() {
	if File_proto_quiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quiz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionsRequest); i {
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
		file_proto_quiz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionsResponse); i {
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
		file_proto_quiz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_quiz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_proto_quiz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_proto_quiz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAnswersRequest); i {
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
		file_proto_quiz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAnswersResponse); i {
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
		file_proto_quiz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionResult); i {
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
		file_proto_quiz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistics); i {
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
			RawDescriptor: file_proto_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quiz_proto_goTypes,
		DependencyIndexes: file_proto_quiz_proto_depIdxs,
		MessageInfos:      file_proto_quiz_proto_msgTypes,
	}.Build()
	File_proto_quiz_proto = out.File
	file_proto_quiz_proto_rawDesc = nil
	file_proto_quiz_proto_goTypes = nil
	file_proto_quiz_proto_depIdxs = nil
}
