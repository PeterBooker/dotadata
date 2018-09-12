// +build ignore

package main

import (
	"log"

	"github.com/PeterBooker/dota2data/internal/files"
	"github.com/PeterBooker/dota2data/internal/tmpls"
	"github.com/shurcooL/vfsgen"
)

func main() {
	var err error

	// Static Assets
	err = vfsgen.Generate(files.Assets, vfsgen.Options{
		Filename:     "internal/files/assets.go",
		PackageName:  "files",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// HTML Templates
	err = vfsgen.Generate(tmpls.Assets, vfsgen.Options{
		Filename:     "internal/tmpls/templates.go",
		PackageName:  "tmpls",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
