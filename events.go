// events.go
package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const frameDelay = 1000 / 60 // Approx. 16.67 milliseconds per frame for 60 FPS

// handleEvents updated with additional controls for playback and settings.
func handleEvents(samples *Samples, renderer *sdl.Renderer, texture *sdl.Texture) {
	var event sdl.Event
	running := true
	beatCounter := 0
	const maxBeatCounter = 16
	bpm := 120.0
	beatPlaying := true
	randomBeatSkip := true
	randomBeatSilence := true

	for running {
		frameStart := sdl.GetTicks()

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					handleKeyPress(e.Keysym.Sym, samples, &beatPlaying, &bpm, &beatCounter)
				}
			}
		}

		if beatPlaying {
			executeBeat(samples, &beatCounter, maxBeatCounter, bpm, randomBeatSkip, randomBeatSilence, "k   k   Kk  k   ", "  s           s ", " h h hhh  hh h h", "        c       ", "t               ", "  r             ", "    o           ")
		}

		renderFrame(renderer, texture)
		maintainFrameRate(frameStart)
	}
}

func renderFrame(renderer *sdl.Renderer, texture *sdl.Texture) {
	renderer.Clear()
	renderer.Copy(texture, nil, nil) // Update to handle dynamic resizing or interface changes
	renderer.Present()
}

// maintainFrameRate ensures that the time per frame does not exceed the expected frame rate delay.
func maintainFrameRate(startTicks uint32) {
	elapsed := sdl.GetTicks() - startTicks
	if elapsed < frameDelay {
		sdl.Delay(frameDelay - elapsed)
	}
}
