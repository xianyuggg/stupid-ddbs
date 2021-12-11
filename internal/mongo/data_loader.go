package mongo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"stupid-ddbs/logutil"
)

type ArticleDoc struct {
	Id          string `json:"id"`
	Timestamp   string `json:"timestamp"`
	Aid         string `json:"aid"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Abstract    string `json:"abstract"`
	ArticleTags string `json:"articleTags"`
	Authors     string `json:"authors"`
	Language    string `json:"language"`
	Text        string `json:"text"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}
type ReadDoc struct {
	Timestamp      string `json:"timestamp"`
	Id             string `json:"id"`
	Uid            string `json:"uid"`
	Aid            string `json:"aid"`
	ReadTimeLength string `json:"readTimeLength"`
	AgreeOrNot     string `json:"agreeOrNot"`
	CommentOrNot   string `json:"commentOrNot"`
	ShareOrNot     string `json:"shareOrNot"`
	CommentDetail  string `json:"commentDetail"`
}

type UserDoc struct {
	Timestamp       string `json:"timestamp"`
	Id              string `json:"id"`
	Uid             string `json:"uid"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Dept            string `json:"dept"`
	Grade           string `json:"grade"`
	Language        string `json:"language"`
	Region          string `json:"region"`
	Role            string `json:"role"`
	PreferTags      string `json:"preferTags"`
	ObtainedCredits string `json:"obtainedCredits"`
}

func LoadArticleDataFromLocal(target string) ([]interface{}, error){
	path := fmt.Sprintf("./dataset/python-generate-3-sized-datasets_new/%v.dat", target)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//stat, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//var size = stat.Size()
	//fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	var result []interface{}

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				log.Info("file read ok")
				break
			} else {
				log.Error("file read error")
				return nil, err
			}
		}

		if target == "article" {
			var tmp ArticleDoc
			err = json.Unmarshal([]byte(line), &tmp)
			result = append(result, tmp)
		} else if target == "read" {
			var tmp ReadDoc
			err = json.Unmarshal([]byte(line), &tmp)
			result = append(result, tmp)
		} else if target == "user" {
			var tmp UserDoc
			err = json.Unmarshal([]byte(line), &tmp)
			result = append(result, tmp)
		}
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	return result, nil
}

