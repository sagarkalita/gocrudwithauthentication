package models

type Data struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Address   string
	State     string
	PhotoPath string
}
