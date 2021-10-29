// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/service/v1/shop.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		return RegisterRequestValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}

	}

	if l := utf8.RuneCountInString(m.GetPass()); l < 6 || l > 18 {
		return RegisterRequestValidationError{
			field:  "Pass",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
	}

	// no validation rules for NickName

	// no validation rules for Age

	return nil
}

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Mobile

	// no validation rules for NickName

	// no validation rules for Age

	return nil
}

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoginRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		return LoginRequestValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}

	}

	if l := utf8.RuneCountInString(m.GetPass()); l < 6 || l > 18 {
		return LoginRequestValidationError{
			field:  "Pass",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
	}

	return nil
}

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Msg

	return nil
}

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

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

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on LoginReply_Data with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoginReply_Data) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// LoginReply_DataValidationError is the validation error returned by
// LoginReply_Data.Validate if the designated constraints aren't met.
type LoginReply_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReply_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReply_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReply_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReply_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReply_DataValidationError) ErrorName() string { return "LoginReply_DataValidationError" }

// Error satisfies the builtin error interface
func (e LoginReply_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReply_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReply_DataValidationError{}

// Validate checks the field values on GetUserReply_Data with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUserReply_Data) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// GetUserReply_DataValidationError is the validation error returned by
// GetUserReply_Data.Validate if the designated constraints aren't met.
type GetUserReply_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReply_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReply_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReply_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReply_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReply_DataValidationError) ErrorName() string {
	return "GetUserReply_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserReply_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReply_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReply_DataValidationError{}
