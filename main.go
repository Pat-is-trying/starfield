package main

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
)

type Star struct {
	x, y int
	sx, sy float32
}
  
  func main() {

	var w, h int = 800, 800
	stars := make([]Star, 25)

	for i := 0; i < len(stars); i++ {
		stars[i] = Star{
			x: rand.Intn(w/2) + (w/4),
			y: rand.Intn(h/2) + (h/4),
			sx: 1.0,
			sy: 1.0,
		}
	}

    var palette = []color.Color{
        color.RGBA{0x00, 0x00, 0x00, 0xff},
        color.RGBA{255, 255, 255, 255},
    }

	var images []*image.Paletted
    var delays []int

	for i := 0; i < 40; i++ {

		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 40)

		for step := 0; step < len(stars); step++ {

			if stars[step].x > 400 {
				stars[step].sx *= 1.2
			} else {
				stars[step].sx /= 1.2
			}

			if stars[step].y > 400 {
				stars[step].sy *= 1.2
			} else {
				stars[step].sy /= 1.2
			}

			deltaX := int(float32(stars[step].x) * stars[step].sx)
			deltaY := int(float32(stars[step].y) * stars[step].sy)
			

			if stars[step].x > 800 || stars[step].y > 800 || stars[step].x < 0 || stars[step].y < 0 {
				stars = append(stars[:step], stars[step+1:]...)
			}
	
			img.Set(deltaX, deltaY, color.RGBA{255,255,255,255})
	
		}

		newStars := rand.Intn(10)
		for i := 0; i < newStars; i++ {
			stars = append(stars, Star{
				x: rand.Intn(w),
				y: rand.Intn(h),
				sx: 1.0,
				sy: 1.0,
			})
		}

	}


    f, _ := os.OpenFile("starfield.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, &gif.GIF{
        Image: images,
        Delay: delays,
    })
  }

 