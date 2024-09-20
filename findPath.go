package main

import (
	"fmt"
	"os"
)

func FindAllPaths(start *Room, end *Room) {
	// File pour stocker les chemins partiels
	queue := [][]*Room{{start}}
	// var paths [][]*Room

	// Parcours de tous les chemins
	for len(queue) > 0 {
		// Récupère le premier chemin
		path := queue[0]
		queue = queue[1:]

		// Dernière salle du chemin
		lastRoom := path[len(path)-1]

		// PrintRoom(*lastRoom)
		// fmt.Println("queue :", queue)

		// Si on atteint la salle de fin, on ajoute ce chemin
		if lastRoom.Name == end.Name {
			Paths = append(Paths, path)
			continue
		}

		// Parcourt les voisins de la dernière salle
		for _, neighbor := range lastRoom.Links {
			if !isVisited(neighbor, path) {
				// Crée un nouveau chemin en ajoutant la salle voisine
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	if len(Paths) == 0 {
		fmt.Println("Invalid data format")
		os.Exit(0)
	}
}

// Fonction pour vérifier si une salle a déjà été visitée dans un chemin
func isVisited(room *Room, path []*Room) bool {
	for _, r := range path {
		if r.Name == room.Name {
			return true
		}
	}
	return false
}

// I would like a function that returns the best paths (the shortest ones) but the maximum disjoints paths in the list of paths depending on the number of ants that will be sent.
func FindBestPaths(paths [][]*Room, numAnts int) [][]*Room {
	BestPaths := [][]*Room{}
	bestScore := float64(^uint(0) >> 1) // Initialize to max float value

	for i := 0; i < len(paths); i++ {
		disjointPaths := [][]*Room{paths[i]}
		for j, room := range paths[i] {
			if j != 0 && j != len(paths[i])-1 {
				room.Visited = true
			}
		}

		for j := 0; j < len(paths); j++ {
			if j != i {
				disjoint := true
				for k, room := range paths[j] {
					if k != 0 && k != len(paths[j])-1 && room.Visited {
						disjoint = false
						break
					}
				}
				if disjoint {
					disjointPaths = append(disjointPaths, paths[j])
					for _, room := range paths[j] {
						if room != paths[j][0] && room != paths[j][len(paths[j])-1] {
							room.Visited = true
						}
					}
				}
			}
		}

		for _, path := range disjointPaths {
			for _, room := range path {
				room.Visited = false
			}
		}

		totalLength := 0
		for _, path := range disjointPaths {
			totalLength += len(path)
		}

		// Calculate score: prioritize number of paths, then total length
		// Use a weighted formula: score = totalLength + weight * (maxPaths - len(disjointPaths))
		weight := 8.0 // Adjust this weight as needed
		score := float64(totalLength) + weight*float64(len(paths)-len(disjointPaths))

		if score < bestScore {
			bestScore = score
			BestPaths = disjointPaths
		}
	}

	if numAnts == 1 {
		BestPaths = [][]*Room{FindShortestPath()}
	}

	return BestPaths
}

// I would like a function that returns the shortest path in the list of paths.
func FindShortestPath() []*Room {
	ShortestPath := Paths[0]
	for i := 0; i < len(Paths); i++ {
		if len(Paths[i]) < len(ShortestPath) {
			PrintPath(Paths[i])
			ShortestPath = Paths[i]
		}
	}
	return ShortestPath
}
