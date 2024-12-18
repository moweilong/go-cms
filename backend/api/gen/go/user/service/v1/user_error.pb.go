// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: user/service/v1/user_error.proto

package servicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type UserErrorReason int32

const (
	UserErrorReason_USER_ERROR_REASON_NOT_LOGGED_IN_UNSPECIFIED UserErrorReason = 0  //401
	UserErrorReason_USER_ERROR_REASON_ACCESS_FORBIDDEN          UserErrorReason = 1  // 403
	UserErrorReason_USER_ERROR_REASON_RESOURCE_NOT_FOUND        UserErrorReason = 2  // 404
	UserErrorReason_USER_ERROR_REASON_METHOD_NOT_ALLOWED        UserErrorReason = 3  // 405
	UserErrorReason_USER_ERROR_REASON_REQUEST_TIMEOUT           UserErrorReason = 4  // 408
	UserErrorReason_USER_ERROR_REASON_INTERNAL_SERVER_ERROR     UserErrorReason = 5  // 500
	UserErrorReason_USER_ERROR_REASON_NOT_IMPLEMENTED           UserErrorReason = 6  // 501
	UserErrorReason_USER_ERROR_REASON_NETWORK_ERROR             UserErrorReason = 7  // 502
	UserErrorReason_USER_ERROR_REASON_SERVICE_UNAVAILABLE       UserErrorReason = 8  // 503
	UserErrorReason_USER_ERROR_REASON_NETWORK_TIMEOUT           UserErrorReason = 9  // 504
	UserErrorReason_USER_ERROR_REASON_REQUEST_NOT_SUPPORT       UserErrorReason = 10 // 505
	UserErrorReason_USER_ERROR_REASON_USER_NOT_FOUND            UserErrorReason = 11
	UserErrorReason_USER_ERROR_REASON_INCORRECT_PASSWORD        UserErrorReason = 12
	UserErrorReason_USER_ERROR_REASON_USER_FREEZE               UserErrorReason = 13
	UserErrorReason_USER_ERROR_REASON_INVALID_USERID            UserErrorReason = 14 // 用户ID无效
	UserErrorReason_USER_ERROR_REASON_INVALID_PASSWORD          UserErrorReason = 15 // 密码无效
	UserErrorReason_USER_ERROR_REASON_TOKEN_EXPIRED             UserErrorReason = 16 // token过期
	UserErrorReason_USER_ERROR_REASON_INVALID_TOKEN             UserErrorReason = 17 // token无效
	UserErrorReason_USER_ERROR_REASON_TOKEN_NOT_EXIST           UserErrorReason = 18 // token不存在
	UserErrorReason_USER_ERROR_REASON_USER_NOT_EXIST            UserErrorReason = 19 // 用户不存在
)

// Enum value maps for UserErrorReason.
var (
	UserErrorReason_name = map[int32]string{
		0:  "USER_ERROR_REASON_NOT_LOGGED_IN_UNSPECIFIED",
		1:  "USER_ERROR_REASON_ACCESS_FORBIDDEN",
		2:  "USER_ERROR_REASON_RESOURCE_NOT_FOUND",
		3:  "USER_ERROR_REASON_METHOD_NOT_ALLOWED",
		4:  "USER_ERROR_REASON_REQUEST_TIMEOUT",
		5:  "USER_ERROR_REASON_INTERNAL_SERVER_ERROR",
		6:  "USER_ERROR_REASON_NOT_IMPLEMENTED",
		7:  "USER_ERROR_REASON_NETWORK_ERROR",
		8:  "USER_ERROR_REASON_SERVICE_UNAVAILABLE",
		9:  "USER_ERROR_REASON_NETWORK_TIMEOUT",
		10: "USER_ERROR_REASON_REQUEST_NOT_SUPPORT",
		11: "USER_ERROR_REASON_USER_NOT_FOUND",
		12: "USER_ERROR_REASON_INCORRECT_PASSWORD",
		13: "USER_ERROR_REASON_USER_FREEZE",
		14: "USER_ERROR_REASON_INVALID_USERID",
		15: "USER_ERROR_REASON_INVALID_PASSWORD",
		16: "USER_ERROR_REASON_TOKEN_EXPIRED",
		17: "USER_ERROR_REASON_INVALID_TOKEN",
		18: "USER_ERROR_REASON_TOKEN_NOT_EXIST",
		19: "USER_ERROR_REASON_USER_NOT_EXIST",
	}
	UserErrorReason_value = map[string]int32{
		"USER_ERROR_REASON_NOT_LOGGED_IN_UNSPECIFIED": 0,
		"USER_ERROR_REASON_ACCESS_FORBIDDEN":          1,
		"USER_ERROR_REASON_RESOURCE_NOT_FOUND":        2,
		"USER_ERROR_REASON_METHOD_NOT_ALLOWED":        3,
		"USER_ERROR_REASON_REQUEST_TIMEOUT":           4,
		"USER_ERROR_REASON_INTERNAL_SERVER_ERROR":     5,
		"USER_ERROR_REASON_NOT_IMPLEMENTED":           6,
		"USER_ERROR_REASON_NETWORK_ERROR":             7,
		"USER_ERROR_REASON_SERVICE_UNAVAILABLE":       8,
		"USER_ERROR_REASON_NETWORK_TIMEOUT":           9,
		"USER_ERROR_REASON_REQUEST_NOT_SUPPORT":       10,
		"USER_ERROR_REASON_USER_NOT_FOUND":            11,
		"USER_ERROR_REASON_INCORRECT_PASSWORD":        12,
		"USER_ERROR_REASON_USER_FREEZE":               13,
		"USER_ERROR_REASON_INVALID_USERID":            14,
		"USER_ERROR_REASON_INVALID_PASSWORD":          15,
		"USER_ERROR_REASON_TOKEN_EXPIRED":             16,
		"USER_ERROR_REASON_INVALID_TOKEN":             17,
		"USER_ERROR_REASON_TOKEN_NOT_EXIST":           18,
		"USER_ERROR_REASON_USER_NOT_EXIST":            19,
	}
)

