package main

import (
	"fmt"
	"math/rand"
	"os"
)

var (
	curdir string
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
	fmt.Println("Up Next:")
	var loopmax int = 5
	if len(songs) < loopmax {
		loopmax = len(songs)
	}
	for i := 0; i < loopmax; i++ {
		fmt.Printf("%d: %s\n", i+1, songs[i])
	}
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

func main() {
	tmpstr, err := os.Getwd()
	if err != nil {
		fmt.Println("An error occured while getting the current directory.")
		os.Exit(1)
	}
	curdir = tmpstr
	var songs []string
	if len(os.Args) <= 1 {
		songs = getallfiles()
	} else {
		songs = getargsongs()
	}
	shuffle(&songs)
	if len(songs) == 0 {
		fmt.Println("No songs found.")
		os.Exit(0)
	}
	showqueue(songs)
}