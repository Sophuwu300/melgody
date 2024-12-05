# MP3 Player

This is a simple MP3 player written in Go. It can play MP3 files from the current directory or from a list of files provided as command-line arguments. The player supports shuffling and skipping songs.

## Features

- Play MP3 files from the current directory or from a list of files provided as arguments
- Shuffle the playlist
- Display the current song and the next few songs in the queue
- Write "skip" to skip the current song

## Limitations
- No volume control
- Exit the player by pressing `Ctrl`+`C`
- No pause or resume functionality
- Only supports MP3 files

## Compilation
Clone the repository and cd into it. Then
run the following command to compile the program and install the dependencies:
```sh
go build -ldflags="-w -s" -trimpath -o build/melgody .
```
## Installation
To install the program into you system bin, run:
```sh
sudo install ./build/melgody /usr/local/bin/melgody
```

## Usage
To play all MP3 files in the current directory, run:
```sh
melgody
```
To play a list of MP3 files, run:
```sh
melgody file1.mp3 file2.mp3 ... 
```

## License
MIT
