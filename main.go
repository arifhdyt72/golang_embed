package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

//go:embed version.txt
var version string

//go:embed download.jpg
var imageData []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("new_image_data.jpg", imageData, fs.ModePerm)
	if err != nil {
		log.Panic(err.Error())
	}

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
