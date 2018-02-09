package main

import (
	"fmt"
	"math"
	"time"
)

type fix struct {
	lat       float64
	lng       float64
	timestamp time.Time
}

func (obj fix) create(lat, lng float64) fix {
	obj.lat = lat
	obj.lng = lng
	obj.timestamp = time.Now()
	return obj
}

func (obj *fix) update(a float64, b float64) {
	obj.lat = a
	obj.lng = b
	obj.timestamp = time.Now()
}

func (obj fix) get_position() (float64, float64) {
	return obj.lat, obj.lng
}

type track struct {
	fix
	crs   float64
	speed float64
}

func (a *track) dead_reckon() {
	heading := a.crs * math.Pi / 180
	lat_rad := a.lat * math.Pi / 180
	elapsed := time.Since(a.timestamp)
	lat := a.lat + (a.speed * elapsed.Hours() / 60 * math.Cos(heading))
	lng := a.lng + a.speed*elapsed.Hours()/60*math.Sin(heading)/math.Cos(lat_rad)
	a.update(lat, lng)

}

func main() {
	var dummy fix
	a := track{fix{45.0, -75.0, time.Now()}, 300, 100.0}
	b := fix.create(dummy, 45, 180)
	c := new(fix)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	time.Sleep(3600 * time.Millisecond)
	a.dead_reckon()
	fmt.Println(a)
}
