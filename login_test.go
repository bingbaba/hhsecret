package hhsecret

import (
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	username := os.Getenv("HH_USERNAME")
	password := os.Getenv("HH_PASSWORD")
	if username != "" && password != "" {
		li := NewLoginInfo(username, password)
		ld, err := li.Do()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", ld)
	}
}
