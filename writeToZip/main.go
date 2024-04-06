package main

import (
	"archive/zip"
	"log"
	"os"
)

var fileData = []struct {
	Name, Body string
}{
	{"files/readme.txt", "This archive contains some text files."},
	{"files/go/gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	{"files/todo/todo.txt", "Get animal handling licence.\nWrite more examples."},
}

func main() {
	zipFile := createFile("archiveFromMemory.zip")
	defer closeZipFile(zipFile)

	zipWriter := zip.NewWriter(zipFile)
	defer closeZipWriter(zipWriter)

	for _, file := range fileData {
		fileInZip, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = fileInZip.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func closeZipFile(zipFile *os.File) {
	err := zipFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func closeZipWriter(zipWriter *zip.Writer) {
	err := zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
