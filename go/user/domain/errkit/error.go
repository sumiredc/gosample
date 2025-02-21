package errkit

import (
	"fmt"
)

const (
	// 0xx - 共通
	ErrUnknownCode        = "E0000"
	ErrInternalServerCode = "E0001"
	ErrUnknownMsg         = "An unknown error occurred"
	ErrInternalServerMsg  = "Internal server error"

	// 1xx - 認証・認可
	ErrAuthorizationCode  = "E1001"
	ErrAuthenticationCode = "E1002"
	ErrTokenExpiredCode   = "E1003"
	ErrAuthorizationMsg   = "Authorization failed"
	ErrAuthenticationMsg  = "Authentication failed"
	ErrTokenExpiredMsg    = "Token has expired"

	// 2xx - バリデーション
	ErrValidationCode       = "E2001"
	ErrInvalidFormatCode    = "E2002"
	ErrMissingParameterCode = "E2003"
	ErrValidationMsg        = "Validation error occurred"
	ErrInvalidFormatMsg     = "Invalid format"
	ErrMissingParameterMsg  = "Missing required parameter"

	// 3xx - リソース関連
	ErrNotFoundCode      = "E3001"
	ErrAlreadyExistsCode = "E3002"
	ErrNotFoundMsg       = "Resource not found"
	ErrAlreadyExistsMsg  = "Resource already exists"

	// 4xx - 外部API
	ErrExternalApiCode = "E4001"
	ErrTimeoutCode     = "E4002"
	ErrExternalApiMsg  = "External API error"
	ErrTimeoutMsg      = "Request timed out"

	// 5xx - データベース
	ErrDatabaseConnectionCode = "E5001"
	ErrUniqueConstraintCode   = "E5002"
	ErrDatabaseConnectionMsg  = "Database connection error"
	ErrUniqueConstraintMsg    = "Unique constraint violation"

	// 6xx - ビジネスロジック
	ErrBusinessRuleCode = "E6001"
	ErrBusinessRuleMsg  = "Business rule violation"

	// 7xx - 権限
	ErrForbiddenCode = "E7001"
	ErrForbiddenMsg  = "Access denied"

	// 8xx - リクエスト
	ErrBadRequestCode = "E8001"
	ErrBadRequestMsg  = "Bad request"

	// 9xx - システム・インフラ
	ErrSystemCode  = "E9001"
	ErrNetworkCode = "E9002"
	ErrSystemMsg   = "System error occurred"
	ErrNetworkMsg  = "Network error occurred"
)

var (
	// 0xx - 共通
	ErrUnknown        = fmt.Errorf("[%s] %s", ErrUnknownCode, ErrUnknownMsg)
	ErrInternalServer = fmt.Errorf("[%s] %s", ErrInternalServerCode, ErrInternalServerMsg)

	// 1xx - 認証・認可
	ErrAuthorization  = fmt.Errorf("[%s] %s", ErrAuthorizationCode, ErrAuthorizationMsg)
	ErrAuthentication = fmt.Errorf("[%s] %s", ErrAuthenticationCode, ErrAuthenticationMsg)
	ErrTokenExpired   = fmt.Errorf("[%s] %s", ErrTokenExpiredCode, ErrTokenExpiredMsg)

	// 2xx - バリデーション
	ErrValidation       = fmt.Errorf("[%s] %s", ErrValidationCode, ErrValidationMsg)
	ErrInvalidFormat    = fmt.Errorf("[%s] %s", ErrInvalidFormatCode, ErrInvalidFormatMsg)
	ErrMissingParameter = fmt.Errorf("[%s] %s", ErrMissingParameterCode, ErrMissingParameterMsg)

	// 3xx - リソース関連
	ErrNotFound      = fmt.Errorf("[%s] %s", ErrNotFoundCode, ErrNotFoundMsg)
	ErrAlreadyExists = fmt.Errorf("[%s] %s", ErrAlreadyExistsCode, ErrAlreadyExistsMsg)

	// 4xx - 外部API
	ErrExternalApi = fmt.Errorf("[%s] %s", ErrExternalApiCode, ErrExternalApiMsg)
	ErrTimeout     = fmt.Errorf("[%s] %s", ErrTimeoutCode, ErrTimeoutMsg)

	// 5xx - データベース
	ErrDatabaseConnection = fmt.Errorf("[%s] %s", ErrDatabaseConnectionCode, ErrDatabaseConnectionMsg)
	ErrUniqueConstraint   = fmt.Errorf("[%s] %s", ErrUniqueConstraintCode, ErrUniqueConstraintMsg)

	// 6xx - ビジネスロジック
	ErrBusinessRule = fmt.Errorf("[%s] %s", ErrBusinessRuleCode, ErrBusinessRuleMsg)

	// 7xx - 権限
	ErrForbidden = fmt.Errorf("[%s] %s", ErrForbiddenCode, ErrForbiddenMsg)

	// 8xx - リクエスト
	ErrBadRequest = fmt.Errorf("[%s] %s", ErrBadRequestCode, ErrBadRequestMsg)

	// 9xx - システム・インフラ
	ErrSystem  = fmt.Errorf("[%s] %s", ErrSystemCode, ErrSystemMsg)
	ErrNetwork = fmt.Errorf("[%s] %s", ErrNetworkCode, ErrNetworkMsg)
)
