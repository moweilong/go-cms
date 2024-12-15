package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	contentV1 "go-cms/api/gen/go/content/service/v1"
	"go-cms/app/core/service/internal/data"
)

type TagService struct {
	contentV1.UnimplementedTagServiceServer

	uc  *data.TagRepo
	log *log.Helper
}

func NewTagService(logger log.Logger, uc *data.TagRepo) *TagService {
	l := log.NewHelper(log.With(logger, "module", "tag/service/core-service"))
	return &TagService{
		log: l,
		uc:  uc,
	}
}

// ListTag 获取标签列表
func (s *TagService) ListTag(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListTagResponse, error) {
	return s.uc.List(ctx, req)
}

// GetTag 获取标签数据
func (s *TagService) GetTag(ctx context.Context, req *contentV1.GetTagRequest) (*contentV1.Tag, error) {
	return s.uc.Get(ctx, req)
}

// CreateTag 创建标签
func (s *TagService) CreateTag(ctx context.Context, req *contentV1.CreateTagRequest) (*contentV1.Tag, error) {
	return s.uc.Create(ctx, req)
}

// UpdateTag 更新标签
func (s *TagService) UpdateTag(ctx context.Context, req *contentV1.UpdateTagRequest) (*contentV1.Tag, error) {
	return s.uc.Update(ctx, req)
}

// DeleteTag 删除标签
func (s *TagService) DeleteTag(ctx context.Context, req *contentV1.DeleteTagRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
