package usecase

// ContactDeleteByIDInputPort は、ContactDeleteByIDUsecase の入力です。
type ContactDeleteByIDInputPort struct {
	// 連絡先 ID
	ContactID int
}

// ContactDeleteByIDOutputPort は、ContactDeleteByIDUsecase の出力です。
type ContactDeleteByIDOutputPort struct{}

// ContactDeleteByIDUsecase は、「連絡先を削除する」ユースケースのインターフェースです。
type ContactDeleteByIDUsecase Usecase[*ContactDeleteByIDInputPort, *ContactDeleteByIDOutputPort]
