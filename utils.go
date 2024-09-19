package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile() string {

	if len(os.Args) != 2 {
		log.Fatal("Please provide a filename as an argument.")
		os.Exit(0)
	}
	filename := os.Args[1]

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

	splittedContent := strings.Split(fileContent, "\n") // Split the entire file by lines

	// The first line is the amount of ant
	antCount, err := strconv.Atoi(splittedContent[0])
	if err != nil {
		log.Fatalf("Error parsing ant count: %v", err)
		os.Exit(0)
	}

	// Name each ant in ascending order starting from 1
	for i := 1; i <= antCount; i++ {
		ant := Ant{Name: fmt.Sprintf("L%d", i)}
		Ants = append(Ants, ant)
	}

	// Loop through all the lines
	for _, line := range splittedContent {
		elements := strings.Split(line, " ") // Split each line by element (separated by spaces)
		// If the program finds "##start" then  the next line is the Start Room
		if startFound {
			startRoom = Room{Name: elements[0]}
			startFound = false
		}
		if line == "##start" {
			startFound = true
		}
		// If the program finds "##end" then  the next line is the End Room
		if endFound {
			endRoom = Room{Name: elements[0]}
			endFound = false
		}
		if line == "##end" {
			endFound = true
		}
		// If the line contains a room name and is not "##start" or "##end" then it's a room
		if len(elements) == 3 {
			room := Room{Name: elements[0]}
			Rooms = append(Rooms, room)
		}
		// If the line contains a link between two rooms then add them to their respective rooms' Links slice
		if strings.Contains(line, "-") {
			links := strings.Split(line, "-")
			room1 := FindRoomByName(links[0])
			room2 := FindRoomByName(links[1])
			room1.Links = append(room1.Links, room2)
			room2.Links = append(room2.Links, room1)
		}

	}
}
