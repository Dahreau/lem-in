package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ant struct {
	Name     string
	Location Room
}

var Ants []Ant

type Room struct {
	Name    string
	IsEmpty bool
	Visited bool
	Links   []string
}

var Rooms []Room

var Path []Room

var Paths [][]Room

var startRoom Room

var endRoom Room

func main() {

	fileContent := ReadFile()

	// fmt.Println(fileContent)

	DecodeFile(fileContent)

	if !IsFileValid() {
		log.Fatal("Invalid file format.")
		os.Exit(0)
	}

	// fmt.Println(startRoom.Name)
	// fmt.Println(endRoom.Name)
	// fmt.Println(Ants)
	// fmt.Println(Rooms[0])

	for _, room := range Rooms {
		PrintRoom(room)
	}
}

func ReadFile() string {

	if len(os.Args) != 2 {
		log.Fatal("Please provide a filename as an argument.")
		os.Exit(0)
	}
	filename := os.Args[1]

	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatalf("Error opening file: %v", err)
	// }

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		os.Exit(0)
	}

	return string(fileContent)
}

func DecodeFile(fileContent string) {
	startFound := false
	endFound := false

	splittedContent := strings.Split(fileContent, "\n")

	antCount, err := strconv.Atoi(splittedContent[0])
	if err != nil {
		log.Fatalf("Error parsing ant count: %v", err)
		os.Exit(0)
	}

	for i := 1; i <= antCount; i++ {
		ant := Ant{Name: fmt.Sprintf("L%d", i)}
		Ants = append(Ants, ant)
	}

	for _, line := range splittedContent {
		elements := strings.Split(line, " ")
		if startFound {
			startRoom = Room{Name: elements[0]}
			startFound = false
		}
		if line == "##start" {
			startFound = true
		}
		if endFound {
			endRoom = Room{Name: elements[0]}
			endFound = false
		}
		if line == "##end" {
			endFound = true
		}
		if len(elements) == 3 {
			room := Room{Name: elements[0]}
			Rooms = append(Rooms, room)
		}
		if strings.Contains(line, "-") {
			links := strings.Split(line, "-")
			room1 := findRoomByName(links[0])
			room2 := findRoomByName(links[1])
			// fmt.Println(room1, room2)
			room1.Links = append(room1.Links, links[1])
			room2.Links = append(room2.Links, links[0])
		}

	}
}

func findRoomByName(name string) *Room {
	for i := range Rooms {
		if Rooms[i].Name == name {
			return &Rooms[i]
		}
	}
	return nil
}

func IsFileValid() bool {

	seen := make(map[string]bool)

	for _, room := range Rooms {
		if seen[room.Name] {
			fmt.Printf("Duplicate room name: %s\n", room.Name)
			return false
		}
		seen[room.Name] = true
	}

	if startRoom.Name != Rooms[0].Name {
		fmt.Println("Start room is not the first room in the file.")
		return false
	}
	if endRoom.Name != Rooms[len(Rooms)-1].Name {
		fmt.Println("End room is the first room in the file.")
		return false
	}
	return true
}

func PrintRoom(room Room) {

	fmt.Printf("Room: %s\n", room.Name)
	fmt.Printf("Links: %v\n", room.Links)
	fmt.Printf("Visited: %v\n", room.Visited)
	fmt.Printf("IsEmpty: %v\n", room.IsEmpty)
	fmt.Println()
}
