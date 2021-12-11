package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"stupid-ddbs/logutil"
)

func bulkLoadDataToMongo(db *mongo.Database, collectionName string, values []interface{}) error{
	collection := db.Collection(collectionName)
	_, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("bulk load ok:", collectionName)
	return nil
}

type OpType = int
const (
	OpDefault OpType = iota
	OpCompEQ
	OpCompLT
	OpCompGT
	OpCompLE
	OpCompGE
	OpCompNE

	OpLogicAND
	OpLogicOR
	OpLogicNOT
)

type Cond struct {
	Field string
	Op    OpType
	Val   string
}

