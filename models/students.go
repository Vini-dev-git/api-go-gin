package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Stundents struct {
	gorm.Model
	Name         string `json:name`
	CPF          string `json:cpf`
	RG           string `json:rg`
	Registration string `json:registration`
}

func ValidateStudent(student *Stundents) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
