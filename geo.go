package geo

import (
	"math"
	"time"
)

//mean radius of earth in NM
const R = 3440

// a geographic position
type Position struct {
	Lat float64
	Lng float64
}

//a position at a given point in time
type Fix struct {
	Position
	TimeStamp time.Time
}

//update position of a fix
func (f *Fix) Update(a float64, b float64) {
	f.Lat = a
	f.Lng = b
	f.TimeStamp = time.Now()
}

// a moving object
type Track struct {
	Fix
	Crs   float64
	Speed float64
}

//update position of track based on its course and speed
func (t *Track) DeadReckon() {
	heading := t.Crs * math.Pi / 180
	lat_rad := t.Lat * math.Pi / 180
	elapsed := time.Since(t.TimeStamp)
	lat := t.Lat + (t.Speed * elapsed.Hours() / 60 * math.Cos(heading))
	lng := t.Lng + t.Speed*elapsed.Hours()/60*math.Sin(heading)/math.Cos(lat_rad)
	t.Update(lat, lng)
}

// great circle distance (NM) and initial bearing between two points
func (p *Position) Haversine(lat, lng float64) (d, b float64) {
	//distance
	dlat := DegToRad(lat) - DegToRad(p.Lat)
	dlng := DegToRad(lng) - DegToRad(p.Lng)
	x := math.Sin(dlat/2) * math.Sin(dlat/2)
	y := math.Sin(dlng/2) * math.Sin(dlng/2)
	a := x + math.Cos(DegToRad(p.Lat))*math.Cos(DegToRad(lat))*y
	min := math.Min(1, math.Sqrt(a))
	c := 2 * math.Asin(min)
	d = R * c
	//initial bearing
	x = math.Sin(dlng) * math.Cos(DegToRad(lat))
	y = math.Cos(DegToRad(p.Lat))*math.Sin(DegToRad(lat)) - math.Sin(DegToRad(p.Lat))*math.Cos(DegToRad(lat))*math.Cos(dlng)
	b = math.Atan2(x, y)
	if b < 0 {
		b = 2*math.Pi + b
	}
	return d, RadToDeg(b)
}

//convert from degress to radians
func DegToRad(degree float64) float64 {
	return (degree * math.Pi / 180)
}

//convert from radians to degrees
func RadToDeg(rad float64) float64 {
	return (rad * 180 / math.Pi)
}

//rhumb line distance and bearing to a point
func (p *Position) RhumbLine(lat, lng float64) (d, b float64) {
	//distance
	var q float64
	lat1 := DegToRad(p.Lat)
	lng1 := DegToRad(p.Lng)
	lat2 := DegToRad(lat)
	lng2 := DegToRad(lng)
	dlat := lat2 - lat1
	dlng := lng2 - lng1
	lat_diff := math.Log(math.Tan(math.Pi/4+lat2/2) / math.Tan(math.Pi/4+lat1/2))
	if lat1 != lat2 {
		q = dlat / lat_diff
	} else {
		q = math.Cos(lat1)
	}
	d = math.Sqrt(dlat*dlat+q*q*dlng*dlng) * R
	//bearing
	if math.Abs(dlng) > math.Pi {
		if dlng > 0 {
			dlng = -(2*math.Pi - dlng)
		} else {
			dlng = 2*math.Pi + dlng
		}
	}
	b = math.Atan2(dlng, lat_diff)
	b = math.Mod((b + 2*math.Pi), 2*math.Pi)
	//if b < 0 {
	//	b = 2*math.Pi + b
	//}
	return d, RadToDeg(b)
}
