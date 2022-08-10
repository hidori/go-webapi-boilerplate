package registry

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/repository"
)

// Repository は、レポジトリのレジストリです。
type Repository struct {
	// Contact レポジトリ
	Contacts repository.ContactRepository
}
