package post

import "context"

type PostRepository interface {
	LikePost(ctx context.Context, postId int64, liekes LikePostDto) error
	Create(ctx context.Context, post PostModel) (int64, error)
	GetPosts(ctx context.Context) ([]PostModelWithUser, error)
	GetMyPosts(ctx context.Context, userId int64) ([]PostModelWithUser, error)
	GetPublishedPosts(ctx context.Context) ([]PostModelWithUser, error)
	GetById(ctx context.Context, postId int64) (PostModel, error)
	GetTotalLikesCountByUser(ctx context.Context, userId int64) (int64, error)
	Update(ctx context.Context, post PostModel) (int64, error)
}
