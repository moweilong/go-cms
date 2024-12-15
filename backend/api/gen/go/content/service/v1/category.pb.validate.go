// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: content/service/v1/category.proto

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

// Validate checks the field values on Category with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Category) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Category with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CategoryMultiError, or nil
// if none found.
func (m *Category) ValidateAll() error {
	return m.validate(true)
}

func (m *Category) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Slug != nil {
		// no validation rules for Slug
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.Thumbnail != nil {
		// no validation rules for Thumbnail
	}

	if m.Password != nil {
		// no validation rules for Password
	}

	if m.FullPath != nil {
		// no validation rules for FullPath
	}

	if m.Priority != nil {
		// no validation rules for Priority
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if m.UpdateTime != nil {
		// no validation rules for UpdateTime
	}

	if m.PostCount != nil {
		// no validation rules for PostCount
	}

	if len(errors) > 0 {
		return CategoryMultiError(errors)
	}

	return nil
}

// CategoryMultiError is an error wrapping multiple validation errors returned
// by Category.ValidateAll() if the designated constraints aren't met.
type CategoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryMultiError) AllErrors() []error { return m }

// CategoryValidationError is the validation error returned by
// Category.Validate if the designated constraints aren't met.
type CategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryValidationError) ErrorName() string { return "CategoryValidationError" }

// Error satisfies the builtin error interface
func (e CategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryValidationError{}

// Validate checks the field values on ListCategoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCategoryResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCategoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCategoryResponseMultiError, or nil if none found.
func (m *ListCategoryResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCategoryResponse) validate(all bool) error {
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
					errors = append(errors, ListCategoryResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCategoryResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCategoryResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListCategoryResponseMultiError(errors)
	}

	return nil
}

// ListCategoryResponseMultiError is an error wrapping multiple validation
// errors returned by ListCategoryResponse.ValidateAll() if the designated
// constraints aren't met.
type ListCategoryResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCategoryResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCategoryResponseMultiError) AllErrors() []error { return m }

// ListCategoryResponseValidationError is the validation error returned by
// ListCategoryResponse.Validate if the designated constraints aren't met.
type ListCategoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCategoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCategoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCategoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCategoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCategoryResponseValidationError) ErrorName() string {
	return "ListCategoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCategoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCategoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCategoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCategoryResponseValidationError{}

// Validate checks the field values on GetCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCategoryRequestMultiError, or nil if none found.
func (m *GetCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetFieldMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCategoryRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCategoryRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCategoryRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCategoryRequestMultiError(errors)
	}

	return nil
}

// GetCategoryRequestMultiError is an error wrapping multiple validation errors
// returned by GetCategoryRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCategoryRequestMultiError) AllErrors() []error { return m }

// GetCategoryRequestValidationError is the validation error returned by
// GetCategoryRequest.Validate if the designated constraints aren't met.
type GetCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCategoryRequestValidationError) ErrorName() string {
	return "GetCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCategoryRequestValidationError{}

// Validate checks the field values on CreateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCategoryRequestMultiError, or nil if none found.
func (m *CreateCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCategoryRequestValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreateCategoryRequestMultiError(errors)
	}

	return nil
}

// CreateCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCategoryRequestMultiError) AllErrors() []error { return m }

// CreateCategoryRequestValidationError is the validation error returned by
// CreateCategoryRequest.Validate if the designated constraints aren't met.
type CreateCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCategoryRequestValidationError) ErrorName() string {
	return "CreateCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCategoryRequestValidationError{}

// Validate checks the field values on UpdateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCategoryRequestMultiError, or nil if none found.
func (m *UpdateCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCategoryRequestValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return UpdateCategoryRequestMultiError(errors)
	}

	return nil
}

// UpdateCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCategoryRequestMultiError) AllErrors() []error { return m }

// UpdateCategoryRequestValidationError is the validation error returned by
// UpdateCategoryRequest.Validate if the designated constraints aren't met.
type UpdateCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCategoryRequestValidationError) ErrorName() string {
	return "UpdateCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCategoryRequestValidationError{}

// Validate checks the field values on DeleteCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCategoryRequestMultiError, or nil if none found.
func (m *DeleteCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteCategoryRequestMultiError(errors)
	}

	return nil
}

// DeleteCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCategoryRequestMultiError) AllErrors() []error { return m }

// DeleteCategoryRequestValidationError is the validation error returned by
// DeleteCategoryRequest.Validate if the designated constraints aren't met.
type DeleteCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCategoryRequestValidationError) ErrorName() string {
	return "DeleteCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCategoryRequestValidationError{}
