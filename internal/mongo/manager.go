package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"stupid-ddbs/logutil"
	"sync"
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
	articles, _ := LoadArticleDataFromLocal("article")
	reads, _ := LoadArticleDataFromLocal("read")
	users, _ := LoadArticleDataFromLocal("user")

	_ = bulkLoadDataToMongo(m.db, "article", articles)
	_ = bulkLoadDataToMongo(m.db, "read", reads)
	_ = bulkLoadDataToMongo(m.db, "user", users)

	return nil
}

func (m* Manager) QueryData(collectionName string, andConditions []Cond) []interface{}{
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

	cursor, err := coll.Find(context.TODO(), condFinalBson)
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
		default:
			log.Error("collection not exists")
		}
	}

	return result
}
