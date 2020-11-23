package services

import (
	"log"

	"github.com/DavidNazareno/h_prueba/controller"
	"github.com/DavidNazareno/h_prueba/utils"
)

type Request struct {
	Token  string `gorm:"primaryKey;autoIncrement:false"`
	Client uint
	Status int `gorm:"default:1"`
	Score  int `gorm:"default:0"`
}

//NewRequest crear una solicitud del servicio
func NewRequest(client Client) (*Request, error) {

	var req = Request{
		Token:  utils.Generate(8),
		Client: client.ID,
		Status: 1,
		Score:  0,
	}

	err := controller.Database.Create(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}

//UpdateRequest Actualizar estado de la solicitud
func UpdateRequest(token string, score int) error {
	var req = Request{}
	err := controller.Database.First(&req, "token = ?", token).Error

	if err != nil {
		log.Println(err)

		return err
	}

	log.Println(req)
	req.Score = score

	controller.Database.Save(&req)

	return nil
}

func GetRequest(token string) (*Request, error) {

	var req = Request{}

	err := controller.Database.First(&req, "token = ?", token).Error

	if err != nil {
		log.Println(err)

		return nil, err
	}
	return &req, nil
}
