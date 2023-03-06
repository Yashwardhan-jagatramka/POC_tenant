package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"

	"project-tenant/pkg/models"
)

type DbMethods interface {
	InsertOne(ctx context.Context, reqBody models.Tenant) (*mongo.InsertOneResult, error)
	TotalDocs() (int, error)
	Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
}
