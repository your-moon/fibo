package impl

import (
	"context"
	sqlS "database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	databaseImpl "fibo/internal/base/database/impl"
	"fibo/internal/base/errors"
	"fibo/internal/post"
)

type PostRepositoryOpts struct {
	ConnManager databaseImpl.ConnManager
}

func NewPostRepository(opts PostRepositoryOpts) post.PostRepository {
	return &postRepository{
		ConnManager: opts.ConnManager,
	}
}

type postRepository struct {
	databaseImpl.ConnManager
}

func (p *postRepository) Update(
	ctx context.Context,
	post post.PostModel,
) (int64, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Update("posts").
		Set(goqu.Record{"title": post.Title, "content": post.Content, "is_published": post.IsPublished, "likes": post.Likes, "category_id": post.CategoryId}).
		Where(goqu.Ex{"id": post.Id}).
		ToSQL()
	fmt.Println(sql)
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	_, err = p.Conn(ctx).Exec(ctx, sql)
	if err != nil {
		return 0, parseUpdatePostError(&post, err)
	}

	return post.Id, nil
}

func (r *postRepository) Create(ctx context.Context, post post.PostModel) (int64, error) {
	sql, _, err := databaseImpl.QueryBuilder.Insert("posts").Rows(databaseImpl.Record{
		"user_id":      post.UserId,
		"title":        post.Title,
		"content":      post.Content,
		"is_published": post.IsPublished,
		"category_id":  post.CategoryId,
	}).Returning("id").ToSQL()
	fmt.Println(sql)
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error post create")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&post.Id); err != nil {
		return 0, parseAddPostError(&post, err)
	}

	return post.Id, nil
}

func (r *postRepository) GetById(ctx context.Context, postId int64) (post.PostModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		From("posts").
		Where(goqu.Ex{"id": postId}).
		ToSQL()
	if err != nil {
		return post.PostModel{}, errors.Wrap(
			err,
			errors.DatabaseError,
			"syntax error get post by id",
		)
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	var p post.PostModel
	var createdAt time.Time
	var updatedAt time.Time
	var deletedAt sqlS.NullTime
	var category sqlS.NullInt64
	if err := row.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.IsPublished, &p.Likes, &createdAt, &updatedAt, &deletedAt, &category); err != nil {
		return post.PostModel{}, errors.Wrap(err, errors.DatabaseError, "scan post failed")
	}
	p.CategoryId = category.Int64
	p.CreatedAt = createdAt.Format(time.RFC3339)
	p.UpdatedAt = updatedAt.Format(time.RFC3339)
	if deletedAt.Valid {
		p.DeletedAt = deletedAt.Time.Format(time.RFC3339)
	} else {
		p.DeletedAt = ""
	}

	return p, nil
}

func (r *postRepository) GetPublishedPosts(ctx context.Context) ([]post.PostModelWithUser, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		From("posts").
		Select("posts.id", "posts.user_id", "posts.title", "posts.content", "posts.is_published", "posts.likes", "posts.created_at", "posts.updated_at", "posts.deleted_at", "posts.category_id", "users.email", "users.firstname").
		InnerJoin(goqu.T("users"), goqu.On(goqu.Ex{"posts.user_id": goqu.I("users.user_id")})).
		Where(databaseImpl.Ex{"is_published": true}).
		ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error get posts")
	}

	rows, err := r.Conn(ctx).Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "get published posts failed")
	}
	defer rows.Close()

	var posts []post.PostModelWithUser
	for rows.Next() {
		var p post.PostModelWithUser
		var createdAt time.Time
		var updatedAt time.Time
		var deletedAt sqlS.NullTime
		var category sqlS.NullInt64
		var userEmail string
		var userName string
		if err := rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.IsPublished, &p.Likes, &createdAt, &updatedAt, &deletedAt, &category, &userEmail, &userName); err != nil {
			return nil, errors.Wrap(err, errors.DatabaseError, "scan post failed")
		}
		p.CategoryId = category.Int64
		p.CreatedAt = createdAt.Format(time.RFC3339)
		p.UpdatedAt = updatedAt.Format(time.RFC3339)
		p.UserEmail = userEmail
		p.UserName = userName
		if deletedAt.Valid {
			p.DeletedAt = deletedAt.Time.Format(time.RFC3339)
		} else {
			p.DeletedAt = ""
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *postRepository) GetTotalLikesCountByUser(
	ctx context.Context,
	userId int64,
) (int64, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		From("posts").
		Select(goqu.SUM("likes").As("total_likes_count")).
		Where(goqu.Ex{"user_id": userId}).
		ToSQL()
	fmt.Println(sql)
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error get total likes count")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)
	var totalLikesCount int64
	if err := row.Scan(&totalLikesCount); err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "scan total likes count failed")
	}

	return totalLikesCount, nil
}

