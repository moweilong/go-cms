package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	commentV1 "go-cms/api/gen/go/comment/service/v1"
	"go-cms/app/core/service/internal/data"
)

type CommentService struct {
	commentV1.UnimplementedCommentServiceServer

	uc  *data.CommentRepo
	log *log.Helper
}

func NewCommentService(logger log.Logger, uc *data.CommentRepo) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "service/comment/core-service"))
	return &CommentService{
		log: l,
		uc:  uc,
	}
}

// ListComment 获取留言列表
func (s *CommentService) ListComment(ctx context.Context, req *pagination.PagingRequest) (*commentV1.ListCommentResponse, error) {
	return s.uc.List(ctx, req)
}

// GetComment 获取留言数据
func (s *CommentService) GetComment(ctx context.Context, req *commentV1.GetCommentRequest) (*commentV1.Comment, error) {
	return s.uc.Get(ctx, req)
}

// CreateComment 创建留言
func (s *CommentService) CreateComment(ctx context.Context, req *commentV1.CreateCommentRequest) (*commentV1.Comment, error) {
	return s.uc.Create(ctx, req)
}

// UpdateComment 更新留言
func (s *CommentService) UpdateComment(ctx context.Context, req *commentV1.UpdateCommentRequest) (*commentV1.Comment, error) {
	return s.uc.Update(ctx, req)
}

// DeleteComment 删除留言
func (s *CommentService) DeleteComment(ctx context.Context, req *commentV1.DeleteCommentRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
