package faq_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type FAQRepository interface {
	FindAll(query any, args ...any) ([]models.FAQ, error)
	FindByID(id any) (models.FAQ, error)
	Create(faq models.FAQ) (models.FAQ, error)
	Update(updates models.FAQ, id any) (models.FAQ, error)
	Delete(id any) error
}

type faqRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) FAQRepository {
	return &faqRepository{db}
}

func (r *faqRepository) FindAll(query any, args ...any) ([]models.FAQ, error) {
	var faqs []models.FAQ

	var err error
	if query != nil {
		err = r.db.Where(query, args...).Find(&faqs).Error
	} else {
		err = r.db.Find(&faqs).Error
	}

	return faqs, err
}

func (r *faqRepository) FindByID(id any) (models.FAQ, error) {
	var faq models.FAQ
	err := r.db.Where("id = ?", id).First(&faq).Error
	return faq, err
}

func (r *faqRepository) Create(faq models.FAQ) (models.FAQ, error) {
	err := r.db.Create(&faq).Error
	return faq, err
}

func (r *faqRepository) Update(updates models.FAQ, id any) (models.FAQ, error) {
	var faq models.FAQ
	err := r.db.Model(&faq).Where("id = ?", id).Updates(&updates).Error
	if err != nil {
		return faq, err
	}
	return r.FindByID(id)
}

func (r *faqRepository) Delete(id any) error {
	var faq models.FAQ
	err := r.db.Where("id = ?", id).Delete(&faq).Error
	return err
}
