package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	contentV1 "go-cms/api/gen/go/content/service/v1"
	v1 "go-cms/api/gen/go/front/service/v1"
)

type CategoryService struct {
	v1.CategoryServiceHTTPServer

	categoryClient contentV1.CategoryServiceClient
	log            *log.Helper
}

func NewCategoryService(logger log.Logger, categoryClient contentV1.CategoryServiceClient) *CategoryService {
	l := log.NewHelper(log.With(logger, "module", "category/service/front-service"))
	return &CategoryService{
		log:            l,
		categoryClient: categoryClient,
	}
}

// ListCategory 获取类别列表
func (s *CategoryService) ListCategory(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListCategoryResponse, error) {
	return s.categoryClient.ListCategory(ctx, req)
}

// GetCategory 获取类别数据
func (s *CategoryService) GetCategory(ctx context.Context, req *contentV1.GetCategoryRequest) (*contentV1.Category, error) {
	return s.categoryClient.GetCategory(ctx, req)
}

// CreateCategory 创建类别
func (s *CategoryService) CreateCategory(ctx context.Context, req *contentV1.CreateCategoryRequest) (*contentV1.Category, error) {
	return s.categoryClient.CreateCategory(ctx, req)
}

// UpdateCategory 更新类别
func (s *CategoryService) UpdateCategory(ctx context.Context, req *contentV1.UpdateCategoryRequest) (*contentV1.Category, error) {
	return s.categoryClient.UpdateCategory(ctx, req)
}

// DeleteCategory 删除类别
func (s *CategoryService) DeleteCategory(ctx context.Context, req *contentV1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	return s.categoryClient.DeleteCategory(ctx, req)
}