func (x UserErrorReason) Enum() *UserErrorReason {
	p := new(UserErrorReason)
	*p = x
	return p
}

func (x UserErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_v1_user_error_proto_enumTypes[0].Descriptor()
}

func (UserErrorReason) Type() protoreflect.EnumType {
	return &file_user_service_v1_user_error_proto_enumTypes[0]
}

func (x UserErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserErrorReason.Descriptor instead.
func (UserErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_user_error_proto_rawDescGZIP(), []int{0}
}

var File_user_service_v1_user_error_proto protoreflect.FileDescriptor

var file_user_service_v1_user_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xab, 0x07, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x2b,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8,
	0x45, 0x91, 0x03, 0x12, 0x2c, 0x0a, 0x22, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x93,
	0x03, 0x12, 0x2e, 0x0a, 0x24, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94,
	0x03, 0x12, 0x2e, 0x0a, 0x24, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x95,
	0x03, 0x12, 0x2b, 0x0a, 0x21, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x31,
	0x0a, 0x27, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x2b, 0x0a, 0x21, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x29,
	0x0a, 0x1f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x2f, 0x0a, 0x25, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x2b, 0x0a, 0x21, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x09, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x12, 0x2f, 0x0a, 0x25, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x12, 0x2a, 0x0a, 0x20, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x1a, 0x04,
	0xa8, 0x45, 0xd8, 0x04, 0x12, 0x2e, 0x0a, 0x24, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52,
	0x45, 0x43, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x0c, 0x1a, 0x04,
	0xa8, 0x45, 0xd7, 0x04, 0x12, 0x27, 0x0a, 0x1d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x46,
	0x52, 0x45, 0x45, 0x5a, 0x45, 0x10, 0x0d, 0x1a, 0x04, 0xa8, 0x45, 0xd6, 0x04, 0x12, 0x29, 0x0a,
	0x20, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x49,
	0x44, 0x10, 0x0e, 0x1a, 0x03, 0xa8, 0x45, 0x65, 0x12, 0x2b, 0x0a, 0x22, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x0f,
	0x1a, 0x03, 0xa8, 0x45, 0x66, 0x12, 0x28, 0x0a, 0x1f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x10, 0x1a, 0x03, 0xa8, 0x45, 0x67, 0x12,
	0x28, 0x0a, 0x1f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x10, 0x11, 0x1a, 0x03, 0xa8, 0x45, 0x68, 0x12, 0x2a, 0x0a, 0x21, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x12,
	0x1a, 0x03, 0xa8, 0x45, 0x69, 0x12, 0x29, 0x0a, 0x20, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x13, 0x1a, 0x03, 0xa8, 0x45, 0x6a,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xb0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2b, 0x67, 0x6f, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x55, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_service_v1_user_error_proto_rawDescOnce sync.Once
	file_user_service_v1_user_error_proto_rawDescData = file_user_service_v1_user_error_proto_rawDesc
)

func file_user_service_v1_user_error_proto_rawDescGZIP() []byte {
	file_user_service_v1_user_error_proto_rawDescOnce.Do(func() {
		file_user_service_v1_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_v1_user_error_proto_rawDescData)
	})
	return file_user_service_v1_user_error_proto_rawDescData
}

var file_user_service_v1_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_v1_user_error_proto_goTypes = []any{
	(UserErrorReason)(0), // 0: user.service.v1.UserErrorReason
}
var file_user_service_v1_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_v1_user_error_proto_init() }
func file_user_service_v1_user_error_proto_init() {
	if File_user_service_v1_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_v1_user_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_service_v1_user_error_proto_goTypes,
		DependencyIndexes: file_user_service_v1_user_error_proto_depIdxs,
		EnumInfos:         file_user_service_v1_user_error_proto_enumTypes,
	}.Build()
	File_user_service_v1_user_error_proto = out.File
	file_user_service_v1_user_error_proto_rawDesc = nil
	file_user_service_v1_user_error_proto_goTypes = nil
	file_user_service_v1_user_error_proto_depIdxs = nil
}
