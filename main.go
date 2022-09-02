package main

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
)

type Star struct {
	x, y, z int
}
  
  func main() {

	var w, h int = 800, 800
	stars := make([]Star, 50)

	for i := 0; i < len(stars); i++ {
		stars[i] = Star{
			x: rand.Intn(w),
			y: rand.Intn(h),
			z: w,
		}
	}

    var palette = []color.Color{
        color.RGBA{0x00, 0x00, 0x00, 0xff},
        color.RGBA{255, 255, 255, 255},
    }

	var images []*image.Paletted
    var delays []int

	for i := 0; i < 20; i++ {

		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for step := 0; step < len(stars); step++ {

			stars[step].z /= 2

			if stars[step].z == 0 {
				break
			}
	
			sx := (stars[step].x) / stars[step].z
			sy := (stars[step].y) / stars[step].z
	
			img.Set(sx, sy, color.RGBA{255,255,255,255})
	
		}
	}


    f, _ := os.OpenFile("starfield.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, &gif.GIF{
        Image: images,
        Delay: delays,
    })
  }

 