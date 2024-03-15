package entities

import "gorm.io/gorm"

type BackofficeRequest struct {
	ID        uint32 `gorm:"primary_key;auto_increment"`
	Nome      string `gorm:"not null"`
	Endereco  string `gorm:"not null"`
	Telefone  string `gorm:"not null"`
	Processed bool   `gorm:"not null"`
	gorm.Model
}
