package Models

import (
	"space-api/src/Config"
)

type Planet struct {
	Id                 string `json:"id" gorm:"column:id"`
	Name               string `json:"name" gorm:"column:name"`
	Type               string `json:"type" gorm:"column:type"`
	DistanceToArrival  string `json:"distanceToArrival" gorm:"column:distanceToArrival"`
	EarthMasses        string `json:"earthMasses" gorm:"column:earthMasses"`
	Radius             string `json:"radius" gorm:"column:radius"`
	SurfaceTemperature string `json:"surfaceTemperature" gorm:"column:surfaceTemperature"`
	AtmosphereType     string `json:"atmosphereType" gorm:"column:atmosphereType"`
	OrbitalPeriod      string `json:"orbitalPeriod" gorm:"column:orbitalPeriod"`
	RotationalPeriod   string `json:"rotationalPeriod" gorm:"column:rotationalPeriod"`
	UpdateTime         string `json:"updateTime" gorm:"column:updateTime"`
	StarId             string `json:"starId" gorm:"column:starId"`
}

func (b *Planet) TableName() string {
	return "planets"
}

func FindAllPlanets(planets *[]Planet) (err error) {
	if err := Config.DB.Find(planets).Error; err != nil {
		return err
	}
	return nil
}

func FindPlanet(planet *Planet, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(planet).Error; err != nil {
		return err
	}
	return nil
}
