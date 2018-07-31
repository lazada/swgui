package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./v2/static")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName: "v2",
		Filename:    "v2/static.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
