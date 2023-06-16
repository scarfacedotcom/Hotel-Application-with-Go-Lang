package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "Scar Face Hotel"
	const totalRooms = 134
	const firstRoomNumber = 110
	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := (float64(roomsOccupied) / float64(totalRooms) * 100)
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "low"
	}
	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms Available", roomsAvailable)
	fmt.Println("Occupancy Level:", occupancyLevel)
	fmt.Println("Occupancy Rate:%0.2f %%n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights")
		}
	} else {
		fmt.Println("No rooms available")
	}
}
