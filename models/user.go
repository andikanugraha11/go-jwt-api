package models

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password string    `gorm:"not null" json:"password" form:"password"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}
