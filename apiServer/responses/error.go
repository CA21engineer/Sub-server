package responses

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NotFoundError 404ステータスエラー
func NotFoundError(message string) error {
	return status.Errorf(
		codes.NotFound,
		fmt.Sprintf(message),
	)
}

// BadRequestError 400エラー
func BadRequestError(message string) error {
	return status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf(message),
	)
}

// InternalServerError 500エラー
func InternalServerError(message string) error {
	return status.Errorf(
		codes.Internal,
		fmt.Sprintf(message),
	)
}
