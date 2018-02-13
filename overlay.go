package main

import (
	"fmt"
	"github.com/mattkasun/gtkmap"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func main() {
	gtk3.Init(&os.Args)

	builder := gtk3.NewBuilder()
	builder.AddFromFile("overlay.ui")
	//builder.ConnectSignals(nil)

	obj := builder.GetObject("window1")
	window := gtk3.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk3.MainQuit)

	bestfit := gtk3.WidgetFromObject(builder.GetObject("BestFitButton"))
	bestfit.Connect("toggled", func() {
		fmt.Println("Best Fit pressed")
	})

	bestfit.Connect("clicked", func() {
		fmt.Println("Best Fit Button clicked")
	})

	radiobutton1 := gtk3.WidgetFromObject(builder.GetObject("radiobutton1"))
	radiobutton1.Connect("toggled", func() {
		fmt.Println("radio1 button selected")
		bestfit.Hide()
		fmt.Println("done")
	})
	radiobutton2 := gtk3.WidgetFromObject(builder.GetObject("radiobutton2"))
	radiobutton2.Connect("toggled", func() {
		fmt.Println("radio2 button selected")
		bestfit.Show()
	})

	errorbutton := gtk3.WidgetFromObject(builder.GetObject("ErrorButton"))
	errorbutton.Connect("activate", func() {
		fmt.Println("Error Button pressed")
	})
	errorbutton.Connect("clicked", func() {
		fmt.Println("Error Button clicked")
	})

	source := gtkmap.SourceVirtualEarthHybrid
	m, err := gtkmap.NewMapOpt(source)
	if err != nil {
		gtk3.MainQuit()
	}

	m.SetCenter(gtkmap.Coord(44.96, -76.02))
	m.SetZoom(10)
	var overlay gtk3.Container
	overlay = gtk3.Container(gtk3.WidgetFromObject(builder.GetObject("overlay")))
	overlay.Add(m)
	fmt.Println("hello")
	gtk3.Main()

}
