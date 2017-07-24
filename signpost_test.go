package hhsecret

import (
	"fmt"
	"os"
	"testing"
)

var (
	TEST_CONSUMER_KEY    = "lRudaAEghEJGEHkw"
	TEST_CONSUMER_SECRET = "0QNOWUHsWcI9i8UyqBFKUarayqBDUsVnxJrumYHEUl"
)

func TestListSign(t *testing.T) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username != "" && password != "" {
		client := NewClient(username, password, TEST_CONSUMER_KEY, TEST_CONSUMER_SECRET)
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

func TestSign(t *testing.T) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username != "" && password != "" {
		client := NewClient(username, password, TEST_CONSUMER_KEY, TEST_CONSUMER_SECRET)
		err := client.Login()
		if err != nil {
			t.Fatal(err)
		}

		data, err := client.Sign()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%+v\n", data)
	}
}
