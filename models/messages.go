package models

type Message struct {
	MessageId string `gorm:"primaryKey"`
	Content   string
	Author    string `gorm:"references:Name"`
	Nonce     string
}
