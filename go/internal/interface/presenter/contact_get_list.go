package presenter

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
)

// FromContactGetListOutputPort は、ユースケースの出力を HTTP 応答に変換します。
func FromContactGetListOutputPort(output *usecase.ContactGetListOutputPort) (interface{}, error) {
	return output.List, nil
}
