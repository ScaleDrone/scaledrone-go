# Scaledrone Go Client

scaledrone-go is a Go client package for accessing the Scaledrone REST API.
![gopher](https://raw.githubusercontent.com/scaledrone/scaledrone-go/master/gopher.png)

# Example

```go
package main

import (
	"fmt"

	"github.com/ScaleDrone/scaledrone-go"
)

func main() {
	// New Scaledrone Client
	channelID := "channel-id"
	secretKey := "secret-key"
	client := scaledrone.NewBasicAuthClient(channelID, secretKey)

	// Publishing to a single room
	_ = client.Publish([]byte("Hello Go"), "golang")

	// Publishing to a multiple rooms
	_ = client.PublishToRooms([]byte("Hello Go"), []string{"golang", "gopher"})

	// Getting the number of users connected to the channel
	count, _ := client.UsersCount()
	fmt.Println(count)

	// Getting the array of users connected to rooms
	users, _ := client.UsersInRooms()
	fmt.Println(users)

	// Getting the array of rooms that are not empty
	rooms, _ := client.ActiveRooms()
	fmt.Println(rooms)

	// Getting the array of users in a room
	users, _ = client.UsersInRoom("empty-room")
	fmt.Println(users)

	// Getting the map of rooms to members
	roomsMap, _ := client.RoomMembers()
	fmt.Println(roomsMap)
}
```
