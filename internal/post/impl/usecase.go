package impl

import (
	"context"
	"fmt"

	"fibo/internal/base/database"
	"fibo/internal/post"
)

type PostUsecaseOpts struct {
	PostRepository post.PostRepository
	TxManager      database.TxManager
}

func NewPostUsecase(opts PostUsecaseOpts) post.PostUseCase {
	return &postUseCase{
		PostRepository: opts.PostRepository,
		TxManager:      opts.TxManager,
	}
}

type postUseCase struct {
	post.PostRepository
	database.TxManager
}

func (p *postUseCase) LikePost(
	ctx context.Context,
	postId int64,
	likes post.LikePostDto,
) (err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		err = p.PostRepository.LikePost(ctx, postId, likes)
		return err
	})

	return err
}

func (p *postUseCase) GetTotalLikesCountByUser(
	ctx context.Context,
	userId int64,
) (likes int64, err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		likes, err = p.PostRepository.GetTotalLikesCountByUser(ctx, userId)
		return err
	})
	return likes, err
}

func (p *postUseCase) GetPublishedPosts(
	ctx context.Context,
) (posts []post.PostModelWithUser, err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		posts, err = p.PostRepository.GetPublishedPosts(ctx)
		return err
	})
	return posts, err
}

func (p *postUseCase) GetMyPosts(
	ctx context.Context,
	userId int64,
) (posts []post.PostModelWithUser, err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		posts, err = p.PostRepository.GetMyPosts(ctx, userId)
		return err
	})
	return posts, err
}

func (p *postUseCase) GetPosts(ctx context.Context) (posts []post.PostModelWithUser, err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		posts, err = p.PostRepository.GetPosts(ctx)
		return err
	})
	return posts, err
}

func (p *postUseCase) GetPostById(ctx context.Context, id int64) (post post.PostModel, err error) {
	err = p.RunTx(ctx, func(ctx context.Context) error {
		post, err = p.PostRepository.GetById(ctx, id)
		return err
	})

	return post, err
}

func (p *postUseCase) UpdatePost(
	ctx context.Context,
	post post.UpdatePostDto,
) (err error) {
	model, err := p.PostRepository.GetById(ctx, post.Id)
	if err != nil {
		return err
	}

	err = model.Update(post.Title, post.Content, post.IsPublished, post.Likes, post.CategoryId)
	if err != nil {
		return err
	}

	modelId, err := p.PostRepository.Update(ctx, model)
	if err != nil {
		return err
	}

	if modelId != model.Id {
		return fmt.Errorf("model id and returned id are different")
	}

	return nil
}

func (p *postUseCase) AddPost(ctx context.Context, post post.AddPostDto) (postId int64, err error) {
	model, err := post.MapToModel()
	if err != nil {
		return 0, err
	}

	err = p.RunTx(ctx, func(ctx context.Context) error {
		postId, err = p.PostRepository.Create(ctx, model)
		if err != nil {
			return err
		}
		model.Id = postId
		return nil
	})

	return postId, err
}
