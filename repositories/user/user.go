package repositories

import (
	"context"
	errCommon "user-service/common/error" //HARUS MASUK KEDALAM FOLDER LAGI JIKA FUNC ATAU CONST BERADA DI DALAM
	errConstant "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.User, error)
	Update(context.Context, *dto.UpdateRequest, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUUID(context.Context, string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error) {
	user := models.User{
		UUID:        uuid.New(),
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		RoleID:      req.RoleID,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, errCommon.WrapError(errConstant.ErrSQLError)
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, req *dto.UpdateRequest, uuid string) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	err := r.db.WithContext(ctx).
		Where("uuid = ?", uuid).
		Updates(&user).Error
	if err != nil {
		return nil, errCommon.WrapError(errConstant.ErrSQLError)
	}

	return &user, nil

}
