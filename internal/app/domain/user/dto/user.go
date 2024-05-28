package dto

type ParamUserDto struct {
	UserId uint `json:"id"`
}

type QueryUserDto struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	Search   string `json:"search" validate:"max=1"`
}

type CreateUserDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserDto struct {
	Name  string `json:"name" updatereq:"omitempty,required"`
	Email string `json:"email" updatereq:"omitempty,required,email"`
}
