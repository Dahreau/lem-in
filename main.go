package main

import (
	"fmt"
	"os"
	"time"
)

type Ant struct {
	Name     string
	Path     []*Room
	Location Room
}

var Ants []Ant

type Room struct {
	Name    string
	Visited bool
	Links   []*Room
}
type PathStruct struct {
	Path      []*Room
	AntsCount int
	Ants      []Ant
}

var DisjointPaths []PathStruct

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
		fmt.Println("ERROR : invalid data format.")
		os.Exit(0)
	}

	// fmt.Println(startRoom.Name)
	// fmt.Println(endRoom.Name)
	// fmt.Println(Ants)
	// fmt.Println(Rooms[0])

	// for _, room := range Rooms {
	// 	PrintRoom(room)
	// }
	timeExec := time.Now()
	FindAllPaths(startRoom, endRoom)
	FindBestPaths(Paths, AntCount)
	AttributesAntsToPaths()
	fmt.Println(fileContent)
	MoveAnts()
	fmt.Printf("Execution time : %v\n", time.Since(timeExec))
}

func PrintRoom(room Room) {

	fmt.Printf("Room: %s\n", room.Name)
	fmt.Printf("Links: %v\n", room.Links)
	fmt.Printf("Visited: %v\n", room.Visited)
	fmt.Println()
}

func PrintPath(Path []*Room) {
	for _, room := range Path {
		fmt.Print(room.Name, " -> ")
	}
	fmt.Println()
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

func AttributesAntsToPaths() {
	antIndex := 0
	for antIndex < AntCount {
		minIndex := FindMinAntsInPathIndex(DisjointPaths)
		Ants[antIndex].Path = DisjointPaths[minIndex].Path
		Ants[antIndex].Location = *DisjointPaths[minIndex].Path[0]
		DisjointPaths[minIndex].AntsCount++
		DisjointPaths[minIndex].Ants = append(DisjointPaths[minIndex].Ants, Ants[antIndex])
		antIndex++
	}
}

func FindMinAntsInPathIndex(disjointPaths []PathStruct) int {
	minIndex := 0
	min := len(disjointPaths[0].Path)
	for i, path := range disjointPaths {
		if len(path.Path)+path.AntsCount < min+disjointPaths[minIndex].AntsCount {
			min = len(path.Path)
			minIndex = i
		}
	}
	return minIndex
}
func FindMinLenPathIndex(disjointPaths []PathStruct) int {
	minIndex := 0
	min := len(disjointPaths[0].Path)
	for i, path := range disjointPaths {
		if len(path.Path) < min {
			min = len(path.Path)
			minIndex = i
		}
	}
	return minIndex
}

func MoveAnts() {
	nbOfIteration := CalculateNbOfIteration()
	result := make([]string, nbOfIteration)
	for _, path := range DisjointPaths {
		for i := 0; i < path.AntsCount; i++ {
			j := 0
			for path.Ants[i].Location.Name != endRoom.Name {
				if FindLocationIndex(path.Ants[i]) < len(path.Ants[i].Path) {
					path.Ants[i].Location = *path.Ants[i].Path[FindLocationIndex(path.Ants[i])+1]
					result[i+j] += path.Ants[i].Name + "-" + path.Ants[i].Location.Name + " "
				}
				j++
			}
		}

	}

	for _, res := range result {
		// pause()
		fmt.Println(res)
	}
	fmt.Printf("\nNumber of iteration(s) : %d\n", nbOfIteration)
}
func FindLocationIndex(ant Ant) int {
	for i, room := range ant.Path {
		if room.Name == ant.Location.Name {
			return i
		}
	}
	return -1
}
func CalculateNbOfIteration() int {
	MinnbOfIteration := 100000000000000
	for _, path := range DisjointPaths {
		if len(path.Path)+path.AntsCount <= MinnbOfIteration {
			MinnbOfIteration = len(path.Path) - 2 + path.AntsCount
		}
	}
	return MinnbOfIteration
}

// func pause() {
// 	fmt.Scan()
// }
