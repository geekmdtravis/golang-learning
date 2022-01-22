package main

import (
	"encoding/json"
	"fmt"
)

type JSONMarshaler interface {
	MarshalToJSON() ([]byte, error)
}

type JSONUnmarshaler interface {
	FromJSON([]byte) error
}

type Note struct {
	Id	int `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (n Note) DisplayToConsole() {
	fmt.Printf("ID: %d\n%s\n=============\n%s\n", n.Id, n.Title, n.Body)
}

func (n Note) MarshalToJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Note) FromJSON(s string) error {
	return json.Unmarshal([]byte(s), n)
}

func NoteFromJSON(s string) (Note, error) {
	var n Note
	err := n.FromJSON(s)
	return n, err
}


func main() {
	json := `{"id":1,"title":"test title","body":"the bodddyyyyy"}`

	n, err := NoteFromJSON(json); if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	defer n.DisplayToConsole()
	defer fmt.Println("Where does this show up?")
	fmt.Println("Let's see what that looks like:")
}