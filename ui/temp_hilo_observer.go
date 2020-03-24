package ui

import (
	"fmt"
	"weather_station/wmsdata"
)

func NewTempHiloUIObserver() TempHiloObserver {
	return TempHiloObserver{}
}

type TempHiloObserver struct {
}

func (observer TempHiloObserver) Update(data wmsdata.HiLoData) {
	fmt.Println("==============================")
	fmt.Print("High Value ", data.GetHighValue())
	fmt.Println(" at ", data.GetHighTime().Format("02/01/2006 15:04:05"))
	fmt.Print("Low Value ", data.GetLowValue())
	fmt.Println(" at ", data.GetLowTime().Format("02/01/2006 15:04:05"))
}
