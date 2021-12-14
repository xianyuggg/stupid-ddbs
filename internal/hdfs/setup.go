package hdfs

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	log "stupid-ddbs/logutil"
)

func (m* Manager) PathInit() {
	// check if path init
	paths, _ := m.client.ListStatus(gowfs.Path{Name: "/"})

	videoPathSkip := false
	imagePathSkip := false
	articlePathSkip := false
	for _, path := range paths {
		if path.PathSuffix == "videos" {
			videoPathSkip = true
		}
		if path.PathSuffix == "images" {
			imagePathSkip = true
		}
		if path.PathSuffix == "articles" {
			articlePathSkip = true
		}
	}

	if !imagePathSkip {
		imagePath := gowfs.Path{Name: "/images"}
		if _, err := m.client.MkDirs(imagePath, 0666); err != nil {
			panic(err)
		}
	}
	if !videoPathSkip {
		videoPath := gowfs.Path{Name: "/videos"}
		if _, err := m.client.MkDirs(videoPath, 0666); err != nil {
			panic(err)
		}
	}
	if !articlePathSkip {
		articlesPath := gowfs.Path{Name: "/articles"}
		if _, err := m.client.MkDirs(articlesPath, 0666); err != nil {
			panic(err)
		}
		for i := 0; i < 10000; i++ {
			articlePath := gowfs.Path{Name: fmt.Sprintf("/articles/%v", i)}
			if _, err := m.client.MkDirs(articlePath, 0666); err != nil {
				panic(err)
			}
		}
	}
	log.Info("path init success")
}
