package dto

import "github.com/google/uuid"

type ParamUserDto struct {
	Id uint `json:"id"`
}

type QueryUserDto struct {
	Query  string `json:"query"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Order  string `json:"order"`
}

type IdParamUserDto struct {
	ID uuid.UUID `json:"id" validate:"uuid"`
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
