package Models

import (
	"space-api/Config"
)

type System struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	X         string `json:"x"`
	Y         string `json:"y"`
	Z         string `json:"z"`
	StarId    string `json:"starId" gorm:"column:starId"`
}

func (b *System) TableName() string {
	return "systems"
}

func FindAllSystems(systems *[]System) (err error) {
	if err := Config.DB.Find(systems).Error; err != nil {
		return err
	}
	return nil
}

func FindSystem(system *System, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(system).Error; err != nil {
		return err
	}
	return nil
}
