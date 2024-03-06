package user

type UserDto struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Reputation int64  `json:"reputation"`
}

func (dto UserDto) MapFromModel(user UserModel) UserDto {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
	dto.Reputation = user.Reputation

	return dto
}

type AddUserDto struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Reputation int64  `json:"reputation"`
}

func (dto AddUserDto) MapToModel() (UserModel, error) {
	return NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
		dto.Reputation,
	)
}

type UpdateUserDto struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type ChangeUserPasswordDto struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
