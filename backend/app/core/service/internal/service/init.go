//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewUserService,
	NewAttachmentService,
	NewCategoryService,
	NewCommentService,
	NewLinkService,
	NewPostService,
	NewTagService,
	NewFileService,
)
