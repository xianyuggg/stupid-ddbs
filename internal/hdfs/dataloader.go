package hdfs

import (
	"bytes"
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"io/ioutil"
	"path"
	log "stupid-ddbs/logutil"
)

type FileType = int
const (
	TypeText FileType = iota
	TypeImage
	TypeVideo
)

// ArticleFiles an article itself includes several files
type ArticleFiles struct {
	fileTypes []FileType
	contents [][]byte
}
type ImageFiles struct {
	fileNames []string
	contents [][]byte
}
type VideoFiles struct {
	fileNames []string
	contents [][]byte
}

func LoadArticlesFromLocal(begin int, end int) ([]ArticleFiles, []int) {
	//articleList := []
	dirFiles := make([]ArticleFiles, 0)
	ids := make([]int, 0)
	for i := begin; i < end; i++ {
		dir := fmt.Sprintf("./dataset/python-generate-3-sized-datasets_new/articles/article%v/", i)
		files, _ := ioutil.ReadDir(dir)

		dirFile := ArticleFiles{
			fileTypes: make([]FileType, 0),
			contents:  make([][]byte, 0),
		}
		for _, f := range files {

			fileSuffix := path.Ext(f.Name())
			if fileSuffix == ".txt" {
				dirFile.fileTypes = append(dirFile.fileTypes, TypeText)
			} else if fileSuffix == ".jpg" {
				dirFile.fileTypes = append(dirFile.fileTypes, TypeImage)
			} else if fileSuffix == ".flv" {
				dirFile.fileTypes = append(dirFile.fileTypes, TypeVideo)
			} else {
				log.Error("met unknown suffix", fileSuffix)
				return nil, nil
			}

			articleBytes, err := ioutil.ReadFile(dir + f.Name())
			if err != nil {
				log.Error("read error:", dir + f.Name())
			}
			dirFile.contents = append(dirFile.contents, articleBytes)
		}
		dirFiles = append(dirFiles, dirFile)
		log.Infof("load %v into memory", dir)
		ids = append(ids, i)
	}
	log.Infof("load articles success")
	return dirFiles, ids
}

func LoadImagesFromLocal() ImageFiles {
	imageFiles := ImageFiles{
		fileNames: make([]string, 0),
		contents:  make([][]byte, 0),
	}
	dirName := fmt.Sprintf("./dataset/python-generate-3-sized-datasets_new/image/")
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		imageFiles.fileNames = append(imageFiles.fileNames ,f.Name())
		imageBytes, err := ioutil.ReadFile(dirName + f.Name())
		if err != nil {
			log.Error("read error:", dirName + f.Name())
		}
		imageFiles.contents = append(imageFiles.contents, imageBytes)
		log.Infof("load %v into memory", f.Name())
	}
	log.Infof("load images success")
	return imageFiles
}

func LoadVideosFromLocal() VideoFiles {
	panic(0)
}

func (m* Manager) loadArticleDataIntoHDFS(begin int, end int)  {
	// article
	articles, ids := LoadArticlesFromLocal(begin, end)
	if len(articles) != len(ids) {
		panic("content length not equal")
	}
	articlesDir := "/articles"
	for _, idx := range ids {
		imageNum := 1
		for i, content := range articles[idx-begin].contents {
			tmpPath := fmt.Sprintf("%v/%v/", articlesDir, idx)
			if articles[idx-begin].fileTypes[i] == TypeText {
				tmpPath = tmpPath + "0.txt"
			} else if articles[idx-begin].fileTypes[i] == TypeImage {
				tmpPath = tmpPath + fmt.Sprintf("%v.jpg", imageNum)
				imageNum += 1
			} else {
				log.Warning("other types not implemented")
				continue
			}
			reader := bytes.NewReader(content)
			if _, err := m.client.Create(reader, gowfs.Path{Name: tmpPath}, true, 0, 0, 0666, 0); err != nil {
				log.Error(err)
			}
			log.Infof("create %v success", tmpPath)
		}
	}

}

func (m* Manager) loadImageDataIntoHDFS() {
	images := LoadImagesFromLocal()
	imagesDir := "/images"
	for i, content := range images.contents {
		tmpPath := fmt.Sprintf("%v/%v", imagesDir, images.fileNames[i])
		reader := bytes.NewReader(content)
		if _, err := m.client.Create(reader, gowfs.Path{Name: tmpPath}, true, 0, 0, 0666, 0); err != nil {
			log.Error(err)
		}
		log.Infof("create %v success", tmpPath)
	}
}

func (m* Manager) LoadDataIntoHDFS() {
	for begin := 1000; begin <= 9500; begin += 500 {
		m.loadArticleDataIntoHDFS(begin, begin + 500)
	}
	m.loadImageDataIntoHDFS()
}