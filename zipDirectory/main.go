package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	zipFile, err := os.Create("archiveFromDirectory.zip")
	checkErr(err)
	defer closeFile(zipFile)

	zipWriter := zip.NewWriter(zipFile)
	defer closeZipWriter(zipWriter)

	visitFile := func(path string, info os.FileInfo, err error) error {
		checkErr(err)

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		checkErr(err)
		defer closeFile(file)

		fileInZip, err := zipWriter.Create(file.Name())
		checkErr(err)

		_, err = io.Copy(fileInZip, file)
		checkErr(err)

		return nil
	}

	filepath.Walk("files", visitFile)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
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
