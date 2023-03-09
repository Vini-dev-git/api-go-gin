package models

import "gorm.io/gorm"

type Stundents struct {
	gorm.Model
	Name         string `json:name`
	CPF          string `json:cpf`
	RG           string `json:rg`
	Registration string `json:registration`
}
