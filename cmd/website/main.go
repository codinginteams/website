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

	err = copyStaticFiles("static/css", "build/static/css")
	if err != nil {
		log.Fatalf("Failed to copy static files: %v", err)
	}

	log.Println("Site generated and static files copied successfully!")
}
