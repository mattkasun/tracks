package geo

import (
	"math"
)

//acceleration of gravity in NM/sec2
const G = 0.005295113391

type Ftp struct {
	Position
}

type Wind struct {
	Speed     float64
	Direction float64
}

type OAC struct {
	Track
	TAS       float64
	Alt       float64
	BankAngle float64
}

func (o *OAC) TurnCircle(w Wind) float64 {
	radius := ((o.TAS + w.Speed) * (o.TAS + w.Speed)) / (G * math.Tan(DegToRad(o.BankAngle)))
	return RadToDeg(radius)
}
