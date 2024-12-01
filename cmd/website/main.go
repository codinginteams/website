package main

import (
	"log"

	"github.com/codinginteams/simple-go-ssg/pkg/generator"
)

func main() {
	err := generator.Run("content", "build", "templates")
	if err != nil {
		log.Fatalf("Failed to generate site: %v", err)
	}
	log.Println("Site generated successfully!")
}
