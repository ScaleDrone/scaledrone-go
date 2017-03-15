package scaledrone

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getClient() *Client {
	return NewBasicAuthClient(os.Getenv("CHANNEL_ID"), os.Getenv("SECRET_KEY"))
}

func TestPublish(t *testing.T) {
	err := getClient().Publish([]byte("Hello Go"), "golang")
	assert.Nil(t, err)
}

func TestPublishToRooms(t *testing.T) {
	err := getClient().PublishToRooms([]byte("Hello Go"), []string{"golang", "gopher"})
	assert.Nil(t, err)
}

func TestUsersCount(t *testing.T) {
	count, err := getClient().UsersCount()
	assert.Nil(t, err)
	assert.True(t, count >= 0)
}

func TestUsersInRooms(t *testing.T) {
	_, err := getClient().UsersInRooms()
	assert.Nil(t, err)
}

func TestActiveRooms(t *testing.T) {
	_, err := getClient().ActiveRooms()
	assert.Nil(t, err)
}

func TestUsersInRoom(t *testing.T) {
	users, err := getClient().UsersInRoom("empty-room")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}

func TestRoomMembers(t *testing.T) {
	_, err := getClient().RoomMembers()
	assert.Nil(t, err)
}
