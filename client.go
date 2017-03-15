package scaledrone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api2.scaledrone.com"

// Scaledrone is the main struct for the API
type Client struct {
	ChannelID string
	SecretKey string
	Bearer    string
}

// NewBasicAuthClient returns a new Scaledrone client that uses Basic Authentication for authentication
func NewBasicAuthClient(channelID, secretKey string) *Client {
	return &Client{
		ChannelID: channelID,
		SecretKey: secretKey,
	}
}

// NewBearerClient returns a new Scaledrone client that uses Bearer token for authentication
func NewBearerClient(channelID, bearer string) *Client {
	return &Client{
		ChannelID: channelID,
		Bearer:    bearer,
	}
}

// Publish sends a message to a single room
func (s *Client) Publish(message []byte, room string) error {
	url := fmt.Sprintf(baseURL+"/%s/%s/publish", s.ChannelID, room)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

// PublishToRooms sends a message to an array of rooms
func (s *Client) PublishToRooms(message []byte, rooms []string) error {
	url := fmt.Sprintf(baseURL+"/%s/publish/rooms", s.ChannelID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	query := req.URL.Query()
	for _, room := range rooms {
		query.Add("r", room)
	}
	req.URL.RawQuery = query.Encode()
	_, err = s.doRequest(req)
	return err
}

type usersCount struct {
	Count int `json:"users_count"`
}

// UsersCount returns how many users have connected to the channel
func (s *Client) UsersCount() (int, error) {
	url := fmt.Sprintf(baseURL+"/%s/stats", s.ChannelID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return -1, err
	}
	var data usersCount
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return -1, err
	}
	return data.Count, nil
}

// UsersInRooms returns an array of members who have joined single or multiple rooms
func (s *Client) UsersInRooms() ([]string, error) {
	url := fmt.Sprintf(baseURL+"/%s/members", s.ChannelID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ActiveRooms returns an array of rooms that are not empty
func (s *Client) ActiveRooms() ([]string, error) {
	url := fmt.Sprintf(baseURL+"/%s/rooms", s.ChannelID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UsersInRoom returns an array of members subscribed to a room
func (s *Client) UsersInRoom(room string) ([]string, error) {
	url := fmt.Sprintf(baseURL+"/%s/%s/members", s.ChannelID, room)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RoomMembers returns a room name to members array map of all non empty rooms
func (s *Client) RoomMembers() (map[string][]string, error) {
	url := fmt.Sprintf(baseURL+"/%s/room-members", s.ChannelID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data map[string][]string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Client) setAuth(req *http.Request) {
	if s.SecretKey != "" {
		req.SetBasicAuth(s.ChannelID, s.SecretKey)
	} else {
		req.Header.Set("Authorization", "Bearer "+s.Bearer)
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	s.setAuth(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
