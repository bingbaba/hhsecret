package hhsecret

import (
	"fmt"
	"os"
	"testing"
)

func TestListSign(t *testing.T) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username != "" && password != "" {
		client := NewClient(username, password, "lRudaAEghEJGEHkw", "0QNOWUHsWcI9i8UyqBFKUarayqBDUsVnxJrumYHEUl")
		err := client.Login()
		if err != nil {
			t.Fatal(err)
		}

		data, err := client.ListSignPost()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%+v\n", data)
	}
}
