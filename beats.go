package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func executeBeat(samples *Samples, beatCounter *int, maxBeatCounter int, bpm float64, randomBeatSkip bool, randomBeatSilence bool, kPat, sPat, hPat, cPat, tPat, rPat, oPat string) {
	if !randomBeatSilence || rand.Float32() > 0.1 { // 10% chance to silence the beat
		index := *beatCounter % len(kPat)
		playPattern(samples, rune(kPat[index]), samples.kicks, samples.currentKick, false)
		playPattern(samples, rune(sPat[index]), samples.snares, samples.currentSnare, false)
		playPattern(samples, rune(hPat[index]), samples.hihats, samples.currentHiHat, false)
		playPattern(samples, rune(cPat[index]), samples.crashes, samples.currentCrash, false)
		playPattern(samples, rune(tPat[index]), samples.toms, samples.currentTom, false)
		playPattern(samples, rune(rPat[index]), samples.rides, samples.currentRide, false)
		playPattern(samples, rune(oPat[index]), samples.ophats, samples.currentOpHat, false)

		*beatCounter++
		if *beatCounter >= maxBeatCounter {
			*beatCounter = 0
		}
	}

	delay := 60000 / bpm // Calculate delay per beat based on BPM
	sdl.Delay(uint32(delay))
}

func playPattern(samples *Samples, char rune, sampleIndexes []SampleIndex, currentIndex SampleIndex, randomize bool) {
	if char != '-' {
		playSample(samples, sampleIndexes, randomize) // Correct usage
	}
}
