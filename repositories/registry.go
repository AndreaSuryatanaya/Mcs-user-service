package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegisty interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegisty {
	return &Registry{db: db}
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
