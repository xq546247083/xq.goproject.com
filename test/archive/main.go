package main

import (
	"archive/tar"
	"log"
	"os"
)

func main() {
	testTar()
}

func testTar() {
	// Create a buffer to write our archive to.
	reader, _ := os.Create("D:\\test.rar")

	// Create a new tar archive.
	tw := tar.NewWriter(reader)
	defer tw.Close()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
}
