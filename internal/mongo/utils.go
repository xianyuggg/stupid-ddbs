package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var CollationConfig = options.Collation {
	Locale: "en_US",
	Strength: 1,
	NumericOrdering: true,
}
// https://docs.mongodb.com/manual/reference/collation/#std-label-collation-document-fields
//{
//locale: <string>,
//caseLevel: <boolean>,
//caseFirst: <string>,
//strength: <int>,
//numericOrdering: <boolean>,
//alternate: <string>,
//maxVariable: <string>,
//backwards: <boolean>
//}

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

