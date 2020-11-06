package Models

import (
	"space-api/src/Config"
)

type Moon struct {
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
	PlanetId           string `json:"planetId" gorm:"column:planetId"`
}

func (b *Moon) TableName() string {
	return "moons"
}

func FindAllMoons(moons *[]Moon) (err error) {
	if err := Config.DB.Find(moons).Error; err != nil {
		return err
	}
	return nil
}

func FindMoon(moon *Moon, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(moon).Error; err != nil {
		return err
	}
	return nil
}
