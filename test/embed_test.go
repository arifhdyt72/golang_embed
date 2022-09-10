package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"testing"

	_ "embed"
)

//go:embed version.txt
var version string

//go:embed download.jpg
var imageData []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

func TestEmbedByteArray(t *testing.T) {
	err := ioutil.WriteFile("new_image_data.jpg", imageData, fs.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}
}

func TestEmbedMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dataEntries, err := path.ReadDir("files")
	if err != nil {
		log.Panic(err.Error())
	}

	for _, entry := range dataEntries {
		fmt.Println(entry.Name())
		if !entry.IsDir() {
			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				log.Panic(err.Error())
				continue
			}

			fmt.Println(string(file))
		}
	}
}
