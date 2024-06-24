package repositories

import (
	models "blog/internal/models"

	"github.com/google/uuid"
	gorm "gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	user.ID = uuid.NewString()
	return r.db.Create(user).Error
}

func (r *UserRepository) FindUserByUserName(userName string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// func (r *UserRepository) FindUsers() ([]*models.User, error) {
// 	if err := r.db.F
// }
