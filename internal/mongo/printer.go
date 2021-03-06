package mongo

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/qeesung/image2ascii/convert"
	"math"
	"os"
	"strconv"
	"stupid-ddbs/internal/hdfs"
	"time"
)

//type ArticleDoc struct {
//	Id          string `json:"id"`
//	Timestamp   string `json:"timestamp"`
//	Aid         string `json:"aid"`
//	Title       string `json:"title"`
//	Category    string `json:"category"`
//	Abstract    string `json:"abstract"`
//	ArticleTags string `json:"articleTags"`
//	Authors     string `json:"authors"`
//	Language    string `json:"language"`
//	Text        string `json:"text"`
//	Image       string `json:"image"`
//	Video       string `json:"video"`
//}
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

// tsp is int64 represent in string
func formatTimestamp(tsp string) string{
	val, _ := strconv.Atoi(tsp)
	return time.Unix(int64(val), 0).Format(DefaultTimeLayout)
}
func formatStringSlice(strings []string) string {
	ret := ""
	for  _, str := range strings {
		ret += str + " "
	}
	return ret
}

func CollectionPrinter(collectionName string, res []interface{}, detailDisplay bool) {
	table := tablewriter.NewWriter(os.Stdout)
	if collectionName == "article" {

		headers := []string{"id", "timestamp", "aid", "title", "category", "abstract", "articleTags", "authors", "language", "text", "image", "video"}
		table.SetHeader(headers)
		images := make([]string, 0)
		contents := make([]string, 0)
		for _, v := range res {
			tmp := v.(ArticleDoc)
			row := make([]string, 0)
			row = append(row, tmp.Id)
			row = append(row, formatTimestamp(tmp.Timestamp))
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

			// show image
			if detailDisplay {
				convertOptions := convert.DefaultOptions
				convertOptions.FixedWidth = 30
				convertOptions.FixedHeight = 30
				convertOptions.Ratio = 1
				convertOptions.Ratio = 1
				convertOptions.FitScreen = false
				image := hdfs.GetArticleImages(tmp.Aid)
				if len(image) > 0 {
					converter := convert.NewImageConverter()
					displayString := converter.Image2ASCIIString(image[0], &convertOptions)
					//row = append(row, displayString)
					//row = append(row, "")
					images = append(images, displayString)
				} else {
					//row = append(row, "")
				}
				content := hdfs.GetArticleContent(tmp.Aid)
				content = content[0: int(math.Min(500, float64(len(content))))]
				contentProcess := ""
				for i := 0; i < len(content); i+=100 {
					if (i + 100) >= len(content) {
						contentProcess += content[i:len(content)] + "\n"
					} else {
						contentProcess += content[i:i+100] + "\n"
					}
				}
				//row = append(row, contentProcess)
				contents = append(contents, contentProcess)
			} else {
				//row = append(row, "")
				//row = append(row, "")
			}

			table.Append(row)
		}
		table.Render()
		if len(images) == 1 && len(contents) == 1{
			for _, img := range(images) {
				fmt.Println(img)
			}
			for _, img := range(contents) {
				fmt.Println(img)
			}
		}

	} else if collectionName == "read" {
		table.SetHeader([]string{"timestamp", "id", "uid", "aid", "readTimeLength", "agreeOrNot", "commentOrNot", "shareOrNot", "commentDetail"})
		for _, v := range res {
			tmp := v.(ReadDoc)
			row := make([]string, 0)
			row = append(row, tmp.Id)
			row = append(row, formatTimestamp(tmp.Timestamp))
			row = append(row, tmp.Uid)
			row = append(row, tmp.Aid)
			row = append(row, tmp.ReadTimeLength)
			row = append(row, tmp.AgreeOrNot)
			row = append(row, tmp.CommentOrNot)
			row = append(row, tmp.ShareOrNot)
			row = append(row, tmp.CommentDetail)
			table.Append(row)
		}
		table.Render()
	} else if collectionName == "user" {
		table.SetHeader([]string{"timestamp", "id", "uid", "name", "gender", "email", "phone", "dept", "grade", "language", "region", "role", "preferTags", "obtainedCredits"})
		for _, v := range res {
			tmp := v.(UserDoc)
			row := make([]string, 0)
			row = append(row, formatTimestamp(tmp.Timestamp))
			row = append(row, tmp.Id)
			row = append(row, tmp.Uid)
			row = append(row, tmp.Name)
			row = append(row, tmp.Gender)
			row = append(row, tmp.Email)
			row = append(row, tmp.Phone)
			row = append(row, tmp.Dept)
			row = append(row, tmp.Grade)
			row = append(row, tmp.Language)
			row = append(row, tmp.Region)
			row = append(row, tmp.Role)
			row = append(row, tmp.PreferTags)
			row = append(row, tmp.ObtainedCredits)
			table.Append(row)
		}
		table.Render()
	} else if collectionName == "beread" {
		table.SetHeader([]string{"timestamp", "aid", "readNum", "readUidList", "commentNum", "commentUidList", "agreeNum", "agreeUidList", "shareNum", "shareUidList"})
		for _, v := range res {
			tmp := v.(BereadDoc)
			row := make([]string, 0)
			row = append(row, formatTimestamp(strconv.Itoa(tmp.Timestamp)))
			row = append(row, tmp.Aid)
			row = append(row, strconv.Itoa(tmp.ReadNum))
			row = append(row, formatStringSlice(tmp.ReadUidList))
			row = append(row, strconv.Itoa(tmp.CommentNum))
			row = append(row, formatStringSlice(tmp.CommentUidList))
			row = append(row, strconv.Itoa(tmp.AgreeNum))
			row = append(row, formatStringSlice(tmp.AgreeUidList))
			row = append(row, strconv.Itoa(tmp.ShareNum))
			row = append(row, formatStringSlice(tmp.ShareUidList))
			table.Append(row)
		}
		table.Render()
	} else if collectionName == "popular" {
		table.SetHeader([]string{"time", "granularity", "article_aids"})
		for _, v := range res {
			tmp := v.(PopularDoc)
			row := make([]string, 0)
			row = append(row, tmp.Time)
			row = append(row, tmp.Granularity)
			row = append(row, formatStringSlice(tmp.ArticleAidList))
			table.Append(row)
		}
		table.Render()
	}
}

func ResultPrinter(header []string, row[][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs

	//table.SetHeader(header)
	//table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(row)
	table.Render()
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