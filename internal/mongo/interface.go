package mongo

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
	ResultPrinter(collectionName, res, showDetails)
}
