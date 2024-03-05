package post

import "context"

type PostUseCase interface {
	AddPost(ctx context.Context, post AddPostDto) (int64, error)
	GetPosts(ctx context.Context) ([]PostModel, error)
	GetMyPosts(ctx context.Context, userId int64) ([]PostModel, error)
	GetPublishedPosts(ctx context.Context) ([]PostModel, error)
	UpdatePost(ctx context.Context, post UpdatePostDto) error
	GetPostById(ctx context.Context, id int64) (PostModel, error)
}
