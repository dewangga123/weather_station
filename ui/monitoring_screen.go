package ui

import (
	"fmt"
)

type monitoringScreen struct{}

func (screen *monitoringScreen) DisplayTemperature(data float32) {
	fmt.Println("Temperature Value", data)
}

func (screen *monitoringScreen) DisplayPressure(data float32) {
	fmt.Println("Pressure Value", data)
}

func (screen *monitoringScreen) DisplayPressureTrend(data float32) {
	fmt.Println("Pressure Value", data)
}
