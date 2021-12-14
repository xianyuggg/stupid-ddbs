package hdfs

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"image"
	"io/ioutil"
	log "stupid-ddbs/logutil"
)

func GetArticleContent(aid string) string{
	m := GetManagerInstance()
	articlePath := gowfs.Path{Name: fmt.Sprintf("/articles/%v/0.txt", aid)}
	stat, err := m.client.GetFileStatus(articlePath)
	if err != nil {
		log.Error(err)
		return ""
	}
	reader, _ := m.client.Open(articlePath, 0, stat.Length, int(stat.Length))
	data, _ := ioutil.ReadAll(reader)
	return string(data)
}

func GetArticleImages(aid string) []image.Image {
	m := GetManagerInstance()
	articlePath := fmt.Sprintf("/articles/%v", aid)
	fileStats, _ := m.client.ListStatus(gowfs.Path{Name: articlePath})
	retImages := make([]image.Image, 0)
	for i := 1; i < len(fileStats); i++ {
		tmpImagePath := fmt.Sprintf("%v/%v.jpg", articlePath, i)
		reader, _ := m.client.Open(gowfs.Path{Name: tmpImagePath}, 0, fileStats[i].Length, int(fileStats[i].Length))
		tmpImage, _, err := image.Decode(reader)
		if err != nil {
			log.Error(err)
			return nil
		}
		retImages = append(retImages, tmpImage)
	}
	return retImages
}
