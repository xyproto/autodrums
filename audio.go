package main

import (
	"fmt"
	"math"
	"os"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	maxChannels = 32
)

func initAudio() error {
	err := mix.OpenAudio(sampleRate, mix.DEFAULT_FORMAT, 2, 512)
	if err != nil {
		return fmt.Errorf("failed to open audio: %v", err)
	}
	mix.AllocateChannels(maxChannels)
	return nil
}

func ADSREnvelope(sampleIdx int, totalSamples int, attackPct float64, decayPct float64, sustainLevel float64, releasePct float64) float64 {
	attackEnd := float64(totalSamples) * attackPct
	decayEnd := float64(totalSamples) * (attackPct + decayPct)
	releaseStart := float64(totalSamples) * (1.0 - releasePct)

	if float64(sampleIdx) < attackEnd {
		return float64(sampleIdx) / attackEnd
	} else if float64(sampleIdx) < decayEnd {
		return 1.0 + (sustainLevel-1.0)*(float64(sampleIdx)-attackEnd)/(decayEnd-attackEnd)
	} else if float64(sampleIdx) < releaseStart {
		return sustainLevel
	} else if float64(sampleIdx) < float64(totalSamples) {
		return sustainLevel * (1.0 - (float64(sampleIdx)-releaseStart)/(float64(totalSamples)-releaseStart))
	}
	return 0.0
}

func generateSawtoothWave(freq int, durationMs int, attackPct float64, decayPct float64, sustainLevel float64, releasePct float64) []byte {
	totalSamples := sampleRate * durationMs / 1000
	wave := make([]byte, totalSamples*2)
	for i := 0; i < totalSamples; i++ {
		envelope := ADSREnvelope(i, totalSamples, attackPct, decayPct, sustainLevel, releasePct)
		t := float64(i) / float64(sampleRate)
		sample := envelope * 32767.0 * 2.0 * (t*float64(freq) - math.Floor(t*float64(freq)+0.5))
		intVal := int16(sample)
		wave[2*i] = byte(intVal & 0xff)
		wave[2*i+1] = byte((intVal >> 8) & 0xff)
	}
	return wave
}

func generateKickDrum(freq, durationMs int) []byte {
	totalSamples := sampleRate * durationMs / 1000
	wave := make([]byte, totalSamples*2)
	amplitude := 32767.0
	frequencyStart := 120.0
	frequencyEnd := 60.0
	envelopeStrength := 0.3
	for i := 0; i < totalSamples; i++ {
		time := float64(i) / float64(sampleRate)
		frequency := frequencyStart + (frequencyEnd-frequencyStart)*float64(i)/float64(totalSamples)
		envelope := amplitude * math.Exp(-envelopeStrength*time)
		sample := envelope * math.Sin(2*math.Pi*frequency*time)
		intVal := int16(sample)
		wave[2*i] = byte(intVal & 0xff)
		wave[2*i+1] = byte((intVal >> 8) & 0xff)
	}
	return wave
}

func playSnareWithDelay(samples *Samples, delayMs int, decayFactor float64) {
	channelCount := 4 // Assuming you want the reverb to affect 4 channels
	if len(samples.snares) > 0 {
		snareSample := samples.chunks[samples.snares[0]]
		for i := 0; i < channelCount; i++ {
			channel := mix.GroupAvailable(-1)
			if channel == -1 {
				fmt.Println("No available channel to play the sample")
				return
			}
			mix.Volume(channel, int(128*math.Pow(decayFactor, float64(i))))
			if _, err := snareSample.Play(channel, 0); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to play snare on channel %d: %v\n", channel, err)
				return
			}
			sdl.Delay(uint32(delayMs))
			if i == channelCount-1 { // Last iteration
				mix.HaltChannel(channel)
			}
		}
	}
}
