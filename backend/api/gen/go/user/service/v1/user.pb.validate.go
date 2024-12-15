// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/service/v1/user.proto

package servicev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.UserName != nil {
		// no validation rules for UserName
	}

	if m.NickName != nil {
		// no validation rules for NickName
	}

	if m.Email != nil {
		// no validation rules for Email
	}

	if m.Avatar != nil {
		// no validation rules for Avatar
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.Password != nil {
		// no validation rules for Password
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if m.UpdateTime != nil {
		// no validation rules for UpdateTime
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.RoleId != nil {
		// no validation rules for RoleId
	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.Authority != nil {
		// no validation rules for Authority
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on ListUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserResponseMultiError, or nil if none found.
func (m *ListUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListUserResponseMultiError(errors)
	}

	return nil
}

// ListUserResponseMultiError is an error wrapping multiple validation errors
// returned by ListUserResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserResponseMultiError) AllErrors() []error { return m }

// ListUserResponseValidationError is the validation error returned by
// ListUserResponse.Validate if the designated constraints aren't met.
type ListUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserResponseValidationError) ErrorName() string { return "ListUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserResponseValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserByUserNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserByUserNameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserByUserNameRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserByUserNameRequestMultiError, or nil if none found.
func (m *GetUserByUserNameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserByUserNameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	if len(errors) > 0 {
		return GetUserByUserNameRequestMultiError(errors)
	}

	return nil
}

// GetUserByUserNameRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserByUserNameRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserByUserNameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserByUserNameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserByUserNameRequestMultiError) AllErrors() []error { return m }

// GetUserByUserNameRequestValidationError is the validation error returned by
// GetUserByUserNameRequest.Validate if the designated constraints aren't met.
type GetUserByUserNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserByUserNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserByUserNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserByUserNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserByUserNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserByUserNameRequestValidationError) ErrorName() string {
	return "GetUserByUserNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserByUserNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserByUserNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserByUserNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserByUserNameRequestValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on DeleteUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserRequestMultiError, or nil if none found.
func (m *DeleteUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return DeleteUserRequestMultiError(errors)
	}

	return nil
}

// DeleteUserRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserRequestMultiError) AllErrors() []error { return m }

// DeleteUserRequestValidationError is the validation error returned by
// DeleteUserRequest.Validate if the designated constraints aren't met.
type DeleteUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRequestValidationError) ErrorName() string {
	return "DeleteUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRequestValidationError{}

// Validate checks the field values on VerifyPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyPasswordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyPasswordRequestMultiError, or nil if none found.
func (m *VerifyPasswordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyPasswordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	// no validation rules for Password

	if len(errors) > 0 {
		return VerifyPasswordRequestMultiError(errors)
	}

	return nil
}

// VerifyPasswordRequestMultiError is an error wrapping multiple validation
// errors returned by VerifyPasswordRequest.ValidateAll() if the designated
// constraints aren't met.
type VerifyPasswordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyPasswordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyPasswordRequestMultiError) AllErrors() []error { return m }

// VerifyPasswordRequestValidationError is the validation error returned by
// VerifyPasswordRequest.Validate if the designated constraints aren't met.
type VerifyPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPasswordRequestValidationError) ErrorName() string {
	return "VerifyPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPasswordRequestValidationError{}

// Validate checks the field values on VerifyPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyPasswordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyPasswordResponseMultiError, or nil if none found.
func (m *VerifyPasswordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyPasswordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return VerifyPasswordResponseMultiError(errors)
	}

	return nil
}

// VerifyPasswordResponseMultiError is an error wrapping multiple validation
// errors returned by VerifyPasswordResponse.ValidateAll() if the designated
// constraints aren't met.
type VerifyPasswordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyPasswordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyPasswordResponseMultiError) AllErrors() []error { return m }

// VerifyPasswordResponseValidationError is the validation error returned by
// VerifyPasswordResponse.Validate if the designated constraints aren't met.
type VerifyPasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPasswordResponseValidationError) ErrorName() string {
	return "VerifyPasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPasswordResponseValidationError{}

// Validate checks the field values on UserExistsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserExistsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserExistsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserExistsRequestMultiError, or nil if none found.
func (m *UserExistsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserExistsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	if len(errors) > 0 {
		return UserExistsRequestMultiError(errors)
	}

	return nil
}

// UserExistsRequestMultiError is an error wrapping multiple validation errors
// returned by UserExistsRequest.ValidateAll() if the designated constraints
// aren't met.
type UserExistsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserExistsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserExistsRequestMultiError) AllErrors() []error { return m }

// UserExistsRequestValidationError is the validation error returned by
// UserExistsRequest.Validate if the designated constraints aren't met.
type UserExistsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserExistsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserExistsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserExistsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserExistsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserExistsRequestValidationError) ErrorName() string {
	return "UserExistsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserExistsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserExistsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserExistsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserExistsRequestValidationError{}

// Validate checks the field values on UserExistsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserExistsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserExistsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserExistsResponseMultiError, or nil if none found.
func (m *UserExistsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserExistsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Exist

	if len(errors) > 0 {
		return UserExistsResponseMultiError(errors)
	}

	return nil
}

// UserExistsResponseMultiError is an error wrapping multiple validation errors
// returned by UserExistsResponse.ValidateAll() if the designated constraints
// aren't met.
type UserExistsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserExistsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserExistsResponseMultiError) AllErrors() []error { return m }

// UserExistsResponseValidationError is the validation error returned by
// UserExistsResponse.Validate if the designated constraints aren't met.
type UserExistsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserExistsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserExistsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserExistsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserExistsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserExistsResponseValidationError) ErrorName() string {
	return "UserExistsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserExistsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserExistsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserExistsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserExistsResponseValidationError{}