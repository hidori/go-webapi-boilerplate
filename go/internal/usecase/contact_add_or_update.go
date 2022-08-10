package usecase

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
)

// ContactAddOrUpdateInputPort は、ContactAddOrUpdateUsecase の入力です。
type ContactAddOrUpdateInputPort struct {
	// 連絡先
	Model *model.Contact
}

// ContactAddOrUpdateOutputPort は、ContactAddOrUpdateUsecase の出力です。
type ContactAddOrUpdateOutputPort struct {
	// 連絡先
	Model *model.Contact
}

// ContactAddOrUpdateUsecase は、「連絡先を追加または更新する」ユースケースのインターフェースです。
type ContactAddOrUpdateUsecase Usecase[*ContactAddOrUpdateInputPort, *ContactAddOrUpdateOutputPort]
