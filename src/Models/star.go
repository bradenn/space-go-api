package Models

import (
	"space-api/src/Config"
)

type Star struct {
	Name               string `json:"name" gorm:"column:name"`
	Id                 string `json:"id" gorm:"column:id"`
	Type               string `json:"type" gorm:"column:type"`
	Age                string `json:"age" gorm:"column:age"`
	SpectralClass      string `json:"spectralClass" gorm:"column:spectralClass"`
	Luminosity         string `json:"luminosity" gorm:"column:luminosity"`
	AbsoluteMagnitude  string `json:"absoluteMagnitude" gorm:"column:absoluteMagnitude"`
	SolarMasses        string `json:"solarMasses" gorm:"column:solarMasses"`
	SolarRadius        string `json:"solarRadius" gorm:"column:solarRadius"`
	SurfaceTemperature string `json:"surfaceTemperature" gorm:"column:surfaceTemperature"`
	RotationalPeriod   string `json:"rotationalPeriod" gorm:"column:rotationalPeriod"`
	UpdateTime         string `json:"updateTime" gorm:"column:updateTime"`
}

func (b *Star) TableName() string {
	return "stars"
}

func FindAllStars(stars *[]Star) (err error) {
	if err := Config.DB.Find(stars).Error; err != nil {
		return err
	}
	return nil
}

func FindStar(star *Star, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(star).Error; err != nil {
		return err
	}
	return nil
}
