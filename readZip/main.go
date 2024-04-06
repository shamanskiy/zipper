package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	zipReader, err := zip.OpenReader("archiveFromMemory.zip")
	check(err)
	defer close(zipReader)

	for _, fileInZip := range zipReader.File {
		fmt.Printf("Contents of %s:\n", fileInZip.Name)

		fileInZipReader, err := fileInZip.Open()
		check(err)
		defer close(fileInZipReader)

		io.CopyN(os.Stdout, fileInZipReader, 1000)

		fmt.Println()
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
