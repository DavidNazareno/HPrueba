package services

import (
	"log"

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
func GetTechOrder(id string) (*int64, error) {

	var count int64
	var ord Order
	err := controller.Database.Model(&ord).Where("technician = ? ", id).Count(&count).Error
	if err != nil {
		return nil, err
	}
	log.Println(id)

	return &count, nil
}
