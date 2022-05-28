package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	window, err := sdl.CreateWindow(
		"Text Rendering",
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

	if err = ttf.Init(); err != nil {
		return fmt.Errorf("could not init ttf: %v", err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("res/fonts/SplineSansMonoBold.ttf", 200)
	if err != nil {
		return fmt.Errorf("could not open font: %v", err)
	}
	defer font.Close()

	fontSurface, err := font.RenderUTF8Solid("Text", sdl.Color{255, 255, 255, 255})
	if err != nil {
		return fmt.Errorf("could not create font surface: %v", err)
	}
	defer fontSurface.Free()

	texture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		return fmt.Errorf("could not create texture from font surface: %v", err)
	}
	defer texture.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.Copy(texture, nil, &sdl.Rect{(800 / 2) - (200 / 2), (600 / 2) - (200 / 2), 200, 200})
		renderer.Present()
	}
	return nil
}
