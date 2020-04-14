package main

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"log"

	"github.com/mholt/archiver"
)

func archiveDir(dirName string, destFileName string) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	//files := readDirFiles(dirName)
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		contents, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		hdr := &tar.Header{
			Name: file.Name(),
			Mode: 0600,
			Size: int64(len(contents)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write(contents); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(destFileName, buf.Bytes(), 0644)

}

func archiveDir2(dirName string, destFileName string) {

	//files := readDirFiles(dirName)
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	allFiles := []string{}
	for _, file := range files {
		allFiles = append(allFiles, file.Name())
	}

	err = archiver.Archive(allFiles, "my_tar.tar")
	if err != nil {
		log.Fatal(err)
	}
}
