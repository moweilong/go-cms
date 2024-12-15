package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	fileV1 "go-cms/api/gen/go/file/service/v1"
	"go-cms/app/core/service/internal/data"
)

type AttachmentService struct {
	fileV1.UnimplementedAttachmentServiceServer

	uc  *data.AttachmentRepo
	log *log.Helper
}

func NewAttachmentService(logger log.Logger, uc *data.AttachmentRepo) *AttachmentService {
	l := log.NewHelper(log.With(logger, "module", "attachment/service/core-service"))
	return &AttachmentService{
		log: l,
		uc:  uc,
	}
}

// ListAttachment 获取附件列表
func (s *AttachmentService) ListAttachment(ctx context.Context, req *pagination.PagingRequest) (*fileV1.ListAttachmentResponse, error) {
	return s.uc.List(ctx, req)
}

// GetAttachment 获取附件数据
func (s *AttachmentService) GetAttachment(ctx context.Context, req *fileV1.GetAttachmentRequest) (*fileV1.Attachment, error) {
	return s.uc.Get(ctx, req)
}

// CreateAttachment 创建附件
func (s *AttachmentService) CreateAttachment(ctx context.Context, req *fileV1.CreateAttachmentRequest) (*fileV1.Attachment, error) {
	return s.uc.Create(ctx, req)
}

// UpdateAttachment 更新附件
func (s *AttachmentService) UpdateAttachment(ctx context.Context, req *fileV1.UpdateAttachmentRequest) (*fileV1.Attachment, error) {
	return s.uc.Update(ctx, req)
}

// DeleteAttachment 删除附件
func (s *AttachmentService) DeleteAttachment(ctx context.Context, req *fileV1.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
