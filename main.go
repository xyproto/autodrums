package main

import (
	"log"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_AUDIO | sdl.INIT_VIDEO); err != nil {
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Autodrums", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 256, 192, sdl.WINDOW_SHOWN)
	if err != nil {
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalln(err)
	}
	defer renderer.Destroy()

	texture, err := loadTexture("img/autodrums.png", renderer)
	if err != nil {
		log.Fatalln(err)
	}
	defer texture.Destroy()

	initAudio()

	samples, err := loadSamples(".")
	if err != nil {
		log.Fatalln(err)
	}

	handleEvents(samples, renderer, texture)
}

func loadTexture(file string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	img.Init(img.INIT_PNG)
	surface, err := img.Load(file)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}
	return texture, nil
}
