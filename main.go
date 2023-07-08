package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	curdir string
	done   = make(chan bool)
)

func getallfiles() []string {
	var files []string
	f, err := os.Open(curdir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files, err = f.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Close()
	var songs []string
	for _, file := range files {
		if file[len(file)-4:] == ".mp3" {
			songs = append(songs, file)
		}
	}
	return songs
}

func getargsongs() []string {
	var songs []string
	for _, song := range os.Args[1:] {
		info, err := os.Stat(song)
		if err == nil && len(song) > 4 && song[len(song)-4:] == ".mp3" && !info.IsDir() && info.Mode().IsRegular() {
			songs = append(songs, song)
		} else {
			fmt.Println(song, "is not a valid mp3 file.")
		}
	}
	return songs
}

func showqueue(songs []string) {
	if len(songs) == 0 {
		fmt.Println("\nEmpty queue.")
		os.Exit(0)
	}
	var tmpstr string
	tmpstr = strings.Split(songs[0], "/")[len(strings.Split(songs[0], "/"))-1]
	tmpstr = tmpstr[:len(tmpstr)-4]
	fmt.Printf("\nNow Playing:\n%s\n\n", tmpstr)
	if len(songs) == 1 {
		return
	}
	fmt.Println("Up Next:")
	var loopmax int = 5
	if len(songs) < loopmax {
		loopmax = len(songs)
	}
	for i := 1; i < loopmax; i++ {
		tmpstr = strings.Split(songs[i], "/")[len(strings.Split(songs[i], "/"))-1]
		tmpstr = tmpstr[:len(tmpstr)-4]
		fmt.Printf("%d: %s\n", i, tmpstr)
	}
	fmt.Println("")
}

func shuffle(songs *[]string) {
	var tmpstr string
	var tmpint int
	for i := 0; i < len(*songs); i++ {
		tmpint = rand.Intn(len(*songs))
		tmpstr = (*songs)[i]
		(*songs)[i] = (*songs)[tmpint]
		(*songs)[tmpint] = tmpstr
	}
}

func playlist(songs []string) {
	for i := range songs {
		showqueue(songs[i:])
		play(songs[i])
	}
}

func play(song string) {
	file, _ := os.Open(song)
	streamer, format, _ := mp3.Decode(file)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done = make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	streamer.Close()
	file.Close()
}

func skipsong() {
	var input string
	for {
		time.Sleep(200 * time.Millisecond)
		fmt.Print("Enter 'skip' to skip the current song: ")
		fmt.Scanln(&input)
		if input == "skip" {
			speaker.Clear()
			done <- true
		}
	}
}

func main() {
	tmpstr, err := os.Getwd()
	if err != nil {
		fmt.Println("An error occured while getting the directory.")
		os.Exit(1)
	}
	curdir = tmpstr
	var songs []string
	if len(os.Args) <= 1 {
		songs = getallfiles()
		shuffle(&songs)
	} else {
		songs = getargsongs()
	}
	if len(songs) == 0 {
		fmt.Println("No songs found.")
		os.Exit(0)
	}
	go playlist(songs)
	skipsong()
}