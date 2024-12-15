package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	v1 "go-cms/api/gen/go/comment/service/v1"
	"go-cms/app/core/service/internal/data/ent"
	"go-cms/app/core/service/internal/data/ent/comment"
)

type CommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) *CommentRepo {
	l := log.NewHelper(log.With(logger, "module", "comment/repo/core-service"))
	return &CommentRepo{
		data: data,
		log:  l,
	}
}

func (r *CommentRepo) convertEntToProto(in *ent.Comment) *v1.Comment {
	if in == nil {
		return nil
	}
	return &v1.Comment{
		Id:                in.ID,
		Author:            in.Author,
		Email:             in.Email,
		IpAddress:         in.IPAddress,
		AuthorUrl:         in.AuthorURL,
		GravatarMd5:       in.GravatarMd5,
		Content:           in.Content,
		Status:            in.Status,
		UserAgent:         in.UserAgent,
		ParentId:          in.ParentID,
		IsAdmin:           in.IsAdmin,
		AllowNotification: in.AllowNotification,
		Avatar:            in.Avatar,
		CreateTime:        util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:        util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *CommentRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Comment.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *CommentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCommentResponse, error) {
	builder := r.data.db.Client().Comment.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), comment.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.Comment, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListCommentResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *CommentRepo) Get(ctx context.Context, req *v1.GetCommentRequest) (*v1.Comment, error) {
	res, err := r.data.db.Client().Comment.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CommentRepo) Create(ctx context.Context, req *v1.CreateCommentRequest) (*v1.Comment, error) {
	res, err := r.data.db.Client().Comment.Create().
		SetNillableAuthor(req.Comment.Author).
		SetNillableEmail(req.Comment.Email).
		SetNillableIPAddress(req.Comment.IpAddress).
		SetNillableAuthorURL(req.Comment.AuthorUrl).
		SetNillableGravatarMd5(req.Comment.GravatarMd5).
		SetNillableContent(req.Comment.Content).
		SetNillableStatus(req.Comment.Status).
		SetNillableUserAgent(req.Comment.UserAgent).
		SetNillableParentID(req.Comment.ParentId).
		SetNillableIsAdmin(req.Comment.IsAdmin).
		SetNillableAllowNotification(req.Comment.AllowNotification).
		SetNillableAvatar(req.Comment.Avatar).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CommentRepo) Update(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.Comment, error) {
	builder := r.data.db.Client().Comment.UpdateOneID(req.Id).
		SetNillableAuthor(req.Comment.Author).
		SetNillableEmail(req.Comment.Email).
		SetNillableIPAddress(req.Comment.IpAddress).
		SetNillableAuthorURL(req.Comment.AuthorUrl).
		SetNillableGravatarMd5(req.Comment.GravatarMd5).
		SetNillableContent(req.Comment.Content).
		SetNillableStatus(req.Comment.Status).
		SetNillableUserAgent(req.Comment.UserAgent).
		SetNillableParentID(req.Comment.ParentId).
		SetNillableIsAdmin(req.Comment.IsAdmin).
		SetNillableAllowNotification(req.Comment.AllowNotification).
		SetNillableAvatar(req.Comment.Avatar).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *CommentRepo) Delete(ctx context.Context, req *v1.DeleteCommentRequest) (bool, error) {
	err := r.data.db.Client().Comment.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
