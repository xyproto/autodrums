.PHONY: clean distclean run

all: autodrums musicradar-drum-samples

musicradar-drum-samples:
	curl -OL 'http://cdn.mos.musicradar.com/audio/samples/musicradar-drum-samples.zip'
	unzip musicradar-drum-samples.zip
	rm musicradar-drum-samples.zip

autodrums.o: main.cpp
	g++ -o $@ -c -std=c++2a -O2 -pipe -fPIC -fno-plt -fstack-protector-strong -Wall -Wshadow -Wpedantic -Wno-parentheses -Wfatal-errors -Wvla -pthread -I/usr/include/SDL2 -D_REENTRANT $<

autodrums: autodrums.o
	g++ -o $@ -Wl,--as-needed $< -lSDL2 -lSDL2_image -lSDL2_mixer -lstdc++fs

run: autodrums musicradar-drum-samples
	./autodrums

clean:
	rm -f autodrums.o *.zip

distclean: clean
	rm -f autodrums musicradar-drum-samples
