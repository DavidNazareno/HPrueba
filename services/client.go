package services

import (
	"log"

	"github.com/DavidNazareno/h_prueba/controller"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Ticket string
	Requests []Request `gorm:"foreignKey:Client"`
}

//Add Agrega in nuevo cliente y retornar tikect
func (client Client) Add() (*Client, error) {

	var cli = Client{
		Ticket: client.Ticket,
	}

	err := controller.Database.Create(&cli).Error

	if err != nil {
		return nil, err
	}
	return &cli, nil

}
//Check validar Ticket del cliente
func (client Client) Check() (*Client, bool, error) {

	var cli Client
	err := controller.Database.First(&cli, "ticket = ?", client.Ticket).Error
	if err != nil {
		log.Println(err)

		return nil, false, err
	}

	return &cli, true, nil

}
