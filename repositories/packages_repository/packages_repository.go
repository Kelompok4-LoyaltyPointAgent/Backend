package packages_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type PackagesRepository interface {
	FindAll() ([]models.Packages, error)
	Create(packages models.Packages) (models.Packages, error)
	FindByProductID(id any) (models.Packages, error)
	FindByProvider(provider string) ([]models.Packages, error)
	UpdateByProductID(packagesUpdate models.Packages, productID any) (models.Packages, error)
	DeleteByProductID(productID any) error
}

type packagesRepository struct {
	db *gorm.DB
}

func NewPackagesRepository(db *gorm.DB) *packagesRepository {
	return &packagesRepository{db}
}

func (r *packagesRepository) FindAll() ([]models.Packages, error) {
	var packages []models.Packages
	err := r.db.Preload("Product").Preload("Product.ProductPicture").Find(&packages).Error
	return packages, err
}

func (r *packagesRepository) Create(packages models.Packages) (models.Packages, error) {
	err := r.db.Create(&packages).Error
	return packages, err
}

func (r *packagesRepository) FindByProductID(id any) (models.Packages, error) {
	var packages models.Packages
	err := r.db.Where("product_id = ?", id).Preload("Product").Preload("Product.ProductPicture").Find(&packages).Error
	return packages, err
}

func (r *packagesRepository) FindByProvider(provider string) ([]models.Packages, error) {
	var packages []models.Packages
	err := r.db.Preload("Product", "provider = ?", provider).Preload("Product.ProductPicture").Find(&packages).Error

	return packages, err
}

func (r *packagesRepository) UpdateByProductID(packagesUpdate models.Packages, productID any) (models.Packages, error) {
	var packages models.Packages
	err := r.db.Model(&packages).Where("product_id = ?", productID).Updates(&packagesUpdate).Error
	if err != nil {
		return packages, err
	}
	return r.FindByProductID(productID)
}

func (r *packagesRepository) DeleteByProductID(productID any) error {
	var packages models.Packages
	err := r.db.Where("product_id = ?", productID).Delete(&packages).Error
	return err
}
