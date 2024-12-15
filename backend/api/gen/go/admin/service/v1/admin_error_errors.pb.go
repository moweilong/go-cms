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
func IsAdminErrorReasonBadRequestUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_BAD_REQUEST_UNSPECIFIED.String() && e.Code == 400
}

// 400
func ErrorAdminErrorReasonBadRequestUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(400, AdminErrorReason_ADMIN_ERROR_REASON_BAD_REQUEST_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

// 401
func IsAdminErrorReasonNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_NOT_LOGGED_IN.String() && e.Code == 401
}

// 401
func ErrorAdminErrorReasonNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, AdminErrorReason_ADMIN_ERROR_REASON_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 403
func IsAdminErrorReasonAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_ACCESS_FORBIDDEN.String() && e.Code == 403
}

// 403
func ErrorAdminErrorReasonAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, AdminErrorReason_ADMIN_ERROR_REASON_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

// 404
func IsAdminErrorReasonResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 404
func ErrorAdminErrorReasonResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, AdminErrorReason_ADMIN_ERROR_REASON_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 405
func IsAdminErrorReasonMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

// 405
func ErrorAdminErrorReasonMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, AdminErrorReason_ADMIN_ERROR_REASON_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

// 408
func IsAdminErrorReasonRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_REQUEST_TIMEOUT.String() && e.Code == 408
}

// 408
func ErrorAdminErrorReasonRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, AdminErrorReason_ADMIN_ERROR_REASON_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 500
func IsAdminErrorReasonInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

// 500
func ErrorAdminErrorReasonInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, AdminErrorReason_ADMIN_ERROR_REASON_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

// 501
func IsAdminErrorReasonNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_NOT_IMPLEMENTED.String() && e.Code == 501
}

// 501
func ErrorAdminErrorReasonNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, AdminErrorReason_ADMIN_ERROR_REASON_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

// 502
func IsAdminErrorReasonNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_NETWORK_ERROR.String() && e.Code == 502
}

// 502
func ErrorAdminErrorReasonNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, AdminErrorReason_ADMIN_ERROR_REASON_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

// 503
func IsAdminErrorReasonServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

// 503
func ErrorAdminErrorReasonServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, AdminErrorReason_ADMIN_ERROR_REASON_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

// 504
func IsAdminErrorReasonNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_NETWORK_TIMEOUT.String() && e.Code == 504
}

// 504
func ErrorAdminErrorReasonNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, AdminErrorReason_ADMIN_ERROR_REASON_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 505
func IsAdminErrorReasonRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

// 505
func ErrorAdminErrorReasonRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, AdminErrorReason_ADMIN_ERROR_REASON_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsAdminErrorReasonInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_INVALID_USERID.String() && e.Code == 101
}

// 用户ID无效
func ErrorAdminErrorReasonInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(101, AdminErrorReason_ADMIN_ERROR_REASON_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsAdminErrorReasonInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_INVALID_PASSWORD.String() && e.Code == 102
}

// 密码无效
func ErrorAdminErrorReasonInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(102, AdminErrorReason_ADMIN_ERROR_REASON_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsAdminErrorReasonTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_TOKEN_EXPIRED.String() && e.Code == 103
}

// token过期
func ErrorAdminErrorReasonTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(103, AdminErrorReason_ADMIN_ERROR_REASON_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// token无效
func IsAdminErrorReasonInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_INVALID_TOKEN.String() && e.Code == 104
}

// token无效
func ErrorAdminErrorReasonInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(104, AdminErrorReason_ADMIN_ERROR_REASON_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

// token不存在
func IsAdminErrorReasonTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_TOKEN_NOT_EXIST.String() && e.Code == 105
}

// token不存在
func ErrorAdminErrorReasonTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(105, AdminErrorReason_ADMIN_ERROR_REASON_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsAdminErrorReasonUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_USER_NOT_EXIST.String() && e.Code == 106
}

// 用户不存在
func ErrorAdminErrorReasonUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(106, AdminErrorReason_ADMIN_ERROR_REASON_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 账号冻结
func IsAdminErrorReasonUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == AdminErrorReason_ADMIN_ERROR_REASON_USER_FREEZE.String() && e.Code == 107
}

// 账号冻结
func ErrorAdminErrorReasonUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(107, AdminErrorReason_ADMIN_ERROR_REASON_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}
