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
	if AntCount < 1 {
		return false
	}
	// Verifies that Rooms are unique
	for _, room := range Rooms {
		if seen[room.Name] {
			fmt.Printf("Duplicate room name: %s\n", room.Name)
			return false
		}
		seen[room.Name] = true
	}
	return true
}
