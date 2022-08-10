package registry

import "github.com/hidori/go-webapi-boilerplate/go/internal/usecase"

// Usecase は、ユースケースのレジストリです。
type Usecase struct {
	ContactAddOrUpdate usecase.ContactAddOrUpdateUsecase
	ContactDeleteByID  usecase.ContactDeleteByIDUsecase
	ContactGetByID     usecase.ContactGetByIDUsecase
	ContactGetList     usecase.ContactGetListUsecase
}
