package post

type LikePostDto struct {
	Likes int64 `json:"likes"`
}

type PostDto struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CategoryId  int64  `json:"category_id"`
	IsPublished bool   `json:"is_published"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Likes       int64  `json:"likes"`
}

type AddPostDto struct {
	UserId      int64  `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
	CategoryId  int64  `json:"category_id"`
}

func (p AddPostDto) MapToModel() (PostModel, error) {
	return NewPost(p.UserId, p.Title, p.Content, p.IsPublished, p.CategoryId)
}

type UpdatePostDto struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
	Likes       int64  `json:"likes"`
	CategoryId  int64  `json:"category_id"`
}

func (p UpdatePostDto) MapToModel() PostModel {
	return PostModel{
		Id:          p.Id,
		Title:       p.Title,
		Content:     p.Content,
		IsPublished: p.IsPublished,
		Likes:       p.Likes,
		CategoryId:  p.CategoryId,
	}
}
