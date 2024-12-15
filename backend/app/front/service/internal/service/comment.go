package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	commentV1 "go-cms/api/gen/go/comment/service/v1"
	v1 "go-cms/api/gen/go/front/service/v1"
)

type CommentService struct {
	v1.CommentServiceHTTPServer

	commentClient commentV1.CommentServiceClient
	log           *log.Helper
}

func NewCommentService(logger log.Logger, commentClient commentV1.CommentServiceClient) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "comment/service/front-service"))
	return &CommentService{
		log:           l,
		commentClient: commentClient,
	}
}

// ListComment 获取留言列表
func (s *CommentService) ListComment(ctx context.Context, req *pagination.PagingRequest) (*commentV1.ListCommentResponse, error) {
	return s.commentClient.ListComment(ctx, req)
}

// GetComment 获取留言数据
func (s *CommentService) GetComment(ctx context.Context, req *commentV1.GetCommentRequest) (*commentV1.Comment, error) {
	return s.commentClient.GetComment(ctx, req)
}

// CreateComment 创建留言
func (s *CommentService) CreateComment(ctx context.Context, req *commentV1.CreateCommentRequest) (*commentV1.Comment, error) {
	return s.commentClient.CreateComment(ctx, req)
}

// UpdateComment 更新留言
func (s *CommentService) UpdateComment(ctx context.Context, req *commentV1.UpdateCommentRequest) (*commentV1.Comment, error) {
	return s.commentClient.UpdateComment(ctx, req)
}

// DeleteComment 删除留言
func (s *CommentService) DeleteComment(ctx context.Context, req *commentV1.DeleteCommentRequest) (*emptypb.Empty, error) {
	return s.commentClient.DeleteComment(ctx, req)
}
