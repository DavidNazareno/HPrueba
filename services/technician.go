package services

import (
	"github.com/DavidNazareno/h_prueba/controller"
	"github.com/jinzhu/gorm"
)

type Technician struct {
	gorm.Model
	Name string
	Orders []Order `gorm:"foreignKey:Technician"`
}

//GetRandomTech Obtener Tecnico aleatorio para agregar a la nueva solicitud
func  GetRandomTech() (*Technician, error) {

	var t = Technician{}
	err := controller.Database.Raw("SELECT * FROM technicians ORDER BY RANDOM() LIMIT 1").Scan(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil


}
