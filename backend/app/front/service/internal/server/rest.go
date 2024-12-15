package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"go-cms/app/front/service/cmd/server/assets"
	"go-cms/app/front/service/internal/service"
	"go-cms/pkg/cache"
	"go-cms/pkg/middleware/auth"

	frontV1 "go-cms/api/gen/go/front/service/v1"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(
	logger log.Logger,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	userToken *cache.UserToken,
) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(userToken),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	userToken *cache.UserToken,
	authnSvc *service.AuthenticationService,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	commentSvc *service.CommentService,
	tagSvc *service.TagService,
	attachmentSvc *service.AttachmentService,
	fileSvc *service.FileService,
) *http.Server {
	srv := rpc.CreateRestServer(cfg, newRestMiddleware(logger, authenticator, authorizer, userToken)...)

	frontV1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	frontV1.RegisterPostServiceHTTPServer(srv, postSvc)
	frontV1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	frontV1.RegisterTagServiceHTTPServer(srv, tagSvc)
	frontV1.RegisterLinkServiceHTTPServer(srv, linkSvc)
	frontV1.RegisterAttachmentServiceHTTPServer(srv, attachmentSvc)
	frontV1.RegisterCommentServiceHTTPServer(srv, commentSvc)
	frontV1.RegisterFileServiceHTTPServer(srv, fileSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Front Service"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
