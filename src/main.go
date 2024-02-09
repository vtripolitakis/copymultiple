package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// CopyFile copies a file from source to destination folder.
//
// src: the path of the source file.
// dstFolder: the path of the destination folder.
// error: returns an error if any operation fails.
func CopyFile(src, dstFolder string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file having the same filename as in src in the specified folder dstPath
	dstPath := filepath.Join(dstFolder, filepath.Base(src))

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Flush the written data to disk
	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

// getFileContents reads the contents of the specified file and returns them as a slice of strings.
// It takes a filename string as a parameter and returns a slice of strings and an error.
func getFileContents(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Close the file when the function returns

	// Check if the file is empty
	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		return []string{}, nil // Return an empty slice if the file is empty
	}

	// Read the file line by line and store the lines in a slice
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil // Return the slice of lines
}

// main function to copy files based on the list provided in the command line arguments
func main() {
	// read command line arguments
	args := os.Args[1:]

	// check if there are enough arguments
	if len(args) < 2 {
		// write a better error message to print here
		fmt.Println("Usage: ./copymultiple <file_to_copy> <destination_list_file>")
		return
	}

	// get the file to copy and the destination list file from command line arguments
	fileToCopy := args[0]
	destinationListFile := args[1]

	// get the contents of the destination list file
	destinationPathList, err := getFileContents(destinationListFile)
	// if destinationPathList is empty array return an error message and exit
	if len(destinationPathList) == 0 {
		fmt.Println("Error: Destination path list file is empty")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	// loop through the destination path list and copy each file
	for _, destinationFolder := range destinationPathList {
		err := CopyFile(fileToCopy, destinationFolder)
		if err != nil {
			fmt.Println("[ERROR] in copying ", fileToCopy, "to ", destinationFolder)
			log.Fatal(err)
			return
		} else {
			fmt.Println("[SUCCESS] Copied", fileToCopy, "to", destinationFolder)
		}
	}
}
