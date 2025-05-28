package services

import (
	blog "campaign-services/gen/go/blog"
	"campaign-services/models"
	"campaign-services/repository"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type blogService struct {
	blog.UnimplementedBlogServiceServer
	repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) *blogService {
	return &blogService{repo: repo}
}

func (s *blogService) toProto(b models.BlogDB) *blog.Blog {
	return &blog.Blog{
		Id:         b.ID,
		UserId:     b.UserID,
		CampaignId: b.CampaignID,
		Content:    b.Content,
		CreatedAt:  timestamppb.New(b.CreatedAt),
		UpdatedAt:  timestamppb.New(b.UpdatedAt),
	}
}

func (s *blogService) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.CreateBlogResponse, error) {
	newBlog := models.BlogDB{
		ID:         uuid.NewString(),
		UserID:     req.UserId,
		CampaignID: req.CampaignId,
		Content:    req.Content,
	}

	b, err := s.repo.Create(newBlog)
	if err != nil {
		return nil, err
	}
	return &blog.CreateBlogResponse{Blog: s.toProto(b)}, nil
}

func (s *blogService) GetBlogsByCampaignID(ctx context.Context, req *blog.GetBlogsByCampaignIDRequest) (*blog.GetBlogsByCampaignIDResponse, error) {
	blogs, err := s.repo.GetByCampaignID(req.CampaignId)
	if err != nil {
		return nil, err
	}
	var list []*blog.Blog
	for _, b := range blogs {
		list = append(list, s.toProto(b))
	}
	return &blog.GetBlogsByCampaignIDResponse{Blogs: list}, nil
}

func (s *blogService) GetBlogsByUserID(ctx context.Context, req *blog.GetBlogsByUserIDRequest) (*blog.GetBlogsByUserIDResponse, error) {
	blogs, err := s.repo.GetByUserID(req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*blog.Blog
	for _, b := range blogs {
		list = append(list, s.toProto(b))
	}
	return &blog.GetBlogsByUserIDResponse{Blogs: list}, nil
}

func (s *blogService) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(req.Id, req.UserId); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
