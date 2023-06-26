.PHONY: clean distclean run

all: autodrums musicradar-drum-samples

musicradar-drum-samples:
	curl -OL 'http://cdn.mos.musicradar.com/audio/samples/musicradar-drum-samples.zip'
	unzip musicradar-drum-samples.zip
	rm musicradar-drum-samples.zip

autodrums.o: main.cpp
	g++ -o $@ -c -std=c++2a -O2 -pipe -fPIC -fstack-protector-strong -Wall -Wshadow -Wpedantic -Wno-parentheses -Wfatal-errors -Wvla -pthread `pkg-config --cflags sdl2` $<

autodrums: autodrums.o
	g++ -o $@ $< `pkg-config --libs sdl2` -lSDL2_image -lSDL2_mixer

run: autodrums musicradar-drum-samples
	./autodrums

clean:
	rm -f autodrums.o *.zip

distclean: clean
	rm -f autodrums musicradar-drum-samples
