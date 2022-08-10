package presenter

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
)

// FromContactAddOrUpdateOutputPort は、ユースケースの出力を HTTP 応答に変換します。
func FromContactAddOrUpdateOutputPort(output *usecase.ContactAddOrUpdateOutputPort) (interface{}, error) {
	return output.Model, nil
}
