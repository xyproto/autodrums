package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/mix"
)

const sampleRate = 44100

type SampleIndex int

type Samples struct {
	kicks, snares, hihats, crashes, toms, rides, ophats                                          []SampleIndex
	chunks                                                                                       []*mix.Chunk
	currentKick, currentSnare, currentHiHat, currentCrash, currentTom, currentRide, currentOpHat SampleIndex
}

func loadSamples(rootPath string) (*Samples, error) {
	files, err := findFiles(rootPath, ".wav")
	if err != nil {
		return nil, err
	}
	samples := &Samples{}
	for _, file := range files {
		chunk, err := mix.LoadWAV(file)
		if err != nil {
			continue
		}
		index := SampleIndex(len(samples.chunks))
		samples.chunks = append(samples.chunks, chunk)
		classifySample(file, index, samples)
	}
	return samples, nil
}

func classifySample(file string, index SampleIndex, samples *Samples) {
	lowerFile := strings.ToLower(file)
	if strings.Contains(lowerFile, "kick") {
		samples.kicks = append(samples.kicks, index)
	} else if strings.Contains(lowerFile, "snare") {
		samples.snares = append(samples.snares, index)
	} else if strings.Contains(lowerFile, "hihat") {
		samples.hihats = append(samples.hihats, index)
	} else if strings.Contains(lowerFile, "crash") {
		samples.crashes = append(samples.crashes, index)
	} else if strings.Contains(lowerFile, "tom") {
		samples.toms = append(samples.toms, index)
	} else if strings.Contains(lowerFile, "ride") {
		samples.rides = append(samples.rides, index)
	} else if strings.Contains(lowerFile, "ophat") {
		samples.ophats = append(samples.ophats, index)
	}
}

func randomizeSamples(samples *Samples) {
	// Ensuring there are samples to randomize
	if len(samples.kicks) > 0 {
		samples.currentKick = samples.kicks[rand.Intn(len(samples.kicks))]
	}
	if len(samples.snares) > 0 {
		samples.currentSnare = samples.snares[rand.Intn(len(samples.snares))]
	}
	if len(samples.hihats) > 0 {
		samples.currentHiHat = samples.hihats[rand.Intn(len(samples.hihats))]
	}
	if len(samples.crashes) > 0 {
		samples.currentCrash = samples.crashes[rand.Intn(len(samples.crashes))]
	}
	if len(samples.toms) > 0 {
		samples.currentTom = samples.toms[rand.Intn(len(samples.toms))]
	}
	if len(samples.rides) > 0 {
		samples.currentRide = samples.rides[rand.Intn(len(samples.rides))]
	}
	if len(samples.ophats) > 0 {
		samples.currentOpHat = samples.ophats[rand.Intn(len(samples.ophats))]
	}
	fmt.Println("Samples have been randomized.")
}

// playSample plays a drum sample; if randomization is enabled, it selects a random sample.
func playSample(samples *Samples, sampleType []SampleIndex, randomize bool) {
	if len(sampleType) == 0 {
		fmt.Println("No samples available to play.")
		return
	}
	var index SampleIndex
	if randomize {
		index = sampleType[rand.Intn(len(sampleType))]
	} else {
		index = sampleType[0] // Default to the first sample if not randomizing
	}
	if channel := mix.GroupAvailable(-1); channel != -1 {
		if _, err := samples.chunks[index].Play(channel, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to play sample on channel %d: %v\n", channel, err)
		}
	} else {
		fmt.Println("No available channel to play the sample")
	}
}

func playRawSample(data []byte) {
	chunk, err := mix.QuickLoadRAW(&data[0], uint32(len(data))) // Pass pointer to the first element of data and its length
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load raw data: %v\n", err)
		return
	}
	defer chunk.Free()
	channel, err := chunk.Play(-1, 0) // Play the chunk once on the first free unreserved channel
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to play channel: %v\n", err)
	} else {
		fmt.Printf("Playing on channel: %d\n", channel)
	}
}

// toggleSampleRandomization toggles the use of randomized samples.
func toggleSampleRandomization(samples *Samples, enabled *bool) {
	*enabled = !*enabled
	if *enabled {
		fmt.Println("Sample randomization enabled.")
	} else {
		fmt.Println("Sample randomization disabled.")
	}
}
