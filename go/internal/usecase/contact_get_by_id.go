package usecase

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
)

// ContactGetByIDInputPort は、ContactGetByIDUsecase の入力です。
type ContactGetByIDInputPort struct {
	// 連絡先 ID
	ContactID int
}

// ContactGetByIDOutputPort は、ContactGetByIDUsecase の出力です。
type ContactGetByIDOutputPort struct {
	// 連絡先
	Model *model.Contact
}

// ContactGetByIDUsecase は、「連絡先を取得する」ユースケースのインターフェースです。
type ContactGetByIDUsecase Usecase[*ContactGetByIDInputPort, *ContactGetByIDOutputPort]
