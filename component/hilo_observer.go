package component

import "weather_station/wmsdata"

type HiloObserver interface {
	Update(data wmsdata.HiLoData)
}
