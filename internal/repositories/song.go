package repositories

// import (
// 	models "blog/internal/models"

// 	gorm "gorm.io/gorm"
// )

// type SongRepository struct {
// 	db *gorm.DB
// }

// func NewSongRepository(db *gorm.DB) *SongRepository {
// 	return &SongRepository{
// 		db: db,
// 	}
// }

// func (r *SongRepository) Create(song *models.Song) error {
// 	return r.db.Create(song).Error
// }

// func (r *SongRepository) FindSongById(id int) (*models.Song, error) {
// 	var user models.User
// 	if err := r.db.Where("id = ?", userName).First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
