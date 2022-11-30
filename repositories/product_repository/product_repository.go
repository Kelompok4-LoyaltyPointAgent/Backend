package product_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id any) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(productUpdate models.Product, id any) (models.Product, error)
	DeleteByID(id any) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("ProductPicture").Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id any) (models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).Preload("ProductPicture").First(&product).Error
	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) Update(productUpdate models.Product, id any) (models.Product, error) {
	var product models.Product
	err := r.db.Model(&product).Where("id = ?", id).Updates(&productUpdate).Error
	if err != nil {
		return product, err
	}
	return r.FindByID(id)
}

func (r *productRepository) DeleteByID(id any) error {
	var product models.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	return err
}
