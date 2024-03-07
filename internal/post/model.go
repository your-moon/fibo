package post

import (
	validation "github.com/go-ozzo/ozzo-validation"

	"fibo/internal/base/errors"
)

type PostModelWithUser struct {
	Id          int64
	UserId      int64
	Title       string
	Content     string
	CategoryId  int64
	Likes       int64
	IsPublished bool
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	UserEmail   string
	UserName    string
}
type PostModel struct {
	Id          int64
	UserId      int64
	Title       string
	Content     string
	CategoryId  int64
	Likes       int64
	IsPublished bool
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func NewPost(
	userId int64,
	title string,
	content string,
	isPublished bool,
	categoryId int64,
) (PostModel, error) {
	post := PostModel{
		UserId:      userId,
		Title:       title,
		Content:     content,
		IsPublished: isPublished,
		CategoryId:  categoryId,
	}

	if err := post.Validate(); err != nil {
		return PostModel{}, err
	}

	return post, nil
}

func (post *PostModel) Update(
	title string,
	content string,
	isPublished bool,
	likes int64,
	categoryId int64,
) error {
	if len(title) > 0 {
		post.Title = title
	}

	if len(content) > 0 {
		post.Content = content
	}

	post.IsPublished = isPublished

	post.CategoryId = categoryId

	if likes > 0 {
		post.Likes = likes
	}

	if err := post.Validate(); err != nil {
		return err
	}

	return nil
}

func (post *PostModel) Validate() error {
	err := validation.ValidateStruct(post,
		validation.Field(&post.Title, validation.Required),
		validation.Field(&post.Content, validation.Required),
		validation.Field(&post.UserId, validation.Required),
	)
	if err != nil {
		return errors.New(errors.ValidationError, err.Error())
	}

	return nil
}
