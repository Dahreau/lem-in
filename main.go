package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Ant struct {
	Name     string
	Path     []*Room
	Location Room
}

var Ants []Ant

type Room struct {
	Name    string
	IsEmpty bool
	Visited bool
	Links   []*Room
}

var Rooms []Room

var Path []Room

var Paths [][]*Room

var AntCount int

var startRoom *Room

var endRoom *Room

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

	// for _, room := range Rooms {
	// 	PrintRoom(room)
	// }

	FindAllPaths(startRoom, endRoom)
	bestPaths := FindBestPaths(Paths, AntCount)
	for _, path := range bestPaths {
		PrintPath(path)
	}
	AttributesAntsToPaths(bestPaths)
}

func PrintRoom(room Room) {

	fmt.Printf("Room: %s\n", room.Name)
	fmt.Printf("Links: %v\n", room.Links)
	fmt.Printf("Visited: %v\n", room.Visited)
	fmt.Printf("IsEmpty: %v\n", room.IsEmpty)
	fmt.Println()
}

func PrintPath(Path []*Room) {
	for _, room := range Path {
		fmt.Print(room.Name, " -> ")
	}
	fmt.Println()
}

func AttributesAntsToPaths(bestPaths [][]*Room) {
	antIndex := 0
	nbOfAntsInPath := make(map[string]int)
	// Sort paths by length
	sort.Slice(bestPaths, func(i, j int) bool {
		return len(bestPaths[i]) < len(bestPaths[j])
	})
	for antIndex < AntCount {
		if len(bestPaths) > 1 {
			for i := 1; i < len(bestPaths); i++ {
				if antIndex >= AntCount {
					break
				}
				if len(bestPaths[i-1])+nbOfAntsInPath[PathToStr(bestPaths[i-1])] <= len(bestPaths[i])+nbOfAntsInPath[PathToStr(bestPaths[i])] {
					Ants[antIndex].Path = bestPaths[i-1]
					Ants[antIndex].Location = *bestPaths[i-1][0]
					nbOfAntsInPath[PathToStr(bestPaths[i-1])]++
					antIndex++
				} else {
					Ants[antIndex].Path = bestPaths[i]
					Ants[antIndex].Location = *bestPaths[i][0]
					nbOfAntsInPath[PathToStr(bestPaths[i])]++
					antIndex++

				}
			}
		} else {
			Ants[antIndex].Path = bestPaths[0]
			Ants[antIndex].Location = *bestPaths[0][0]
			nbOfAntsInPath[PathToStr(bestPaths[0])]++
			antIndex++
		}
	}
	for pathStr, nbOfAnts := range nbOfAntsInPath {
		fmt.Printf("Path: %v, Number of ants: %d\n", pathStr, nbOfAnts)
	}
}

func PrintAnt(ant Ant) {
	fmt.Printf("Ant: %s\n", ant.Name)
	fmt.Printf("Path: %v\n", ant.Path)
	fmt.Printf("Location: %v\n", ant.Location)
	fmt.Println()
}

func PathToStr(path []*Room) string {
	str := ""
	for _, room := range path {
		str += room.Name + "-"
	}
	return str[:len(str)-1]
}
