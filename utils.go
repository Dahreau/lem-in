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
			room1 := FindRoomByName(links[0])
			room2 := FindRoomByName(links[1])
			// fmt.Println(room1, room2)
			room1.Links = append(room1.Links, links[1])
			room2.Links = append(room2.Links, links[0])
		}

	}
}
