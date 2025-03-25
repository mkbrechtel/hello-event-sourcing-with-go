package main

import (
	"fmt"
	"hello-event-sourcing-with-go/example"
)

func main() {
	events := example.CreateSampleEvents()
	for _, e := range events {
		fmt.Println(e)
	}
}
