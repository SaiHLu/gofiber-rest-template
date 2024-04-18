package user_model

type UserModel struct {
	ID    uint   `gorm:"primaryKey" json:"id" validate:"omitempty,required"`
	Name  string `json:"name,omitempty" validate:"omitempty,required"`
	Email string `gorm:"unique" json:"email" validate:"required,email"`
}
