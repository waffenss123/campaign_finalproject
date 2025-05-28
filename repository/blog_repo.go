package repository

import (
	"campaign-services/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(blog models.BlogDB) (models.BlogDB, error)
	GetByCampaignID(campaignID string) ([]models.BlogDB, error)
	GetByUserID(userID int32) ([]models.BlogDB, error)
	Delete(id string, userID int32) error
}

type blogRepository struct{ db *gorm.DB }

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db}
}

func (r *blogRepository) Create(blog models.BlogDB) (models.BlogDB, error) {
	if err := r.db.Create(&blog).Error; err != nil {
		return models.BlogDB{}, status.Error(codes.Internal, err.Error())
	}
	return blog, nil
}

func (r *blogRepository) GetByCampaignID(campaignID string) ([]models.BlogDB, error) {
	var blogs []models.BlogDB
	if err := r.db.Where("campaign_id = ?", campaignID).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return blogs, nil
}

func (r *blogRepository) GetByUserID(userID int32) ([]models.BlogDB, error) {
	var blogs []models.BlogDB
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return blogs, nil
}

func (r *blogRepository) Delete(id string, userID int32) error {
	res := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.BlogDB{})
	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "blog not found or not owned by user")
	}
	return res.Error
}
