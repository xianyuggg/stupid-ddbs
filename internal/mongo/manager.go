package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"strconv"
	"stupid-ddbs/logutil"
	"sync"
	"time"
)

type Manager struct {
	db *mongo.Database
}

var instance *Manager
var once sync.Once

func GetManagerInstance() *Manager {
	once.Do(func() {
		log.Info("mongo manager starts to initialize.")
		defer log.Info("mongo manager has been initialized.")
		db, _ := mongoGetDatabase("")
		instance = &Manager{
			db: db,
		}
	})
	return instance
}

func (m* Manager) Close() {
	_ = mongoCloseDatabase(m.db)
}


func (m* Manager) LoadAllData() error{
	articles, _ := loadCollectionFromLocal("article")
	reads, _ := loadCollectionFromLocal("read")
	users, _ := loadCollectionFromLocal("user")

	_ = bulkLoadDataToMongo(m.db, "article", articles)
	_ = bulkLoadDataToMongo(m.db, "read", reads)
	_ = bulkLoadDataToMongo(m.db, "user", users)

	return nil
}

func (m* Manager) CheckCollections(collectionName string) bool{
	colls, _ := m.db.ListCollectionNames(context.TODO(), bson.D{})
	retValue := false
	for _, name := range colls {
		if collectionName == name {
			retValue = true
		}
	}
	return retValue
}
func (m* Manager) QueryData(collectionName string, andConditions []Cond) ([]interface{}, error){
	coll := m.db.Collection(collectionName)
	//db.col.find({"likes": {$gt:50}, $or: [{"by": "菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty(
	attrMapCond := make(map[string][]Cond)
	if andConditions == nil {
		andConditions = make([]Cond, 0)
	}
	for _, cond := range andConditions {
		if _, ok := attrMapCond[cond.Field]; !ok {
			attrMapCond[cond.Field] = make([]Cond, 0)
		}
		attrMapCond[cond.Field] = append(attrMapCond[cond.Field], cond)
	}

	condFinalBson := bson.D{}

	for field, conds := range attrMapCond {
		condBSON := bson.D{}
		skipFlag := false
		for _, cond := range conds {
			switch cond.Op {
			case OpCompEQ:
				condFinalBson = append(condFinalBson, bson.E{
					cond.Field,
					cond.Val,
				})
				skipFlag = true
				break
			case OpCompGE:
				condBSON = append(condBSON, bson.E{
					"$gte", cond.Val,
				})
			case OpCompGT:
				condBSON = append(condBSON, bson.E{
					"$gt", cond.Val,
				})
			case OpCompLE:
				condBSON = append(condBSON, bson.E{
					"$lte", cond.Val,
				})
			case OpCompLT:
				condBSON = append(condBSON, bson.E{
					"$lt", cond.Val,
				})
			case OpCompNE:
				condBSON = append(condBSON, bson.E{
					"$ne", cond.Val,
				})
			}
		}
		if !skipFlag {
			condFinalBson = append(condFinalBson, bson.E{
				field, condBSON,
			})
		}
	}

	cursor, err := coll.Find(context.TODO(), condFinalBson, &options.FindOptions{Collation: &CollationConfig})
	if err != nil {
		log.Error(err)
	}
	result := make([]interface{}, 0)
	for cursor.Next(context.TODO()) {
		switch collectionName {
		case "article":
			var article ArticleDoc
			if err = cursor.Decode(&article); err != nil {
				log.Error(err)
			}
			result = append(result, article)
		case "read":
			var read ReadDoc
			if err = cursor.Decode(&read); err != nil {
				log.Error(err)
			}
			result = append(result, read)

		case "user":
			var user UserDoc
			if err = cursor.Decode(&user); err != nil {
				log.Error(err)
			}
			result = append(result, user)
		case "beread":
			var user BereadDoc
			if err = cursor.Decode(&user); err != nil {
				log.Error(err)
			}
			result = append(result, user)
		case "popular":
			var user PopularDoc
			if err = cursor.Decode(&user); err != nil {
				log.Error(err)
			}
			result = append(result, user)
		default:
			log.Error("collection not exists")
		}
	}

	return result, err
}

