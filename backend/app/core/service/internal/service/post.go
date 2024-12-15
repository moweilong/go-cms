package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	contentV1 "go-cms/api/gen/go/content/service/v1"
	"go-cms/app/core/service/internal/data"
)

type PostService struct {
	contentV1.UnimplementedPostServiceServer

	uc  *data.PostRepo
	log *log.Helper
}

func NewPostService(logger log.Logger, uc *data.PostRepo) *PostService {
	l := log.NewHelper(log.With(logger, "module", "post/service/core-service"))
	return &PostService{
		log: l,
		uc:  uc,
	}
}

// ListPost 获取帖子列表
func (s *PostService) ListPost(ctx context.Context, req *pagination.PagingRequest) (*contentV1.ListPostResponse, error) {
	return s.uc.List(ctx, req)
}

// GetPost 获取帖子数据
func (s *PostService) GetPost(ctx context.Context, req *contentV1.GetPostRequest) (*contentV1.Post, error) {
	return s.uc.Get(ctx, req)
}

// CreatePost 创建帖子
func (s *PostService) CreatePost(ctx context.Context, req *contentV1.CreatePostRequest) (*contentV1.Post, error) {
	return s.uc.Create(ctx, req)
}

// UpdatePost 更新帖子
func (s *PostService) UpdatePost(ctx context.Context, req *contentV1.UpdatePostRequest) (*contentV1.Post, error) {
	return s.uc.Update(ctx, req)
}

// DeletePost 删除帖子
func (s *PostService) DeletePost(ctx context.Context, req *contentV1.DeletePostRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
