package mongodb

/*
const (
	database = "starwars"
)

// Connect do a connection with running mongodb server

func Connect(url string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return client.Database(database, nil), err
}


// ToDocument will convert any struct into bson.D primitive struct
func ToDocument(doc interface{}) *bson.D {
	data, err := bson.Marshal(doc)
	var docD *bson.D
	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("toDocument Marshal %+v", doc))
		return docD
	}
	err = bson.Unmarshal(data, &docD)
	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("toDocument Unmarshal %+v", doc))
	}
	return docD
}*/
