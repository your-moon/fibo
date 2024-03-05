package post

type PostDto struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type AddPostDto struct {
	UserId      int64  `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
}

func (p AddPostDto) MapToModel() (PostModel, error) {
	return NewPost(p.UserId, p.Title, p.Content, p.IsPublished)
}

type UpdatePostDto struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
}

func (p UpdatePostDto) MapToModel() PostModel {
	return PostModel{
		Id:          p.Id,
		Title:       p.Title,
		Content:     p.Content,
		IsPublished: p.IsPublished,
	}
}
