package presenter

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
)

// FromContactDeleteByIDOutputPort は、ユースケースの出力を HTTP 応答に変換します。
func FromContactDeleteByIDOutputPort(output *usecase.ContactDeleteByIDOutputPort) (interface{}, error) {
	return &empty{}, nil
}
