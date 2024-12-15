// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"go-cms/app/admin/service/internal/data"
	"go-cms/app/admin/service/internal/server"
	"go-cms/app/admin/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *v1.Bootstrap) (*kratos.App, func(), error) {
	authenticator := data.NewAuthenticator(bootstrap)
	engine := data.NewAuthorizer()
	client := data.NewRedisClient(bootstrap, logger)
	dataData, cleanup, err := data.NewData(client, authenticator, engine, logger)
	if err != nil {
		return nil, nil, err
	}
	userToken := data.NewUserTokenRepo(dataData, authenticator, logger)
	discovery := data.NewDiscovery(bootstrap)
	userServiceClient := data.NewUserServiceClient(discovery, bootstrap)
	authenticationService := service.NewAuthenticationService(logger, userServiceClient, userToken)
	userService := service.NewUserService(logger, userServiceClient)
	postServiceClient := data.NewPostServiceClient(discovery, bootstrap)
	postService := service.NewPostService(logger, postServiceClient)
	linkServiceClient := data.NewLinkServiceClient(discovery, bootstrap)
	linkService := service.NewLinkService(logger, linkServiceClient)
	categoryServiceClient := data.NewCategoryServiceClient(discovery, bootstrap)
	categoryService := service.NewCategoryService(logger, categoryServiceClient)
	commentServiceClient := data.NewCommentServiceClient(discovery, bootstrap)
	commentService := service.NewCommentService(logger, commentServiceClient)
	tagServiceClient := data.NewTagServiceClient(discovery, bootstrap)
	tagService := service.NewTagService(logger, tagServiceClient)
	attachmentServiceClient := data.NewAttachmentServiceClient(discovery, bootstrap)
	attachmentService := service.NewAttachmentService(logger, attachmentServiceClient)
	httpServer := server.NewHTTPServer(bootstrap, logger, authenticator, engine, userToken, authenticationService, userService, postService, linkService, categoryService, commentService, tagService, attachmentService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
