package category

import "context"

type CatRepository interface {
	Add(ctx context.Context, post CategoryModel) (int64, error)
	GetById(ctx context.Context, id int64) (*CategoryModel, error)
	GetCategories(ctx context.Context) ([]CategoryModel, error)
}
