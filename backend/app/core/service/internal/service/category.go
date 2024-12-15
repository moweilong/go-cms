package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	contentV1 "go-cms/api/gen/go/content/service/v1"
	"go-cms/app/core/service/internal/data"
)

type CategoryService struct {
	contentV1.UnimplementedCategoryServiceServer

	uc  *data.CategoryRepo
	log *log.Helper
}

func NewCategoryService(logger log.Logger, uc *data.CategoryRepo) *CategoryService {
	l := log.NewHelper(log.With(logger, "module", "category/service/core-service"))
	return &CategoryService{
		log: l,
		uc:  uc,
	}
}

// ListCategory 获取类别列表
func (s *CategoryService) ListCategory(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListCategoryResponse, error) {
	return s.uc.List(ctx, req)
}

// GetCategory 获取类别数据
func (s *CategoryService) GetCategory(ctx context.Context, req *contentV1.GetCategoryRequest) (*contentV1.Category, error) {
	return s.uc.Get(ctx, req)
}

// CreateCategory 创建类别
func (s *CategoryService) CreateCategory(ctx context.Context, req *contentV1.CreateCategoryRequest) (*contentV1.Category, error) {
	return s.uc.Create(ctx, req)
}

// UpdateCategory 更新类别
func (s *CategoryService) UpdateCategory(ctx context.Context, req *contentV1.UpdateCategoryRequest) (*contentV1.Category, error) {
	return s.uc.Update(ctx, req)
}

// DeleteCategory 删除类别
func (s *CategoryService) DeleteCategory(ctx context.Context, req *contentV1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
