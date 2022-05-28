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
		"Collison Detection",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("could not create renderer: %v", err)
	}
	running := true
	var x, y int32 = 250, 250
	i := 0
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				x, y = t.X, t.Y
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		rectA := sdl.Rect{20, 20, 200, 200}
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&rectA)

		rectB := sdl.Rect{x, y, 200, 200}
		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.FillRect(&rectB)
		if rectA.HasIntersection(&rectB) {
			fmt.Println("Collison Occured ", i)
		}
		i++

		renderer.Present()
	}
	return nil
}
