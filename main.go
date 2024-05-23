package main

import (
	"fmt"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
) 

func main() {
	myApp := app.New()
    w := myApp.NewWindow("Simple Map")

    mapWidget := getMapImage(144642, 157324, 18) 
    w.SetContent(mapWidget)
    w.ShowAndRun()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}


func  getMapImage(latitute int, longitute int,zoom int)*canvas.Image {
    // https://tile.openstreetmap.org/18/144642/157324.png?lang=en
    map_image_url := fmt.Sprintf("https://tile.openstreetmap.org/%d/%d/%d.png?lang=en", zoom, latitute, longitute)
    fmt.Println(map_image_url)

    map_image_uri, err := storage.ParseURI(map_image_url)
    if err != nil {
        panic(err)
    }
    
    map_image_widget:= canvas.NewImageFromURI(map_image_uri)
    map_image_widget.FillMode = canvas.ImageFillOriginal
    return map_image_widget
}



