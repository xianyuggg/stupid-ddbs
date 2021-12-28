package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
	log "stupid-ddbs/logutil"
)

func QueryData(collectionName string, andConditions []Cond, showDetails bool) {
	m := GetManagerInstance()
	if !m.CheckCollections(collectionName) {
		println("collection does not exist")
	}
	res, err := m.QueryData(collectionName, andConditions)
	if err != nil {
		println(err)
		return
	}
	CollectionPrinter(collectionName, res, showDetails)
}

/*
	Load Data
*/
func LoadData(target string) {
	m := GetManagerInstance()
	//collections := GetCollections()
	//collMap := make(map[string]int)
	//for _,1 coll := range collections {
	//	collMap[coll] = 1
	//}
	//if _, ok := collMap[target]; ok {
	//	println("target exists, plz drop fist")
	//	return
	//}
	//if len(collections) != 0 && target == "local"{
	//	println("collection exists, plz drop first")
	//	return
	//}

	switch target {
	case "article":
		tmp, _ := loadCollectionFromLocal("article")
		_ = bulkLoadDataToMongo(m.db, "article", tmp)
	case "read":
		tmp, _ := loadCollectionFromLocal("read")
		_ = bulkLoadDataToMongo(m.db, "read", tmp)
	case "user":
		tmp, _ := loadCollectionFromLocal("user")
		_ = bulkLoadDataToMongo(m.db, "user", tmp)
	case "local":
		_ = m.LoadAllData()
	case "beread":
		if err := m.ComputeBeRead(false); err != nil {
			println(err)
		}
	case "popular":
		if err := m.ComputePopular(); err != nil {
			println(err)
		}
	case "all":
		ShardingSetup()
		if err := m.LoadAllData(); err != nil {
			println(err)
		}
		if err := m.ComputeBeRead(false); err != nil {
			println(err)
		}
		if err := m.ComputePopular(); err != nil {
			println(err)
		}
	default:
		println("target not valid")
	}
}

/*
	Drop collections
 */
func DropCollection(target string) {
	m := GetManagerInstance()
	if target == "all" {
		if err := m.db.Drop(context.TODO()); err != nil {
			println(err)
		}
	}

	collections := GetCollections()
	collMap := make(map[string]int)
	for _, coll := range collections {
		collMap[coll] = 1
	}
	if _, ok := collMap[target]; !ok {
		println("target not exists")
	}
	if err := m.db.Collection(target).Drop(context.TODO()); err != nil {
		println(err)
	}
}



/*
	Monitoring functions
 */


type RetListShards struct {
	Shards []struct {
		ID    string `bson:"_id"`
		Host  string `bson:"host"`
		State int `bson:"state"`
	} `bson:"shards"`
}

func PrintShards() {
	db, err := mongoGetDatabase("admin")
	if err != nil {
		log.Error(err)
	}
	res := db.RunCommand(context.TODO(), bson.D{
		//{"balancerCollectionStatus", "ProjectDB.article"},
		{"listShards", "1"},
		//{"collStats", "article"},
	})
	if res.Err() != nil {
		log.Error(res.Err().Error())
		return
	}
	var retVal RetListShards
	if err := res.Decode(&retVal); err != nil {
		log.Error(err)
	}
	//fmt.Print(retVal)
	headers := []string{"shard", "host", "state"}
	rows := make([][]string, 0)
	for _, shard := range retVal.Shards {
		row := make([]string, 0)
		row = append(row, shard.ID)
		row = append(row, shard.Host)
		row = append(row, strconv.Itoa(shard.State))
		rows = append(rows, row)
	}
	ResultPrinter(headers, rows)
}

func getWrappedValue(tmp string, wrappedType string) string{
	//something like this {"$numberInt":"3120594"}
	var result map[string]interface{}
	err := json.Unmarshal([]byte(tmp), &result)
	if err != nil {
		log.Warning(err)
		return ""
	}
	if wrappedType == "int"{
		return result["$numberInt"].(string)
	} else if wrappedType == "double" {
		return result["$numberDouble"].(string)
	} else {
		panic("unimplemented")
	}

}
func GetCollections() []string {
	db,err := mongoGetDatabase(DefaultDbName)
	if err != nil {
		panic(err)
	}
	colls, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	return colls
}


func PrintDbStats() {
	db, err := mongoGetDatabase(DefaultDbName)
	if err != nil {
		log.Error(err)
	}
	res := db.RunCommand(context.TODO(), bson.D{
		{"dbStats", 1},
	})
	if res.Err() != nil {
		log.Error(res.Err().Error())
		return
	}
	bsonRaw, _ := res.DecodeBytes()
	fmt.Println(bsonRaw)
}

func PrintCollectionStats(colls []string) {
	db, err := mongoGetDatabase(DefaultDbName)
	if err != nil {
		log.Error(err)
	}
	// storageSize The total amount of storage allocated to this collection for document storage. The scale argument affects this value.
	// size The total uncompressed size in memory of all records in a collection

	headers := []string{"namespace", "sharded", "shards", "size", "count", "avgObjSize", "storageSize"}
	rows := [][]string{}

	for _, coll := range colls {
		res := db.RunCommand(context.TODO(), bson.D{
			{"collStats", coll},
		})
		if res.Err() != nil {
			log.Error(res.Err().Error())
			return
		}
		bsonRaw, _ := res.DecodeBytes()


		if fmt.Sprintf("%v", bsonRaw.Lookup("sharded")) == "false" {
			row := make([]string, 0)

			row = append(row, strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("ns")), "\""))
			row = append(row, strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("sharded")), "\""))
			row = append(row, strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("primary")), "\""))

			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("size")), "\""), "int"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("count")), "\""), "int"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("avgObjSize")), "\""), "double"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("storageSize")), "\""), "int"))
			rows = append(rows, row)
		} else {
			//rows := make([][]string, 0)
			row := make([]string, 0)
			//shardedHeaders := []string{"namespace", "sharded", "shards", "size", "count", "avgObjSize", "storageSize"}
			row = append(row, strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("ns")), "\""))
			row = append(row, strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("sharded")), "\""))
			row = append(row, "Total")
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("size")), "\""), "int"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("count")), "\""), "int"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("avgObjSize")), "\""), "double"))
			row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", bsonRaw.Lookup("storageSize")), "\""), "int"))
			rows = append(rows, row)

			shards, err := bsonRaw.Lookup("shards").Document().Elements()

			if err != nil {
				panic(err)
			}
			for _, shard := range shards {
				doc := shard.Value().Document()
				row := make([]string, 0)
				row = append(row, "")
				row = append(row, "")
				row = append(row, shard.Key())
				row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", doc.Lookup("size")), "\""), "int"))
				row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", doc.Lookup("count")), "\""), "int"))
				row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", doc.Lookup("avgObjSize")), "\""), "int"))
				row = append(row, getWrappedValue(strings.Trim(fmt.Sprintf("%v", doc.Lookup("storageSize")), "\""), "int"))
				rows = append(rows, row)
			}
			//ResultPrinter(shardedHeaders, rows)
		}

	}

	ResultPrinter(headers, rows)



}

