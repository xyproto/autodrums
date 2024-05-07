package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyPress(key sdl.Keycode, samples *Samples, beatPlaying *bool, bpm *float64, beatCounter *int) {
	switch key {
	case sdl.K_SPACE:
		*beatPlaying = !*beatPlaying // Toggle beat playing
	case sdl.K_ESCAPE:
		os.Exit(0) // Exit the program
	case sdl.K_UP:
		*bpm += 10 // Increase BPM
	case sdl.K_DOWN:
		*bpm -= 10 // Decrease BPM
	case sdl.K_a:
		playSample(samples, samples.kicks, false) // Play the first kick sample
	case sdl.K_s:
		playSnareWithDelay(samples, 100, 0.75) // Play snare with delay
	case sdl.K_d:
		playSample(samples, samples.crashes, false) // Play the first crash sample
	case sdl.K_f:
		playSample(samples, samples.hihats, false) // Play the first hihat sample
	case sdl.K_q:
		playSample(samples, samples.toms, false) // Play the first tom sample
	case sdl.K_w:
		playSample(samples, samples.rides, false) // Play the first ride sample
	case sdl.K_e:
		playSample(samples, samples.ophats, false) // Play the first open hat sample
	case sdl.K_r:
		randomizeSamples(samples) // Re-randomize all samples
	case sdl.K_v:
		playRawSample(generateKickDrum(120, 500)) // Play a generated kick drum sound
	case sdl.K_b:
		playRawSample(generateSawtoothWave(440, 1000, 0.1, 0.1, 0.6, 0.1)) // Play a generated sawtooth wave
	default:
		fmt.Printf("Key %v pressed, no associated action\n", key)
	}
}