func (r *postRepository) GetMyPosts(
	ctx context.Context,
	userId int64,
) ([]post.PostModelWithUser, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		From("posts").
		Select("posts.id", "posts.user_id", "posts.title", "posts.content", "posts.is_published", "posts.likes", "posts.created_at", "posts.updated_at", "posts.deleted_at", "posts.category_id", "users.email", "users.firstname").
		InnerJoin(goqu.T("users"), goqu.On(goqu.Ex{"posts.user_id": goqu.I("users.user_id")})).
		Where(databaseImpl.Ex{"posts.user_id": userId}).
		ToSQL()
	fmt.Println(sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error get posts")
	}

	rows, err := r.Conn(ctx).Query(ctx, sql)
	fmt.Println(err)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "get my posts failed")
	}
	defer rows.Close()

	var posts []post.PostModelWithUser
	for rows.Next() {
		var p post.PostModelWithUser
		var createdAt time.Time
		var updatedAt time.Time
		var deletedAt sqlS.NullTime
		var category sqlS.NullInt64
		var userEmail string
		var userName string
		if err := rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.IsPublished, &p.Likes, &createdAt, &updatedAt, &deletedAt, &category, &userEmail, &userName); err != nil {
			return nil, errors.Wrap(err, errors.DatabaseError, "scan post failed")
		}
		p.CategoryId = category.Int64
		p.CreatedAt = createdAt.Format(time.RFC3339)
		p.UpdatedAt = updatedAt.Format(time.RFC3339)
		p.UserEmail = userEmail
		p.UserName = userName
		if deletedAt.Valid {
			p.DeletedAt = deletedAt.Time.Format(time.RFC3339)
		} else {
			p.DeletedAt = ""
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *postRepository) GetPosts(
	ctx context.Context,
) ([]post.PostModelWithUser, error) {
	sql, _, err := databaseImpl.QueryBuilder.From("posts").
		Select("posts.id", "posts.user_id", "posts.title", "posts.content", "posts.is_published", "posts.likes", "posts.created_at", "posts.updated_at", "posts.deleted_at", "posts.category_id", "users.email", "users.firstname").
		InnerJoin(goqu.T("users"), goqu.On(goqu.Ex{"posts.user_id": goqu.I("users.user_id")})).
		ToSQL()
	fmt.Println(sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error get posts")
	}

	rows, err := r.Conn(ctx).Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "get posts failed")
	}

	defer rows.Close()

	var posts []post.PostModelWithUser
	for rows.Next() {
		var p post.PostModelWithUser
		var createdAt time.Time
		var updatedAt time.Time
		var deletedAt sqlS.NullTime
		var category sqlS.NullInt64
		var userEmail string
		var userName string
		if err := rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Content, &p.IsPublished, &p.Likes, &createdAt, &updatedAt, &deletedAt, &category, &userEmail, &userName); err != nil {
			fmt.Println(err)
			return nil, errors.Wrap(err, errors.DatabaseError, "scan post failed")
		}
		fmt.Println(err)
		p.CategoryId = category.Int64
		p.CreatedAt = createdAt.Format(time.RFC3339)
		p.UpdatedAt = updatedAt.Format(time.RFC3339)
		p.UserEmail = userEmail
		p.UserName = userName
		if deletedAt.Valid {
			p.DeletedAt = deletedAt.Time.Format(time.RFC3339)
		} else {
			p.DeletedAt = ""
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func parseUpdatePostError(post *post.PostModel, err error) error {
	pgErr, isPgErr := err.(*pgconn.PgError)

	if isPgErr && pgErr.Code == pgerrcode.UniqueViolation {
		return errors.Wrapf(err, errors.DatabaseError, "unique violation")
	}
	return errors.Wrapf(err, errors.DatabaseError, "update post failed")
}

func parseAddPostError(post *post.PostModel, err error) error {
	pgErr, isPgErr := err.(*pgconn.PgError)

	if isPgErr && pgErr.Code == pgerrcode.UniqueViolation {
		return errors.Wrapf(err, errors.DatabaseError, "unique violation")
	}
	return errors.Wrapf(err, errors.DatabaseError, "add post failed")
}
