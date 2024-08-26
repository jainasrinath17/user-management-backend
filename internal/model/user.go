package model

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName   string `gorm:"type:varchar(50);not null;unique" json:"user_name"`
	FirstName  string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName   string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email      string `gorm:"type:varchar(255);unique;not null" json:"email"`
	UserStatus string `gorm:"type:char(1);not null" json:"user_status"`
	Department string `gorm:"type:varchar(255)" json:"department,omitempty"`
}
