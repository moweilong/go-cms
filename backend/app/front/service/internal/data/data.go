package data

import (
	"context"

	"github.com/redis/go-redis/v9"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"
	bRegistry "github.com/tx7do/kratos-bootstrap/registry"
	"github.com/tx7do/kratos-bootstrap/rpc"

	commentV1 "go-cms/api/gen/go/comment/service/v1"
	contentV1 "go-cms/api/gen/go/content/service/v1"
	fileV1 "go-cms/api/gen/go/file/service/v1"
	userV1 "go-cms/api/gen/go/user/service/v1"

	"go-cms/pkg/cache"
	"go-cms/pkg/service"
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine
}

// NewData .
func NewData(redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	logger log.Logger,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/front-service"))

	d := &Data{
		rdb:           redisClient,
		log:           l,
		authenticator: authenticator,
		authorizer:    authorizer,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	// l := log.NewHelper(log.With(logger, "module", "redis/data/front-service"))
	return redisClient.NewClient(cfg.Data)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bRegistry.NewDiscovery(cfg.Registry)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Rest.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Rest.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer 创建权鉴器
func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}

func NewUserTokenRepo(data *Data, authenticator authnEngine.Authenticator, logger log.Logger) *cache.UserToken {
	const (
		userAccessTokenKeyPrefix  = "uat_"
		userRefreshTokenKeyPrefix = "urt_"
	)
	return cache.NewUserToken(data.rdb, authenticator, logger, userAccessTokenKeyPrefix, userRefreshTokenKeyPrefix)
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewAttachmentServiceClient(r registry.Discovery, c *conf.Bootstrap) fileV1.AttachmentServiceClient {
	return fileV1.NewAttachmentServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewCommentServiceClient(r registry.Discovery, c *conf.Bootstrap) commentV1.CommentServiceClient {
	return commentV1.NewCommentServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewCategoryServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.CategoryServiceClient {
	return contentV1.NewCategoryServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewLinkServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.LinkServiceClient {
	return contentV1.NewLinkServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewPostServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.PostServiceClient {
	return contentV1.NewPostServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewTagServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.TagServiceClient {
	return contentV1.NewTagServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewFileServiceClient(r registry.Discovery, cfg *conf.Bootstrap) fileV1.FileServiceClient {
	return fileV1.NewFileServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, cfg))
}
