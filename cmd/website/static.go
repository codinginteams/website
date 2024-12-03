package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ensureDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	return nil
}

func copyFile(sourcePath, destPath string) error {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", sourcePath, err)
	}
	err = os.WriteFile(destPath, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", destPath, err)
	}
	log.Printf("Copied %s to %s\n", sourcePath, destPath)
	return nil
}

func processFile(path string, info os.FileInfo, destDir string) error {
	if info.IsDir() {
		return nil
	}
	destPath := filepath.Join(destDir, info.Name())
	return copyFile(path, destPath)
}

func walkSourceDir(sourceDir, destDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %s: %w", path, err)
		}
		return processFile(path, info, destDir)
	})
}

func copyStaticFiles(sourceDir, destDir string) error {
	if err := ensureDir(destDir); err != nil {
		return err
	}
	return walkSourceDir(sourceDir, destDir)
}
