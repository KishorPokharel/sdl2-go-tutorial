package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	window, err := sdl.CreateWindow(
		"Image Render",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
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

	dinoSprites := make([]*sdl.Texture, 10)
	for i := 1; i <= 10; i++ {
		image, err := img.Load(fmt.Sprintf("res/walk/Walk (%d).png", i))
		if err != nil {
			return fmt.Errorf("could not load image: %v", err)
		}
		texture, err := renderer.CreateTextureFromSurface(image)
		if err != nil {
			return fmt.Errorf("could not create texture from surface: %v", err)
		}
		dinoSprites[i-1] = texture
	}
	defer func() {
		for i := 1; i <= 10; i++ {
			dinoSprites[i-1].Destroy()
		}
	}()

	running := true
	sprite := 0
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		rect := sdl.Rect{0, 0, 680, 472}
		renderer.Copy(dinoSprites[sprite%10], nil, &rect)
		renderer.Present()
		sdl.Delay(70)
		sprite++
	}

	return nil
}
