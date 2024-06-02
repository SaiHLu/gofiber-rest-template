package dto

type ParamUserDto struct {
	Id uint `json:"id"`
}

type QueryUserDto struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	Search   string `json:"search"`
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserDto struct {
	Name  string `json:"name" updatereq:"omitempty,required"`
	Email string `json:"email" updatereq:"omitempty,required,email"`
}
