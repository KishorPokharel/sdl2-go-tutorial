package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var title string = "Go and SDL2"
var width, height int32 = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var rect sdl.Rect

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width,
		height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatal("failed to create window", err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal("failed to create renderer", err)
	}
	defer renderer.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		rect = sdl.Rect{250, 250, 200, 200}
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&rect)

		renderer.Present()
		sdl.Delay(16)
	}
}
