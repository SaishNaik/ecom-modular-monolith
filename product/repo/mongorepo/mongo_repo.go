package mongorepo

import (
	"context"
	"github.com/SaishNaik/ecom-modular-monolith/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	// todo check if this needs to be in config
	_databaseName      = ""
	_productcollection = "product"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDBRepo(mongoUrl string) (domain.IProductRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//set up options
	clientOptions := options.Client().ApplyURI(mongoUrl)
	//todo set up auth

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return &MongoDB{client: client}, nil
}

func (m *MongoDB) GetAllProducts(ctx context.Context) ([]*domain.ProductModel, error) {
	collection := m.client.Database(_databaseName).Collection(_productcollection)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) //todo change timeout,based on config
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []*domain.ProductModel
	for cur.Next(ctx) {
		var product *domain.ProductModel
		err := cur.Decode(product)
		if err != nil {
			log.Println("Not able to decode for GetAllProducts", err)
			return nil, err
		} else {
			products = append(products, product)
		}
	}
	return products, nil
}
