package post

import "context"

type PostRepository interface {
	Create(ctx context.Context, post PostModel) (int64, error)
	GetPosts(ctx context.Context) ([]PostModel, error)
	GetMyPosts(ctx context.Context, userId int64) ([]PostModel, error)
	GetPublishedPosts(ctx context.Context) ([]PostModel, error)
	GetById(ctx context.Context, postId int64) (PostModel, error)
	Update(ctx context.Context, post PostModel) (int64, error)
}
