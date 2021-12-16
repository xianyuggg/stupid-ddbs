package moniter

import (
	"strconv"
	"stupid-ddbs/internal/hdfs"
	"stupid-ddbs/internal/mongo"
	"time"
)

func PrintAllCollectionsStats()  {
	colls := mongo.GetCollections()
	//colls = []string{"read"}
	mongo.PrintCollectionStats(colls)
}

func PrintDbStats() {
	//mongo.PrintDbStats()
	// TODO
}

func PrintShards() {
	mongo.PrintShards()
}

func ShowHdfsPathStatus(path string) {
	files, err := hdfs.GetPathInfo(path)
	if err != nil {
		println("path not valid")
	}

	header := []string{"PathSuffix", "Type",  "Length", "AccessTime", "BlockSize", "Group", "ModificationTime", "Owner",  "Permission", "Replication"}
	rows := make([][]string, 0)
	for _, f := range files {
		row := make([]string, 0)
		row = append(row, f.PathSuffix)
		row = append(row, f.Type)
		row = append(row, strconv.FormatInt(f.Length, 10))


		row = append(row, time.Unix(f.AccesTime, 0).Format("2006-1-2"))
		row = append(row, strconv.FormatInt(f.BlockSize, 10))
		row = append(row, f.Group)

		row = append(row, time.Unix(f.ModificationTime, 0).Format("2006-1-2"))
		row = append(row, f.Owner)

		row = append(row, f.Permission)
		row = append(row, strconv.FormatInt(f.Replication, 10))
		rows = append(rows, row)

	}
	mongo.ResultPrinter(header, rows)
}