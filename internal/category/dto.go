package category

type CatDto struct {
	Id        int64  `json:"Id"`
	Name      string `json:"Name"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func (c CategoryModel) MapToDto() CatDto {
	return CatDto{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

type AddCatDto struct {
	Name string `json:"name"`
}

func (c AddCatDto) MapToModel() (CategoryModel, error) {
	return NewCategory(c.Name)
}

type UpdateCatDto struct {
	Name string `json:"name"`
}

func (c UpdateCatDto) MapToModel() CategoryModel {
	return CategoryModel{
		Name: c.Name,
	}
}