func (m* Manager) ComputeBeRead(overwrite bool) error{
	articleCollection := m.db.Collection("article")

	//lookupStage := bson.D{
	//	{"$lookup",
	//		bson.D{
	//			{"from", "read"},
	//			{"as", "tmp"},
	//			{"pipeline", bson.A{
	//				bson.D{{"$limit", 20000}},
	//			}},
	//			},
	//		},
	//}
	//boolTmp := true
	//opt := options.AggregateOptions{
	//	AllowDiskUse: &boolTmp,
	//}
	projectStage := bson.D{
		{
			"$project",
			bson.D{
				{"_id", 0},
				{"aid", 1},
				{"timestamp", 1},
			},
		},
	}
	cursor, err := articleCollection.Aggregate(context.TODO(), mongo.Pipeline{projectStage})
	if err != nil {
		log.Error(err)
	}
	var tmpArt ArticleDoc

	bereadId := 0
	//jsonSchema := bson.M{
	//	"bsonType": "object",
	//	"required": []string{"aid", "timestamp", "readNum", "readUidList", "commentNum", "commentUidList", "agreeNum", "agreeUidList", "shareNum", "shareUidList"},
	//	"properties": bson.M{
	//		"aid": bson.M{
	//			"bsonType": "string",
	//		},
	//		"timestamp": bson.M{
	//			"bsonType": "int",
	//		},
	//		"readNum": bson.M{
	//			"bsonType": "int",
	//		},
	//		"readUidList": bson.M{
	//			"bsonType": "array",
	//		},
	//		"commentNum": bson.M{
	//			"bsonType": "int",
	//		},
	//		"commentUidList": bson.M{
	//			"bsonType": "array",
	//		},
	//		"agreeNum": bson.M{
	//			"bsonType": "int",
	//		},
	//		"agreeUidList": bson.M{
	//			"bsonType": "array",
	//		},
	//		"shareNum": bson.M{
	//			"bsonType": "int",
	//		},
	//		"shareUidList": bson.M{
	//			"bsonType": "array",
	//		},
	//	},
	//}
	//validator := bson.M{
	//	"$jsonSchema": jsonSchema,
	//}
	//opts := options.CreateCollection().SetValidator(validator)

	bereadCollection := m.db.Collection("beread")

	if overwrite {
		_, _ = bereadCollection.DeleteMany(context.TODO(), bson.D{})
		log.Info("overwrite and delete all")
	} else {
		num, _ := bereadCollection.CountDocuments(context.TODO(), bson.M{})
		if num > 0 {
			log.Info("document exists, not overwrite and return")
			return nil
		}
		log.Info("no documents, continue computing")
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&tmpArt)
		if err != nil {
			log.Error(err)
			break
		}
		//fmt.Println(tmpArt)
		res, _ := m.QueryData("read", []Cond{{"aid", OpCompEQ, tmpArt.Aid}})
		//CollectionPrinter("read", res)
		tsp, _ := strconv.Atoi(tmpArt.Timestamp)

		tmpBeread := BereadDoc{
			Aid:            tmpArt.Aid,
			Timestamp:      tsp,
			ReadNum:        0,
			ReadUidList:    make([]string, 0),
			CommentNum:     0,
			CommentUidList: make([]string, 0),
			AgreeNum:       0,
			AgreeUidList:   make([]string, 0),
			ShareNum:       0,
			ShareUidList:   make([]string, 0),
		}
		for _, read := range res {
			tmpRead := read.(ReadDoc)
			tmpBeread.ReadNum += 1
			tmpBeread.ReadUidList = append(tmpBeread.ReadUidList, tmpRead.Uid)
			if tmpRead.CommentOrNot == "1" {
				tmpBeread.CommentNum += 1
				tmpBeread.CommentUidList = append(tmpBeread.CommentUidList, tmpRead.Uid)
			}
			if tmpRead.AgreeOrNot == "1" {
				tmpBeread.AgreeNum += 1
				tmpBeread.AgreeUidList = append(tmpBeread.AgreeUidList, tmpRead.Uid)
			}
			if tmpRead.ShareOrNot == "1" {
				tmpBeread.ShareNum += 1
				tmpBeread.ShareUidList = append(tmpBeread.ShareUidList, tmpRead.Uid)
			}
		}
		//tmpBeread.ReadNum = fmt.Sprintf("%v", readNum)
		//tmpBeread.CommentNum = fmt.Sprintf("%v", commentNum)
		//tmpBeread.AgreeNum = fmt.Sprintf("%v", agreeNum)
		//tmpBeread.ShareNum = fmt.Sprintf("%v", shareNum)
		//fmt.Println(tmpBeread)
		bereadId += 1

		if res, err := bereadCollection.InsertOne(context.Background(), tmpBeread); err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("insert", res)
		}

	}

	return nil
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }


