package category

import "context"

type CatUseCase interface {
	AddCategory(ctx context.Context, category AddCatDto) (int64, error)
	GetByID(ctx context.Context, id int64) (*CategoryModel, error)
	GetCategories(ctx context.Context) ([]CatDto, error)
}
