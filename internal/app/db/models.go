package db

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name      string
	Fleets    []Fleet
	Flights   []Fleet
	CreatedOn time.Time `gorm:"autoCreateTime"`
	UpdatedOn time.Time `gorm:"autoUpdateTime"`
}

type Fleet struct {
	gorm.Model
	Nickname        string
	AircraftType    string `gorm:"size:4"`
	Registration    string
	CurrentLocation string `gorm:"size:4"`
	PlayerId        uint
	Flights         []Flight
	CreatedOn       time.Time `gorm:"autoCreateTime"`
	UpdatedOn       time.Time `gorm:"autoUpdateTime"`
}

type Flight struct {
	gorm.Model
	Departure      string `gorm:"size:4"`
	Arrival        string `gorm:"size:4"`
	JobDescription string
	FleetId        uint
	PlayerId       uint
	CreatedOn      time.Time `gorm:"autoCreateTime"`
	UpdatedOn      time.Time `gorm:"autoUpdateTime"`
}
