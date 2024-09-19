package main

import "fmt"

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

		PrintRoom(*lastRoom)
		fmt.Println("queue :", queue)

		// Si on atteint la salle de fin, on ajoute ce chemin
		if lastRoom.Name == end.Name {
			fmt.Println("lalalaalaa")
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
