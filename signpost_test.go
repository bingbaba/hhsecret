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
	username := os.Getenv("HH_USERNAME")
	password := os.Getenv("HH_PASSWORD")
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

		if len(data.Signs) > 0 {
			fmt.Printf("%+v\n", data.Signs[0])
		}

		//lat := fmt.Sprintf("%0.6f", 36.128+random.Float64()/1000)
		//lng := fmt.Sprintf("%0.6f", 120.418+random.Float64()/1000)
		//configId, err := client.signConfigId(lat, lng)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//fmt.Printf("configId: %s\n", configId)
	}
}

func TestSign(t *testing.T) {
	username := os.Getenv("HH_USERNAME")
	password := os.Getenv("HH_PASSWORD")
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
