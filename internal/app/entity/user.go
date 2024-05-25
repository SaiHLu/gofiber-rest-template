package entity

type User struct {
	ID    uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  *string `json:"name"`
	Email string  `json:"email" gorm:"not null"`
}
