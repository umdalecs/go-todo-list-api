package main

import (
	"log"
)

func main() {
	s := NewAPIServer(":8080", nil)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
