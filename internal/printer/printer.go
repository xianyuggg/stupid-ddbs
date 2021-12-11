package printer

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"stupid-ddbs/internal/mongo"
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
//type ReadDoc struct {
//	Timestamp      string `json:"timestamp"`
//	Id             string `json:"id"`
//	Uid            string `json:"uid"`
//	Aid            string `json:"aid"`
//	ReadTimeLength string `json:"readTimeLength"`
//	AgreeOrNot     string `json:"agreeOrNot"`
//	CommentOrNot   string `json:"commentOrNot"`
//	ShareOrNot     string `json:"shareOrNot"`
//	CommentDetail  string `json:"commentDetail"`
//}
//
//type UserDoc struct {
//	Timestamp       string `json:"timestamp"`
//	Id              string `json:"id"`
//	Uid             string `json:"uid"`
//	Name            string `json:"name"`
//	Gender          string `json:"gender"`
//	Email           string `json:"email"`
//	Phone           string `json:"phone"`
//	Dept            string `json:"dept"`
//	Grade           string `json:"grade"`
//	Language        string `json:"language"`
//	Region          string `json:"region"`
//	Role            string `json:"role"`
//	PreferTags      string `json:"preferTags"`
//	ObtainedCredits string `json:"obtainedCredits"`
//}

func ResultPrinter(collectionName string, res []interface{}) {
	table := tablewriter.NewWriter(os.Stdout)
	if collectionName == "article" {
		table.SetHeader([]string{"id", "timestamp", "aid", "title", "category", "abstract", "articleTags", "authors", "language", "text", "image", "video"})
		for _, v := range res {
			tmp := v.(mongo.ArticleDoc)
			row := make([]string, 0)
			row = append(row, tmp.Id)
			row = append(row, tmp.Timestamp)
			row = append(row, tmp.Aid)
			row = append(row, tmp.Title)
			row = append(row, tmp.Category)
			row = append(row, tmp.Abstract)
			row = append(row, tmp.ArticleTags)
			row = append(row, tmp.Authors)
			row = append(row, tmp.Language)
			row = append(row, tmp.Text)
			row = append(row, tmp.Image)
			row = append(row, tmp.Video)
			table.Append(row)
		}
		table.Render()
	}
}
