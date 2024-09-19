package main

import "fmt"

func FindRoomByName(name string) *Room {
	for i := range Rooms {
		if Rooms[i].Name == name {
			return &Rooms[i]
		}
	}
	return nil
}

func IsFileValid() bool {

	seen := make(map[string]bool)

	// Verifies that Rooms are unique
	for _, room := range Rooms {
		if seen[room.Name] {
			fmt.Printf("Duplicate room name: %s\n", room.Name)
			return false
		}
		seen[room.Name] = true
	}

	// Verifies that Start Room and End Room are respectively the first and last given rooms in the file
	if startRoom.Name != Rooms[0].Name {
		fmt.Println("Start room is not the first room in the file.")
		return false
	}
	if endRoom.Name != Rooms[len(Rooms)-1].Name {
		fmt.Println("End room is not the last room in the file.")
		return false
	}
	startRoom = Rooms[0]
	endRoom = Rooms[len(Rooms)-1]
	return true
}
