package services

import (
	"github.com/DavidNazareno/h_prueba/controller"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Technician uint
	Token      string
}

//CreateOrder  Asignar order a la tabla de ordenes relacionada con la tabla tecnicos
func CreateOrder(token string, tech Technician) (*Order, error) {

	var ord = Order{
		Technician: tech.ID,
		Token:      token,
	}
	err := controller.Database.Create(&ord).Error

	if err != nil {
		return nil, err
	}

	return &ord, nil

}
//GetTechOrder  Obtener todas las ordenes
func GetTechOrder(id string) ([]Request, error) {

	var t []Request
	err := controller.Database.Where("id = ?", id).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
