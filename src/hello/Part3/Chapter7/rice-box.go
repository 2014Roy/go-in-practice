package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "appstore.pdf",
		FileModTime: time.Unix(1506323025, 0),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1507863526, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "appstore.pdf"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`./files/`, &embedded.EmbeddedBox{
		Name: `./files/`,
		Time: time.Unix(1507863526, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"appstore.pdf": file2,
		},
	})
}