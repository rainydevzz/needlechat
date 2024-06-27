package models

type User struct {
	UserId    string `gorm:"primaryKey"`
	Name      string
	PublicKey string
}
