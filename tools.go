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
