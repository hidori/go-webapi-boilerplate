package api

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/api/middleware"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/api/router"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/database"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/controller"
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/gateway/mysql"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase/interactor"
	"github.com/labstack/echo/v4"
)

// NewServer は、Server の新規インスタンスを返します。
func NewServer(config *config.Config) (*echo.Echo, error) {
	db, err := NewDatabase(config.DB)
	if err != nil {
		logger.Errorf("fail to NewDatabase(): err=%v", err)
		return nil, err
	}

	rp := NewRepository()
	sv := NewService()
	uc := NewUsecase(config, db, rp, sv)
	ct := NewController(uc)

	server := echo.New()
	middleware.Attach(server, config)
	router.Attach(server.Group("api/v1"), ct)

	return server, nil
}

// NewDatabase は、registry.Database の新規インスタンスを返します。
func NewDatabase(config *config.DBConfig) (*registry.Database, error) {
	reader, err := database.Open(config.ReaderDSN)
	if err != nil {
		logger.Errorf("fail to database.Open(config.DB.ReaderDSN): err=%v", err)
		return nil, err
	}

	writer, err := database.Open(config.WriterDSN)
	if err != nil {
		logger.Errorf("fail to database.Open(config.DB.WriterDSN): err=%v", err)
		return nil, err
	}

	return &registry.Database{
		Reader: reader,
		Writer: writer,
	}, nil
}

// NewRepository は、registry.Repository の新規インスタンスを返します。
func NewRepository() *registry.Repository {
	return &registry.Repository{
		Contacts: mysql.NewContactRepository(),
	}
}

// NewService は、registry.Service の新規インスタンスを返します。
func NewService() *registry.Service {
	return &registry.Service{}
}

// NewUsecase は、registry.Usecase の新規インスタンスを返します。
func NewUsecase(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) *registry.Usecase {
	return &registry.Usecase{
		ContactAddOrUpdate: interactor.NewContactAddOrUpdateInteractor(config, db, rp, sv),
		ContactDeleteByID:  interactor.NewContactDeleteByIDInteractor(config, db, rp, sv),
		ContactGetByID:     interactor.NewContactGetByIDInteractor(config, db, rp, sv),
		ContactGetList:     interactor.NewContactGetListInteractor(config, db, rp, sv),
	}
}

// NewController は、registry.Controller の新規インスタンスを返します。
func NewController(uc *registry.Usecase) *registry.Controller {
	return &registry.Controller{
		ContactAddOrUpdate: controller.NewContactAddOrUpdateController(uc.ContactAddOrUpdate),
		ContactDeleteByID:  controller.NewContactDeleteByIDController(uc.ContactDeleteByID),
		ContactGetByID:     controller.NewContactGetByIDController(uc.ContactGetByID),
		ContactGetList:     controller.NewContactGetListController(uc.ContactGetList),
	}
}
