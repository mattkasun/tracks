package tracks

import (
	"math"
	"time"
)

type Fix struct {
	Lat       float64
	Lng       float64
	timestamp time.Time
}

func (junk Fix) Create(lat, lng float64) Fix {
	var obj Fix
	obj.Lat = lat
	obj.Lng = lng
	obj.timestamp = time.Now()
	return obj
}

func (obj *Fix) Update(a float64, b float64) {
	obj.Lat = a
	obj.Lng = b
	obj.timestamp = time.Now()
}

func (obj Fix) GetPosition() (float64, float64) {
	return obj.Lat, obj.Lng
}

type track struct {
	Fix
	crs   float64
	speed float64
}

func (a track) create(lat, lng, course, speed float64) track {
	a.Lat = lat
	a.Lng = lng
	a.timestamp = time.Now()
	a.crs = course
	a.speed = speed
	return a
}

func (a *track) DeadReckon() {
	heading := a.crs * math.Pi / 180
	lat_rad := a.Lat * math.Pi / 180
	elapsed := time.Since(a.timestamp)
	lat := a.Lat + (a.speed * elapsed.Hours() / 60 * math.Cos(heading))
	lng := a.Lng + a.speed*elapsed.Hours()/60*math.Sin(heading)/math.Cos(lat_rad)
	a.Update(lat, lng)

}

func CreateFix(lat, lng float64) Fix {
	var obj Fix
	obj.Lat = lat
	obj.Lng = lng
	obj.timestamp = time.Now()
	return obj
}

func CreateTrack(lat, lng, course, speed float64) track {
	var a track
	a.Lat = lat
	a.Lng = lng
	a.timestamp = time.Now()
	a.crs = course
	a.speed = speed
	return a
}
