<img alt="autodrums" width="250" src="img/keybindings.png">

A randomized drum machine.

Play non-repeating drum beats.

## Requirements

* SDL2
* Download and extract [the drum samples](http://cdn.mos.musicradar.com/audio/samples/musicradar-drum-samples.zip), either manually or by using `make run`.
* g++ that supports C++20.
* `pkg-config`
* `make`

## Build

Tested on Arch Linux and macOS.

* Build with `make`.

## Keybindings

* Press `a` to play a kick drum sound.
* Press `p` to play a snare sound with a tiny bit of delay added.
* Press `w` to play a snare sound.
* Press `r` to play a snare sound.
* Press `d` to play a crash sound.
* Press `s` to play a closed hi-hat sound.
* Press `q` to play a tom sound.
* Press `e` to play a ride sound.
* Press `x` to play an open hi-hat sound.
* Press `o` to output the current sample indices.
* Press `f` to randomize the samples.
* Press `k` to play a beat.
* Press `m` to increase the tempo.
* Press `n` to decrease the tempo.
* Press `y` to use the current settings, don't change the samples.
* Press `i` to toggle "random beat skip".
* Press `j` to toggle "use random beat silence".
* Press `esc` to quit.
* Press `space` to lock the current samples.

Note that playing too many sounds at the same time does not always work.

* [keydrums](https://github.com/xyproto/keydrums) has a few improvements if the goal is to play drums with the keyboard.

## General info

* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;rodseth@gmail.com&gt;
* License: BSD-3
