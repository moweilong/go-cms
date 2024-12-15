// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package servicev1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 400
func IsFrontErrorReasonBadRequestUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_BAD_REQUEST_UNSPECIFIED.String() && e.Code == 400
}

// 400
func ErrorFrontErrorReasonBadRequestUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(400, FrontErrorReason_FRONT_ERROR_REASON_BAD_REQUEST_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

// 401
func IsFrontErrorReasonNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_NOT_LOGGED_IN.String() && e.Code == 401
}

// 401
func ErrorFrontErrorReasonNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, FrontErrorReason_FRONT_ERROR_REASON_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 403
func IsFrontErrorReasonAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_ACCESS_FORBIDDEN.String() && e.Code == 403
}

// 403
func ErrorFrontErrorReasonAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, FrontErrorReason_FRONT_ERROR_REASON_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

// 404
func IsFrontErrorReasonResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 404
func ErrorFrontErrorReasonResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, FrontErrorReason_FRONT_ERROR_REASON_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 405
func IsFrontErrorReasonMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

// 405
func ErrorFrontErrorReasonMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, FrontErrorReason_FRONT_ERROR_REASON_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

// 408
func IsFrontErrorReasonRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_REQUEST_TIMEOUT.String() && e.Code == 408
}

// 408
func ErrorFrontErrorReasonRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, FrontErrorReason_FRONT_ERROR_REASON_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 500
func IsFrontErrorReasonInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

// 500
func ErrorFrontErrorReasonInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, FrontErrorReason_FRONT_ERROR_REASON_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

// 501
func IsFrontErrorReasonNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_NOT_IMPLEMENTED.String() && e.Code == 501
}

// 501
func ErrorFrontErrorReasonNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, FrontErrorReason_FRONT_ERROR_REASON_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

// 502
func IsFrontErrorReasonNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_NETWORK_ERROR.String() && e.Code == 502
}

// 502
func ErrorFrontErrorReasonNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, FrontErrorReason_FRONT_ERROR_REASON_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

// 503
func IsFrontErrorReasonServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

// 503
func ErrorFrontErrorReasonServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, FrontErrorReason_FRONT_ERROR_REASON_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

// 504
func IsFrontErrorReasonNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_NETWORK_TIMEOUT.String() && e.Code == 504
}

// 504
func ErrorFrontErrorReasonNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, FrontErrorReason_FRONT_ERROR_REASON_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 505
func IsFrontErrorReasonRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

// 505
func ErrorFrontErrorReasonRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, FrontErrorReason_FRONT_ERROR_REASON_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsFrontErrorReasonInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_INVALID_USERID.String() && e.Code == 101
}

// 用户ID无效
func ErrorFrontErrorReasonInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(101, FrontErrorReason_FRONT_ERROR_REASON_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsFrontErrorReasonInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_INVALID_PASSWORD.String() && e.Code == 102
}

// 密码无效
func ErrorFrontErrorReasonInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(102, FrontErrorReason_FRONT_ERROR_REASON_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsFrontErrorReasonTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_TOKEN_EXPIRED.String() && e.Code == 103
}

// token过期
func ErrorFrontErrorReasonTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(103, FrontErrorReason_FRONT_ERROR_REASON_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// token无效
func IsFrontErrorReasonInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_INVALID_TOKEN.String() && e.Code == 104
}

// token无效
func ErrorFrontErrorReasonInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(104, FrontErrorReason_FRONT_ERROR_REASON_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

// token不存在
func IsFrontErrorReasonTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_TOKEN_NOT_EXIST.String() && e.Code == 105
}

// token不存在
func ErrorFrontErrorReasonTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(105, FrontErrorReason_FRONT_ERROR_REASON_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsFrontErrorReasonUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_USER_NOT_EXIST.String() && e.Code == 106
}

// 用户不存在
func ErrorFrontErrorReasonUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(106, FrontErrorReason_FRONT_ERROR_REASON_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 账号冻结
func IsFrontErrorReasonUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == FrontErrorReason_FRONT_ERROR_REASON_USER_FREEZE.String() && e.Code == 107
}

// 账号冻结
func ErrorFrontErrorReasonUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(107, FrontErrorReason_FRONT_ERROR_REASON_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}
