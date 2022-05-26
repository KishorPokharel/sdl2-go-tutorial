package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	window, err := sdl.CreateWindow(
		"Window",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("could not create renderer: %v", err)
	}
	defer renderer.Destroy()

	running := true
	var x, y int32 = 250, 250
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_UP, sdl.K_k:
						y = y - 10
					case sdl.K_DOWN, sdl.K_j:
						y = y + 10
					case sdl.K_RIGHT, sdl.K_l:
						x = x + 10
					case sdl.K_LEFT, sdl.K_h:
						x = x - 10
					}
				}
			case *sdl.MouseMotionEvent:
				fmt.Printf("(%d, %d)\n", t.X, t.Y)
			}
		}

		renderer.SetDrawColor(255, 0, 0, 255)
		rect := sdl.Rect{x, y, 200, 200}
		renderer.FillRect(&rect)
		renderer.Present()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
	}

	return nil
}
