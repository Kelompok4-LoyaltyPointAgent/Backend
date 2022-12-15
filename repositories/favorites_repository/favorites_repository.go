package favorites_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type FavoritesRepository interface {
	FindAll(args ...any) ([]models.Favorites, error)
	Create(favorites models.Favorites) (models.Favorites, error)
	Delete(userID string, productID string) error
}

type favoritesRepository struct {
	db *gorm.DB
}

func NewFavoritesRepository(db *gorm.DB) FavoritesRepository {
	return &favoritesRepository{db}
}

func (r *favoritesRepository) FindAll(args ...any) ([]models.Favorites, error) {
	var favorites []models.Favorites
	if len(args) == 0 {
		err := r.db.Preload("Product").Preload("Product.Icon").Preload("Product.ProductPicture").Find(&favorites).Error
		return favorites, err
	} else {
		err := r.db.Where("user_id = ?", args[0]).Preload("Product").Preload("Product.Icon").Preload("Product.ProductPicture").Find(&favorites).Error
		return favorites, err
	}

}

func (r *favoritesRepository) Create(favorites models.Favorites) (models.Favorites, error) {
	err := r.db.Create(&favorites).Error
	return favorites, err
}

func (r *favoritesRepository) Delete(userID string, productID string) error {
	var favorites models.Favorites
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&favorites).Error
	return err
}
