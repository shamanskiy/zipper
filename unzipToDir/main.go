package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	zipReader, err := zip.OpenReader("archiveFromMemory.zip")
	check(err)
	defer close(zipReader)

	for _, fileInZip := range zipReader.File {
		fileInZipReader, err := fileInZip.Open()
		check(err)
		defer close(fileInZipReader)

		unzippedFile, err := os.Create(fileInZip.Name)
		check(err)
		defer close(unzippedFile)

		_, err = io.Copy(unzippedFile, fileInZipReader)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func close(closer io.Closer) {
	err := closer.Close()
	check(err)
}
