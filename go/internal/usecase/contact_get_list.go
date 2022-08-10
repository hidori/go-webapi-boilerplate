package usecase

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
)

// ContactGetListInputPort は、ContactGetListUsecase の入力です。
type ContactGetListInputPort struct{}

// ContactGetListOutputPort は、ContactGetListUsecase の出力です。
type ContactGetListOutputPort struct {
	// 連絡先の一覧
	List []model.Contact
}

// ContactGetListUsecase は、「連絡先の一覧を取得する」ユースケースのインターフェースです。
type ContactGetListUsecase Usecase[*ContactGetListInputPort, *ContactGetListOutputPort]
