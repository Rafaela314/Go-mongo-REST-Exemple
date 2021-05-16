package dao

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/xerrors"
)

//dao is the service template for db connections
type dao struct {
	db   *mongo.Database
	ctx  context.Context // context for client requests
	name string          // collection name
}

func New(name string) (*dao, error) {
	var err error
	newDao := &dao{
		ctx:  context.Background(),
		name: name,
	}

	newDao.db, err = Connect()
	if err != nil {
		log.Info("connect error in %s database", name)
		return nil, err
	}

	return newDao, err
}

// DB get mongo database
func (d *dao) DB() *mongo.Database {
	return d.db
}

// Context get the context
func (d *dao) Context() context.Context {
	return d.ctx
}

// Name get the name
func (d *dao) Name() string {
	return d.name
}

// DropDatabase from database
func (d *dao) DropDatabase() error {
	return d.DB().Collection(d.name).Drop(d.Context())
}

// getCollection returns a collection
func (d *dao) GetCollection() *mongo.Collection {
	return d.DB().Collection(d.name)
}

// Save Insert a new register
func (d *dao) Save(data interface{}, col string) error {

	_, err := d.GetCollection().InsertOne(d.ctx, data)
	if err != nil {

		return xerrors.Errorf("Unable to create a new planet: %w", err)
	}
	return nil
}
