package impl

import (
	"context"

	"fibo/internal/base/database"
	"fibo/internal/category"
)

type CatUsecaseOpts struct {
	CatRepository category.CatRepository
	TxManager     database.TxManager
}

func NewCatUsecase(opts CatUsecaseOpts) category.CatUseCase {
	return &catUseCase{
		CatRepository: opts.CatRepository,
		TxManager:     opts.TxManager,
	}
}

type catUseCase struct {
	category.CatRepository
	database.TxManager
}

func (c *catUseCase) GetCategories(
	ctx context.Context,
) ([]category.CatDto, error) {
	models, err := c.CatRepository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	var result []category.CatDto
	for _, model := range models {
		result = append(result, model.MapToDto())
	}

	return result, nil
}

func (c *catUseCase) GetByID(
	ctx context.Context,
	id int64,
) (*category.CategoryModel, error) {
	return c.CatRepository.GetById(ctx, id)
}

func (c *catUseCase) AddCategory(
	ctx context.Context,
	category category.AddCatDto,
) (catId int64, err error) {
	model, err := category.MapToModel()
	if err != nil {
		return 0, err
	}

	err = c.RunTx(ctx, func(ctx context.Context) error {
		_, err := c.CatRepository.Add(ctx, model)
		if err != nil {
			return err
		}
		model.Id = catId
		return nil
	})
	if err != nil {
		return 0, err
	}

	return catId, err
}
