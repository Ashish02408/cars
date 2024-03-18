package cars

import (
	"github.com/Ashish02408/vehicle"
)

func Sound() string {
	return "go"
}

func Sounds() string {
	return "go go go"
}

func SoundsCrazy() string {
	return vehicle.WhenComesUp(Sound())
}

func SoundsCrazys() string {
	return vehicle.WhenComesUp(Sounds())
}
