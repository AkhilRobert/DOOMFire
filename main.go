package main

import (
	"image/color"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

var palette = []color.RGBA{
	{R: 7, G: 7, B: 7, A: 255},       //  0
	{R: 31, G: 7, B: 7, A: 255},      //  1
	{R: 47, G: 15, B: 7, A: 255},     //  2
	{R: 71, G: 15, B: 7, A: 255},     //  3
	{R: 87, G: 23, B: 7, A: 255},     //  4
	{R: 103, G: 31, B: 7, A: 255},    //  5
	{R: 119, G: 31, B: 7, A: 255},    //  6
	{R: 143, G: 39, B: 7, A: 255},    //  7
	{R: 159, G: 47, B: 7, A: 255},    //  8
	{R: 175, G: 63, B: 7, A: 255},    //  9
	{R: 191, G: 71, B: 7, A: 255},    // 10
	{R: 199, G: 71, B: 7, A: 255},    // 11
	{R: 223, G: 79, B: 7, A: 255},    // 12
	{R: 223, G: 87, B: 7, A: 255},    // 13
	{R: 223, G: 87, B: 7, A: 255},    // 14
	{R: 215, G: 95, B: 7, A: 255},    // 15
	{R: 215, G: 95, B: 7, A: 255},    // 16
	{R: 215, G: 103, B: 15, A: 255},  // 17
	{R: 207, G: 111, B: 15, A: 255},  // 18
	{R: 207, G: 119, B: 15, A: 255},  // 19
	{R: 207, G: 127, B: 15, A: 255},  // 20
	{R: 207, G: 135, B: 23, A: 255},  // 21
	{R: 199, G: 135, B: 23, A: 255},  // 22
	{R: 199, G: 143, B: 23, A: 255},  // 23
	{R: 199, G: 151, B: 31, A: 255},  // 24
	{R: 191, G: 159, B: 31, A: 255},  // 25
	{R: 191, G: 159, B: 31, A: 255},  // 26
	{R: 191, G: 167, B: 39, A: 255},  // 27
	{R: 191, G: 167, B: 39, A: 255},  // 28
	{R: 191, G: 175, B: 47, A: 255},  // 29
	{R: 183, G: 175, B: 47, A: 255},  // 30
	{R: 183, G: 183, B: 47, A: 255},  // 31
	{R: 183, G: 183, B: 55, A: 255},  // 32
	{R: 207, G: 207, B: 111, A: 255}, // 33
	{R: 223, G: 223, B: 159, A: 255}, // 34
	{R: 239, G: 239, B: 199, A: 255}, // 35
	{R: 255, G: 255, B: 255, A: 255}, // 36
}

const (
	WIDTH  = 108
	HEIGHT = 70
)

func main() {
	size := 8
	screenWidth := (WIDTH * size) + 1   // Additional space for the last one
	screenHeight := (HEIGHT * size) + 1 // Additional space for the last one
	pixels := [HEIGHT * WIDTH]int{}

	// Setting the bottom line to white
	for i := WIDTH; i > 0; i-- {
		pixels[(HEIGHT*WIDTH)-i] = len(palette) - 1
	}

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic("unable to intialize sdl")
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("DOOMFire", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(screenWidth), int32(screenHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic("unable to create window")
	}
	window.SetResizable(true)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic("unable to create renderer")
	}
	defer renderer.Destroy()

main:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				break main

			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_q:
					for i := WIDTH; i > 0; i-- {
						reduced := pixels[(HEIGHT*WIDTH)-i] - 1
						if reduced >= 0 {
							pixels[(HEIGHT*WIDTH)-i] = reduced
						}
					}

				case sdl.K_e:
					for i := WIDTH; i > 0; i-- {
						increased := pixels[(HEIGHT*WIDTH)-i] + 1
						if increased < len(palette) {
							pixels[(HEIGHT*WIDTH)-i] = increased
						}
					}
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// Handling fire propogation
		for i := 0; i < WIDTH; i++ {
			for j := 0; j < HEIGHT; j++ {
				current := i + (j * WIDTH)

				below := current + WIDTH

				if below > (WIDTH*HEIGHT)-1 {
					continue
				}

				r := rand.Int() % 3
				c := current - rand.Int()%2
				intensity := pixels[below] - r
				if intensity >= 0 && c >= 0 {
					pixels[c] = intensity
				}
			}
		}

		for i := 0; i < WIDTH; i++ {
			for j := 0; j < HEIGHT; j++ {
				c := palette[pixels[i+(j*WIDTH)]]
				renderer.SetDrawColor(c.R, c.G, c.B, c.A)
				rect := sdl.Rect{X: int32(i * size), Y: int32(j * size), W: int32(size), H: int32(size)}
				renderer.FillRect(&rect)
			}
		}

		renderer.Present()
		sdl.Delay(1000 / 60) // 60 FPS
	}
}
