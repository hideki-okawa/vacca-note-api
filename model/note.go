package model

import (
	"time"
)

type Note struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Gender             string    `json:"gender"`
	Age                string    `json:"age"`
	VaccineType        string    `json:"vaccine_type"`
	SecondVaccineType  string    `json:"second_vaccine_type"`
	NuberOfVaccination int       `json:"number_of_vaccination"`
	MaxTemperature     string    `json:"max_temperature"`
	Log                string    `json:"log"`
	Remarks            string    `json:"remarks"`
	GoodCount          int       `json:"good_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type TemperatureData struct {
	Name string             `json:"name"`
	List []TemperatureCount `json:"list"`
}

type TemperatureCount struct {
	Num   string `json:"num"`
	Count int    `json:"count"`
}
