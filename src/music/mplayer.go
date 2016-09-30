package main

import (
	"bufio"
	"fmt"
	"library"
	"os"
	"play"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{string(strconv.Itoa(id)), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: library add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.Remove(3) // 有问题
		} else {
			fmt.Println("USAGE: library remove <name>")
		}
	default:
		fmt.Println("Unrecongnized library command:", tokens[1])
	}
}

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	play.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
    Enter following commands to control the player:
    lib add <name><artist><source><type>
    lib remove <name>
    play <name>
  `)

	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command->")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unrecongnized command:", tokens[0])
		}
	}
}

/*
go build library

go build play

go build music



go test library



go install library

go install play

go install music
*/
