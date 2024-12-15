package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	fileV1 "go-cms/api/gen/go/file/service/v1"
	"go-cms/app/core/service/internal/data"
)

type FileService struct {
	fileV1.UnimplementedFileServiceServer

	mc  *data.MinIOClient
	log *log.Helper
}

func NewFileService(logger log.Logger, mc *data.MinIOClient) *FileService {
	l := log.NewHelper(log.With(logger, "module", "file/service/core-service"))
	return &FileService{
		log: l,
		mc:  mc,
	}
}

func (s *FileService) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	return s.mc.OssUploadUrl(ctx, req)
}
