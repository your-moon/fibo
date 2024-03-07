package category

import (
	validation "github.com/go-ozzo/ozzo-validation"

	"fibo/internal/base/errors"
)

type CategoryModel struct {
	Id        int64
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type PublicCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) (CategoryModel, error) {
	category := CategoryModel{
		Name: name,
	}

	if err := category.Validate(); err != nil {
		return CategoryModel{}, err
	}

	return category, nil
}

func (cat *CategoryModel) Update(
	name string,
) error {
	if len(name) > 0 {
		cat.Name = name
	}

	if err := cat.Validate(); err != nil {
		return err
	}

	return nil
}

func (cat *CategoryModel) Validate() error {
	err := validation.ValidateStruct(cat,
		validation.Field(&cat.Name, validation.Required, validation.Length(3, 100)),
	)
	if err != nil {
		return errors.New(errors.ValidationError, err.Error())
	}

	return nil
}
