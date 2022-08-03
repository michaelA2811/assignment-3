package model

import (
	"math/rand"
	"time"
)

type WeatherStatus struct {
	Status Status `json:"status"`
}

type CompiledWeatherStatus struct {
	Status         Status `json:"status"`
	StatusCompiled string `json:"status_compiled"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomValueStatus() (status WeatherStatus) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	status.Status.Water = rand.Intn(max-min+1) + min
	status.Status.Wind = rand.Intn(max-min+1) + min
	return status
}

func (s *Status) CheckStatus() string {
	overallStatus := ""
	if s.Water <= 5 && s.Wind <= 6 {
		overallStatus = "SAFE"
	} else if s.Water <= 8 && s.Wind <= 15 {
		overallStatus = "STANDBY"
	} else {
		overallStatus = "DANGER"
	}
	return overallStatus
}
