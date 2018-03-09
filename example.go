package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	gtkmap "github.com/mattkasun/gtk-map"
	"math"
	"sandbox/geo"
	"time"
)

type OwnShip struct {
	geo.Track
	desired_crs float64
}

func main() {
	var (
		ownship geo.Track
		wp      geo.Position
		dist    float64
		bearing float64
	)

	gtk.Init(nil)
	builder, _ := gtk.BuilderNewFromFile("info.glade")

	lat, _ := builder.GetObject("lat")
	lng, _ := builder.GetObject("lng")
	crs, _ := builder.GetObject("course")
	speed, _ := builder.GetObject("speed")
	wplat, _ := builder.GetObject("wplat")
	wplng, _ := builder.GetObject("wplng")
	dtg, _ := builder.GetObject("dtg")
	heading, _ := builder.GetObject("heading")

	wp = geo.Position{45, -76}
	ownship = geo.Track{geo.Fix{geo.Position{44.80, -75.20}, time.Now()}, 0, 250}
	lat.(*gtk.Entry).SetText(fmt.Sprintf("%4.2f", ownship.Lat))
	lng.(*gtk.Entry).SetText(fmt.Sprintf("%4.2f", ownship.Lng))
	crs.(*gtk.Entry).SetText(fmt.Sprintf("%4.2f", ownship.Crs))
	speed.(*gtk.Entry).SetText(fmt.Sprintf("%4.2f", ownship.Speed))

	tick1 := time.NewTicker(time.Second)
	defer tick1.Stop()
	tick2 := time.NewTicker(time.Second)
	defer tick2.Stop()
	tick3 := time.NewTicker(time.Second * 10)
	defer tick3.Stop()

	go func() {

		for {
			_ = <-tick1.C
			dist, bearing = ownship.Haversine(wp.Lat, wp.Lng)
			_, _ = glib.IdleAdd(wplat.(*gtk.Entry).SetText, fmt.Sprintf("%4.2f", wp.Lat))
			_, _ = glib.IdleAdd(wplng.(*gtk.Entry).SetText, fmt.Sprintf("%4.2f", wp.Lng))
			_, _ = glib.IdleAdd(dtg.(*gtk.Entry).SetText, fmt.Sprintf("%6.2f", dist))
			_, _ = glib.IdleAdd(heading.(*gtk.Entry).SetText, fmt.Sprintf("%4.2f", bearing))
		}
	}()

	go func() {
		for {
			_ = <-tick2.C
			//check course
			fmt.Println("cource bearing", ownship.Crs, bearing)
			if ownship.Crs > bearing {
				if ownship.Crs-bearing < 180 {
					ownship.Crs = ownship.Crs - 1
				} else {
					ownship.Crs = ownship.Crs + 1
				}
			} else if ownship.Crs < bearing {
				if math.Abs(ownship.Crs-bearing) > 180 {
					ownship.Crs = ownship.Crs - 1
				} else {
					ownship.Crs = ownship.Crs + 1
				}
			}
			if ownship.Crs < 0 {
				ownship.Crs = ownship.Crs + 360
			}
			if math.Abs(ownship.Crs-bearing) < 1 {
				ownship.Crs = bearing
			}
			_, _ = glib.IdleAdd(crs.(*gtk.Entry).SetText, fmt.Sprintf("%3.0f", ownship.Crs))
			fmt.Println("update crs ", ownship.Crs, bearing)

			//deadreckon
			ownship.DeadReckon()
			//s := fmt.Sprintf("%4.2f", ownship.Lat)
			_, _ = glib.IdleAdd(lat.(*gtk.Entry).SetText, fmt.Sprintf("%4.2f", ownship.Lat))
			_, _ = glib.IdleAdd(lng.(*gtk.Entry).SetText, fmt.Sprintf("%4.2f", ownship.Lng))

		}
	}()

	m, err := gtkmap.MapNew()
	if err != nil {
		fmt.Println("error creating map")
	}
	err = m.SetProperty("map-source", gtkmap.SourceVirtualEarthHybrid)
	if err != nil {
		fmt.Println("failed to set source", err)
	}
	m.SetCenterAndZoom(45.18, -75.93, 8)
	osd := gtkmap.OsdNewFull(true, true, true, true, true, false, false, false, 30.5)
	m.LayerAdd(osd)
	m.GpsAdd(44.80, -75.20, 0)

	go func() {
		for {
			_ = <-tick3.C
			m.GpsAdd(ownship.Lat, ownship.Lng, ownship.Crs)
		}
	}()

	box, _ := builder.GetObject("mapbox")
	box.(*gtk.Box).PackStart(m, true, true, 0)
	w, _ := builder.GetObject("window")
	w.(*gtk.Window).ShowAll()
	w.(*gtk.Window).Connect("destroy", gtk.MainQuit)
	gtk.Main()
}
