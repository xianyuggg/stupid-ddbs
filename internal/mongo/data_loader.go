package mongo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadArticleDataFromLocal() {
	file, err := os.Open("dataset/python-generate-3-sized-datasets_new/article.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}

		fmt.Println(line)
		var result map[string]interface{}
		err = json.Unmarshal([]byte(line), &result)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)

	}

}