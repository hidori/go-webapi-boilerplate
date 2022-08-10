package presenter

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
)

// FromContactGetByIDOutputPort は、ユースケースの出力を HTTP 応答に変換します。
func FromContactGetByIDOutputPort(output *usecase.ContactGetByIDOutputPort) (interface{}, error) {
	return output.Model, nil
}
