# Scaledrone Go Client

scaledrone-go is a Go client package for accessing the Scaledrone REST API.

# Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/ScaleDrone/scaledrone-go"
)

func main() {
	// New Scaledrone Client
	channelId := "channel-id"
	secretKey := "secret-key"
	client := scaledrone.NewBasicAuthClient(appKey, secretKey)

	// Publishing to a single room
	_ = client.Publish([]byte("Hello Go"), "golang")

	// Publishing to a multiple rooms
	_ = client.PublishToRooms([]byte("Hello Go"), []string{"golang", "gopher"})

	// Getting the number of users connected to the channel
	count, _ := client.UsersCount()

	// Getting the array of users connected to rooms
	users, _ := client.UsersInRooms()

	// Getting the array of rooms that are not empty
	rooms, _ := client.ActiveRooms()

	// Getting the array of users in a room
	users, _ = client.UsersInRoom("empty-room")

	// Getting the map of rooms to members
	roomsMap, _ := client.RoomMembers()
}
```
