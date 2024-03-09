package post

import "context"

type PostUseCase interface {
	LikePost(ctx context.Context, postId int64, likes LikePostDto) error
	AddPost(ctx context.Context, post AddPostDto) (int64, error)
	GetPosts(ctx context.Context) ([]PostModelWithUser, error)
	GetMyPosts(ctx context.Context, userId int64) ([]PostModelWithUser, error)
	GetPublishedPosts(ctx context.Context) ([]PostModelWithUser, error)
	GetTotalLikesCountByUser(ctx context.Context, userId int64) (int64, error)
	UpdatePost(ctx context.Context, post UpdatePostDto) error
	GetPostById(ctx context.Context, id int64) (PostModel, error)
}
