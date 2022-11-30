package product_picture_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type ProductPictureRepository interface {
	Create(productPicture models.ProductPicture) (models.ProductPicture, error)
	Update(productPicture models.ProductPicture, id string) (models.ProductPicture, error)
	Delete(id string) error
	FindByID(id string) (models.ProductPicture, error)
	FindByName(name string) (models.ProductPicture, error)
}

type productPictureRepository struct {
	db *gorm.DB
}

func NewProductPictureRepository(db *gorm.DB) *productPictureRepository {
	return &productPictureRepository{db}
}

func (r *productPictureRepository) Create(productPicture models.ProductPicture) (models.ProductPicture, error) {
	err := r.db.Create(&productPicture).Error
	return productPicture, err
}

func (r *productPictureRepository) Update(productPicture models.ProductPicture, id string) (models.ProductPicture, error) {
	err := r.db.Model(&productPicture).Where("id = ?", id).Updates(&productPicture).Error
	if err != nil {
		return productPicture, err
	}
	return r.FindByID(id)
}

func (r *productPictureRepository) Delete(id string) error {
	var productPicture models.ProductPicture
	err := r.db.Where("id = ?", id).Delete(&productPicture).Error
	return err
}

func (r *productPictureRepository) FindByID(id string) (models.ProductPicture, error) {
	var productPicture models.ProductPicture
	err := r.db.Where("id = ?", id).Find(&productPicture).Error
	return productPicture, err
}

func (r *productPictureRepository) FindByName(name string) (models.ProductPicture, error) {
	var productPicture models.ProductPicture
	err := r.db.Where("name = ?", name).First(&productPicture).Error
	return productPicture, err
}