func (m* Manager) ComputePopular() error {
	//id, timestamp, temporalGranularity, articleAidList
	// Query the top-5 daily/weekly/monthly popular articles with articles details (text, image, and video if existing)
	// (involving the join of Be-Read table and Article table)
	bereadCollection := m.db.Collection("beread")
	popularCollection := m.db.Collection("popular")
	cursor, err := bereadCollection.Find(context.TODO(), bson.M{})

	yearMap := make(map[time.Time]map[string]int)
	weekMap := make(map[time.Time]map[string]int)
	monthMap := make(map[time.Time]map[string]int)

	// id time temporal granularity, articleAidList
	for cursor.Next(context.TODO()) {
		var beread BereadDoc
		if err = cursor.Decode(&beread); err == nil {
			tsp := beread.Timestamp
			t := time.Unix(int64(tsp), 0)

			// yearly
			yt := time.Date(t.Year(), 1, 1, 0, 0, 0, 0,  time.Local)
			if _, ok := yearMap[yt]; !ok {
				yearMap[yt] = make(map[string]int)
			} else {
				if _, ok := yearMap[yt]; !ok {
					yearMap[yt][beread.Aid] = beread.ReadNum
				} else {
					yearMap[yt][beread.Aid] += beread.ReadNum
				}
			}
			// monthly
			mt := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0,  time.Local)
			if _, ok := monthMap[mt]; !ok {
				monthMap[mt] = make(map[string]int)
			} else {
				if _, ok := monthMap[mt]; !ok {
					monthMap[mt][beread.Aid] = beread.ReadNum
				} else {
					monthMap[mt][beread.Aid] += beread.ReadNum
				}
			}

			// weekly
			wt := t.AddDate(0, 0, -int(t.Weekday()))
			wt = time.Date(wt.Year(), wt.Month(), wt.Day(), 0, 0, 0, 0, time.Local)
			if _, ok := weekMap[wt]; !ok {
				weekMap[wt] = make(map[string]int)
			} else {
				if _, ok := weekMap[wt]; !ok {
					weekMap[wt][beread.Aid] = beread.ReadNum
				} else {
					weekMap[wt][beread.Aid] += beread.ReadNum
				}
			}

			//fmt.Println(unixTime.Format("2006-1-2"))
		} else {
			log.Error(err)
			return err
		}
	}

	for k, v := range yearMap {
		kvPairs := make(PairList, 0)
		for aid, count := range v {
			kvPairs = append(kvPairs, Pair{
				Key:  aid ,
				Value: count,
			})
		}
		sort.Sort(kvPairs)
		kvPairs = kvPairs[0:5]
		tmpPopularDoc := PopularDoc{
			Time:           k.Format(DefaultTimeLayout),
			Granularity:    "year",
			ArticleAidList: make([]string, 0),
		}
		for i, _ := range kvPairs {
			tmpPopularDoc.ArticleAidList = append(tmpPopularDoc.ArticleAidList, kvPairs[i].Key)
		}
		if res, err := popularCollection.InsertOne(context.Background(), tmpPopularDoc); err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("insert popular", res)
		}
	}

	for k, v := range monthMap {
		kvPairs := make(PairList, 0)
		for aid, count := range v {
			kvPairs = append(kvPairs, Pair{
				Key:  aid ,
				Value: count,
			})
		}
		sort.Sort(kvPairs)
		kvPairs = kvPairs[0:5]
		tmpPopularDoc := PopularDoc{
			Time:           k.Format(DefaultTimeLayout),
			Granularity:    "month",
			ArticleAidList: make([]string, 0),
		}
		for i, _ := range kvPairs {
			tmpPopularDoc.ArticleAidList = append(tmpPopularDoc.ArticleAidList, kvPairs[i].Key)
		}
		if res, err := popularCollection.InsertOne(context.Background(), tmpPopularDoc); err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("insert popular", res)
		}
	}
	for k, v := range weekMap {
		kvPairs := make(PairList, 0)
		for aid, count := range v {
			kvPairs = append(kvPairs, Pair{
				Key:  aid ,
				Value: count,
			})
		}
		sort.Sort(kvPairs)
		kvPairs = kvPairs[0:5]
		tmpPopularDoc := PopularDoc{
			Time:           k.Format(DefaultTimeLayout),
			Granularity:    "week",
			ArticleAidList: make([]string, 0),
		}
		for i, _ := range kvPairs {
			tmpPopularDoc.ArticleAidList = append(tmpPopularDoc.ArticleAidList, kvPairs[i].Key)
		}
		if res, err := popularCollection.InsertOne(context.Background(), tmpPopularDoc); err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("insert popular", res)
		}
	}



	return nil
}