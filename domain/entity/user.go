package entity

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}
